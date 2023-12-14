package controllers

import (
	"github.com/AwespireTech/RXCA-Backend/blockchain"
	"github.com/gin-gonic/gin"
)

func IsAdmin(c *gin.Context) {
	addr := c.Param("address")
	isAdmin, err := blockchain.IsAdmin(addr)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"isAdmin": isAdmin,
	})
}
