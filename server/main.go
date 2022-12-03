package main

import (
	"context"
	"go-server/controllers"
	"go-server/models"
	"log"

	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000",
	}))

	client, ctx := models.InitDataLayer()
	client.Disconnect(context.Background())

	auth := controllers.NewAuthsController(client, ctx)
	chats := controllers.NewChatsController(client, ctx)
	likes := controllers.NewLikesController(client, ctx)
	posts := controllers.NewPostsController(client, ctx)
	admin := controllers.NewAdminsController(client, ctx)
	notifs := controllers.NewNotifsController(client, ctx)
	search := controllers.NewSearchController(client, ctx)
	friends := controllers.NewFriendsController(client, ctx)
	comments := controllers.NewCommentsController(client, ctx)
	bookmark := controllers.NewBookmarksController(client, ctx)
	updateProf := controllers.NewUpdateUserController(client, ctx)

	// admin
	app.Get("/admin/topics", admin.GetTopics)
	app.Post("/admin/getUsers", admin.GetUsers)
	app.Post("/admin/getUser", admin.GetOneUser)
	app.Post("/admin/deletePost", admin.DeletePost)
	app.Post("/admin/deleteUser", admin.DeleteUser)
	app.Post("/admin/deleteTopic", admin.DeleteTopic)
	app.Post("/admin/updateTopic", admin.UpdateTopic)
	app.Post("/admin/createTopic", admin.CreateTopic)

	// auth
	app.Get("/auth", auth.GetUser)
	app.Post("/auth/login", auth.Login)
	app.Get("/auth/admin", auth.GetAdmin)
	app.Post("/auth/signup", auth.Signup)
	app.Post("/auth/adminLogin", auth.AdminLogin)
	app.Post("/auth/other-user", auth.GetOneOtherUser)

	// user actions
	app.Post("/user/update-avatar", updateProf.UpdateAvatar)
	app.Post("/user/update-profile", updateProf.UpdateProfile)
	app.Post("/user/update-password", updateProf.UpdatePassword)

	// post actions
	app.Post("/post/add", posts.AddPost)
	app.Post("/post/one", posts.GetOnePost)
	app.Post("/post/all", posts.GetAllPosts)
	app.Post("/post/fromTopic", posts.GetPostsByTopic)

	// like actions
	app.Post("/like/add", likes.AddLike)
	app.Post("/like/remove", likes.RemoveLike)

	// comment actions
	app.Post("/comment/add", comments.AddComment)
	app.Post("/comment/edit", comments.EditComment)
	app.Post("/comment/delete", comments.DeleteComment)

	// friendship actions
	app.Post("/friendship/block", friends.BlockUser)
	app.Post("/friendship/send", friends.SendRequest)
	app.Post("/friendship/deny", friends.DenyRequest)
	app.Post("/friendship/unfriend", friends.Unfriend)
	app.Post("/friendship/unblock", friends.UnblockUser)
	app.Post("/friendship/unsend", friends.UnsendRequest)
	app.Post("/friendship/accept", friends.AcceptRequest)

	// bookmark actions
	app.Post("/bookmark/add", bookmark.AddBookmark)
	app.Post("/bookmark/all", bookmark.GetAllBookmarks)
	app.Post("/bookmark/remove", bookmark.RemoveBookmark)

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
