package routes

import (
	laporandanasantunananakasuh "server/controllers/laporan-dana-santunan-anak-asuh"
	"server/middlewares"

	"github.com/gofiber/fiber/v2"
)

func LaporanDanaSantunanAnakAsuhRoute(app *fiber.App) {
	app.Get("/api/laporan-dana-santunan-anak-asuh", middlewares.AuthMiddleware([]string{"BENDAHARA", "ADMIN", "HUMAS", "KETUA_DKM", "ORANG_TUA_ASUH", "DONATUR"}), laporandanasantunananakasuh.LaporanDanaSantunanAnakAsuhGet)
	app.Get("/api/laporan-dana-santunan-anak-asuh/:id", middlewares.AuthMiddleware([]string{"BENDAHARA", "ADMIN", "HUMAS", "KETUA_DKM", "ORANG_TUA_ASUH", "DONATUR"}), laporandanasantunananakasuh.LaporanDanaSantunanAnakAsuhGetById)
	app.Post("/api/laporan-dana-santunan-anak-asuh", middlewares.AuthMiddleware([]string{"BENDAHARA", "ADMIN"}), laporandanasantunananakasuh.LaporanDanaSantunanAnakAsuhCreate)
	app.Patch("/api/laporan-dana-santunan-anak-asuh/:id", middlewares.AuthMiddleware([]string{"BENDAHARA", "ADMIN"}), laporandanasantunananakasuh.LaporanDanaSantunanAnakAsuhUpdate)
	app.Delete("/api/laporan-dana-santunan-anak-asuh/:id", middlewares.AuthMiddleware([]string{"BENDAHARA", "ADMIN"}), laporandanasantunananakasuh.LaporanDanaSantunanAnakAsuhDelete)

	app.Patch("/api/laporan-dana-santunan-anak-asuh/assign/:id", middlewares.AuthMiddleware([]string{"BENDAHARA", "ADMIN"}), laporandanasantunananakasuh.AssignDanaSantunanAnakAsuh)
	app.Patch("/api/laporan-dana-santunan-anak-asuh/tanda-tangan/:id", middlewares.AuthMiddleware([]string{"KETUA_DKM", "ADMIN"}), laporandanasantunananakasuh.AssignTandaTanganAnakAsuh)
	app.Patch("/api/laporan-dana-santunan-anak-asuh/bukti-penggunaan/:id", middlewares.AuthMiddleware([]string{"HUMAS", "ADMIN"}), laporandanasantunananakasuh.AssignBuktiPenggunaanAnakAsuh)
}
