package controllers

import (
	"github.com/gofiber/fiber/v2"
	"gopkg.in/mgo.v2"
)

type ChatsController struct {
	session *mgo.Session
}

func NewChatsController(s *mgo.Session) *ChatsController {
	return &ChatsController{s}
}

func (uc ChatsController) GetAllChats(c *fiber.Ctx) error {}
