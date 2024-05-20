package donatur

import (
	"server/config"
	"server/models"
	"strings"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm/clause"
)

func DonaturGetById(c *fiber.Ctx) error {
	id := c.Params("id")
	donatur := models.Donatur{}

	// * Check if data exist
	if err := config.DB.Preload(clause.Associations).First(&donatur, "donatur_id = ?", id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Data donatur tidak ditemukan.",
		})
	}
	donaturJSON := map[string]interface{}{
		"donatur_id":              donatur.DonaturId,
		"nama":                    donatur.Nama,
		"email":                   donatur.Email,
		"password":                donatur.Password,
		"jenis_kelamin":           donatur.JenisKelamin,
		"no_handphone":            donatur.NoHandphone,
		"validasi":                donatur.Validasi,
		"nama_jalan":              donatur.NamaJalan,
		"rt":                      donatur.Rt,
		"rw":                      donatur.Rw,
		"kelurahan":               donatur.Kelurahan,
		"kecamatan":               donatur.Kecamatan,
		"kota":                    donatur.Kota,
		"provinsi":                donatur.Provinsi,
		"kode_pos":                donatur.KodePos,
		"anak_yatim":              donatur.AnakYatim,
		"dana_santunan_anak_asuh": donatur.DanaSantunanAnakAsuh,
		"created_at":              donatur.CreatedAt,
		"updated_at":              donatur.UpdatedAt,
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
