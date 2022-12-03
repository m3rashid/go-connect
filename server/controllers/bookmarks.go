package controllers

import (
	"context"
	"encoding/json"
	"go-server/models"
	"go-server/utils"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type BookmarksController struct {
	db  *mongo.Client
	ctx context.Context
}

func NewBookmarksController(db *mongo.Client, ctx context.Context) *BookmarksController {
	return &BookmarksController{db, ctx}
}

func (uc BookmarksController) RemoveBookmark(c *fiber.Ctx) error {
	bookmark := models.Bookmarks{}
	if err := c.BodyParser(&bookmark); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	collection := models.GetCollection(uc.db, models.BookmarksCollectionName)
	err := collection.FindOneAndDelete(uc.ctx, bson.M{"_id": bookmark.BookmarkID}).Decode(&bookmark)
	if err != nil {
		return utils.HandleControllerError(c, err)
	}

	return c.SendStatus(fiber.StatusOK)
}

func (uc BookmarksController) AddBookmark(c *fiber.Ctx) error {
	bookmark := models.Bookmarks{}
	if err := c.BodyParser(&bookmark); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	bookmark.BookmarkID = primitive.NewObjectID()

	collection := models.GetCollection(uc.db, models.BookmarksCollectionName)
	_, err := collection.InsertOne(context.Background(), bson.M{
		"BookmarkID": bookmark.BookmarkID,
		"User":       bookmark.User,
		"Post":       bookmark.Post,
	})

	if err != nil {
		return utils.HandleControllerError(c, err)
	}

	back, err := json.Marshal(bookmark)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.JSON(back)
}

func (uc BookmarksController) GetAllBookmarks(c *fiber.Ctx) error {
	bookmarks := []models.Bookmarks{}
	collection := models.GetCollection(uc.db, models.BookmarksCollectionName)

	err := collection.FindOne(context.Background(), bson.M{}).Decode(&bookmarks)

	if err != nil {
		return utils.HandleControllerError(c, err)
	}

	back, err := json.Marshal(bookmarks)
	if err != nil {
		return utils.HandleControllerError(c, err)
	}

	return c.JSON(back)
}
