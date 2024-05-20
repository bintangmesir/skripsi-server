package datacount

import (
	"server/config"
	"server/models"

	"github.com/gofiber/fiber/v2"
)

func CountDanaSantunan(c *fiber.Ctx) error {

	var totalNominalDiterima int
	if err := config.DB.Model(&models.DanaSantunan{}).Where("validasi = ? AND tipe = ?", models.Diverifikasi, models.Debit).Select("COALESCE(SUM(nominal), 0)").Scan(&totalNominalDiterima).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Mohon maaf terjadi kesalahan pada server.",
		})
	}

	var totalNominalDiterima2 int
	if err := config.DB.Model(&models.DanaSantunanAnakAsuh{}).Where("validasi = ?", models.Diverifikasi).Select("COALESCE(SUM(nominal), 0)").Scan(&totalNominalDiterima2).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Mohon maaf terjadi kesalahan pada server.",
		})
	}

	var countPending int64
	if err := config.DB.Model(&models.DanaSantunan{}).Where("validasi = ?", models.Pending).Count(&countPending).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Mohon maaf terjadi kesalahan pada server.",
		})
	}

	var countDiverifikasi int64
	if err := config.DB.Model(&models.DanaSantunan{}).Where("validasi = ?", models.Diverifikasi).Count(&countDiverifikasi).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Mohon maaf terjadi kesalahan pada server.",
		})
	}

	var countDitolak int64
	if err := config.DB.Model(&models.DanaSantunan{}).Where("validasi = ?", models.Ditolak).Count(&countDitolak).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Mohon maaf terjadi kesalahan pada server.",
		})
	}

	dataCountDanaSantunanJSON := map[string]interface{}{
		"jumlah_nominal_keseluruhan_diterima":           totalNominalDiterima,
		"jumlah_nominal_keseluruhan_diterima_anak_asuh": totalNominalDiterima2,
		"jumlah_dana_santunan_pending":                  countPending,
		"jumlah_dana_santunan_diverifikasi":             countDiverifikasi,
		"jumlah_dana_santunan_ditolak":                  countDitolak,
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": dataCountDanaSantunanJSON,
	})
}
