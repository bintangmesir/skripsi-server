package routes

import (
	datacount "server/controllers/data-count"

	"github.com/gofiber/fiber/v2"
)

func DataCountRoute(app *fiber.App) {
	app.Get("/api/count-data-anak-yatim", datacount.CountDataAnakYatim)
	app.Get("/api/count-data-pengurus", datacount.CountDataPengurus)
	app.Get("/api/count-data-donatur", datacount.CountDataDonatur)
	app.Get("/api/count-dana-santunan", datacount.CountDanaSantunan)
	app.Get("/api/count-dana-santunan-anak-asuh", datacount.CountDanaSantunanAnakAsuh)
}
