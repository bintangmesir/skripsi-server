package routes

import (
	danasantunananakasuh "server/controllers/dana-santunan-anak-asuh"
	"server/middlewares"

	"github.com/gofiber/fiber/v2"
)

func DanaSantunanAnakAsuhRoute(app *fiber.App) {
	app.Get("/api/dana-santunan-anak-asuh", middlewares.AuthMiddleware([]string{}), danasantunananakasuh.DanaSantunaAnakAsuhnGet)
	app.Get("/api/dana-santunan-anak-asuh/:id", middlewares.AuthMiddleware([]string{}), danasantunananakasuh.DanaSantunanAnakAsuhGetById)
	app.Post("/api/dana-santunan-anak-asuh", middlewares.AuthMiddleware([]string{"ADMIN", "BENDAHARA"}), danasantunananakasuh.DanaSantunanAnakAsuhCreate)
	app.Post("/api/dana-santunan-anak-asuh/transfer", middlewares.AuthMiddleware([]string{"BENDAHARA", "ADMIN", "ORANG_TUA_ASUH", "DONATUR"}), danasantunananakasuh.DanaSantunanAnakAsuhCreateTransfer)
	app.Patch("/api/dana-santunan-anak-asuh/:id", middlewares.AuthMiddleware([]string{"BENDAHARA", "ADMIN"}), danasantunananakasuh.DanaSantunanAnakAsuhUpdate)
	app.Delete("/api/dana-santunan-anak-asuh/:id", middlewares.AuthMiddleware([]string{"BENDAHARA", "ADMIN"}), danasantunananakasuh.DanaSantunanAnakAsuhDelete)
}
