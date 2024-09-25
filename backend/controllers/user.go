package controllers

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/mohithchintu/College_Book/backend/db"
	"github.com/mohithchintu/College_Book/backend/models"
	"github.com/mohithchintu/College_Book/backend/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateUser(c *fiber.Ctx) error {
	UserCollection := db.GetCollection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return utils.SendErrorResponse(c, fiber.StatusBadRequest, "Invalid request body", err.Error())
	}

	if user.Email == "" {
		return utils.SendErrorResponse(c, fiber.StatusBadRequest, "Email is required", "One or more required fields are empty")
	}

	if user.Username == "" {
		return utils.SendErrorResponse(c, fiber.StatusBadRequest, "Username is required", "One or more required fields are empty")
	}

	if user.Password == "" {
		return utils.SendErrorResponse(c, fiber.StatusBadRequest, "Password is required", "One or more required fields are empty")
	}
	existingUser := &models.User{}
	err := UserCollection.FindOne(ctx, bson.M{"email": user.Email}).Decode(existingUser)
	if err == nil {
		return utils.SendErrorResponse(c, fiber.StatusConflict, "User already exists", "A user with this email already exists")
	}

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return utils.SendErrorResponse(c, fiber.StatusInternalServerError, "Error hashing password", err.Error())
	}
	user.Password = hashedPassword

	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	data, err := UserCollection.InsertOne(ctx, user)
	if err != nil {
		return utils.SendErrorResponse(c, fiber.StatusInternalServerError, "Error while inserting user", err.Error())
	}
	user.ID = data.InsertedID.(primitive.ObjectID)
	return utils.SendSuccessResponse(c, "User created successfully", user)
}

func UpdateUserProfile(c *fiber.Ctx) error {
	UserCollection := db.GetCollection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objId, err := utils.GetObjectId(c)
	if err != nil {
		return utils.SendErrorResponse(c, fiber.StatusBadRequest, "Invalid user ID", err.Error())
	}

	var profile models.Profile
	if err := c.BodyParser(&profile); err != nil {
		return utils.SendErrorResponse(c, fiber.StatusBadRequest, "Invalid request body", err.Error())
	}

	update := bson.M{
		"$set": bson.M{
			"profile":   profile,
			"updatedAt": time.Now(),
		},
	}

	result, err := UserCollection.UpdateOne(ctx, bson.M{"_id": objId}, update)
	if err != nil {
		return utils.SendErrorResponse(c, fiber.StatusInternalServerError, "Error updating profile", err.Error())
	}
	if result.MatchedCount == 0 {
		return utils.SendErrorResponse(c, fiber.StatusNotFound, "User not found", "No user found with the provided ID")
	}

	return utils.SendSuccessResponse(c, "Profile updated successfully", nil)
}

func UpdateUserPreferences(c *fiber.Ctx) error {
	UserCollection := db.GetCollection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objId, err := utils.GetObjectId(c)
	if err != nil {
		return utils.SendErrorResponse(c, fiber.StatusBadRequest, "Invalid user ID", err.Error())
	}

	var preferences models.Preferences
	if err := c.BodyParser(&preferences); err != nil {
		return utils.SendErrorResponse(c, fiber.StatusBadRequest, "Invalid request body", err.Error())
	}

	update := bson.M{
		"$set": bson.M{
			"preferences": preferences,
			"updatedAt":   time.Now(),
		},
	}

	result, err := UserCollection.UpdateOne(ctx, bson.M{"_id": objId}, update)
	if err != nil {
		return utils.SendErrorResponse(c, fiber.StatusInternalServerError, "Error updating preferences", err.Error())
	}
	if result.MatchedCount == 0 {
		return utils.SendErrorResponse(c, fiber.StatusNotFound, "User not found", "No user found with the provided ID")
	}

	return utils.SendSuccessResponse(c, "Preferences updated successfully", nil)
}

func DeleteUserAccount(c *fiber.Ctx) error {
	UserCollection := db.GetCollection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objId, err := utils.GetObjectId(c)
	if err != nil {
		return utils.SendErrorResponse(c, fiber.StatusBadRequest, "Invalid user ID", err.Error())
	}

	result, err := UserCollection.DeleteOne(ctx, bson.M{"_id": objId})
	if err != nil {
		return utils.SendErrorResponse(c, fiber.StatusInternalServerError, "Error deleting user", err.Error())
	}
	if result.DeletedCount == 0 {
		return utils.SendErrorResponse(c, fiber.StatusNotFound, "User not found", "No user found with the provided ID")
	}

	return utils.SendSuccessResponse(c, "User deleted successfully", nil)
}
