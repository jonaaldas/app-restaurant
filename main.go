package main

import (
	"context"
	"log"

	"github.com/gofiber/fiber/v2"

	"github.com/joho/godotenv"
	"github.com/jonaaldas/go-restaurant-crud/database"
	"github.com/jonaaldas/go-restaurant-crud/handlers"
)

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
