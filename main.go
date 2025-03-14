package main

import (
	"log"
	"main/database"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	router := gin.Default()

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	log.Println("Database Intializing Started...")
	database.Initialize()
	log.Println("Database Intializing Stopped...")

	port := os.Getenv("PORT")
	if port == "" {
		port = os.Getenv("PORT")
	}
	router.Run(":" + port)
}
