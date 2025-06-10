package handlers

import (
	"net/http"
	"todo-list-backend/database"
	"todo-list-backend/models"

	"github.com/gin-gonic/gin"
)

func GetTodos(c *gin.Context) {
	var todos []models.Todo
	database.DB.Find(&todos)
	c.JSON(http.StatusOK, todos)
}

func CreateTodo(c *gin.Context) {
	var newTodo models.Todo
	if err := c.ShouldBindBodyWithJSON(&newTodo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Couldn't create the task."})
		return
	}
	database.DB.Create(&newTodo)
	c.JSON(http.StatusCreated, newTodo)
}

func UpdateTodo(c *gin.Context) {
	var updatedTodo models.Todo
	if err := c.ShouldBindBodyWithJSON(&updatedTodo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Couldn't update the task."})
		return
	}
	database.DB.Save(&updatedTodo)
	c.JSON(http.StatusOK, updatedTodo)
}

func DeleteTodo(c *gin.Context) {
	id := c.Param("id")
	var todo models.Todo

	if err := database.DB.First(&todo, id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Couldn't delete the task."})
		return
	}
	database.DB.Delete(&todo)
	c.JSON(http.StatusOK, todo)
}

func DeleteAllTodos(c *gin.Context) {
	database.DB.Where("1 = 1").Delete(&models.Todo{})
	c.Status(http.StatusNoContent)
}
