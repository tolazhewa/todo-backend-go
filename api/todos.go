package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddTodoRoutes(r *gin.RouterGroup) {
	todos := r.Group("/todos")

	todos.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "HELLO THERE!")
	})
}

func GetTodos(c *gin.Context) {

}
