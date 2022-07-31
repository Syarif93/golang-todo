package controllers

import (
	"fmt"
	"net/http"

	"todo/config"
	"todo/helpers"
	"todo/models"

	"github.com/gin-gonic/gin"
)

type ITodoController interface {
	GetTodos(ctx *gin.Context)
	CreateTodo(ctx *gin.Context)
	GetOneTodo(ctx *gin.Context)
}

type STodoController struct{}

func TodoController() ITodoController {
	return &STodoController{}
}

var db = config.Connect()

func (controller *STodoController) GetTodos(ctx *gin.Context) {
	var todos []models.Todo
	if err := db.Find(&todos).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"todos": "Empty"})
		return
	}
	ctx.IndentedJSON(http.StatusOK, helpers.SuccessResponse("GetTodosSuccess", todos))
	return
}

func (controller *STodoController) CreateTodo(ctx *gin.Context) {
	todo := []models.Todo{{Item: "Test", Completed: false}}
	result := db.Create(&todo)
	if result != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"todos": "Empty"})
		return
	} else {
		ctx.JSON(http.StatusCreated, gin.H{"todo": result})
		return
	}
}

func (controller *STodoController) GetOneTodo(ctx *gin.Context) {
	id := ctx.Param("id")
	var todo []models.Todo
	result := db.First(&todo, id)
	fmt.Sprintln(result)
	if result != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"todos": "Empty"})
		return
	} else {
		ctx.JSON(http.StatusCreated, gin.H{"todo": result})
		return
	}
}
