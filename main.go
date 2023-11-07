package main

import (
	"auth-users-aws/services/cache"
	"auth-users-aws/services/database"
	"auth-users-aws/services/http"
	"fmt"

	"github.com/Valgard/godotenv"
)

func init() {
	dotenv := godotenv.New()
	if err := dotenv.Load(".env"); err != nil {
		panic(err)
	}

	sqlConn := database.Connection()
	if sqlConn != nil {
		fmt.Println("Database Connection Success!")
		sqlConn.Close()
	}

	redisConn := cache.RedisConnection()
	if redisConn != nil {
		redisConn.Close()
		fmt.Println("Redis Connection Success!")
	}

	// database.ValidTableCliente()

	// usecase.UpdateCache()

	http.HttpClient()

}

func main() {

}
