package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadConfig() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("An error has ocurred while loading .env file. ", err)
	}
}

func GetEnvVariable(key string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		log.Fatal("Error finding the env variable: " + key)
	}
	return value
}
