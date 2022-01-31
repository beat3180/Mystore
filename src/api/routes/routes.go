package routes

import (
	"api/controllers"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Post("/api/register", controllers.Register)
	app.Post("/api/login", controllers.Login)
	app.Get("/api/user", controllers.User)
	app.Get("/api/logout", controllers.Logout)
	app.Post("/api/forgot", controllers.Forgot)
	app.Post("/api/reset", controllers.Reset)
	app.Post("/api/post", controllers.Post)
	app.Get("/api/post/get", controllers.GetPost)
	app.Get("/api/post/get/content/:id", controllers.GetPostID)
	app.Post("/api/post/get/mystore", controllers.GetMystoreID)
	app.Post("/api/post/mystore/public/update", controllers.PublicUpdate)
	app.Post("/api/post/mystore/hide/update", controllers.HideUpdate)
	app.Post("/api/post/mystore/delete", controllers.MyStoreDelet)
	//app.Get("/api/post/image/get", controllers.GetImage)
}
