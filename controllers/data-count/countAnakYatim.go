package datacount

import (
	"server/config"
	"server/models"

	"github.com/gofiber/fiber/v2"
)

func CountDataAnakYatim(c *fiber.Ctx) error {

	//* Count jumlah anak yatim
	var countAnakYatim int64
	result := config.DB.Model(&models.AnakYatim{}).Count(&countAnakYatim)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Mohon maaf terjadi kesalahan pada server.",
		})
	}

	//* Count jumlah anak asuh
	var countAnakAsuh int64
	result2 := config.DB.Model(&models.AnakYatim{}).Where("donatur_id IS NOT NULL").Count(&countAnakAsuh)
	if result2.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Mohon maaf terjadi kesalahan pada server.",
		})
	}

	//* Count status anak yatim
	var countYatim int64
	if err := config.DB.Model(&models.AnakYatim{}).Where("status = ?", "Yatim").Count(&countYatim).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Mohon maaf terjadi kesalahan pada server.",
		})
	}

	//* Count status anak yatim	piatu
	var countYatimPiatu int64
	if err := config.DB.Model(&models.AnakYatim{}).Where("status = ?", "Yatim Piatu").Count(&countYatimPiatu).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Mohon maaf terjadi kesalahan pada server.",
		})
	}

	//* Count jenis kelamin anak yatim
	var countLakiLaki int64
	var countPerempuan int64
	if err := config.DB.Model(&models.AnakYatim{}).Where("jenis_kelamin = ?", "L").Count(&countLakiLaki).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Mohon maaf terjadi kesalahan pada server.",
		})
	}
	if err := config.DB.Model(&models.AnakYatim{}).Where("jenis_kelamin = ?", "P").Count(&countPerempuan).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Mohon maaf terjadi kesalahan pada server.",
		})
	}

	//* Count jumlah orang tua asuh
	var donatur []models.Donatur
	if err := config.DB.Find(&donatur).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Mohon maaf terjadi kesalahan pada server.",
		})
	}

	var countOrangTuaAsuh int
	var countCalonOrangTuaAsuh int
	for _, donatur := range donatur {
		var orangTuaAsuh string
		if err := config.DB.Model(&models.Donatur{}).Select("validasi").Where("donatur_id = ?", donatur.DonaturId).Scan(&orangTuaAsuh).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Mohon maaf terjadi kesalahan pada server.",
			})
		}
		if orangTuaAsuh == "ORANG_TUA_ASUH" {
			countOrangTuaAsuh++
		} else if orangTuaAsuh == "DONATUR" {
			countCalonOrangTuaAsuh++
		}
	}

	//* Count jumlah pendidikan anak yatim
	var countSD int64
	var countSMP int64
	var countSMA int64
	var countTidakBersekolah int64
	if err := config.DB.Model(&models.AnakYatim{}).Where("pendidikan = ?", "SD").Count(&countSD).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Mohon maaf terjadi kesalahan pada server.",
		})
	}
	if err := config.DB.Model(&models.AnakYatim{}).Where("pendidikan = ?", "SMP").Count(&countSMP).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Mohon maaf terjadi kesalahan pada server.",
		})
	}
	if err := config.DB.Model(&models.AnakYatim{}).Where("pendidikan = ?", "SMA").Count(&countSMA).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Mohon maaf terjadi kesalahan pada server.",
		})
	}
	if err := config.DB.Model(&models.AnakYatim{}).Where("pendidikan = ?", "Tidak bersekolah").Count(&countTidakBersekolah).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Mohon maaf terjadi kesalahan pada server.",
		})
	}

	dataCountAnakYatimJSON := map[string]interface{}{
		"jumlah_anak_asuh":                   countAnakAsuh,
		"jumlah_anak_yatim":                  countAnakYatim,
		"jumlah_status_yatim":                countYatim,
		"jumlah_status_yatim_piatu":          countYatimPiatu,
		"jumlah_anak_yatim_sd":               countSD,
		"jumlah_anak_yatim_smp":              countSMP,
		"jumlah_anak_yatim_sma":              countSMA,
		"jumlah_anak_yatim_tidak_bersekolah": countTidakBersekolah,
		"jumlah_anak_yatim_laki_laki":        countLakiLaki,
		"jumlah_anak_yatim_perempuan":        countPerempuan,
		"jumlah_orang_tua_asuh":              countOrangTuaAsuh,
		"jumlah_calon_orang_tua_asuh":        countCalonOrangTuaAsuh,
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": dataCountAnakYatimJSON,
	})
}
