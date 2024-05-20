package routes

import (
	"server/controllers/donatur"
	"server/middlewares"

	"github.com/gofiber/fiber/v2"
)

func DonaturRoute(app *fiber.App) {
	app.Get("/api/donatur", middlewares.AuthMiddleware([]string{"ADMIN", "KETUA_DKM", "BENDAHARA", "HUMAS"}), donatur.DonaturGet)
	app.Get("/api/donatur/:id", middlewares.AuthMiddleware([]string{"ADMIN", "KETUA_DKM", "BENDAHARA", "HUMAS", "DONATUR", "ORANG_TUA_ASUH"}), donatur.DonaturGetById)
	app.Patch("/api/donatur/:id", middlewares.AuthMiddleware([]string{"ADMIN"}), donatur.DonaturUpdate)
	app.Delete("/api/donatur/:id", middlewares.AuthMiddleware([]string{"ADMIN"}), donatur.DonaturDelete)
}
