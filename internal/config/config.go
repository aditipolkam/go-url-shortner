package config

import (
	"os"
	// "github.com/joho/godotenv"
)

type Config struct {
	DB_URL        string
	ServerAddress string
}

func Load() Config {
	// if err := godotenv.Load(); err != nil {
	// 	log.Println("No .env file found, loading environment variables directly")
	// }

	return Config{
		DB_URL:        os.Getenv("DB_URL"),
		ServerAddress: ":8080",
	}
}
