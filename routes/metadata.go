package routes

import (
	"github.com/AwespireTech/dXCA-Backend/controllers"
	"github.com/gin-gonic/gin"
)

func SetMetadataRoute(router *gin.RouterGroup) {
	router.GET("/metadata/:address", controllers.GetMetadataByID)
	router.GET("/metadata-image", controllers.GetMetadataImage)
}
