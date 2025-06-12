package main

import (
	"time"
	"todo-list-backend/config"
	"todo-list-backend/database"
	"todo-list-backend/models"
	"todo-list-backend/routes"

	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
)

func main() {

	//Load environment
	config.LoadConfig()
	database.Connect()
	database.DB.AutoMigrate(&models.Todo{})

	// Initialize gin
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"}, // Cambia esto según tu frontend
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	routes.SetUpRoutes(r)
	r.Run(":8080")

}
