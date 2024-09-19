package controllers

import (
	"context"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/mohithchintu/College_Book/backend/db"
	"github.com/mohithchintu/College_Book/backend/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func getObjectId(c *fiber.Ctx) (primitive.ObjectID, error) {
	userId := c.Params("id")
	objId, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return primitive.NilObjectID, fmt.Errorf("invalid user ID")
	}
	return objId, nil
}

func sendErrorResponse(c *fiber.Ctx, status int, message string, details string) error {
	return c.Status(status).JSON(fiber.Map{"error": message, "details": details})
}

func sendSuccessResponse(c *fiber.Ctx, message string, data interface{}) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": message, "data": data})
}

func CreateUser(c *fiber.Ctx) error {
	UserCollection := db.GetCollection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return sendErrorResponse(c, fiber.StatusBadRequest, "Invalid request body", err.Error())
	}

	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	_, err := UserCollection.InsertOne(ctx, user)
	if err != nil {
		return sendErrorResponse(c, fiber.StatusInternalServerError, "Error while inserting user", err.Error())
	}

	return sendSuccessResponse(c, "User created successfully", user)
}

func UpdateUserProfile(c *fiber.Ctx) error {
	UserCollection := db.GetCollection("users")
	objId, err := getObjectId(c)
	if err != nil {
		return sendErrorResponse(c, fiber.StatusBadRequest, "Invalid user ID", err.Error())
	}

	var profile models.Profile
	if err := c.BodyParser(&profile); err != nil {
		return sendErrorResponse(c, fiber.StatusBadRequest, "Invalid request body", err.Error())
	}

	update := bson.M{
		"$set": bson.M{
			"profile":   profile,
			"updatedAt": time.Now(),
		},
	}

	_, err = UserCollection.UpdateOne(context.Background(), bson.M{"id": objId}, update)
	if err != nil {
		return sendErrorResponse(c, fiber.StatusInternalServerError, "Error updating profile", err.Error())
	}

	return sendSuccessResponse(c, "Profile updated successfully", nil)
}

func UpdateUserPreferences(c *fiber.Ctx) error {
	UserCollection := db.GetCollection("users")
	objId, err := getObjectId(c)
	if err != nil {
		return sendErrorResponse(c, fiber.StatusBadRequest, "Invalid user ID", err.Error())
	}

	var preferences models.Preferences
	if err := c.BodyParser(&preferences); err != nil {
		return sendErrorResponse(c, fiber.StatusBadRequest, "Invalid request body", err.Error())
	}

	update := bson.M{
		"$set": bson.M{
			"preferences": preferences,
			"updatedAt":   time.Now(),
		},
	}

	_, err = UserCollection.UpdateOne(context.Background(), bson.M{"id": objId}, update)
	if err != nil {
		return sendErrorResponse(c, fiber.StatusInternalServerError, "Error updating preferences", err.Error())
	}

	return sendSuccessResponse(c, "Preferences updated successfully", nil)
}

func DeleteUserAccount(c *fiber.Ctx) error {
	UserCollection := db.GetCollection("users")
	objId, err := getObjectId(c)
	if err != nil {
		return sendErrorResponse(c, fiber.StatusBadRequest, "Invalid user ID", err.Error())
	}

	_, err = UserCollection.DeleteOne(context.Background(), bson.M{"id": objId})
	if err != nil {
		return sendErrorResponse(c, fiber.StatusInternalServerError, "Error deleting user", err.Error())
	}

	return sendSuccessResponse(c, "User deleted successfully", nil)
}
