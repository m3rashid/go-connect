package controllers

import (
	"encoding/json"
	"go-server/models"

	"github.com/gofiber/fiber/v2"
	"gopkg.in/mgo.v2"
)

type PostsController struct {
	session *mgo.Session
}

func NewPostsController(s *mgo.Session) *PostsController {
	return &PostsController{s}
}

func (uc PostsController) GetAllPosts(c *fiber.Ctx) error {
	posts := []models.Posts{}
	if err := uc.session.DB(models.DatabaseName).C(models.PostsCollectionName).Find(nil).All(&posts); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	back, err := json.Marshal(posts)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.JSON(back)
}

func (uc PostsController) GetOnePost(c *fiber.Ctx) error {
	posts := []models.Posts{}
	if err := uc.session.DB(models.DatabaseName).C(models.PostsCollectionName).Find(nil).All(&posts); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	back, err := json.Marshal(posts)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.JSON(back)
}

func (uc PostsController) GetPostsByTopic(c *fiber.Ctx) error {
	posts := []models.Posts{}
	if err := uc.session.DB(models.DatabaseName).C(models.PostsCollectionName).Find(nil).All(&posts); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	back, err := json.Marshal(posts)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.JSON(back)
}

func (uc PostsController) AddPost(c *fiber.Ctx) error {
	posts := []models.Posts{}
	if err := uc.session.DB(models.DatabaseName).C(models.PostsCollectionName).Find(nil).All(&posts); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	back, err := json.Marshal(posts)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.JSON(back)
}
