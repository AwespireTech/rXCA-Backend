package controllers

import (
	"net/http"

	"github.com/AwespireTech/dXCA-Backend/database"
	"github.com/AwespireTech/dXCA-Backend/models"
	"github.com/gin-gonic/gin"
)

func GetMetadataByID(c *gin.Context) {
	id := c.Param("address")
	dao, err := database.GetDAOByAddress(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	metadata := models.Metadata{
		DAOData:     dao,
		Name:        "Token For " + dao.Name,
		Image:       "https://rxca.imlab.app/api/metadata-image",
		Description: "This is a token for DAO \"" + dao.Name + "\".\n",
	}
	c.JSON(200, metadata)
}
func GetMetadataImage(c *gin.Context) {
	//Send Static Image
	c.File("./static/RXCA.png")
}
