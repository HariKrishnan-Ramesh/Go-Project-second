package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	router := gin.Default()

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v",err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = os.Getenv("PORT")
	}
	router.Run(":" + port)
}