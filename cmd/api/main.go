package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"ToDo-App/internal/todo"
)

func main() {
	log.Println("[INFO] Server running on localhost:8080")

	fmt.Println(`
  _____     ____                _                
 |_   _|__ |  _ \  ___         / \   _ __  _ __  
   | |/ _ \| | | |/ _ \ _____ / _ \ | '_ \| '_ \ 	
   | | (_) | |_| | (_) |_____/ ___ \| |_) | |_) |
   |_|\___/|____/ \___/     /_/   \_\ .__/| .__/ 
                                    |_|   |_|  
	`)

	fmt.Println(`
		╔══════════════════════════════════════╗
		║             TODO API                 ║
		║                                      ║
		║   🚀 Server started successfully     ║
		║   🌐 http://localhost:8080           ║
		╚══════════════════════════════════════╝
	`)

	router := gin.Default()

	router.GET("/todos", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Hello from Gin!"})
	})

	router.POST("/todos", func(c *gin.Context) {
		var todoItem todo.Todo

		if err := c.ShouldBindJSON(&todoItem); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusCreated, todoItem)
	})

	err := router.Run(":8080")
	if err != nil {
		println("The router encountered an error while running\n", err)
	}
}
