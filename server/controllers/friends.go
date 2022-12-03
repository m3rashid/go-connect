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

type FriendsController struct {
	db  *mongo.Client
	ctx context.Context
}

func NewFriendsController(db *mongo.Client, ctx context.Context) *FriendsController {
	return &FriendsController{db, ctx}
}

func (uc FriendsController) SendRequest(c *fiber.Ctx) error {
	request := models.Friends{}
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	request.FriendshipID = primitive.NewObjectID()

	collection := models.GetCollection(uc.db, models.FriendsCollectionName)
	_, err := collection.InsertOne(context.Background(), bson.M{
		"FromID": request.FromID,
		"ToID":   request.ToID,
		"Status": request.Status,
	})
	if err != nil {
		return utils.HandleControllerError(c, err)
	}

	return c.SendStatus(fiber.StatusOK)
}

func (uc FriendsController) UnsendRequest(c *fiber.Ctx) error {
	request := models.Friends{}
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	collection := models.GetCollection(uc.db, models.FriendsCollectionName)
	_, err := collection.DeleteOne(context.Background(), bson.M{
		"FromID": request.FromID,
		"ToID":   request.ToID,
	})

	if err != nil {
		return utils.HandleControllerError(c, err)
	}

	return c.SendStatus(fiber.StatusOK)
}

func (uc FriendsController) AcceptRequest(c *fiber.Ctx) error {
	request := models.Friends{}
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	collection := models.GetCollection(uc.db, models.FriendsCollectionName)
	_, err := collection.UpdateOne(context.Background(), bson.M{
		"FromID": request.FromID,
		"ToID":   request.ToID,
	}, bson.M{
		"$set": bson.M{
			"Status": "friends",
		},
	})

	if err != nil {
		return utils.HandleControllerError(c, err)
	}

	return c.SendStatus(fiber.StatusOK)
}

func (uc FriendsController) DenyRequest(c *fiber.Ctx) error {
	request := models.Friends{}
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	collection := models.GetCollection(uc.db, models.FriendsCollectionName)
	_, err := collection.DeleteOne(context.Background(), bson.M{
		"FromID": request.FromID,
		"ToID":   request.ToID,
	})

	if err != nil {
		return utils.HandleControllerError(c, err)
	}

	return c.SendStatus(fiber.StatusOK)
}

func (uc FriendsController) Unfriend(c *fiber.Ctx) error {
	request := models.Friends{}
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	collection := models.GetCollection(uc.db, models.FriendsCollectionName)
	_, err := collection.DeleteOne(context.Background(), bson.M{
		"FromID": request.FromID,
		"ToID":   request.ToID,
	})

	if err != nil {
		return utils.HandleControllerError(c, err)
	}

	return c.SendStatus(fiber.StatusOK)
}

func (uc FriendsController) BlockUser(c *fiber.Ctx) error {
	request := models.Friends{}
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	collection := models.GetCollection(uc.db, models.FriendsCollectionName)
	_, err := collection.DeleteOne(context.Background(), bson.M{
		"FromID": request.FromID,
		"ToID":   request.ToID,
	})

	if err != nil {
		return utils.HandleControllerError(c, err)
	}

	return c.SendStatus(fiber.StatusOK)
}

func (uc FriendsController) UnblockUser(c *fiber.Ctx) error {
	request := models.Friends{}
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	collection := models.GetCollection(uc.db, models.FriendsCollectionName)
	_, err := collection.DeleteOne(context.Background(), bson.M{
		"FromID": request.FromID,
		"ToID":   request.ToID,
	})

	if err != nil {
		return utils.HandleControllerError(c, err)
	}

	return c.SendStatus(fiber.StatusOK)
}
