package laporandanasantunananakasuh

import (
	"log"
	"server/config"
	danasantunananakasuh "server/controllers/dana-santunan-anak-asuh"
	"server/models"
	"server/pkg"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm/clause"
)

func AssignDanaSantunanAnakAsuh(c *fiber.Ctx) error {

	idParams := c.Params("id")
	// * Check if data exist
	laporanDanaSantunanAnakAsuh := models.LaporanDanaSantunanAnakAsuh{LaporanDanaSantunanAnakAsuhId: idParams}
	if err := config.DB.Preload(clause.Associations).First(&laporanDanaSantunanAnakAsuh).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Data laporan dana santunan anak asuh tidak ditemukan.",
		})
	}

	form, err := c.MultipartForm()
	if err != nil {
		return err
	}

	danaSantunanAnakAsuhId := form.Value["danaSantunanAnakAsuhId"][0]
	tanggal := form.Value["tanggal"][0]
	keterangan := form.Value["keterangan"][0]
	tipe := form.Value["tipe"][0]
	validasi := form.Value["validasi"][0]
	index := form.Value["index"][0]
	nominal := form.Value["nominal"][0]

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

	if danaSantunanAnakAsuhId == "" {
		danaSantunanAnakAsuhId, err := danasantunananakasuh.HandleDanaSantunanAnakAsuhId()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Mohon maaf terjadi kesalahan pada server.",
			})
		}

		newDanaSantunanAnakAsuh := models.DanaSantunanAnakAsuh{
			DanaSantunanAnakAsuhId:        danaSantunanAnakAsuhId,
			Index:                         &indexConverted,
			Tipe:                          models.PembayaranEnum(tipe),
			Keterangan:                    &keterangan,
			Tanggal:                       tanggalParsedDate,
			Nominal:                       nominalConverted,
			Validasi:                      models.ValidationEnum(validasi),
			UpdatedAt:                     time.Now(),
			PengurusId:                    &id,
			LaporanDanaSantunanAnakAsuhId: &idParams,
		}

		config.DB.Create(&newDanaSantunanAnakAsuh)
	} else {
		danaSantunanAnakAsuh := models.DanaSantunanAnakAsuh{DanaSantunanAnakAsuhId: danaSantunanAnakAsuhId}
		if err := config.DB.First(&danaSantunanAnakAsuh).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "Data dana santunan anak asuh tidak ditemukan.",
			})
		}
		newDanaSantunanAnakAsuh := models.DanaSantunanAnakAsuh{
			Index:                         &indexConverted,
			Tipe:                          models.PembayaranEnum(tipe),
			Tanggal:                       tanggalParsedDate,
			Nominal:                       nominalConverted,
			Validasi:                      models.ValidationEnum(validasi),
			UpdatedAt:                     time.Now(),
			LaporanDanaSantunanAnakAsuhId: &idParams,
			PengurusId:                    &id,
		}

		config.DB.Model(&danaSantunanAnakAsuh).Updates(&newDanaSantunanAnakAsuh)
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Lpaoran dana santunan anak asuh berhasil dibuat",
	})
}
