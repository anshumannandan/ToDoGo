package main

import (
	"github.com/anshumannandan/ToDoGo/initializers"
	"github.com/anshumannandan/ToDoGo/models"
)

func init() {
	initializers.LoadENV()
	initializers.ConnectDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.ToDo{})
}