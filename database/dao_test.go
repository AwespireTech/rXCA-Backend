package database

import (
	"testing"

	"github.com/AwespireTech/dXCA-Backend/models"
)

func InsertRandomDAO(t *testing.T) models.DAO {
	t.Helper()
	dao := models.RandomDAO(t)
	return dao
}
