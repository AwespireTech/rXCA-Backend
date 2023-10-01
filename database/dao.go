package database

import (
	"context"

	"github.com/AwespireTech/dXCA-Backend/models"
)

func GetDAOByAddress(address string) (*models.DAO, error) {
	db := GetClient().Database("dXCA").Collection("DAO")
	var dao *models.DAO
	err := db.FindOne(context.TODO(), models.DAOFilter{Address: address}).Decode(dao)
	if err != nil {
		return nil, err
	}
	return dao, nil
}
