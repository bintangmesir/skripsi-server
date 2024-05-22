package routes

import (
	"server/controllers/santunan"
	"server/middlewares"

	"github.com/gofiber/fiber/v2"
)

func SantunanRoute(app *fiber.App) {
	app.Patch("/api/santunan/non-aktif/:id", middlewares.AuthMiddleware([]string{"ORANG_TUA_ASUH", "DONATUR"}), santunan.NonAktif)
	app.Patch("/api/santunan/berhenti/:id", middlewares.AuthMiddleware([]string{"ORANG_TUA_ASUH", "DONATUR"}), santunan.Berhenti)
}
