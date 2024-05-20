package datacount

import (
	"server/config"
	"server/models"

	"github.com/gofiber/fiber/v2"
)

func CountDataPengurus(c *fiber.Ctx) error {
	var countPengurus int64
	if err := config.DB.Model(&models.Pengurus{}).Count(&countPengurus).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to count pengurus",
		})
	}

	var countAdmin int64
	var countKetuaDkm int64
	var countBendahara int64
	var countHumas int64

	if err := config.DB.Model(&models.Pengurus{}).Where("jabatan = ?", models.AdminRole).Count(&countAdmin).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Mohon maaf terjadi kesalahan pada server.",
		})
	}
	if err := config.DB.Model(&models.Pengurus{}).Where("jabatan = ?", models.KetuaDkmRole).Count(&countKetuaDkm).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Mohon maaf terjadi kesalahan pada server.",
		})
	}
	if err := config.DB.Model(&models.Pengurus{}).Where("jabatan = ?", models.BendaharaRole).Count(&countBendahara).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Mohon maaf terjadi kesalahan pada server.",
		})
	}
	if err := config.DB.Model(&models.Pengurus{}).Where("jabatan = ?", models.HumasRole).Count(&countHumas).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Mohon maaf terjadi kesalahan pada server.",
		})
	}

	dataCountPengurusJSON := map[string]interface{}{
		"jumlah_pengurus":  countPengurus,
		"jumlah_admin":     countAdmin,
		"jumlah_ketua_dkm": countKetuaDkm,
		"jumlah_bendahara": countBendahara,
		"jumlah_humas":     countHumas,
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": dataCountPengurusJSON,
	})
}
