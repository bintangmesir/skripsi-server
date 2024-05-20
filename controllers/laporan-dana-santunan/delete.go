package laporandanasantunan

import (
	"server/config"
	"server/models"
	"server/pkg"

	"github.com/gofiber/fiber/v2"
)

func LaporanDanaSantunanDelete(c *fiber.Ctx) error {

	// * Check if data exist
	id := c.Params("id")
	laporanDanaSantunan := models.LaporanDanaSantunan{LaporanDanaSantunanId: id}
	if err := config.DB.First(&laporanDanaSantunan).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Data laporan dana santunan tidak ditemukan.",
		})
	}

	if laporanDanaSantunan.TandaTangan != nil {
		// * Handle If File	Doesn't Exist
		if err := pkg.DeleteFile(laporanDanaSantunan.TandaTangan, pkg.DIR_IMG_TANDA_TANGAN); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Mohon maaf terjadi kesalahan pada server.",
			})
		}
	}

	if laporanDanaSantunan.File != nil {
		// * Handle If File	Doesn't Exist
		if err := pkg.DeleteFile(laporanDanaSantunan.File, pkg.DIR_IMG_BUKTI_PENGGUNAAN); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Mohon maaf terjadi kesalahan pada server.",
			})
		}
	}

	var danaSantunan []models.DanaSantunan
	if err := config.DB.Where("laporan_dana_santunan_id = ? AND tipe = ?", id, "KREDIT").Find(&danaSantunan).Error; err == nil {
		for _, dana := range danaSantunan {
			if dana.File != nil {
				// * Handle If File Doesn't Exist
				if err := pkg.DeleteFile(dana.File, pkg.DIR_IMG_BUKTI_PENGGUNAAN); err != nil {
					return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
						"message": "Mohon maaf terjadi kesalahan pada server.",
					})
				}
			}
		}
		config.DB.Where("laporan_dana_santunan_id = ? AND tipe = ?", id, "KREDIT").Delete(&models.DanaSantunan{})
	}

	config.DB.Delete(&laporanDanaSantunan)
	return c.SendStatus(fiber.StatusNoContent)
}
