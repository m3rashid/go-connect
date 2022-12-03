package controllers

import (
	"context"
	"encoding/json"
	"go-server/models"
	"go-server/utils"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ChatsController struct {
	db  *mongo.Client
	ctx context.Context
}

func NewChatsController(db *mongo.Client, ctx context.Context) *ChatsController {
	return &ChatsController{db, ctx}
}

func (uc ChatsController) GetAllChats(c *fiber.Ctx) error {
	chats := []models.Chats{}
	collection := models.GetCollection(uc.db, models.ChatsCollectionName)

	curr, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		return utils.HandleControllerError(c, err)
	}

	defer curr.Close(context.Background())
	for curr.Next(context.Background()) {
		err := curr.Decode(&chats)
		if err != nil {
			return utils.HandleControllerError(c, err)
		}
	}

	back, err := json.Marshal(chats)
	if err != nil {
		return utils.HandleControllerError(c, err)
	}

	return c.JSON(back)
}
