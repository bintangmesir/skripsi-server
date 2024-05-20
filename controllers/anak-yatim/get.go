package anakyatim

import (
	"server/config"
	"server/models"
	"strings"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm/clause"
)

func AnakYatimGet(c *fiber.Ctx) error {
	// * Check if data exist
	anakYatim := []models.AnakYatim{}
	if err := config.DB.Order("created_at DESC").Preload(clause.Associations).Find(&anakYatim).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Data anak yatim tidak ditemukan.",
		})
	}

	var anakYatimJSON []map[string]interface{}
	for _, ay := range anakYatim {
		ayJSON := map[string]interface{}{
			"anak_yatim_id":           ay.AnakYatimId,
			"nama":                    ay.Nama,
			"status":                  ay.Status,
			"jenis_kelamin":           ay.JenisKelamin,
			"tanggal_lahir":           ay.TanggalLahir,
			"pendidikan":              ay.Pendidikan,
			"pekerjaan_orang_tua":     ay.PekerjaanOrangTua,
			"penghasilan_orang_tua":   ay.PenghasilanOrangTua,
			"tanggungan_orang_tua":    ay.TanggunganOrangTua,
			"status_santunan":         ay.StatusSantunan,
			"nominal_santunan":        ay.NominalSantunan,
			"pengurus_id":             ay.PengurusId,
			"pengurus":                ay.Pengurus,
			"donatur_id":              ay.DonaturId,
			"donatur":                 ay.Donatur,
			"dana_santunan_anak_asuh": ay.DanaSantunanAnakAsuh,
			"created_at":              ay.CreatedAt,
			"updated_at":              ay.UpdatedAt,
		}

		if ay.Foto != nil {
			ayJSON["foto"] = strings.Split(*ay.Foto, ";")
		} else {
			ayJSON["foto"] = []string{}
		}
		anakYatimJSON = append(anakYatimJSON, ayJSON)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": anakYatimJSON,
	})
}
