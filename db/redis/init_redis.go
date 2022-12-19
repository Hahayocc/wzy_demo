package redis

import (
	"os"

	"github.com/go-redis/redis/v8"
)

var redisInstance *redis.Client

func InitRedis() {
	user := os.Getenv("REDIS_USERNAME")
	pwd := os.Getenv("REDIS_PASSWORD")
	addr := os.Getenv("REDIS_ADDRESS")

	redisInstance = redis.NewClient(&redis.Options{
		Addr:     addr,
		Username: user,
		Password: pwd, // no password set
		DB:       0,   // use default DB
	})
}

func GetRedis() *redis.Client {
	return redisInstance
}
