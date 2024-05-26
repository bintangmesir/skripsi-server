package laporandanasantunananakasuh

import (
	"server/config"
	"server/models"
	"server/pkg"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm/clause"
)

func LaporanDanaSantunanAnakAsuhUpdate(c *fiber.Ctx) error {

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

	judul := form.Value["judul"][0]
	keterangan := form.Value["keterangan"][0]
	saldoAwal := form.Value["saldoAwal"][0]
	saldoSisa := form.Value["saldoSisa"][0]
	donaturId := form.Value["donaturId"][0]
	tandaTangan := laporanDanaSantunanAnakAsuh.TandaTangan
	buktiPenggunaan := laporanDanaSantunanAnakAsuh.File

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

	id, err := pkg.GetUserActive(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Mohon login terlebih dahulu.",
		})
	}

	var danaSantunanAnakAsuh []models.DanaSantunanAnakAsuh
	if err := config.DB.Where("laporan_dana_santunan_anak_asuh_id = ? AND tipe = ?", idParams, "KREDIT").Find(&danaSantunanAnakAsuh).Error; err == nil {
		for _, dana := range danaSantunanAnakAsuh {
			if dana.File != nil {
				// * Handle If File Doesn't Exist
				if err := pkg.DeleteFile(dana.File, pkg.DIR_IMG_BUKTI_PENGGUNAAN); err != nil {
					return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
						"message": "Mohon maaf terjadi kesalahan pada server.",
					})
				}
			}
		}
		config.DB.Where("laporan_dana_santunan_anak_asuh_id = ? AND tipe = ?", idParams, "KREDIT").Delete(&models.DanaSantunanAnakAsuh{})
	}

	config.DB.Delete(&laporanDanaSantunanAnakAsuh)

	newLaporanDanaSantunanAnakAsuh := models.LaporanDanaSantunanAnakAsuh{
		LaporanDanaSantunanAnakAsuhId: idParams,
		Judul:                         judul,
		Keterangan:                    &keterangan,
		SaldoAwal:                     saldoAwalConverted,
		SaldoSisa:                     saldoSisaConverted,
		TandaTangan:                   tandaTangan,
		File:                          buktiPenggunaan,
		PengurusId:                    &id,
		DonaturId:                     &donaturId,
	}

	config.DB.Create(&newLaporanDanaSantunanAnakAsuh)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"data": idParams,
	})
}
