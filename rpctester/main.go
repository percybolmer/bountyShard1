package main

import (
	"log"
	"os"
	"percybolmer/rpc-shard-testing/rpctester/cmd"

	"github.com/joho/godotenv"
)

var (
	address string
	url     string
)

func init() {
	// Load .dotenv file
	environment := os.Getenv("RPCTESTER_ENVIRONMENT")

	if environment == "production" {
		if err := godotenv.Load(".env-prod"); err != nil {
			log.Fatal("Error loading .env-prod file")
		}
	} else {
		if err := godotenv.Load(".env"); err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	address = os.Getenv("ADDRESS")
	url = os.Getenv("NET_URL")
}
func main() {
	cmd.Execute()
}
