package datacount

import (
	"server/config"
	"server/models"

	"github.com/gofiber/fiber/v2"
)

func CountDataDonatur(c *fiber.Ctx) error {
	//* Count jumlah donatur
	var countDonatur int64
	if err := config.DB.Model(&models.Donatur{}).Count(&countDonatur).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to count donatur",
		})
	}

	//* Count jumlah donatur dan calon donatur
	var countOrangTuaAsuhRole int64
	var countDonaturRole int64

	if err := config.DB.Model(&models.Donatur{}).Where("validasi = ?", models.OrangTuaAsuhRole).Count(&countOrangTuaAsuhRole).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Mohon maaf terjadi kesalahan pada server.",
		})
	}
	if err := config.DB.Model(&models.Donatur{}).Where("validasi = ?", models.DonaturRole).Count(&countDonaturRole).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Mohon maaf terjadi kesalahan pada server.",
		})
	}

	dataCountDonaturJSON := map[string]interface{}{
		"jumlah_donatur":        countDonatur,
		"jumlah_orang_tua_asuh": countOrangTuaAsuhRole,
		"jumlah_calon_donatur":  countDonaturRole,
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": dataCountDonaturJSON,
	})
}
