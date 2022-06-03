package main

import (
	"go-server/controllers"
	"log"

	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"gopkg.in/mgo.v2"

	"github.com/gofiber/fiber/v2"
)

const mongodbUri = "mongodb://localhost:27017"

func main() {
	app := fiber.New()
	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000",
	}))

	s, err := mgo.Dial(mongodbUri)
	if err != nil {
		panic(err)
	}

	admin := controllers.NewAdminsController(s)
	auth := controllers.NewAuthsController(s)
	bookmark := controllers.NewBookmarksController(s)
	chats := controllers.NewChatsController(s)
	comments := controllers.NewCommentsController(s)
	friends := controllers.NewFriendsController(s)
	likes := controllers.NewLikesController(s)
	notifs := controllers.NewNotifsController(s)
	posts := controllers.NewPostsController(s)
	search := controllers.NewSearchController(s)
	updateProf := controllers.NewUpdateUserController(s)

	// admin
	app.Post("/admin/createTopic", admin.CreateTopic)
	app.Post("/admin/updateTopic", admin.UpdateTopic)
	app.Post("/admin/deleteTopic", admin.DeleteTopic)
	app.Post("/admin/deletePost", admin.DeletePost)
	app.Post("/admin/deleteUser", admin.DeleteUser)
	app.Post("/admin/getUser", admin.GetOneUser)
	app.Post("/admin/getUsers", admin.GetUsers)
	app.Get("/admin/topics", admin.GetTopics)

	// auth
	app.Get("/auth", auth.GetUser)
	app.Get("/auth/admin", auth.GetAdmin)
	app.Post("/auth/adminLogin", auth.AdminLogin)
	app.Post("/auth/login", auth.Login)
	app.Post("/auth/signup", auth.Signup)
	app.Post("/auth/other-user", auth.GetOneOtherUser)

	// user actions
	app.Post("/user/update-avatar", updateProf.UpdateAvatar)
	app.Post("/user/update-profile", updateProf.UpdateProfile)
	app.Post("/user/update-password", updateProf.UpdatePassword)

	// post actions
	app.Post("/post/fromTopic", posts.GetPostsByTopic)
	app.Post("/post/all", posts.GetAllPosts)
	app.Post("/post/one", posts.GetOnePost)
	app.Post("/post/add", posts.AddPost)

	// like actions
	app.Post("/like/add", likes.AddLike)
	app.Post("/like/remove", likes.RemoveLike)

	// comment actions
	app.Post("/comment/add", comments.AddComment)
	app.Post("/comment/edit", comments.EditComment)
	app.Post("/comment/delete", comments.DeleteComment)

	// friendship actions
	app.Post("/friendship/send", friends.SendRequest)
	app.Post("/friendship/unsend", friends.UnsendRequest)
	app.Post("/friendship/accept", friends.AcceptRequest)
	app.Post("/friendship/deny", friends.DenyRequest)
	app.Post("/friendship/unfriend", friends.Unfriend)
	app.Post("/friendship/block", friends.BlockUser)
	app.Post("/friendship/unblock", friends.UnblockUser)

	// bookmark actions
	app.Post("/bookmark/remove", bookmark.RemoveBookmark)
	app.Post("/bookmark/add", bookmark.AddBookmark)
	app.Post("/bookmark/all", bookmark.GetAllBookmarks)

	// chat actions
	app.Post("/chats/all", chats.GetAllChats)

	// search actions
	app.Post("/search", search.SearchUserAndTopic)

	// notifications actions
	app.Post("/notifs", notifs.GetNotifications)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	log.Fatal(app.Listen(":5000"))
}
