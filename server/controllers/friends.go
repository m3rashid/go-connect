package controllers

import (
	"github.com/gofiber/fiber/v2"
	"gopkg.in/mgo.v2"
)

type FriendsController struct {
	session *mgo.Session
}

func NewFriendsController(s *mgo.Session) *FriendsController {
	return &FriendsController{s}
}

func (uc FriendsController) SendRequest(c *fiber.Ctx) error {}

func (uc FriendsController) UnsendRequest(c *fiber.Ctx) error {}

func (uc FriendsController) AcceptRequest(c *fiber.Ctx) error {}

func (uc FriendsController) DenyRequest(c *fiber.Ctx) error {}

func (uc FriendsController) Unfriend(c *fiber.Ctx) error {}

func (uc FriendsController) BlockUser(c *fiber.Ctx) error {}

func (uc FriendsController) UnblockUser(c *fiber.Ctx) error {}
