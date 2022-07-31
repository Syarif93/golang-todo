package main

import (
	"todo/config"
	"todo/models"
	"todo/routes"

	"gorm.io/gorm"
)

var db *gorm.DB = config.Connect()

// func addTodo(context *gin.Context) {
// 	var newTodo Todo

// 	if err := context.BindJSON(&newTodo); err != nil {
// 		return
// 	}

// 	todos = append(todos, newTodo)

// 	context.IndentedJSON(http.StatusCreated, newTodo)
// }

// func getTodo(context *gin.Context) {
// 	id := context.Param("id")
// 	todo, err := getTodoById(id)
// 	if err != nil {
// 		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
// 		return
// 	}

// 	context.IndentedJSON(http.StatusOK, todo)
// }

// func toggleTodoStatus(context *gin.Context) {
// 	id := context.Param("id")
// 	todo, err := getTodoById(id)
// 	if err != nil {
// 		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
// 		return
// 	}

// 	todo.Completed = !todo.Completed

// 	context.IndentedJSON(http.StatusOK, todo)
// }

// func getTodoById(id string) (*Todo, error) {
// 	for i, todo := range todos {
// 		if todo.Id == id {
// 			return &todos[i], nil
// 		}
// 	}

// 	return nil, errors.New("todo not found")
// }

func main() {
	defer config.Disconnect(db)

	db.AutoMigrate(&models.Todo{})

	router := routes.Router()

	router.Run()
}
