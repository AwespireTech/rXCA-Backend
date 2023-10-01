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

func GetAllDAOs(fil models.DAOFilter) ([]models.DAO, int, error) {
	// Filter not supported yet
	db := GetClient().Database("dXCA").Collection("DAO")
	var daos []models.DAO
	cnt, err := db.CountDocuments(context.Background(), fil)
	if err != nil {
		return nil, 0, err
	}
	cur, err := db.Find(context.Background(), fil)
	if err != nil {
		return nil, 0, err
	}
	err = cur.All(context.Background(), &daos)
	if err != nil {
		return nil, 0, err
	}
	return daos, int(cnt), nil
}
