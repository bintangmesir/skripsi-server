package danasantunananakasuh

import (
	"server/config"
	"server/models"
	"strings"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm/clause"
)

// ! Get Data Dana Zis
func DanaSantunaAnakAsuhnGet(c *fiber.Ctx) error {

	// * Check if data exist
	danaSantunanAnakAsuh := []models.DanaSantunanAnakAsuh{}
	if err := config.DB.Preload(clause.Associations).Order("created_at DESC").Find(&danaSantunanAnakAsuh).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Data santunan tidak ditemukan.",
		})
	}

	var danaSantunanAnakAsuhJSON []map[string]interface{}
	for _, ds := range danaSantunanAnakAsuh {

		dsJSON := map[string]interface{}{
			"dana_santunan_anak_asuh_id": ds.DanaSantunanAnakAsuhId,
			"index":                      ds.Index,
			"tanggal":                    ds.Tanggal,
			"keterangan":                 ds.Keterangan,
			"tipe":                       ds.Tipe,
			"nominal":                    ds.Nominal,
			"validasi":                   ds.Validasi,
			"pengurus_id":                ds.PengurusId,
			"anak_yatim_id":              ds.AnakYatimId,
			"anak_yatim":                 ds.AnakYatim,
			"donatur_id":                 ds.DonaturId,
			"donatur":                    ds.Donatur,
			"created_at":                 ds.CreatedAt,
			"updated_at":                 ds.UpdatedAt,
		}

		if ds.File != nil {
			dsJSON["file"] = strings.Split(*ds.File, ";")
		} else {
			dsJSON["file"] = []string{}
		}
		danaSantunanAnakAsuhJSON = append(danaSantunanAnakAsuhJSON, dsJSON)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": danaSantunanAnakAsuhJSON,
	})
}
