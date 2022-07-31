package controllers

import (
	"fmt"
	"net/http"

	"todo/config"
	"todo/models"

	"github.com/gin-gonic/gin"
)

type ITodoController interface {
	GetTodos(ctx *gin.Context)
	CreateTodo(ctx *gin.Context)
}

type STodoController struct{}

func TodoController() ITodoController {
	return &STodoController{}
}

var db = config.Connect()

func (controller *STodoController) GetTodos(ctx *gin.Context) {
	var todos []models.Todo
	result := db.Find(&todos)
	fmt.Sprintln("result: ", result)
	if result != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"todos": "Empty"})
		return
	} else {
		ctx.JSON(http.StatusOK, gin.H{"todos": result})
		return
	}
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
