package controllers

import (
	"github.com/gofiber/fiber/v2"
	"gopkg.in/mgo.v2"
)

type AdminsController struct {
	session *mgo.Session
}

func NewAdminsController(s *mgo.Session) *AdminsController {
	return &AdminsController{s}
}

func (uc AdminsController) GetUsers(c *fiber.Ctx) error {
}

func (uc AdminsController) GetOneUser(c *fiber.Ctx) error {}

func (uc AdminsController) DeleteUser(c *fiber.Ctx) error {}

func (uc AdminsController) GetTopics(c *fiber.Ctx) error {}

func (uc AdminsController) CreateTopic(c *fiber.Ctx) error {}

func (uc AdminsController) UpdateTopic(c *fiber.Ctx) error {}

func (uc AdminsController) DeleteTopic(c *fiber.Ctx) error {}

func (uc AdminsController) DeletePost(c *fiber.Ctx) error {}
