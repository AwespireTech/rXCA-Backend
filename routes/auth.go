package routes

import (
	"github.com/AwespireTech/dXCA-Backend/controllers"
	"github.com/gin-gonic/gin"
)

func SetAuthRoutes(router *gin.RouterGroup) {
	router.GET("/auth/:address", controllers.IsAdmin)
}