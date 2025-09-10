package database

import (
	"context"
	"fmt"
	"os"

	"github.com/jonaaldas/go-restaurant-crud/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InitMongo() (*mongo.Client, error) {
func InitMongo() (*mongo.Client, error) {
    mongoURL := os.Getenv("MONGO_URL")
    serverAPI := options.ServerAPI(options.ServerAPIVersion1)
    // …
}
	opts := options.Client().ApplyURI(mongoURL).SetServerAPIOptions(serverAPI)
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		return nil, err
	}
	// Send a ping to confirm a successful connection
	if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Err(); err != nil {
		return nil, err
	}
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")
	return client, nil
}

func InsertRestaurant(ctx context.Context, collection *mongo.Collection, restaurant types.Restaurant) error {
	var _, err = collection.InsertOne(ctx, restaurant)

	if err != nil {
		fmt.Println("Error inserting restaurant:", err)
		return err
	}

	return nil
}
