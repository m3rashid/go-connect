package controllers

import (
	"go-server/models"

	"github.com/gofiber/fiber/v2"
	"gopkg.in/mgo.v2"
)

type LikesController struct {
	session *mgo.Session
}

func NewLikesController(s *mgo.Session) *LikesController {
	return &LikesController{s}
}

func (uc LikesController) AddLike(c *fiber.Ctx) error {
	request := models.Likes{}
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	if err := uc.session.DB(models.DatabaseName).C(models.LikesCollectionName).Insert(request); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.SendStatus(fiber.StatusOK)
}

func (uc LikesController) RemoveLike(c *fiber.Ctx) error {
	request := models.Likes{}
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	if err := uc.session.DB(models.DatabaseName).C(models.LikesCollectionName).Remove(request); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.SendStatus(fiber.StatusOK)
}
