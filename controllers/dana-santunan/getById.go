package danasantunan

import (
	"server/config"
	"server/models"
	"strings"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm/clause"
)

func DanaSantunanGetById(c *fiber.Ctx) error {

	id := c.Params("id")
	DanaSantunan := models.DanaSantunan{}

	// * Check if data exist
	if err := config.DB.Preload(clause.Associations).First(&DanaSantunan, "dana_santunan_id= ?", id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Data dana santunan tidak ditemukan.",
		})
	}

	DanaSantunanJSON := map[string]interface{}{
		"dana_santunan_id": DanaSantunan.DanaSantunanId,
		"tanggal":          DanaSantunan.Tanggal,
		"nama":             DanaSantunan.Nama,
		"nominal":          DanaSantunan.Nominal,
		"keterangan":       DanaSantunan.Keterangan,
		"validasi":         DanaSantunan.Validasi,
		"pengurus_id":      DanaSantunan.PengurusId,
		"pengurus":         DanaSantunan.Pengurus,
		"created_at":       DanaSantunan.CreatedAt,
		"updated_at":       DanaSantunan.UpdatedAt,
	}

	if DanaSantunan.File != nil {
		DanaSantunanJSON["file"] = strings.Split(*DanaSantunan.File, ";")
	} else {
		DanaSantunanJSON["file"] = []string{}
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": DanaSantunanJSON,
	})
}
