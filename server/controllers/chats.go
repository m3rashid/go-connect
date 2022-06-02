package controllers

import (
	"encoding/json"
	"go-server/models"

	"github.com/gofiber/fiber/v2"
	"gopkg.in/mgo.v2"
)

type ChatsController struct {
	session *mgo.Session
}

func NewChatsController(s *mgo.Session) *ChatsController {
	return &ChatsController{s}
}

func (uc ChatsController) GetAllChats(c *fiber.Ctx) error {
	chats := []models.Chats{}
	if err := uc.session.DB(models.DatabaseName).C(models.ChatsCollectionName).Find(nil).All(&chats); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	back, err := json.Marshal(chats)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.JSON(back)
}
