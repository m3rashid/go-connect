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

type PostsController struct {
	db  *mongo.Client
	ctx context.Context
}

func NewPostsController(db *mongo.Client, ctx context.Context) *PostsController {
	return &PostsController{db, ctx}
}

func (uc PostsController) GetAllPosts(c *fiber.Ctx) error {
	posts := []models.Posts{}
	collection := models.GetCollection(uc.db, models.PostsCollectionName)
	curr, err := collection.Find(context.Background(), bson.M{})

	if err != nil {
		return utils.HandleControllerError(c, err)
	}

	defer curr.Close(context.Background())
	for curr.Next(context.Background()) {
		err := curr.Decode(&posts)
		if err != nil {
			return utils.HandleControllerError(c, err)
		}
	}

	back, err := json.Marshal(posts)
	if err != nil {
		return utils.HandleControllerError(c, err)
	}

	return c.JSON(back)
}

func (uc PostsController) GetOnePost(c *fiber.Ctx) error {
	post := models.Posts{}
	collection := models.GetCollection(uc.db, models.PostsCollectionName)

	err := collection.FindOne(context.Background(), bson.M{
		// criteria
	}).Decode(&post)

	if err != nil {
		return utils.HandleControllerError(c, err)
	}
	back, err := json.Marshal(post)

	if err != nil {
		return utils.HandleControllerError(c, err)
	}

	return c.JSON(back)
}

func (uc PostsController) GetPostsByTopic(c *fiber.Ctx) error {
	posts := []models.Posts{}
	collection := models.GetCollection(uc.db, models.PostsCollectionName)
	curr, err := collection.Find(context.Background(), bson.M{})

	if err != nil {
		return utils.HandleControllerError(c, err)
	}

	defer curr.Close(context.Background())
	for curr.Next(context.Background()) {
		err := curr.Decode(&posts)
		if err != nil {
			return utils.HandleControllerError(c, err)
		}
	}

	back, err := json.Marshal(posts)
	if err != nil {
		return utils.HandleControllerError(c, err)
	}

	return c.JSON(back)
}

func (uc PostsController) AddPost(c *fiber.Ctx) error {
	post := models.Posts{}
	collection := models.GetCollection(uc.db, models.PostsCollectionName)
	_, err := collection.InsertOne(context.Background(), bson.M{})

	if err != nil {
		return utils.HandleControllerError(c, err)
	}
	back, err := json.Marshal(post)
	if err != nil {
		return utils.HandleControllerError(c, err)
	}

	return c.Send(back)
}
