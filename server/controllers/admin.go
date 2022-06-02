package controllers

import (
	"encoding/json"
	"go-server/models"

	"github.com/gofiber/fiber/v2"
	"gopkg.in/mgo.v2"
)

type AdminsController struct {
	session *mgo.Session
}

func NewAdminsController(s *mgo.Session) *AdminsController {
	return &AdminsController{s}
}

func (uc AdminsController) GetUsers(c *fiber.Ctx) error {
	users := []models.Auths{}
	if err := uc.session.DB(models.DatabaseName).C(models.AuthCollectionName).Find(nil).All(&users); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	back, err := json.Marshal(users)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.JSON(back)
}

func (uc AdminsController) GetOneUser(c *fiber.Ctx) error {
	user := models.Auths{}
	if err := uc.session.DB(models.DatabaseName).C(models.AuthCollectionName).Find(nil).One(&user); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	back, err := json.Marshal(user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.JSON(back)
}

func (uc AdminsController) DeleteUser(c *fiber.Ctx) error {
	user := models.Auths{}
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	if err := uc.session.DB(models.DatabaseName).C(models.AuthCollectionName).Remove(user); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.SendStatus(fiber.StatusOK)
}

func (uc AdminsController) GetTopics(c *fiber.Ctx) error {
	topics := []models.Topics{}
	if err := uc.session.DB(models.DatabaseName).C(models.TopicsCollectionName).Find(nil).All(&topics); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
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
	if err := uc.session.DB(models.DatabaseName).C(models.TopicsCollectionName).Insert(topic); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.SendStatus(fiber.StatusOK)
}

func (uc AdminsController) UpdateTopic(c *fiber.Ctx) error {
	topic := models.Topics{}
	if err := c.BodyParser(&topic); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	if err := uc.session.DB(models.DatabaseName).C(models.TopicsCollectionName).Update(topic.TopicID, topic); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.SendStatus(fiber.StatusOK)
}

func (uc AdminsController) DeleteTopic(c *fiber.Ctx) error {
	topic := models.Topics{}
	if err := c.BodyParser(&topic); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	if err := uc.session.DB(models.DatabaseName).C(models.TopicsCollectionName).Remove(topic); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.SendStatus(fiber.StatusOK)
}

func (uc AdminsController) DeletePost(c *fiber.Ctx) error {
	post := models.Posts{}
	if err := c.BodyParser(&post); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	if err := uc.session.DB(models.DatabaseName).C(models.PostsCollectionName).Remove(post); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.SendStatus(fiber.StatusOK)
}
