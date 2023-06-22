package main

import (
	"github.com/anshumannandan/ToDoGo/initializers"
	"github.com/gin-gonic/gin"
)

func init(){
	initializers.LoadENV()
	initializers.ConnectDB()
}

func main(){

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Let's GO",
		})
	})

	r.Run()
}