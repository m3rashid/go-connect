package controllers

import (
	"go-server/models"

	"github.com/gofiber/fiber/v2"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type CommentsController struct {
	session *mgo.Session
}

func NewCommentsController(s *mgo.Session) *CommentsController {
	return &CommentsController{s}
}

func (uc CommentsController) AddComment(c *fiber.Ctx) error {
	comment := models.Comments{}
	if err := c.BodyParser(&comment); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	comment.CommentID = bson.NewObjectId()
	if err := uc.session.DB(models.DatabaseName).C(models.CommentsCollectionName).Insert(comment); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.SendStatus(fiber.StatusOK)
}

func (uc CommentsController) DeleteComment(c *fiber.Ctx) error {
	comment := models.Comments{}
	if err := c.BodyParser(&comment); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	if err := uc.session.DB(models.DatabaseName).C(models.CommentsCollectionName).Remove(comment); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.SendStatus(fiber.StatusOK)
}

func (uc CommentsController) EditComment(c *fiber.Ctx) error {
	comment := models.Comments{}
	if err := c.BodyParser(&comment); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	if err := uc.session.DB(models.DatabaseName).C(models.CommentsCollectionName).Update(comment.CommentID, comment); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.SendStatus(fiber.StatusOK)
}
