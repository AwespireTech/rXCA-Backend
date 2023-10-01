package database

import (
	"testing"
	"math/rand"
	"github.com/AwespireTech/dXCA-Backend/models"
)

func InsertDAO(t *testing.T) models.DAO {
	t.Helper()
	dao := models.RandomDAO(t)
	return dao
}