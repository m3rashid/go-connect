package controllers

import (
	"github.com/gofiber/fiber/v2"
	"gopkg.in/mgo.v2"
)

type SearchController struct {
	session *mgo.Session
}

func NewSearchController(s *mgo.Session) *SearchController {
	return &SearchController{s}
}

func (uc SearchController) SearchUserAndTopic(c *fiber.Ctx) error {
	return c.JSON("Hello World")
}
