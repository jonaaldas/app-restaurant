package main

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"

	"github.com/joho/godotenv"
	"github.com/jonaaldas/go-restaurant-crud/database"
	"github.com/jonaaldas/go-restaurant-crud/places"
	"github.com/jonaaldas/go-restaurant-crud/types"
)

func main() {
	envErr := godotenv.Load()
	if envErr != nil {
		log.Fatal("Error loading .env file")
	}
	rdb := database.InitRedis()
	ctx := context.Background()
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Printf("Redis connection failed: %v", err)
	} else {
		log.Println("Redis connected successfully")
	}

	defer rdb.Close()

	mongoClient, err := database.InitMongo()

	if err != nil {
		log.Printf("Mongo connection failed: %v", err)
	} else {
		log.Println("Mongo connected successfully")
	}

	var dbName = "restaurant-app"
	var collectionName = "restaurants"
	collection := mongoClient.Database(dbName).Collection(collectionName)

	defer mongoClient.Disconnect(context.Background())

	app := fiber.New()

	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("pong")
	})

	app.Get("/api/search", func(c *fiber.Ctx) error {
		query := c.Query("query")
		lat := c.Query("lat")
		lng := c.Query("lng")
		rad := c.Query("radius")

		if query == "" || lat == "" || lng == "" || rad == "" {
			return c.Status(400).JSON(fiber.Map{
				"message": "Missing required parameters: query, lat, lng, radius",
			})
		}

		latitude, err := strconv.ParseFloat(lat, 64)
		if err != nil {
			return c.Status(400).JSON(fiber.Map{
				"message": "Invalid latitude value",
			})
		}

		longitude, err := strconv.ParseFloat(lng, 64)
		if err != nil {
			return c.Status(400).JSON(fiber.Map{
				"message": "Invalid longitude value",
			})
		}

		radius, err := strconv.ParseFloat(rad, 64)
		if err != nil {
			return c.Status(400).JSON(fiber.Map{
				"message": "Invalid radius value",
			})
		}

		restaurants, err := places.GetPlacesByText(query, latitude, longitude, radius)

		if err != nil {
			fmt.Print(err)
			return c.Status(400).JSON(fiber.Map{
				"message": "Error searching",
			})
		}

		return c.JSON(fiber.Map{
			"data": restaurants,
		})
	})

	app.Post("/api/save", func(c *fiber.Ctx) error {
		r := new(types.RestaurantId)

		if err := c.BodyParser(r); err != nil {
			return err
		}

		restaurant, err := database.GetRestaurant(context.Background(), rdb, r.PlaceID)
		if err != nil {
			return c.Status(400).JSON(fiber.Map{
				"message": "Restaurant not found",
			})
		}

		err = database.InsertRestaurant(context.Background(), collection, types.Restaurant{
			Name:     restaurant.Name,
			Rating:   restaurant.Rating,
			Photos:   restaurant.Photos,
			Location: restaurant.Location,
			PlaceID:  restaurant.PlaceID,
			WouldTry: restaurant.WouldTry,
			Reviews:  restaurant.Reviews,
		})

		if err != nil {
			log.Println("Error saving restaurant:", err)
			return c.Status(500).JSON(fiber.Map{"message": "Error saving restaurant"})
		}

		return c.JSON(fiber.Map{
			"message": "Restaurant saved successfully",
		})
	})

	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	app.Listen(":3000")
}
