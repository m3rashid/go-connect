package controllers

import (
	"github.com/gofiber/fiber/v2"
	"gopkg.in/mgo.v2"
)

type NotifsController struct {
	session *mgo.Session
}

func NewNotifsController(s *mgo.Session) *NotifsController {
	return &NotifsController{s}
}

func (uc NotifsController) GetNotifications(c *fiber.Ctx) error {}
