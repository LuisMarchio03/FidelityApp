package main

import (
	"log"

	"github.com/LuisMarchio03/acim-backend/config"
	"github.com/LuisMarchio03/acim-backend/router"
	"github.com/joho/godotenv"
)

var (
	logger *config.Logger
)

func main() {
	// Load environment variables
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading environment variables file")
	}

	logger = config.GetLogger("main")

	// Initialize the config
	err = config.Init()
	if err != nil {
		logger.Errorf("config initialization error: %v", err)
		return
	}

	// Initialize the router
	router.Initialize()
}
