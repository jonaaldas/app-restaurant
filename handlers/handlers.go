package handlers

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/jonaaldas/go-restaurant-crud/database"
	"github.com/jonaaldas/go-restaurant-crud/places"
	"github.com/jonaaldas/go-restaurant-crud/types"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
)

type Handlers struct {
	RedisClient     *redis.Client
	MongoCollection *mongo.Collection
}

func NewHandlers(redisClient *redis.Client, mongoCollection *mongo.Collection) *Handlers {
	return &Handlers{
		RedisClient:     redisClient,
		MongoCollection: mongoCollection,
	}
}

// SearchRestaurants handles GET /api/search
func (h *Handlers) SearchRestaurants(c *fiber.Ctx) error {
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

	restaurants, err := places.GetPlacesByText(query, latitude, longitude, radius, h.RedisClient)

	if err != nil {
		fmt.Print(err)
		return c.Status(400).JSON(fiber.Map{
			"message": "Error searching",
		})
	}

	return c.JSON(fiber.Map{
		"data": restaurants,
	})
}

// GetRestaurant handles GET /api/restaurant/:placeId
func (h *Handlers) GetRestaurant(c *fiber.Ctx) error {
	placeId := c.Params("placeId")

	var restaurant *types.Restaurant
	var err error
	var source string

	// Step 1: Try Redis cache first
	restaurant, err = database.GetRestaurant(context.Background(), h.RedisClient, placeId)
	if err != nil {
		log.Printf("Restaurant not found in Redis cache: %v", err)

		// Step 2: Try MongoDB if not in cache
		restaurant, err = database.GetRestaurantFromMongo(context.Background(), h.MongoCollection, placeId)
		if err != nil {
			log.Printf("Restaurant not found in MongoDB: %v", err)
			return c.Status(404).JSON(fiber.Map{
				"message": "Restaurant not found",
				"placeID": placeId,
			})
		}

		// Step 3: Cache the restaurant from MongoDB to Redis for next time
		go func() {
			success := database.SetRestaurant(context.Background(), h.RedisClient, restaurant.PlaceID, *restaurant)
			if success {
				log.Printf("Successfully cached restaurant %s from MongoDB to Redis", restaurant.PlaceID)
			}
		}()

		log.Printf("Found restaurant in MongoDB: %s", placeId)
		source = "mongodb"
	} else {
		log.Printf("Found restaurant in Redis cache: %s", placeId)
		source = "redis"
	}

	return c.JSON(fiber.Map{
		"data":   restaurant,
		"source": source,
	})
}

// SaveRestaurant handles POST /api/save
func (h *Handlers) SaveRestaurant(c *fiber.Ctx) error {
	r := new(types.RestaurantId)

	if err := c.BodyParser(r); err != nil {
		log.Printf("Error parsing request body: %v", err)
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	log.Printf("Attempting to save restaurant with PlaceID: %s", r.PlaceID)

	if r.PlaceID == "" {
		return c.Status(400).JSON(fiber.Map{
			"message": "PlaceID is required",
		})
	}

	var restaurant *types.Restaurant
	var err error

	// Step 1: Try Redis cache first
	restaurant, err = database.GetRestaurant(context.Background(), h.RedisClient, r.PlaceID)
	if err != nil {
		log.Printf("Restaurant not found in Redis cache: %v", err)

		// Step 2: Try MongoDB if not in cache
		restaurant, err = database.GetRestaurantFromMongo(context.Background(), h.MongoCollection, r.PlaceID)
		if err != nil {
			log.Printf("Restaurant not found in MongoDB: %v", err)
			return c.Status(404).JSON(fiber.Map{
				"message": "Restaurant not found",
				"placeID": r.PlaceID,
			})
		}

		// Step 3: Cache the restaurant from MongoDB to Redis for next time
		go func() {
			success := database.SetRestaurant(context.Background(), h.RedisClient, restaurant.PlaceID, *restaurant)
			if success {
				log.Printf("Successfully cached restaurant %s from MongoDB to Redis", restaurant.PlaceID)
			} else {
				log.Printf("Failed to cache restaurant %s to Redis", restaurant.PlaceID)
			}
		}()

		log.Printf("Found restaurant in MongoDB and cached to Redis: %s", r.PlaceID)
	} else {
		log.Printf("Found restaurant in Redis cache: %s", r.PlaceID)
	}

	// Update the WouldTry field with the user's preference
	restaurant.WouldTry = r.WouldTry

	err = database.UpsertRestaurant(context.Background(), h.MongoCollection, *restaurant)

	if err != nil {
		log.Println("Error saving restaurant:", err)
		return c.Status(500).JSON(fiber.Map{"message": "Error saving restaurant"})
	}

	return c.JSON(fiber.Map{
		"message": "Restaurant saved successfully",
	})
}
