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

	todos.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "HELLO THERE!")
	})
}

func GetTodos(c *gin.Context) {
	db := db.GetDB()
	queryStatement := `SELECT * FROM todos`
	rows, err := db.Query(queryStatement)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Failed to retrieve todos"})
		return
	}

	todos := make([]models.Todo, 0)
	for rows.Next() {
		var todo models.Todo
 		err = rows.Scan(&todo)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed getting todos"})
			return
		}
		todos = append(todos, todo)
	}

	if err = rows.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed getting todos"})
		return
	}
	
	c.JSON(http.StatusOK, todos)
}
