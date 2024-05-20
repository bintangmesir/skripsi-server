package auth

import (
	"server/config"
	"server/models"
	"server/pkg"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"gorm.io/gorm/clause"
)

func User(c *fiber.Ctx) error {

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
	var pengurus models.Pengurus
	var donatur models.Donatur

	if err := config.DB.Preload(clause.Associations).Where("pengurus_id = ?", claims.Issuer).First(&pengurus).Error; err == nil {
		pengurusJSON := map[string]interface{}{
			"pengurus_id":  pengurus.PengurusId,
			"nama":         pengurus.Nama,
			"email":        pengurus.Email,
			"no_handphone": pengurus.NoHandphone,
			"jabatan":      pengurus.Jabatan,
			"admin_id":     pengurus.AdminId,
			"pengurus":     pengurus.Pengurus,
			"created_at":   pengurus.CreatedAt,
			"updated_at":   pengurus.UpdatedAt,
		}

		if pengurus.Foto != nil {
			pengurusJSON["foto"] = strings.Split(*pengurus.Foto, ";")
		} else {
			pengurusJSON["foto"] = []string{}
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"data": pengurusJSON,
		})
	}

	if err := config.DB.Preload(clause.Associations).Where("donatur_id = ?", claims.Issuer).First(&donatur).Error; err == nil {
		donaturJSON := map[string]interface{}{
			"donatur_id":    donatur.DonaturId,
			"nama":          donatur.Nama,
			"email":         donatur.Email,
			"password":      donatur.Password,
			"jenis_kelamin": donatur.JenisKelamin,
			"no_handphone":  donatur.NoHandphone,
			"validasi":      donatur.Validasi,
			"nama_jalan":    donatur.NamaJalan,
			"rt":            donatur.Rt,
			"rw":            donatur.Rw,
			"kelurahan":     donatur.Kelurahan,
			"kecamatan":     donatur.Kecamatan,
			"kota":          donatur.Kota,
			"provinsi":      donatur.Provinsi,
			"kode_pos":      donatur.KodePos,
			"anak_yatim":    donatur.AnakYatim,
			"created_at":    donatur.CreatedAt,
			"updated_at":    donatur.UpdatedAt,
		}

		if donatur.Foto != nil {
			donaturJSON["foto"] = strings.Split(*donatur.Foto, ";")
		} else {
			donaturJSON["foto"] = []string{}
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"data": donaturJSON,
		})
	}

	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
		"message": "Pengguna tidak ditemukan.",
	})
}
