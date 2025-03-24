package main

import (
	"log"
	"main/database"
	"main/handlers"
	"main/managers"
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

	otpManager := managers.NewOtpManager()
	otpHandler := handlers.NewOtpHandler(otpManager)
	otpHandler.RegisterOtpApis(router)

	adminManager := managers.NewAdminManager()
	adminHandler := handlers.NewAdminHandler(adminManager)
	adminHandler.RegisterAdminApis(router)


	port := os.Getenv("PORT")
	if port == "" {
		port = os.Getenv("PORT")
	}
	router.Run(":" + port)
}
