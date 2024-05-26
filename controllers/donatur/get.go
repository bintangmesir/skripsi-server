package donatur

import (
	"server/config"
	"server/models"
	"strings"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm/clause"
)

func DonaturGet(c *fiber.Ctx) error {
	donatur := []models.Donatur{}
	result := config.DB.Preload(clause.Associations).Find(&donatur)

	// * Check if data exist
	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Data donatur tidak ditemukan.",
		})
	}

	var donaturJSON []map[string]interface{}
	for _, p := range donatur {
		pJSON := map[string]interface{}{
			"donatur_id":              p.DonaturId,
			"nama":                    p.Nama,
			"email":                   p.Email,
			"jenis_kelamin":           p.JenisKelamin,
			"no_handphone":            p.NoHandphone,
			"validasi":                p.Validasi,
			"pengurus_id":             p.PengurusId,
			"anak_yatim":              p.AnakYatim,
			"dana_santunan_anak_asuh": p.DanaSantunanAnakAsuh,
			"created_at":              p.CreatedAt,
			"updated_at":              p.UpdatedAt,
		}
		if p.Foto != nil {
			pJSON["foto"] = strings.Split(*p.Foto, ";")
		} else {
			pJSON["foto"] = []string{}
		}
		donaturJSON = append(donaturJSON, pJSON)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": donaturJSON,
	})
}
