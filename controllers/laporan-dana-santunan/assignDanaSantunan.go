package laporandanasantunan

import (
	"log"
	"server/config"
	danasantunan "server/controllers/dana-santunan"
	"server/models"
	"server/pkg"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm/clause"
)

func AssignDanaSantunan(c *fiber.Ctx) error {

	idParams := c.Params("id")
	// * Check if data exist
	laporanDanaSantunan := models.LaporanDanaSantunan{LaporanDanaSantunanId: idParams}
	if err := config.DB.Preload(clause.Associations).First(&laporanDanaSantunan).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Data laporan dana santunan tidak ditemukan.",
		})
	}

	form, err := c.MultipartForm()
	if err != nil {
		return err
	}

	danaSantunanId := form.Value["danaSantunanId"][0]
	tanggal := form.Value["tanggal"][0]
	tipe := form.Value["tipe"][0]
	validasi := form.Value["validasi"][0]
	index := form.Value["index"][0]
	nominal := form.Value["nominal"][0]
	keterangan := form.Value["keterangan"][0]

	//* Handle tanggal
	tanggalParsedDate, err := time.Parse("Mon Jan 2 2006 15:04:05 GMT+0700 (Western Indonesia Time)", tanggal)
	if err != nil {
		log.Fatal(err)
	}

	indexConverted, err := strconv.Atoi(index)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Mohon maaf terjadi kesalahan pada server.",
		})
	}

	nominalConverted, err := strconv.Atoi(nominal)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Mohon maaf terjadi kesalahan pada server.",
		})
	}

	id, err := pkg.GetUserActive(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Mohon login terlebih dahulu.",
		})
	}

	if danaSantunanId == "" {
		danaSantunanId, err := danasantunan.HandleDanaSantunanId()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Mohon maaf terjadi kesalahan pada server.",
			})
		}

		newDanaSantunan := models.DanaSantunan{
			DanaSantunanId:        danaSantunanId,
			Index:                 &indexConverted,
			Tipe:                  models.PembayaranEnum(tipe),
			Tanggal:               tanggalParsedDate,
			Keterangan:            &keterangan,
			Nominal:               nominalConverted,
			Validasi:              models.ValidationEnum(validasi),
			UpdatedAt:             time.Now(),
			PengurusId:            &id,
			LaporanDanaSantunanId: &idParams,
		}

		config.DB.Create(&newDanaSantunan)
	} else {
		danaSantunan := models.DanaSantunan{DanaSantunanId: danaSantunanId}
		if err := config.DB.First(&danaSantunan).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "Data dana santunan tidak ditemukan.",
			})
		}
		newDanaSantunan := models.DanaSantunan{
			Index:                 &indexConverted,
			Tipe:                  models.PembayaranEnum(tipe),
			Tanggal:               tanggalParsedDate,
			Keterangan:            &keterangan,
			Nominal:               nominalConverted,
			Validasi:              models.ValidationEnum(validasi),
			UpdatedAt:             time.Now(),
			LaporanDanaSantunanId: &idParams,
			PengurusId:            &id,
		}

		config.DB.Model(&danaSantunan).Updates(&newDanaSantunan)
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Lpaoran dana santunan berhasil dibuat",
	})
}
