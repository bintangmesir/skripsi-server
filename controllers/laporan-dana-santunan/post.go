package laporandanasantunan

import (
	"server/config"
	"server/models"
	"server/pkg"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func LaporanDanaSantunanCreate(c *fiber.Ctx) error {
	form, err := c.MultipartForm()
	if err != nil {
		return err
	}

	judul := form.Value["judul"][0]
	keterangan := form.Value["keterangan"][0]
	saldoAwal := form.Value["saldoAwal"][0]
	saldoSisa := form.Value["saldoSisa"][0]

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
	laporanDanaSantunanId, err := HandleLaporanDanaSantunanId()
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

	newLaporanDanaSantunan := models.LaporanDanaSantunan{
		LaporanDanaSantunanId: laporanDanaSantunanId,
		Judul:                 judul,
		Keterangan:            &keterangan,
		SaldoAwal:             saldoAwalConverted,
		SaldoSisa:             saldoSisaConverted,
		PengurusId:            &id,
	}

	config.DB.Create(&newLaporanDanaSantunan)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"data": laporanDanaSantunanId,
	})
}
