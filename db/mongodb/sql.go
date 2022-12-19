package mongodb

import (
	"context"
	"strconv"
)

func InsertOne(name string, count string) (interface{}, error){
	db := GetMongo()
	coll := db.Database(DataBase).Collection(Collection)
	cnt, _ := strconv.ParseInt(count, 10, 64)
	doc := &MongoCount{
		Name:  name,
		Count: cnt,
	}
	res, err := coll.InsertOne(context.Background(), doc)
	if err != nil {
		return nil, err
	}
	return res.InsertedID, err
}

func FindOne(name string, count string) (*MongoCount, error){
	db := GetMongo()
	coll := db.Database(DataBase).Collection(Collection)
	cnt, _ := strconv.ParseInt(count, 10, 64)
	doc := &MongoCount{
		Name:  name,
		Count: cnt,
	}
	res := coll.FindOne(context.Background(), doc)
	var mongoCount MongoCount
	if err := res.Decode(&mongoCount); err == nil {
		return &mongoCount, nil
	}
	return nil, res.Err()
}

func DeleteOne(name string, count string) (interface{}, error){
	db := GetMongo()
	coll := db.Database(DataBase).Collection(Collection)
	cnt, _ := strconv.ParseInt(count, 10, 64)
	doc := &MongoCount{
		Name:  name,
		Count: cnt,
	}
	res, err := coll.DeleteOne(context.Background(), doc)
	if err != nil {
		return nil, err
	}
	return res.DeletedCount, err
}
