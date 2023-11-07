package cache

import (
	"fmt"
	"os"

	"github.com/redis/go-redis/v9"
)

type redis_credentials struct {
	DB       int
	Password string
	Addr     string
}

func RedisConnection() *redis.Client {
	var credentials redis_credentials
	credentials.Addr = fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT"))
	credentials.Password = ""
	credentials.DB = 0

	client := redis.NewClient(&redis.Options{
		Addr:     credentials.Addr,
		Password: credentials.Password,
		DB:       credentials.DB,
	})

	return client
}
