package anakyatim

import (
	"server/config"
	"server/models"
	"strings"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm/clause"
)

func AnakYatimGetById(c *fiber.Ctx) error {
	id := c.Params("id")
	anakYatim := models.AnakYatim{}

	// * Check if data exist
	if err := config.DB.Preload(clause.Associations).First(&anakYatim, "anak_yatim_id = ?", id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Data anak yatim tidak ditemukan.",
		})
	}

	anakYatimJSON := map[string]interface{}{
		"anak_yatim_id":           anakYatim.AnakYatimId,
		"nama":                    anakYatim.Nama,
		"status":                  anakYatim.Status,
		"jenis_kelamin":           anakYatim.JenisKelamin,
		"tanggal_lahir":           anakYatim.TanggalLahir,
		"pendidikan":              anakYatim.Pendidikan,
		"pekerjaan_orang_tua":     anakYatim.PekerjaanOrangTua,
		"penghasilan_orang_tua":   anakYatim.PenghasilanOrangTua,
		"tanggungan_orang_tua":    anakYatim.TanggunganOrangTua,
		"deskripsi":               anakYatim.Deskripsi,
		"status_santunan":         anakYatim.StatusSantunan,
		"nominal_santunan":        anakYatim.NominalSantunan,
		"donatur_id":              anakYatim.DonaturId,
		"donatur":                 anakYatim.Donatur,
		"pengurus_id":             anakYatim.Pengurus,
		"dana_santunan_anak_asuh": anakYatim.DanaSantunanAnakAsuh,
		"created_at":              anakYatim.CreatedAt,
		"updated_at":              anakYatim.UpdatedAt,
	}

	if anakYatim.Foto != nil {
		anakYatimJSON["foto"] = strings.Split(*anakYatim.Foto, ";")
	} else {
		anakYatimJSON["foto"] = []string{}
	}
	anakYatimJSON["kebutuhan"] = strings.Split(anakYatim.Kebutuhan, ",")

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": anakYatimJSON,
	})
}
