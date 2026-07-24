package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseURl string
	ServerPort string
	RedisHost string
	RedisMasterName string
	RedisPassword string
	RedisDB int
}

func Load() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found")
		return nil, err
	}

	RedisDB, _ := strconv.Atoi(os.Getenv("REDIS_DB"))

	return &Config{
		DatabaseURl: os.Getenv("DATABASE_URL"),
		ServerPort: os.Getenv("SERVER_PORT"),
		RedisHost: os.Getenv("REDIS_HOST"),
		RedisMasterName: os.Getenv("REDIS_MASTER_NAME"),
		RedisPassword: os.Getenv("REDIS_PASSWORD"),
		RedisDB: RedisDB,
	}, nil
}