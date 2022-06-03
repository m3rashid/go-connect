package controllers

import (
	"encoding/json"
	"go-server/models"

	"github.com/gofiber/fiber/v2"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type AuthsController struct {
	session *mgo.Session
}

func NewAuthsController(s *mgo.Session) *AuthsController {
	return &AuthsController{s}
}

func (uc AuthsController) GetUser(c *fiber.Ctx) error {
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

func (uc AuthsController) GetAdmin(c *fiber.Ctx) error {
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

func (uc AuthsController) Signup(c *fiber.Ctx) error {
	user := models.Auths{}
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	user.UserID = bson.NewObjectId()
	user.AvatarID = bson.NewObjectId()
	// create a new avatar also with default config
	if err := uc.session.DB(models.DatabaseName).C(models.AuthCollectionName).Insert(user); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	back, err := json.Marshal(user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.JSON(back)
}

func (uc AuthsController) Login(c *fiber.Ctx) error {
	user := models.Auths{}
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	if err := uc.session.DB(models.DatabaseName).C(models.AuthCollectionName).Find(user).One(&user); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	back, err := json.Marshal(user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	// TODO send jwt also

	return c.JSON(back)
}

func (uc AuthsController) AdminLogin(c *fiber.Ctx) error {
	user := models.Auths{}
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	if err := uc.session.DB(models.DatabaseName).C(models.AuthCollectionName).Find(user).One(&user); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	back, err := json.Marshal(user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	// TODO send jwt also

	return c.JSON(back)
}

func (uc AuthsController) GetOneOtherUser(c *fiber.Ctx) error {
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
