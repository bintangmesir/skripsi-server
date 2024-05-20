package routes

import (
	"server/controllers/profile"
	"server/middlewares"

	"github.com/gofiber/fiber/v2"
)

func ProfileRoute(app *fiber.App) {
	app.Patch("/api/profile", middlewares.AuthMiddleware([]string{}), profile.UpdateProfile)
	app.Patch("/api/password", middlewares.AuthMiddleware([]string{}), profile.UpdatePassword)
	app.Patch("/api/password/:id", middlewares.AuthMiddleware([]string{}), profile.ResetPassword)
}
