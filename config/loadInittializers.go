package config

import (
	"github.com/joho/godotenv"
)

func LoadInitializers() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
}