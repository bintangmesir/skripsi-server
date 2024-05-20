package routes

import (
	danasantunan "server/controllers/dana-santunan"
	"server/middlewares"

	"github.com/gofiber/fiber/v2"
)

func DanaSantunanRoute(app *fiber.App) {
	app.Get("/api/dana-santunan", middlewares.AuthMiddleware([]string{"ADMIN", "KETUA_DKM", "BENDAHARA"}), danasantunan.DanaSantunanGet)
	app.Get("/api/dana-santunan/:id", middlewares.AuthMiddleware([]string{"ADMIN", "KETUA_DKM", "BENDAHARA"}), danasantunan.DanaSantunanGetById)
	app.Post("/api/dana-santunan", middlewares.AuthMiddleware([]string{"ADMIN", "BENDAHARA"}), danasantunan.DanaSantunanCreate)
	app.Post("/api/dana-santunan/transfer", danasantunan.DanaSantunanCreateTransfer)
	app.Patch("/api/dana-santunan/:id", middlewares.AuthMiddleware([]string{"ADMIN", "BENDAHARA"}), danasantunan.DanaSantunanUpdate)
	app.Delete("/api/dana-santunan/:id", middlewares.AuthMiddleware([]string{"ADMIN", "BENDAHARA"}), danasantunan.DanaSantunanDelete)
}
