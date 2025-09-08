package database

import (
	"github.com/redis/go-redis/v9"
)

func InitRedis() *redis.Client {
	var Rdb = redis.NewClient(&redis.Options{
		// TODO: Make redis secure before production
		Addr:     "65.109.128.10:6379",
		Password: "",
		DB:       0, // use default DB
	})
	return Rdb
}

// get all restaurants
// func GetAll(rdb *redis.Client) {
// 	allRes := rdb.Get()
// }

// delete a restaurant
// set a new restaurant
// update a new restaurant comments?
