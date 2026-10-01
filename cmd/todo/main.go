package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	log.Println("[INFO] Server running on localhost:8080")
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
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	router.POST("/todos", func(c *gin.Context) {
		c.JSON(http.StatusCreated, gin.H{
			"status": "created",
		})
	})

	router.Run() // listens on 0.0.0.0:8080 by default

}
