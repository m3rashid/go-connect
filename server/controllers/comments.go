package controllers

import (
	"github.com/gofiber/fiber/v2"
	"gopkg.in/mgo.v2"
)

type CommentsController struct {
	session *mgo.Session
}

func NewCommentsController(s *mgo.Session) *CommentsController {
	return &CommentsController{s}
}

func (uc CommentsController) AddComment(c *fiber.Ctx) error {}

func (uc CommentsController) DeleteComment(c *fiber.Ctx) error {}

func (uc CommentsController) EditComment(c *fiber.Ctx) error {}
