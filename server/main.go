package main

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"

	"github.com/joho/godotenv"
	"github.com/jonaaldas/go-restaurant-crud/database"
	"github.com/jonaaldas/go-restaurant-crud/handlers"
)

type GooglePhotoResponse struct {
	Name     string `json:"name"`
	PhotoUri string `json:"photoUri"`
}

func main() {
	envErr := godotenv.Load()
	if envErr != nil {
		log.Fatal("Error loading .env file")
	}

	mongoClient, err := database.InitMongo()

	if err != nil {
		log.Fatalf("Mongo connection failed: %v", err)
	}
	log.Println("Mongo connected successfully")

	dbName := "restaurant-app"
	collectionName := "restaurants"
	saved_restaurants := "saved_restaurants"
	collection := mongoClient.Database(dbName).Collection(collectionName)
	saved_restaurants_collection := mongoClient.Database(dbName).Collection(saved_restaurants)

	defer func() {
		if derr := mongoClient.Disconnect(context.Background()); derr != nil {
			log.Printf("Mongo disconnect error: %v", derr)
		}
	}()

	// Initialize handlers
	h := handlers.NewHandlers(collection, saved_restaurants_collection)

	app := fiber.New()

	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("pong")
	})

	app.Get("/img", func(c *fiber.Ctx) error {
		url := "https://places.googleapis.com/v1/places/ChIJ2fzCmcW7j4AR2JzfXBBoh6E/photos/ATKogpeivkIjQ1FT7QmbeT33nBSwqLhdPvIWHfrG1WfmgrFjeZYpS_Ls7c7rj8jejN9QGzlx4GoAH0atSvUzATDrgrZic_tTEJdeITdWL-oG3TWi5HqZoLozrjTaxoAIxmROHfV5KXVcLeTdCC6kmZExSy0CLVIG3lAPIgmvUiewNf-ZHYE4-jXYwPQpWHJgqVosvZJ6KWEgowEA-qRAzNTu9VH6BPFqHakGQ7EqBAeYOiU8Dh-xIQC8FcBJiTi0xB4tr-MYXUaF0p_AqzAhJcDE6FAgLqG1s7EsME0o36w2nDRHA-IuoISBC3SIahINE3Xwq2FzEZE6TpNTFVfgTpdPhV8CGLeqrauHn2I6ePm-2hA8-87aO7aClXKJJVzlQ1dc_JuHz6Ks07d2gglw-ZQ3ibCTF5lMtCF9O-9JHyRQXsfuXw/media?key=AIzaSyCDj-cCWpGL2Zi7hut3Uppa40-A7kCJO8Q&maxHeightPx=400&maxWidthPx=400&skipHttpRedirect=true"

		res, err := http.Get(url)

		if err != nil {
			return c.JSON(fiber.Map{
				"success": false,
				"error":   err.Error(),
			})
		}

		defer res.Body.Close()

		// read body
		body, _ := io.ReadAll(res.Body)

		var jsonResponse GooglePhotoResponse
		if err := json.Unmarshal(body, &jsonResponse); err != nil {
			log.Printf("Failed to parse JSON: %v", err)
			return c.JSON(fiber.Map{
				"success": false,
				"error":   err.Error(),
			})
		}

		return c.JSON(fiber.Map{
			"data":    jsonResponse,
			"success": true,
		})

	})

	// search restaurants
	app.Get("/api/search", h.SearchRestaurants)

	// get a restaurant by id
	app.Get("/api/restaurant/:placeId", h.GetRestaurant)

	// save a restaurant by ID
	app.Post("/api/save", h.SaveRestaurant)

	// get all saved restaurants
	app.Get("/api/restaurants", h.GetAllRestaurants)

	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	app.Listen(":3000")
}
