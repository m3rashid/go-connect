package controllers

import (
	"github.com/gofiber/fiber/v2"
	"gopkg.in/mgo.v2"
)

type AuthsController struct {
	session *mgo.Session
}

func NewAuthsController(s *mgo.Session) *AuthsController {
	return &AuthsController{s}
}

func (uc AuthsController) GetUser(c *fiber.Ctx) error {}

func (uc AuthsController) GetAdmin(c *fiber.Ctx) error {}

func (uc AuthsController) Signup(c *fiber.Ctx) error {}

func (uc AuthsController) Login(c *fiber.Ctx) error {}

func (uc AuthsController) AdminLogin(c *fiber.Ctx) error {}

func (uc AuthsController) GetOneOtherUser(c *fiber.Ctx) error {}
