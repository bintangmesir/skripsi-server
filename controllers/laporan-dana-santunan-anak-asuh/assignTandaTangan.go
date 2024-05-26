package laporandanasantunananakasuh

import (
	"log"
	"server/config"
	"server/models"
	"server/pkg"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm/clause"
)

func AssignTandaTanganAnakAsuh(c *fiber.Ctx) error {

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

	namaTandaTangan := form.Value["namaTandaTangan"][0]
	tanggalTandaTangan := form.Value["tanggalTandaTangan"][0]
	tandaTangan := form.File["tandaTangan"]

	tanggalTandaTanganParsedDate, err := time.Parse("Mon Jan 2 2006 15:04:05 GMT+0700 (Western Indonesia Time)", tanggalTandaTangan)
	if err != nil {
		log.Fatal(err)
	}

	newLaporanDanaSantunanAnakAsuh := models.LaporanDanaSantunanAnakAsuh{
		Validasi:           models.Diterima,
		NamaTandaTangan:    &namaTandaTangan,
		TanggalTandaTangan: tanggalTandaTanganParsedDate,
		UpdatedAt:          time.Now(),
	}

	// * Check if foto exist
	if len(tandaTangan) > 0 {

		// * Handle if foto exist
		if laporanDanaSantunanAnakAsuh.TandaTangan != nil {
			err := pkg.DeleteFile(laporanDanaSantunanAnakAsuh.TandaTangan, pkg.DIR_IMG_TANDA_TANGAN)
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
		newLaporanDanaSantunanAnakAsuh.TandaTangan = &uploadedFileNames
	}
	config.DB.Model(&laporanDanaSantunanAnakAsuh).Updates(&newLaporanDanaSantunanAnakAsuh)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Laporan dana santunan anak asuh berhasil di edit",
	})
}
