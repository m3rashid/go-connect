package controllers

import (
	"github.com/gofiber/fiber/v2"
	"gopkg.in/mgo.v2"
)

type LikesController struct {
	session *mgo.Session
}

func NewLikesController(s *mgo.Session) *LikesController {
	return &LikesController{s}
}

func (uc LikesController) AddLike(c *fiber.Ctx) error {}

func (uc LikesController) RemoveLike(c *fiber.Ctx) error {}
