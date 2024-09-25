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

func CreateMaterial(c *fiber.Ctx) error {
	MaterialCollection := db.GetCollection("materials")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var material models.Material
	if err := c.BodyParser(&material); err != nil {
		return utils.SendErrorResponse(c, fiber.StatusBadRequest, "Invalid request body", err.Error())
	}

	userExists, err := utils.CheckUserExists(material.UserID)
	if err != nil {
		return utils.SendErrorResponse(c, fiber.StatusInternalServerError, "Error while checking user existence", err.Error())
	}

	if !userExists {
		return utils.SendErrorResponse(c, fiber.StatusNotFound, "User not found", "The specified user does not exist")
	}

	if len(material.Notes) == 0 && len(material.PDFs) == 0 && len(material.Goals) == 0 {
		return utils.SendErrorResponse(c, fiber.StatusBadRequest, "At least one Note, PDF, or Goal must be provided", "Empty notes, pdfs, and goals")
	}

	material.CreatedAt = time.Now()
	material.UpdatedAt = time.Now()

	data, err := MaterialCollection.InsertOne(ctx, material)
	if err != nil {
		return utils.SendErrorResponse(c, fiber.StatusInternalServerError, "Error while inserting material", err.Error())
	}

	material.ID = data.InsertedID.(primitive.ObjectID)

	return utils.SendSuccessResponse(c, "Material created successfully", material)
}

func GetMaterials(c *fiber.Ctx) error {
	MaterialCollection := db.GetCollection("materials")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var materials []models.Material

	cursor, err := MaterialCollection.Find(ctx, bson.M{})
	if err != nil {
		return utils.SendErrorResponse(c, fiber.StatusInternalServerError, "Error fetching materials", err.Error())
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var material models.Material
		if err := cursor.Decode(&material); err != nil {
			return utils.SendErrorResponse(c, fiber.StatusInternalServerError, "Error decoding material", err.Error())
		}
		materials = append(materials, material)
	}

	if len(materials) == 0 {
		return utils.SendErrorResponse(c, fiber.StatusNotFound, "No materials found", "Empty materials list")
	}

	return utils.SendSuccessResponse(c, "Materials fetched successfully", materials)
}

func GetMaterialByID(c *fiber.Ctx) error {
	MaterialCollection := db.GetCollection("materials")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	materialID := c.Params("id")
	mid, err := primitive.ObjectIDFromHex(materialID)
	if err != nil {
		return utils.SendErrorResponse(c, fiber.StatusBadRequest, "Invalid material ID", err.Error())
	}

	var material models.Material

	err = MaterialCollection.FindOne(ctx, bson.M{"_id": mid}).Decode(&material)
	if err != nil {
		return utils.SendErrorResponse(c, fiber.StatusNotFound, "Material not found", err.Error())
	}

	return utils.SendSuccessResponse(c, "Material fetched successfully", material)
}
