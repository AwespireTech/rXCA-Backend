package database

import (
	"context"

	"github.com/AwespireTech/dXCA-Backend/models"
	"go.mongodb.org/mongo-driver/bson"
)

func InsertDAO(dao models.DAO) error {
	db := GetClient().Database("dXCA").Collection("DAO")
	_, err := db.InsertOne(context.Background(), dao)
	return err
}
func GetDAOByAddress(address string) (models.DAO, error) {
	db := GetClient().Database("dXCA").Collection("DAO")
	dao := models.DAO{}
	err := db.FindOne(context.TODO(), models.DAOFilter{Address: address}).Decode(&dao)
	if err != nil {
		return dao, err
	}
	return dao, nil
}
func DeleteDAOByAddress(address string) error {
	db := GetClient().Database("dXCA").Collection("DAO")
	_, err := db.DeleteOne(context.Background(), models.DAOFilter{Address: address})
	return err
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
func UpdateDAOByAddress(address string, dao models.DAO) error {
	db := GetClient().Database("dXCA").Collection("DAO")
	update := bson.M{
		"$set": dao,
	}
	_, err := db.UpdateOne(context.Background(), models.DAOFilter{Address: address}, update)
	return err
}
