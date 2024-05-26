package laporandanasantunananakasuh

import (
	"server/config"
	"server/models"
	"server/pkg"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func LaporanDanaSantunanAnakAsuhCreate(c *fiber.Ctx) error {
	form, err := c.MultipartForm()
	if err != nil {
		return err
	}

	judul := form.Value["judul"][0]
	keterangan := form.Value["keterangan"][0]
	saldoAwal := form.Value["saldoAwal"][0]
	saldoSisa := form.Value["saldoSisa"][0]
	donaturId := form.Value["donaturId"][0]

	donatur := models.Donatur{DonaturId: donaturId}
	if err := config.DB.First(&donatur).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Data donatur tidak ditemukan.",
		})
	}

	saldoAwalConverted, err := strconv.Atoi(saldoAwal)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Mohon maaf terjadi kesalahan pada server.",
		})
	}

	saldoSisaConverted, err := strconv.Atoi(saldoSisa)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Mohon maaf terjadi kesalahan pada server.",
		})
	}

	// * Handle Anak Yatim Id
	laporanDanaSantunanAnakAsuhId, err := HandleLaporanDanaSantunanAnakAsuhId()
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

	newLaporanDanaSantunanAnakAsuh := models.LaporanDanaSantunanAnakAsuh{
		LaporanDanaSantunanAnakAsuhId: laporanDanaSantunanAnakAsuhId,
		Judul:                         judul,
		Keterangan:                    &keterangan,
		SaldoAwal:                     saldoAwalConverted,
		SaldoSisa:                     saldoSisaConverted,
		PengurusId:                    &id,
		DonaturId:                     &donaturId,
	}

	config.DB.Create(&newLaporanDanaSantunanAnakAsuh)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"data": laporanDanaSantunanAnakAsuhId,
	})
}
