package datacount

import (
	"server/config"
	"server/models"
	"server/pkg"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

func CountDanaSantunanAnakAsuh(c *fiber.Ctx) error {

	// * Check if cookies valid
	cookie := c.Cookies("jwt")
	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(pkg.SECRET_KEY), nil
	})
	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "Tolong login terlebih dahulu.",
		})
	}

	// * Get values of Cookies
	claims := token.Claims.(*jwt.StandardClaims)

	var donatur models.Donatur
	if err := config.DB.Where("donatur_id = ?", claims.Issuer).First(&donatur).Error; err == nil {

		var countPending int64
		if err := config.DB.Model(&models.DanaSantunanAnakAsuh{}).Where("validasi = ? AND donatur_id = ?", models.Pending, donatur.DonaturId).Count(&countPending).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Mohon maaf terjadi kesalahan pada server.",
			})
		}

		var countDiverifikasi int64
		if err := config.DB.Model(&models.DanaSantunanAnakAsuh{}).Where("validasi = ? AND donatur_id = ?", models.Diverifikasi, donatur.DonaturId).Count(&countDiverifikasi).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Mohon maaf terjadi kesalahan pada server.",
			})
		}

		var countDitolak int64
		if err := config.DB.Model(&models.DanaSantunanAnakAsuh{}).Where("validasi = ? AND donatur_id = ?", models.Ditolak, donatur.DonaturId).Count(&countDitolak).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Mohon maaf terjadi kesalahan pada server.",
			})
		}

		var totalNominalDonatur int
		if err := config.DB.Model(&models.DanaSantunanAnakAsuh{}).Where("validasi = ? AND donatur_id = ?", models.Diverifikasi, donatur.DonaturId).Select("COALESCE(SUM(nominal), 0)").Scan(&totalNominalDonatur).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Mohon maaf terjadi kesalahan pada server.",
			})
		}
		dataCountDanaSantunanAnakAsuhJSON := map[string]interface{}{
			"jumlah_nominal_dana_santunan_anak_asuh_donatur": totalNominalDonatur,
			"jumlah_dana_santunan_anak_asuh_pending":         countPending,
			"jumlah_dana_santunan_anak_asuh_diverifikasi":    countDiverifikasi,
			"jumlah_dana_santunan_anak_asuh_ditolak":         countDitolak,
		}
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"data": dataCountDanaSantunanAnakAsuhJSON,
		})
	}

	var totalNominalDiterima int
	if err := config.DB.Model(&models.DanaSantunanAnakAsuh{}).Where("validasi = ?", models.Diverifikasi).Select("COALESCE(SUM(nominal), 0)").Scan(&totalNominalDiterima).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Mohon maaf terjadi kesalahan pada server.",
		})
	}

	var countPending int64
	if err := config.DB.Model(&models.DanaSantunanAnakAsuh{}).Where("validasi = ?", models.Pending).Count(&countPending).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Mohon maaf terjadi kesalahan pada server.",
		})
	}

	var countDiverifikasi int64
	if err := config.DB.Model(&models.DanaSantunanAnakAsuh{}).Where("validasi = ?", models.Diverifikasi).Count(&countDiverifikasi).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Mohon maaf terjadi kesalahan pada server.",
		})
	}

	var countDitolak int64
	if err := config.DB.Model(&models.DanaSantunanAnakAsuh{}).Where("validasi = ?", models.Ditolak).Count(&countDitolak).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Mohon maaf terjadi kesalahan pada server.",
		})
	}

	dataCountDanaSantunanAnakAsuhJSON := map[string]interface{}{
		"jumlah_nominal_keseluruhan_diterima":         totalNominalDiterima,
		"jumlah_dana_santunan_anak_asuh_pending":      countPending,
		"jumlah_dana_santunan_anak_asuh_diverifikasi": countDiverifikasi,
		"jumlah_dana_santunan_anak_asuh_ditolak":      countDitolak,
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": dataCountDanaSantunanAnakAsuhJSON,
	})
}
