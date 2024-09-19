package db

import (
	"context"
	"fmt"
	"time"

	"github.com/mohithchintu/College_Book/backend/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

func MongoConnect() error {
	clientOptions := options.Client().ApplyURI(config.AppConfig.MongoURI)

	var err error
	Client, err = mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return fmt.Errorf("error in MongoDB connection: %v", err)
	}

	err = Client.Ping(context.Background(), nil)
	if err != nil {
		return fmt.Errorf("error in MongoDB ping: %v", err)
	}

	fmt.Println("MongoDB connection established")
	return nil
}

func Geth(a string) int {
	fmt.Println(Client)
	fmt.Println(a)
	return 4
}

func GetCollection(collectionName string) *mongo.Collection {
	if Client == nil {
		fmt.Println("MongoDB client is not initialized")
	}
	collection := Client.Database("cb-database").Collection(collectionName)
	return collection
}

func MongoDisconnect() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := Client.Disconnect(ctx); err != nil {
		return fmt.Errorf("error disconnecting from MongoDB: %v", err)
	}
	fmt.Println("MongoDB connection closed")
	return nil
}
