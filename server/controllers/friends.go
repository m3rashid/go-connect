package controllers

import (
	"go-server/models"

	"github.com/gofiber/fiber/v2"
	"gopkg.in/mgo.v2"
)

type FriendsController struct {
	session *mgo.Session
}

func NewFriendsController(s *mgo.Session) *FriendsController {
	return &FriendsController{s}
}

func (uc FriendsController) SendRequest(c *fiber.Ctx) error {
	request := models.Friends{}
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	if err := uc.session.DB(models.DatabaseName).C(models.FriendsCollectionName).Insert(request); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.SendStatus(fiber.StatusOK)
}

func (uc FriendsController) UnsendRequest(c *fiber.Ctx) error {
	request := models.Friends{}
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	if err := uc.session.DB(models.DatabaseName).C(models.FriendsCollectionName).Remove(request); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.SendStatus(fiber.StatusOK)
}

func (uc FriendsController) AcceptRequest(c *fiber.Ctx) error {
	request := models.Friends{}
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	if err := uc.session.DB(models.DatabaseName).C(models.FriendsCollectionName).Update(request.FriendshipID, request); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.SendStatus(fiber.StatusOK)
}

func (uc FriendsController) DenyRequest(c *fiber.Ctx) error {
	request := models.Friends{}
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	if err := uc.session.DB(models.DatabaseName).C(models.FriendsCollectionName).Remove(request); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.SendStatus(fiber.StatusOK)
}

func (uc FriendsController) Unfriend(c *fiber.Ctx) error {
	request := models.Friends{}
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	if err := uc.session.DB(models.DatabaseName).C(models.FriendsCollectionName).Remove(request); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.SendStatus(fiber.StatusOK)
}

func (uc FriendsController) BlockUser(c *fiber.Ctx) error {
	request := models.Friends{}
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	if err := uc.session.DB(models.DatabaseName).C(models.FriendsCollectionName).Remove(request); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.SendStatus(fiber.StatusOK)
}

func (uc FriendsController) UnblockUser(c *fiber.Ctx) error {
	request := models.Friends{}
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	if err := uc.session.DB(models.DatabaseName).C(models.FriendsCollectionName).Remove(request); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.SendStatus(fiber.StatusOK)
}
