package configport

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Portconfig() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")

	return port
}
