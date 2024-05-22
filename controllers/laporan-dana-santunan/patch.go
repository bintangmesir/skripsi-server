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

	var danaSantunan []models.DanaSantunan
	if err := config.DB.Where("laporan_dana_santunan_id = ? AND tipe = ?", id, "KREDIT").Find(&danaSantunan).Error; err == nil {
		for _, dana := range danaSantunan {
			if dana.File != nil {
				// * Handle If File Doesn't Exist
				if err := pkg.DeleteFile(dana.File, pkg.DIR_IMG_BUKTI_PENGGUNAAN); err != nil {
					return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
						"message": "Mohon maaf terjadi kesalahan pada server.",
					})
				}
			}
		}
		config.DB.Where("laporan_dana_santunan_id = ? AND tipe = ?", id, "KREDIT").Delete(&models.DanaSantunan{})
	}

	config.DB.Delete(&laporanDanaSantunan)

	newLaporanDanaSantunan := models.LaporanDanaSantunan{
		LaporanDanaSantunanId: laporanDanaSantunanId,
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
		"data": laporanDanaSantunanId,
	})
}
