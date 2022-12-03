package controllers

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

type SearchController struct {
	db  *mongo.Client
	ctx context.Context
}

func NewSearchController(db *mongo.Client, ctx context.Context) *SearchController {
	return &SearchController{db, ctx}
}

func (uc SearchController) SearchUserAndTopic(c *fiber.Ctx) error {
	return c.JSON("Hello World")
}
