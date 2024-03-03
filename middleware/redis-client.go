package middleware

import (
	"github.com/go-redis/redis/v8"
	"github.com/honestbank/tech-assignment-backend-engineer/constants"
	"github.com/joho/godotenv"
	"log"
	"os"
	"time"
)

func RedisClient() *redis.Client {
	var redisURL string
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	if os.Getenv("CLOUD") == "true" {
		redisURL = constants.REDIS_BASE_URL + ":" + constants.REDIS_PORT
	} else {
		redisURL = "redis:" + constants.REDIS_PORT
	}
	return redis.NewClient(&redis.Options{
		Addr:         redisURL,        // replace with your Redis server address
		Password:     "",              // replace with your password if any
		DB:           0,               // use default DB
		DialTimeout:  5 * time.Second, // set custom timeout
		ReadTimeout:  5 * time.Second, // set custom timeout for reading from the connection
		WriteTimeout: 5 * time.Second, // set custom timeout for writing to the connection
	})
}
