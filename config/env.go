package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

func Init() {
	if err := godotenv.Load(); err != nil {
		log.Println("Error loading .env file")
		return
	}
	log.Println("Loading .env file")
}

func getString(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		value = defaultValue
	}
	return value
}

func getInt(key string, defaultValue int) int {
	value := os.Getenv(key)
	i, err := strconv.Atoi(value)
	if err != nil {
		return defaultValue
	}
	return i
}

func getBool(key string, defaultValue bool) bool {
	value := os.Getenv(key)
	i, err := strconv.ParseBool(value)
	if err != nil {
		return defaultValue
	}
	return i
}

type DatabaseConfig struct {
	url string
}

// Default

func NewDatabaseConfig() *DatabaseConfig {
	return &DatabaseConfig{
		url: getString("DATABASE_URL", ""),
	}
}
