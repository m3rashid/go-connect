package controllers

import (
	"github.com/gofiber/fiber/v2"
	"gopkg.in/mgo.v2"
)

type PostsController struct {
	session *mgo.Session
}

func NewPostsController(s *mgo.Session) *PostsController {
	return &PostsController{s}
}

func (uc PostsController) GetAllPosts(c *fiber.Ctx) error {}

func (uc PostsController) GetOnePost(c *fiber.Ctx) error {}

func (uc PostsController) GetPostsByTopic(c *fiber.Ctx) error {}

func (uc PostsController) AddPost(c *fiber.Ctx) error {}
