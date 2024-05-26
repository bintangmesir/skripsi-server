package laporandanasantunananakasuh

import (
	"server/config"
	"server/models"
	"server/pkg"

	"github.com/gofiber/fiber/v2"
)

func LaporanDanaSantunanAnakAsuhDelete(c *fiber.Ctx) error {

	// * Check if data exist
	id := c.Params("id")
	laporanDanaSantunanAnakAsuh := models.LaporanDanaSantunanAnakAsuh{LaporanDanaSantunanAnakAsuhId: id}
	if err := config.DB.First(&laporanDanaSantunanAnakAsuh).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Data laporan dana santunan anak asuh tidak ditemukan.",
		})
	}

	if laporanDanaSantunanAnakAsuh.TandaTangan != nil {
		// * Handle If File	Doesn't Exist
		if err := pkg.DeleteFile(laporanDanaSantunanAnakAsuh.TandaTangan, pkg.DIR_IMG_TANDA_TANGAN); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Mohon maaf terjadi kesalahan pada server.",
			})
		}
	}

	if laporanDanaSantunanAnakAsuh.File != nil {
		// * Handle If File	Doesn't Exist
		if err := pkg.DeleteFile(laporanDanaSantunanAnakAsuh.File, pkg.DIR_IMG_BUKTI_PENGGUNAAN); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Mohon maaf terjadi kesalahan pada server.",
			})
		}
	}

	var danaSantunanAnakAsuh []models.DanaSantunanAnakAsuh
	if err := config.DB.Where("laporan_dana_santunan_anak_asuh_id = ? AND tipe = ?", id, "KREDIT").Find(&danaSantunanAnakAsuh).Error; err == nil {
		for _, dana := range danaSantunanAnakAsuh {
			if dana.File != nil {
				// * Handle If File Doesn't Exist
				if err := pkg.DeleteFile(dana.File, pkg.DIR_IMG_BUKTI_PENGGUNAAN); err != nil {
					return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
						"message": "Mohon maaf terjadi kesalahan pada server.",
					})
				}
			}
		}
		config.DB.Where("laporan_dana_santunan_anak_asuh_id = ? AND tipe = ?", id, "KREDIT").Delete(&models.DanaSantunanAnakAsuh{})
	}

	config.DB.Delete(&laporanDanaSantunanAnakAsuh)
	return c.SendStatus(fiber.StatusNoContent)
}
