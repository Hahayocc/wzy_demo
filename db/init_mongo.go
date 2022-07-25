package db

import (
	"context"
	"demo/db/model"
	"fmt"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongoInstance *mongo.Client

func InitMongoDB() error {
	user := os.Getenv("MONGO_USERNAME")
	pwd := os.Getenv("MONGO_PASSWORD")
	authSource := os.Getenv("MONGO_AUTHSOURCE")
	mongoUrl := os.Getenv("MONGO_URL")
	mongoDataBase := os.Getenv("MONGO_DATABASE")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	credential := options.Credential{
		AuthSource: authSource,
		Username:   user,
		Password:   pwd,
	}


	clientOpts := options.Client().ApplyURI(mongoUrl).SetAuth(credential)

	client, err := mongo.Connect(ctx, clientOpts)

	if err != nil {
		return err
	}

	coll := client.Database(mongoDataBase).Collection("count")
	doc := &model.MongoCount{
		Type:  "mongodb",
		Count: 2022,
	}
	result, err := coll.InsertOne(context.TODO(), doc)
	if err != nil {
		panic(err)
	}
	fmt.Printf("documents inserted with result:%v\n", result)

	mongoInstance = client
	return err
}

func GetMongo() *mongo.Client {
	return mongoInstance
}
