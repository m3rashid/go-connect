package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"go-server/models"
	"go-server/utils"
	"reflect"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type AdminsController struct {
	db  *mongo.Client
	ctx context.Context
}

func NewAdminsController(db *mongo.Client, ctx context.Context) *AdminsController {
	return &AdminsController{db, ctx}
}

func (uc AdminsController) GetUsers(c *fiber.Ctx) error {
	users := []models.Auths{}
	collection := models.GetCollection(uc.db, models.AuthCollectionName)

	curr, err := collection.Find(context.Background(), bson.D{})
	if err != nil {
		return utils.HandleControllerError(c, err)
	}

	defer curr.Close(context.Background())
	for curr.Next(context.Background()) {
		err := curr.Decode(&users)
		if err != nil {
			return utils.HandleControllerError(c, err)
		}
	}

	back, err := json.Marshal(users)
	if err != nil {
		return utils.HandleControllerError(c, err)
	}

	return c.JSON(back)
}

func (uc AdminsController) GetOneUser(c *fiber.Ctx) error {
	user := models.Auths{}
	collection := models.GetCollection(uc.db, models.AuthCollectionName)

	curr, err := collection.Find(context.Background(), bson.D{})

	if err != nil {
		return utils.HandleControllerError(c, err)
	}

	defer curr.Close(context.Background())
	for curr.Next(context.Background()) {
		err := curr.Decode(&user)
		if err != nil {
			return utils.HandleControllerError(c, err)
		}
	}

	back, err := json.Marshal(user)
	if err != nil {
		return utils.HandleControllerError(c, err)
	}

	return c.JSON(back)
}

func (uc AdminsController) DeleteUser(c *fiber.Ctx) error {
	var userId string

	if err := c.BodyParser(&userId); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	_id, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return utils.HandleControllerError(c, err)
	}

	collection := models.GetCollection(uc.db, models.AuthCollectionName)

	res, err := collection.DeleteOne(uc.ctx, bson.M{"_id": _id})

	if err != nil {
		return utils.HandleControllerError(c, err)
	}

	fmt.Println("DeleteOne Result TYPE:", reflect.TypeOf(res))

	return c.SendStatus(fiber.StatusOK)
}

func (uc AdminsController) GetTopics(c *fiber.Ctx) error {
	topics := []models.Topics{}
	collection := models.GetCollection(uc.db, models.TopicsCollectionName)

	curr, err := collection.Find(context.Background(), bson.D{})
	if err != nil {
		return utils.HandleControllerError(c, err)
	}

	defer curr.Close(context.Background())
	for curr.Next(context.Background()) {
		err := curr.Decode(&topics)
		if err != nil {
			return utils.HandleControllerError(c, err)
		}
	}

	back, err := json.Marshal(topics)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.JSON(back)
}

func (uc AdminsController) CreateTopic(c *fiber.Ctx) error {
	topic := models.Topics{}
	if err := c.BodyParser(&topic); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	topic.TopicID = primitive.NewObjectID()
	collection := models.GetCollection(uc.db, models.TopicsCollectionName)

	res, err := collection.InsertOne(context.Background(), topic)

	if err != nil {
		return utils.HandleControllerError(c, err)
	}

	back, err := json.Marshal(res)
	if err != nil {
		return utils.HandleControllerError(c, err)
	}

	return c.JSON(back)
}

func (uc AdminsController) UpdateTopic(c *fiber.Ctx) error {
	topic := models.Topics{}
	if err := c.BodyParser(&topic); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	collection := models.GetCollection(uc.db, models.TopicsCollectionName)
	res, err := collection.UpdateOne(context.Background(), bson.M{"TopicID": topic.TopicID}, bson.D{
		{"set", bson.M{
			"Name": topic.Name,
		}},
	})

	if err != nil {
		return utils.HandleControllerError(c, err)
	}

	back, err := json.Marshal(res)
	if err != nil {
		return utils.HandleControllerError(c, err)
	}

	return c.JSON(back)
}

func (uc AdminsController) DeleteTopic(c *fiber.Ctx) error {
	topic := models.Topics{}

	if err := c.BodyParser(&topic); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	collection := models.GetCollection(uc.db, models.TopicsCollectionName)
	res, err := collection.DeleteOne(context.Background(), bson.M{"TopicID": topic.TopicID})

	if err != nil {
		return utils.HandleControllerError(c, err)
	}

	back, err := json.Marshal(res)
	if err != nil {
		return utils.HandleControllerError(c, err)
	}

	return c.JSON(back)
}

func (uc AdminsController) DeletePost(c *fiber.Ctx) error {
	post := models.Posts{}

	if err := c.BodyParser(&post); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	collection := models.GetCollection(uc.db, models.PostsCollectionName)
	res, err := collection.DeleteOne(context.Background(), bson.M{"PostID": post.PostID})

	if err != nil {
		return utils.HandleControllerError(c, err)
	}

	back, err := json.Marshal(res)
	if err != nil {
		return utils.HandleControllerError(c, err)
	}

	return c.JSON(back)
}
