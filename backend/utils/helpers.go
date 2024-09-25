package utils

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

func GetObjectId(c *fiber.Ctx) (primitive.ObjectID, error) {
	userId := c.Params("id")
	objId, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return primitive.NilObjectID, fmt.Errorf("invalid ID")
	}
	return objId, nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func CheckPasswordHash(Password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(Password))
	return err == nil
}

func SendErrorResponse(c *fiber.Ctx, status int, message string, details string) error {
	return c.Status(status).JSON(fiber.Map{"error": message, "details": details})
}

func SendSuccessResponse(c *fiber.Ctx, message string, data interface{}) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": message, "data": data})
}
