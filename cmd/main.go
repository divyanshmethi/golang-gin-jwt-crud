package main

import (
	"fmt"
	"log"

	"RESTAPI_Gin/controllers"
	"RESTAPI_Gin/db"
	"RESTAPI_Gin/middleware"
	"RESTAPI_Gin/models"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	loadEnv()
	loadDatabase()
	serveApplication()
}

func serveApplication() {
	router := gin.Default()
	publicRoutes := router.Group("/auth")
	publicRoutes.POST("/register", controllers.Register)
	publicRoutes.POST("/login", controllers.Login)
	protectedRoutes := router.Group("/api")
	protectedRoutes.Use(middleware.JWTAuthMiddleware())
	protectedRoutes.POST("/entry", controllers.AddEntry)
	protectedRoutes.GET("/entry", controllers.GetAllEntries)
	router.Run(":8000")
	fmt.Println("Server running on port 8000")
}

func loadDatabase() {
	db.Connect()
	err := db.Database.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal("Unable to create User table")
	}
	err = db.Database.AutoMigrate(&models.Entry{})
	if err != nil {
		log.Fatal("Unable to create Entry table")
	}
}

func loadEnv() {
	err := godotenv.Load(".env.local")
	if err != nil {
		log.Fatal("Error while loading .env file")
	}
}
