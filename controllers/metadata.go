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
		Name:        "miXed organization Certification Authority - Soulbound Token - " + dao.Name,
		Image:       "https://rxca.imlab.app/api/metadata-image",
		Description: "This is a concept verification token issued by the Taiwan Ministry of Digital Development. The token is used to symbolize the identity of a DAO that has been officially approved by the moda",
	}
	c.JSON(200, metadata)
}
func GetMetadataImage(c *gin.Context) {
	//Send Static Image
	c.File("./static/RXCA.png")
}
