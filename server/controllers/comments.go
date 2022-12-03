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

type CommentsController struct {
	db  *mongo.Client
	ctx context.Context
}

func NewCommentsController(db *mongo.Client, ctx context.Context) *CommentsController {
	return &CommentsController{db, ctx}
}

func (uc CommentsController) AddComment(c *fiber.Ctx) error {
	comment := models.Comments{}
	if err := c.BodyParser(&comment); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	comment.CommentID = primitive.NewObjectID()
	collection := models.GetCollection(uc.db, models.CommentsCollectionName)
	_, err := collection.InsertOne(context.Background(), bson.M{
		"CommentID": comment.CommentID,
		"UserID":    comment.UserID,
		"PostID":    comment.PostID,
		"Text":      comment.Text,
	})

	if err != nil {
		return utils.HandleControllerError(c, err)
	}

	return c.SendStatus(fiber.StatusOK)
}

func (uc CommentsController) DeleteComment(c *fiber.Ctx) error {
	comment := models.Comments{}
	if err := c.BodyParser(&comment); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	collection := models.GetCollection(uc.db, models.CommentsCollectionName)
	_, err := collection.DeleteOne(context.Background(), bson.M{
		"CommentID": comment.CommentID,
	})

	if err != nil {
		return utils.HandleControllerError(c, err)
	}
	return c.SendStatus(fiber.StatusOK)
}

func (uc CommentsController) EditComment(c *fiber.Ctx) error {
	comment := models.Comments{}
	if err := c.BodyParser(&comment); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	collection := models.GetCollection(uc.db, models.CommentsCollectionName)
	_, err := collection.UpdateOne(context.Background(), bson.M{
		"CommentID": comment.CommentID,
	}, bson.M{
		"$set": bson.M{
			"Text": comment.Text,
		},
	})

	if err != nil {
		return utils.HandleControllerError(c, err)
	}

	return c.SendStatus(fiber.StatusOK)
}
