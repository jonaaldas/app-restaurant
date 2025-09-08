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

// places, err := places.GetPlaces("41.45,2.2474", 1000, "restaurant")

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
		rad := c.Query("rad")
		latlong := c.Query("latlong")
		resType := c.Query("resType")
		radNumber, err := strconv.Atoi(rad)

		if err != nil {
			fmt.Printf("Error converting string to int: %v\n", err)
			return c.Status(400).JSON(fiber.Map{
				"message": "There was an error please reload the page.",
			})
		}

		if rad == "" || latlong == "" || resType == "" {
			return c.Status(400).JSON(fiber.Map{
				"message": "Missing required parameters",
			})
		}

		places, err := places.GetPlaces(latlong, radNumber, resType)
		if err != nil {
			fmt.Print(err)
			return c.Status(400).JSON(fiber.Map{
				"message": "Error searching",
			})
		}

		return c.JSON(fiber.Map{
			"data": places,
		})
	})

	app.Post("/api/save", func(c *fiber.Ctx) error {
		r := new(types.Restaurant)

		if err := c.BodyParser(r); err != nil {
			return err
		}

		if c.Query("wouldTry") == "true" {
			r.WouldTry = true
		}

		err := database.InsertRestaurant(context.Background(), collection, *r)

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
