package routes

import (
	"net/http"

	"todo/controllers"

	"github.com/gin-gonic/gin"
)

var todoController = controllers.TodoController

func ApiRoutes(router *gin.Engine) *gin.RouterGroup {
	apiRoutes := router.Group("/api")
	{
		apiRoutes.GET("/", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{"message": "Api online!"})
		})

		apiRoutes.GET("/todos", todoController().GetTodos)
		apiRoutes.POST("/todo", todoController().CreateTodo)
		// router.GET("/todo/:id", getTodo)
		// router.PUT("/todo/:id/toggle-status", toggleTodoStatus)
	}

	return apiRoutes
}
