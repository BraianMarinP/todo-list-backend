package routes

import (
	"todo-list-backend/handlers"

	"github.com/gin-gonic/gin"
)

var endpoint = "/todolistbackend/v1/todo"

func SetUpRoutes(router *gin.Engine) {
	router.GET(endpoint, handlers.GetTodos)
	router.POST(endpoint, handlers.CreateTodo)
	router.PUT(endpoint+"/:id", handlers.UpdateTodo)
	router.DELETE(endpoint+"/:id", handlers.DeleteTodo)
	router.DELETE(endpoint+"/all", handlers.DeleteAllTodos)
}
