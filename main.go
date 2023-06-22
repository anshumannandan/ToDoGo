package main

import (
	"github.com/anshumannandan/ToDoGo/controllers"
	"github.com/anshumannandan/ToDoGo/initializers"
	"github.com/gin-gonic/gin"
)

func init(){
	initializers.LoadENV()
	initializers.ConnectDB()
}

func main(){

	r := gin.Default()

	r.POST("/todo", controllers.TodoCreate)
	r.GET("/todo", controllers.TodoList)
	r.GET("/todo/:id", controllers.TodoGet)
	r.PUT("/todo/:id", controllers.TodoUpdate)
	r.DELETE("/todo/:id", controllers.TodoDelete)

	r.Run()
}