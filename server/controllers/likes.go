package controllers

import (
	"context"
	"go-server/models"
	"go-server/utils"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type LikesController struct {
	db  *mongo.Client
	ctx context.Context
}

func NewLikesController(db *mongo.Client, ctx context.Context) *LikesController {
	return &LikesController{db, ctx}
}

func (uc LikesController) AddLike(c *fiber.Ctx) error {
	request := models.Likes{}
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	request.LikeID = primitive.NewObjectID()

	collection := models.GetCollection(uc.db, models.LikesCollectionName)
	_, err := collection.InsertOne(context.Background(), bson.M{
		"UserID": request.UserID,
		"PostID": request.PostID,
	})

	if err != nil {
		return utils.HandleControllerError(c, err)
	}

	return c.SendStatus(fiber.StatusOK)
}

func (uc LikesController) RemoveLike(c *fiber.Ctx) error {
	request := models.Likes{}
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	collection := models.GetCollection(uc.db, models.LikesCollectionName)
	_, err := collection.DeleteOne(context.Background(), bson.M{
		"UserID": request.UserID,
		"PostID": request.PostID,
	})

	if err != nil {
		return utils.HandleControllerError(c, err)
	}
	return c.SendStatus(fiber.StatusOK)
}
