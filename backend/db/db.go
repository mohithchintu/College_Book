package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/mohithchintu/College_Book/backend/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

func MongoConnect() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI(config.AppConfig.MongoURI)

	var err error
	Client, err = mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal("Error in MongoDB connection:", err)
	}

	err = Client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("Error in MongoDB ping:", err)
	}

	fmt.Println("MongoDB connection established")
}

func MongoDisconnect() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := Client.Disconnect(ctx); err != nil {
		log.Fatal("Error disconnecting from MongoDB:", err)
	}
	fmt.Println("MongoDB connection closed")
}
