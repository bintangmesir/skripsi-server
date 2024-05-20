package laporandanasantunan

import (
	"server/config"
	"server/models"
	"server/pkg"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm/clause"
)

func AssignBuktiPenggunaan(c *fiber.Ctx) error {

	id := c.Params("id")
	// * Check if data exist
	laporanDanaSantunan := models.LaporanDanaSantunan{LaporanDanaSantunanId: id}
	if err := config.DB.Preload(clause.Associations).First(&laporanDanaSantunan).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Data laporan dana santunan tidak ditemukan.",
		})
	}

	form, err := c.MultipartForm()
	if err != nil {
		return err
	}

	file := form.File["file"]

	newLaporanDanaSantunan := models.LaporanDanaSantunan{
		UpdatedAt: time.Now(),
	}

	// * Check if foto exist
	if len(file) > 0 {

		// * Handle if foto exist
		if laporanDanaSantunan.File != nil {
			err := pkg.DeleteFile(laporanDanaSantunan.File, pkg.DIR_IMG_BUKTI_PENGGUNAAN)
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"message": "Mohon maaf terjadi kesalahan pada server.",
				})
			}
		}

		// * Handle foto
		uploadedFileNames, err := pkg.UploadFile(file, pkg.DIR_IMG_BUKTI_PENGGUNAAN)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Mohon maaf terjadi kesalahan pada server.",
			})
		}
		newLaporanDanaSantunan.File = &uploadedFileNames
	}
	config.DB.Model(&laporanDanaSantunan).Updates(&newLaporanDanaSantunan)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Laporan dana santunan berhasil di edit",
	})
}
