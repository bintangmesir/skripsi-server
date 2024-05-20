package laporandanasantunan

import (
	"log"
	"server/config"
	"server/models"
	"server/pkg"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm/clause"
)

func AssignTandaTangan(c *fiber.Ctx) error {

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

	namaTandaTangan := form.Value["namaTandaTangan"][0]
	tanggalTandaTangan := form.Value["tanggalTandaTangan"][0]
	tandaTangan := form.File["tandaTangan"]

	tanggalTandaTanganParsedDate, err := time.Parse("Mon Jan 2 2006 15:04:05 GMT+0700 (Western Indonesia Time)", tanggalTandaTangan)
	if err != nil {
		log.Fatal(err)
	}

	newLaporanDanaSantunan := models.LaporanDanaSantunan{
		Validasi:           models.Diterima,
		NamaTandaTangan:    &namaTandaTangan,
		TanggalTandaTangan: tanggalTandaTanganParsedDate,
		UpdatedAt:          time.Now(),
	}

	// * Check if foto exist
	if len(tandaTangan) > 0 {

		// * Handle if foto exist
		if laporanDanaSantunan.TandaTangan != nil {
			err := pkg.DeleteFile(laporanDanaSantunan.TandaTangan, pkg.DIR_IMG_TANDA_TANGAN)
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"message": "Mohon maaf terjadi kesalahan pada server.",
				})
			}
		}

		// * Handle foto
		uploadedFileNames, err := pkg.UploadFile(tandaTangan, pkg.DIR_IMG_TANDA_TANGAN)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Mohon maaf terjadi kesalahan pada server.",
			})
		}
		newLaporanDanaSantunan.TandaTangan = &uploadedFileNames
	}
	config.DB.Model(&laporanDanaSantunan).Updates(&newLaporanDanaSantunan)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Laporan dana santunan berhasil di edit",
	})
}
