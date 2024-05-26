package laporandanasantunananakasuh

import (
	"server/config"
	"server/models"
	"server/pkg"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm/clause"
)

func AssignBuktiPenggunaanAnakAsuh(c *fiber.Ctx) error {

	id := c.Params("id")
	// * Check if data exist
	laporanDanaSantunanAnakAsuh := models.LaporanDanaSantunanAnakAsuh{LaporanDanaSantunanAnakAsuhId: id}
	if err := config.DB.Preload(clause.Associations).First(&laporanDanaSantunanAnakAsuh).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Data laporan dana santunan anak asuh tidak ditemukan.",
		})
	}

	form, err := c.MultipartForm()
	if err != nil {
		return err
	}

	file := form.File["file"]

	newLaporanDanaSantunanAnakAsuh := models.LaporanDanaSantunanAnakAsuh{
		UpdatedAt: time.Now(),
	}

	// * Check if foto exist
	if len(file) > 0 {

		// * Handle if foto exist
		if laporanDanaSantunanAnakAsuh.File != nil {
			err := pkg.DeleteFile(laporanDanaSantunanAnakAsuh.File, pkg.DIR_IMG_BUKTI_PENGGUNAAN)
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
		newLaporanDanaSantunanAnakAsuh.File = &uploadedFileNames
	}
	config.DB.Model(&laporanDanaSantunanAnakAsuh).Updates(&newLaporanDanaSantunanAnakAsuh)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Laporan dana santunan anak asuh berhasil di edit",
	})
}
