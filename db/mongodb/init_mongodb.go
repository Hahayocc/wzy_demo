package mongodb

import (
	"context"
	"fmt"
	"net/url"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongoInstance *mongo.Client

func InitMongoDB() {
	user := os.Getenv("MONGO_USERNAME")
	pwd := os.Getenv("MONGO_PASSWORD")
	mongoAddress := os.Getenv("MONGO_ADDRESS")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	tmp, err := url.Parse(mongoAddress)
	if err != nil {
		fmt.Printf("mongoAddress parse error err %v", err)
	}
	authSource := tmp.Query().Get("authSource")

	credential := options.Credential{
		AuthSource: authSource,
		Username:   user,
		Password:   pwd,
	}

	mongoUrl := fmt.Sprintf("mongodb://%s", mongoAddress)

	clientOpts := options.Client().ApplyURI(mongoUrl).SetAuth(credential)

	client, err := mongo.Connect(ctx, clientOpts)

	if err != nil {
		fmt.Printf("mongo connect error err %v", err)
	}
	mongoInstance = client
}

func GetMongo() *mongo.Client {
	return mongoInstance
}
