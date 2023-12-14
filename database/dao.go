package database

import (
	"context"
	"errors"

	"github.com/AwespireTech/RXCA-Backend/config"
	"github.com/AwespireTech/RXCA-Backend/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InsertDAO(dao models.DAO) error {
	db := GetClient().Database(config.DATABASE_NAME).Collection("DAO")
	//Check dup address
	filter := models.DAOAddressFilter{
		Address: dao.Address,
	}
	cnt, err := db.CountDocuments(context.Background(), filter)
	if err != nil {
		return err
	}
	if cnt != 0 {
		return errors.New("DAO already exists")
	}
	dao.ID, err = AutoIncreamentID()
	if err != nil {
		return err
	}
	_, err = db.InsertOne(context.Background(), dao)
	return err
}
func GetDAOByAddress(address string) (models.DAO, error) {
	db := GetClient().Database(config.DATABASE_NAME).Collection("DAO")
	dao := models.DAO{}
	filter := models.DAOAddressFilter{
		Address: address,
	}
	err := db.FindOne(context.TODO(), filter).Decode(&dao)
	if err != nil {
		return dao, err
	}
	return dao, nil
}
func DeleteDAOByAddress(address string) error {
	db := GetClient().Database(config.DATABASE_NAME).Collection("DAO")
	filter := models.DAOAddressFilter{
		Address: address,
	}
	_, err := db.DeleteOne(context.Background(), filter)
	return err
}
func GetAllDAOs(fil models.DAOFilter, opt *options.FindOptions) ([]models.DAO, int, error) {
	db := GetClient().Database(config.DATABASE_NAME).Collection("DAO")
	var daos []models.DAO
	cnt, err := db.CountDocuments(context.Background(), fil)
	if err != nil {
		return nil, 0, err
	}

	cur, err := db.Find(context.Background(), fil, opt)
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
	db := GetClient().Database(config.DATABASE_NAME).Collection("DAO")
	update := bson.M{
		"$set": dao,
	}
	filter := models.DAOAddressFilter{
		Address: address,
	}
	_, err := db.UpdateOne(context.Background(), filter, update)
	return err
}

func AutoIncreamentID() (int, error) {
	db := GetClient().Database(config.DATABASE_NAME).Collection("DAOid")
	id := models.DAOid{}
	err := db.FindOne(context.Background(), bson.M{}).Decode(&id)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			id.ID = 2
			_, err = db.InsertOne(context.Background(), id)
			if err != nil {
				return -1, err
			}
			return 1, nil
		}
		return -1, err
	}
	res := id.ID
	update := bson.M{
		"$inc": bson.M{"id": 1},
	}
	_, err = db.UpdateOne(context.Background(), bson.M{}, update)
	if err != nil {
		return -1, err
	}
	return res, nil

}
