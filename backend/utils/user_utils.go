package utils

import (
	"context"
	"time"

	"github.com/mohithchintu/College_Book/backend/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func CheckUserExists(userID primitive.ObjectID) (bool, error) {
	UserCollection := db.GetCollection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err := UserCollection.FindOne(ctx, bson.M{"_id": userID}).Err()
	if err == mongo.ErrNoDocuments {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return true, nil
}
