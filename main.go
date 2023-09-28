package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := createRouter()
	router.Run(":8080")
}

func createRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to dXCA API",
		})
	})
	return router
}
