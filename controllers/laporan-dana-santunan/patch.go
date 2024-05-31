package laporandanasantunan

import (
	"server/config"
	"server/models"
	"server/pkg"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm/clause"
)

func LaporanDanaSantunanUpdate(c *fiber.Ctx) error {

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

	judul := form.Value["judul"][0]
	keterangan := form.Value["keterangan"][0]
	saldoAwal := form.Value["saldoAwal"][0]
	saldoSisa := form.Value["saldoSisa"][0]
	tandaTangan := laporanDanaSantunan.TandaTangan
	buktiPenggunaan := laporanDanaSantunan.File

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

	id, err := pkg.GetUserActive(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Mohon login terlebih dahulu.",
		})
	}

	config.DB.Delete(&laporanDanaSantunan)

	newLaporanDanaSantunan := models.LaporanDanaSantunan{
		LaporanDanaSantunanId: idParams,
		Judul:                 judul,
		Keterangan:            &keterangan,
		SaldoAwal:             saldoAwalConverted,
		SaldoSisa:             saldoSisaConverted,
		TandaTangan:           tandaTangan,
		File:                  buktiPenggunaan,
		PengurusId:            &id,
	}

	config.DB.Create(&newLaporanDanaSantunan)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"data": idParams,
	})
}
