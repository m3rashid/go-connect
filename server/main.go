package main

import (
	"go-server/controllers"
	"log"

	"github.com/gofiber/fiber/v2/middleware/logger"
	"gopkg.in/mgo.v2"

	"github.com/gofiber/fiber/v2"
)

const mongodbUri = "mongodb://localhost:27017"

func main() {
	app := fiber.New()
	app.Use(logger.New())

	uc := controllers.NewBookmarksController(getSession())
	app.Post("/bookmarks/add", uc.AddBookmark)
	app.Delete("/bookmarks/remove", uc.RemoveBookmark)
	app.Get("/bookmarks/all", uc.GetAllBookmarks)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	log.Fatal(app.Listen(":5000"))
}

func getSession() *mgo.Session {
	s, err := mgo.Dial(mongodbUri)
	if err != nil {
		panic(err)
	}
	return s
}
