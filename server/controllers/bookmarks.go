package controllers

import (
	"encoding/json"
	"go-server/models"

	"github.com/gofiber/fiber/v2"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type BookmarksController struct {
	session *mgo.Session
}

func NewBookmarksController(s *mgo.Session) *BookmarksController {
	return &BookmarksController{s}
}

func (uc BookmarksController) RemoveBookmark(c *fiber.Ctx) error {
	bookmark := models.Bookmarks{}
	if err := c.BodyParser(&bookmark); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	if err := uc.session.DB(models.DatabaseName).C(models.BookmarksCollectionName).Remove(bookmark); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.SendStatus(fiber.StatusOK)
}

func (uc BookmarksController) AddBookmark(c *fiber.Ctx) error {
	bookmark := models.Bookmarks{}
	if err := c.BodyParser(&bookmark); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	// add some validator functions
	// if err := bookmark.Validate(); err != nil {}

	bookmark.BookmarkID = bson.NewObjectId()
	uc.session.DB(models.DatabaseName).C(models.BookmarksCollectionName).Insert(bookmark)
	back, err := json.Marshal(bookmark)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.JSON(back)
}

func (uc BookmarksController) GetAllBookmarks(c *fiber.Ctx) error {
	bookmarks := []models.Bookmarks{}
	if err := uc.session.DB(models.DatabaseName).C(models.BookmarksCollectionName).Find(nil).All(&bookmarks); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	back, err := json.Marshal(bookmarks)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.JSON(back)
}
