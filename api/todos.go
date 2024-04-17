package api

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tolazhewa/todo-backend-go/db"
	"github.com/tolazhewa/todo-backend-go/models"
)

func AddTodoRoutes(r *gin.RouterGroup) {
	todos := r.Group("/todos")

	todos.GET("/", getTodos)
	todos.POST("/", createTodo)
	todos.PUT("/", updateTodo)
	todos.DELETE("/:id", deleteTodo)
	todos.GET("/:id", getTodo)
}

func getTodos(c *gin.Context) {
	todos, err := db.GetTodos()
	if err != nil {
		log.Printf("failed to retrieve todos: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve todos"})
	}
	c.JSON(http.StatusOK, todos)
}

func createTodo(c *gin.Context) {
	var todo models.Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		log.Printf("failed to parse todo during create: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse todo"})
		return
	}

	todo, err := db.CreateTodo(todo)
	if err != nil {
		log.Printf("failed to create todo: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create todo"})
		return
	}

	c.JSON(http.StatusCreated, todo)
}

func updateTodo(c *gin.Context) {
	var todo models.Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		log.Printf("failed to parse todo during update: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "can't parse input"})
		return
	}

	todo, err := db.UpdateTodo(todo)
	if err != nil {
		log.Printf("failed to update todo: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update todo"})
		return
	}

	c.JSON(http.StatusOK, todo)
}

func deleteTodo(c *gin.Context) {
	todoId := c.Param("id")
	err := db.DeleteTodo(todoId)
	if err != nil {
		log.Printf("failed to delete todo: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete todo"})
		return
	}
	c.Status(http.StatusOK)
}

func getTodo(c *gin.Context) {
	todoId := c.Param("id")
	todo, err := db.GetTodo(todoId)
	if err != nil {
		log.Printf("failed to get todo: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve todo"})
		return
	}
	c.JSON(http.StatusOK, todo)
}
