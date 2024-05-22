package danasantunan

import (
	"server/config"
	"server/models"
	"strings"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm/clause"
)

// ! Get Data Dana Zis
func DanaSantunanGet(c *fiber.Ctx) error {

	// * Check if data exist
	danaSantunan := []models.DanaSantunan{}
	if err := config.DB.Preload(clause.Associations).Order("created_at DESC").Find(&danaSantunan).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Data dana santunan tidak ditemukan.",
		})
	}

	var danaSantunanJSON []map[string]interface{}
	for _, ds := range danaSantunan {

		dsJSON := map[string]interface{}{
			"dana_santunan_id":         ds.DanaSantunanId,
			"index":                    ds.Index,
			"tanggal":                  ds.Tanggal,
			"nama":                     ds.Nama,
			"keterangan":               ds.Keterangan,
			"tipe":                     ds.Tipe,
			"nominal":                  ds.Nominal,
			"validasi":                 ds.Validasi,
			"pengurus_id":              ds.PengurusId,
			"laporan_dana_santunan_id": ds.LaporanDanaSantunanId,
			"created_at":               ds.CreatedAt,
			"updated_at":               ds.UpdatedAt,
		}

		if ds.File != nil {
			dsJSON["file"] = strings.Split(*ds.File, ";")
		} else {
			dsJSON["file"] = []string{}
		}
		danaSantunanJSON = append(danaSantunanJSON, dsJSON)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": danaSantunanJSON,
	})
}
