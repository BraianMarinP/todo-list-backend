package main

import (
	"todo-list-backend/config"
	"todo-list-backend/database"
	"todo-list-backend/models"
	"todo-list-backend/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	//Load environment
	config.LoadConfig()
	database.Connect()
	database.DB.AutoMigrate(&models.Todo{})

	// Initialize gin
	r := gin.Default()
	routes.SetUpRoutes(r)
	r.Run(":8080")

}
