package controllers

import (
	"net/http"

	"github.com/AwespireTech/dXCA-Backend/database"
	"github.com/AwespireTech/dXCA-Backend/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetDAOByAddr(c *gin.Context) {
	address := c.Param("address")
	dao, err := database.GetDAOByAddress(address)
	if err != nil {
		if err.Error() == mongo.ErrNoDocuments.Error() {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "DAO not found",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, dao)
}

func GetAllDAOs(c *gin.Context) {
	fil := models.DAOFilter{}
	daos, cnt, err := database.GetAllDAOs(fil)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	resp := models.DAOsResponse{
		Count: cnt,
		DAOs:  daos,
	}
	c.JSON(http.StatusOK, resp)
}
