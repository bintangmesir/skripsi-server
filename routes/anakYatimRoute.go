package routes

import (
	anakyatim "server/controllers/anak-yatim"
	"server/middlewares"

	"github.com/gofiber/fiber/v2"
)

func AnakYatimRoute(app *fiber.App) {
	app.Get("/api/anak-yatim", anakyatim.AnakYatimGet)
	app.Get("/api/anak-yatim/:id", anakyatim.AnakYatimGetById)
	app.Post("/api/anak-yatim", middlewares.AuthMiddleware([]string{"HUMAS", "ADMIN"}), anakyatim.AnakYatimCreate)
	app.Patch("/api/anak-yatim/:id", middlewares.AuthMiddleware([]string{"HUMAS", "ADMIN"}), anakyatim.AnakYatimUpdate)
	app.Delete("/api/anak-yatim/:id", middlewares.AuthMiddleware([]string{"HUMAS", "ADMIN"}), anakyatim.AnakYatimDelete)
}
