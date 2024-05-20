package routes

import (
	"server/controllers/auth"
	"server/middlewares"

	"github.com/gofiber/fiber/v2"
)

func AuthRoute(app *fiber.App) {
	app.Get("/api/auth/user", middlewares.AuthMiddleware([]string{}), auth.User)
	app.Post("/api/auth/register", auth.Register)
	app.Post("/api/auth/login/donatur", auth.LoginDonatur)
	app.Post("/api/auth/login/pengurus", auth.LoginPengurus)
	app.Post("/api/auth/logout", middlewares.AuthMiddleware([]string{}), auth.Logout)
}
