package routes

import (
	laporandanasantunan "server/controllers/laporan-dana-santunan"
	"server/middlewares"

	"github.com/gofiber/fiber/v2"
)

func LaporanDanaSantunanRoute(app *fiber.App) {
	app.Get("/api/laporan-dana-santunan", middlewares.AuthMiddleware([]string{"BENDAHARA", "ADMIN", "HUMAS", "KETUA_DKM"}), laporandanasantunan.LaporanDanaSantunanGet)
	app.Get("/api/laporan-dana-santunan/:id", middlewares.AuthMiddleware([]string{"BENDAHARA", "ADMIN", "HUMAS", "KETUA_DKM"}), laporandanasantunan.LaporanDanaSantunanGetById)
	app.Post("/api/laporan-dana-santunan", middlewares.AuthMiddleware([]string{"BENDAHARA", "ADMIN"}), laporandanasantunan.LaporanDanaSantunanCreate)
	app.Patch("/api/laporan-dana-santunan/:id", middlewares.AuthMiddleware([]string{"BENDAHARA", "ADMIN"}), laporandanasantunan.LaporanDanaSantunanUpdate)
	app.Delete("/api/laporan-dana-santunan/:id", middlewares.AuthMiddleware([]string{"BENDAHARA", "ADMIN"}), laporandanasantunan.LaporanDanaSantunanDelete)

	app.Patch("/api/laporan-dana-santunan/assign/:id", middlewares.AuthMiddleware([]string{"BENDAHARA", "ADMIN"}), laporandanasantunan.AssignDanaSantunan)
	app.Patch("/api/laporan-dana-santunan/tanda-tangan/:id", middlewares.AuthMiddleware([]string{"KETUA_DKM", "ADMIN"}), laporandanasantunan.AssignTandaTangan)
	app.Patch("/api/laporan-dana-santunan/bukti-penggunaan/:id", middlewares.AuthMiddleware([]string{"HUMAS", "ADMIN"}), laporandanasantunan.AssignBuktiPenggunaan)
}
