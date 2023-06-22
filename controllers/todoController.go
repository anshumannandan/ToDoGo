package controllers

import (
	"strings"
	"github.com/anshumannandan/ToDoGo/initializers"
	"github.com/anshumannandan/ToDoGo/models"
	"github.com/gin-gonic/gin"
)

func TodoCreate(c *gin.Context) {
	
	var request struct{
		Title string
		Description string
	}
	c.Bind(&request)

	if ( strings.TrimSpace(request.Title) == "" || strings.TrimSpace(request.Description) == "") {
		c.JSON(400, gin.H{"message": "Title and Description are required fields"})
		return
	}

	todo := models.ToDo{Title: request.Title, Description: request.Description}
	initializers.DB.Create(&todo)

	c.JSON(201, todo)
}

func TodoList(c *gin.Context) {

	var todos []models.ToDo
	initializers.DB.Find(&todos)

	c.JSON(200, todos)
}

func TodoGet(c *gin.Context) {

	todo_id := c.Param("id")

	var todo models.ToDo
	initializers.DB.Find(&todo, todo_id)

	if todo.ID == 0 {
		c.JSON(404, gin.H{"message": "Todo with requested ID not found"})
		return
	}

	c.JSON(200, todo)
}

func TodoUpdate(c *gin.Context) {

	todo_id := c.Param("id")

	var todo models.ToDo
	initializers.DB.Find(&todo, todo_id)

	if todo.ID == 0 {
		c.JSON(404, gin.H{"message": "Todo with requested ID not found"})
		return
	}

	var request struct{
		Title string
		Description string
	}
	c.Bind(&request)

	initializers.DB.Model(&todo).Updates(models.ToDo{Title: strings.TrimSpace(request.Title), 
		Description: strings.TrimSpace(request.Description)})

	c.JSON(200, todo)
}

func TodoDelete(c *gin.Context) {

	todo_id := c.Param("id")

	var todo models.ToDo
	initializers.DB.Find(&todo, todo_id)

	if todo.ID == 0 {
		c.JSON(404, gin.H{"message": "Todo with requested ID not found"})
		return
	}

	initializers.DB.Delete(&todo)

	c.Status(200)
}