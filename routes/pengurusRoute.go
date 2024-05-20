package routes

import (
	"server/controllers/pengurus"
	"server/middlewares"

	"github.com/gofiber/fiber/v2"
)

func PengurusRoute(app *fiber.App) {
	app.Get("/api/pengurus", middlewares.AuthMiddleware([]string{"ADMIN", "KETUA_DKM", "BENDAHARA", "HUMAS"}), pengurus.PengurusGet)
	app.Get("/api/pengurus/:id", middlewares.AuthMiddleware([]string{"ADMIN"}), pengurus.PengurusGetById)
	app.Post("/api/pengurus", middlewares.AuthMiddleware([]string{"ADMIN"}), pengurus.PengurusPost)
	app.Patch("/api/pengurus/:id", middlewares.AuthMiddleware([]string{"ADMIN"}), pengurus.PengurusPatch)
	app.Delete("/api/pengurus/:id", middlewares.AuthMiddleware([]string{"ADMIN"}), pengurus.PengurusDelete)
}
