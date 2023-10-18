package main

import (
	"github.com/AwespireTech/dXCA-Backend/blockchain"
	"github.com/AwespireTech/dXCA-Backend/config"
	"github.com/AwespireTech/dXCA-Backend/database"
	"github.com/AwespireTech/dXCA-Backend/routes"
	"github.com/AwespireTech/dXCA-Backend/utils"

	"github.com/gin-gonic/gin"
)

func main() {

	config.PrintConfig()
	err := database.Init(config.DATABASE_URL)
	if err != nil {
		panic(err)
	}
	err = blockchain.Init(config.ETH_RPC_URL)
	if err != nil {
		panic(err)
	}

	router := createRouter()
	router.Run()
}

func createRouter() *gin.Engine {
	router := gin.Default()
	router.Use(utils.CORSMiddleware())
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to dXCA API",
		})
	})
	api := router.Group("/api")
	routes.SetDAORoutes(api)
	routes.SetAuthRoutes(api)

	return router
}
