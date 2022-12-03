package controllers

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

type UpdateUserController struct {
	db  *mongo.Client
	ctx context.Context
}

func NewUpdateUserController(db *mongo.Client, ctx context.Context) *UpdateUserController {
	return &UpdateUserController{db, ctx}
}

func (uc UpdateUserController) UpdateAvatar(c *fiber.Ctx) error {
	return c.JSON("Hello World")
}

func (uc UpdateUserController) UpdatePassword(c *fiber.Ctx) error {
	return c.JSON("Hello World")
}

func (uc UpdateUserController) UpdateProfile(c *fiber.Ctx) error {
	return c.JSON("Hello World")
}
