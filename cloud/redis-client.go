package cloud

import (
	"github.com/go-redis/redis/v8"
	"github.com/honestbank/tech-assignment-backend-engineer/constants"
	"time"
)

var redisURL string = constants.REDIS_BASE_URL + ":" + constants.REDIS_PORT

// Initialize a new Redis client with a custom timeout
func RedisClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:         redisURL,        // replace with your Redis server address
		Password:     "",              // replace with your password if any
		DB:           0,               // use default DB
		DialTimeout:  5 * time.Second, // set custom timeout
		ReadTimeout:  5 * time.Second, // set custom timeout for reading from the connection
		WriteTimeout: 5 * time.Second, // set custom timeout for writing to the connection
	})
}
