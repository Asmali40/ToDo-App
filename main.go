package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	log.Println("Server running on localhost:8080")

  router := gin.Default()
  router.GET("/todos", func(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
      "status": "ok",
    })
  })

	router.POST("/todos/:id", func(c *gin.Context){
		c.JSON(http.StatusCreated, gin.H{
			"status": "created",
		})
	})

	router.Run() // listens on 0.0.0.0:8080 by default

}
