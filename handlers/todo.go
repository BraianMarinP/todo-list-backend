package handlers

import (
	"net/http"
	"strconv"
	"todo-list-backend/database"
	"todo-list-backend/models"

	"github.com/gin-gonic/gin"
)

func GetTodos(c *gin.Context) {

	var todos []models.Todo
	result := database.DB.Find(&todos)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't fetch all tasks."})
		return
	}

	c.JSON(http.StatusOK, todos)
}

func CreateTodo(c *gin.Context) {

	var newTodo models.Todo

	if err := c.ShouldBindBodyWithJSON(&newTodo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Couldn't create the task."})
		return
	}

	if err := database.DB.Create(&newTodo).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Counldn't create the task."})
	}

	c.JSON(http.StatusCreated, newTodo)
}

func UpdateTodo(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var updatedTodoState models.UpdateTodo
	if err := c.ShouldBindBodyWithJSON(&updatedTodoState); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Couldn't update the task."})
		return
	}

	var todo models.Todo
	if err := database.DB.First(&todo, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
	}

	if updatedTodoState.Completed != nil {
		todo.Completed = *updatedTodoState.Completed
	}

	if err := database.DB.Save(&todo).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Couldn't update the task."})
		return
	}
	c.JSON(http.StatusOK, todo)
}

func DeleteTodo(c *gin.Context) {
	id := c.Param("id")
	var todo models.Todo

	if err := database.DB.First(&todo, id).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Couldn't delete the task."})
		return
	}
	if err := database.DB.Delete(&todo).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't delete the task."})
		return
	}
	c.Status(http.StatusNoContent)
}

func DeleteAllTodos(c *gin.Context) {
	if err := database.DB.Where("1 = 1").Delete(&models.Todo{}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't delete all tasks."})
		return
	}
	c.Status(http.StatusNoContent)
}
