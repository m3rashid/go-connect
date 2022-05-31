package controllers

import (
	"github.com/gofiber/fiber/v2"
	"gopkg.in/mgo.v2"
)

type UpdateUserController struct {
	session *mgo.Session
}

func NewUpdateUserController(s *mgo.Session) *UpdateUserController {
	return &UpdateUserController{s}
}

func (uc UpdateUserController) UpdateAvatar(c *fiber.Ctx) error {}

func (uc UpdateUserController) UpdatePassword(c *fiber.Ctx) error {}

func (uc UpdateUserController) UpdateProfile(c *fiber.Ctx) error {}
