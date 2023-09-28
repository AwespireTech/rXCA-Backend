package main
import (
	"github.com/gin-gonic/gin"
)
func main() {
	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to dXCA API",
		})
	})
	router.Run(":8080")
}
