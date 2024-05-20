package routes

import (
	midtranscontroller "server/controllers/midtrans-controller"

	"github.com/gofiber/fiber/v2"
)

func MidtransRoute(app *fiber.App) {
	app.Get("/api/midtrans-dana-santunan/:id", midtranscontroller.MidtransDanaSantunanGetById)
	app.Post("/api/midtrans-dana-santunan", midtranscontroller.MidtransDanaSantunanCreate)
	app.Get("/api/midtrans-dana-santunan-anak-asuh/:id", midtranscontroller.MidtransDanaSantunanAnakAsuhGetById)
	app.Post("/api/midtrans-dana-santunan-anak-asuh", midtranscontroller.MidtransDanaSantunanAnakAsuhCreate)
}
