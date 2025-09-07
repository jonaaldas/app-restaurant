package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"

	"github.com/joho/godotenv"
)

// places, err := places.GetPlaces("41.45,2.2474", 1000, "restaurant")

func main() {
	app := fiber.New()

	app.Get("/api/search", func(c *fiber.Ctx) error {
		rad := c.Query("rad")
		location := c.Query("location")
		latlong := c.Query("latlong")
		resType := c.Query("resType")

		fmt.Println("empty", resType)

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

	app.Listen(":3000")

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
