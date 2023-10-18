package controllers

import (
	"net/http"
	"strings"

	"github.com/AwespireTech/dXCA-Backend/blockchain"
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
func CancelDAO(c *gin.Context) {
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
	if dao.State != models.DAO_STATE_PENDING {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "DAO is not pending",
		})
		return
	}
	err = database.DeleteDAOByAddress(address)
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
	c.JSON(http.StatusOK, gin.H{
		"message": "DAO successfully deleted",
	})
}
func CreateDAO(c *gin.Context) {
	dao := models.DAO{}
	err := c.ShouldBindJSON(&dao)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	dao.Address = strings.ToLower(dao.Address)

	err = database.InsertDAO(dao)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			c.JSON(http.StatusConflict, gin.H{
				"error": "DAO already exists",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, dao)
}
func ValidateDAOByAddr(c *gin.Context) {
	address := c.Param("address")
	val := models.DAOVerifyRequest{}
	err := c.ShouldBindJSON(&val)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	if val.Validate {
		addr, tid, err := blockchain.DecodeMintTransaction(val.TxHash)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		if addr != address {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Transaction is not minting to the correct address",
			})
			return
		}

		dao := models.DAO{
			State:   models.DAO_STATE_APPROVED,
			TokenId: tid,
		}
		err = database.UpdateDAOByAddress(address, dao)
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
		c.JSON(http.StatusOK, gin.H{
			"message": "DAO successfully validated",
		})
		return
	} else {
		dao := models.DAO{
			State:   models.DAO_STATE_DENIED,
			TokenId: -1,
		}
		err = database.UpdateDAOByAddress(address, dao)
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
		c.JSON(http.StatusOK, gin.H{
			"message": "DAO successfully denied",
		})
		return
	}
}
func RevokeDAOByAddr(c *gin.Context) {
	address := c.Param("address")
	oriDAO, err := database.GetDAOByAddress(address)
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
	val := models.DAORevokeRequest{}
	err = c.ShouldBindJSON(&val)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	addr, tid, err := blockchain.DecodeBurnTransaction(val.TxHash)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	if addr != address || tid != oriDAO.TokenId {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Transaction is not buring of correct address",
		})
		return
	}
	update := models.DAO{
		State:   models.DAO_STATE_DENIED,
		TokenId: -1,
	}
	err = database.UpdateDAOByAddress(address, update)
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
	c.JSON(http.StatusOK, gin.H{
		"message": "DAO successfully revoked",
	})
	return

}
