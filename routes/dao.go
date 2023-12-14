package routes

import (
	"github.com/AwespireTech/RXCA-Backend/controllers"
	"github.com/gin-gonic/gin"
)

func SetDAORoutes(router *gin.RouterGroup) {
	router.GET("/dao/:address", controllers.GetDAOByAddr)
	router.GET("/dao", controllers.GetAllDAOs)
	router.POST("/dao", controllers.CreateDAO)
	router.DELETE("/dao/:address", controllers.CancelDAO)
	router.POST("/dao/:address", controllers.ValidateDAOByAddr)
	router.PUT("/dao/:address/revoke", controllers.RevokeDAOByAddr)
}
