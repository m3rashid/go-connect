package controllers

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

type NotifsController struct {
	db  *mongo.Client
	ctx context.Context
}

func NewNotifsController(db *mongo.Client, ctx context.Context) *NotifsController {
	return &NotifsController{db, ctx}
}

func (uc NotifsController) GetNotifications(c *fiber.Ctx) error {
	return c.JSON("Hello World")
}
