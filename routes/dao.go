package routes
import (
	"github.com/AwespireTech/dXCA-Backend/controllers"
	"github.com/gin-gonic/gin"
)
func SetDAORoutes(router *gin.RouterGroup) {
	router.GET("/dao/:address", controllers.GetDAOByAddr)
	router.GET("/dao", controllers.GetAllDAOs)
}