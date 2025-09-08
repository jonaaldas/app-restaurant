package main

import (
	"context"
	"log"

	"github.com/gofiber/fiber/v2"

	"github.com/joho/godotenv"
	"github.com/jonaaldas/go-restaurant-crud/database"
)

// places, err := places.GetPlaces("41.45,2.2474", 1000, "restaurant")

func main() {
	rdb := database.InitRedis()

	ctx := context.Background()
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Printf("Redis connection failed: %v", err)
	} else {
		log.Println("Redis connected successfully")
	}

	defer rdb.Close()
	app := fiber.New()

	app.Get("/api/search", func(c *fiber.Ctx) error {
		rad := c.Query("rad")
		location := c.Query("location")
		latlong := c.Query("latlong")
		resType := c.Query("resType")

		if rad == "" || location == "" || latlong == "" || resType == "" {
			return c.Status(400).JSON(fiber.Map{
				"message": "Missing required parameters",
			})
		}

		return c.JSON(fiber.Map{
			"rad":             rad,
			"location":        location,
			"latlong":         latlong,
			"restaurant type": resType,
		})
	})

	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	app.Listen(":3000")
}
