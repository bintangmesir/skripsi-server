package pengurus

import (
	"server/config"
	"server/models"
	"strings"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm/clause"
)

func PengurusGet(c *fiber.Ctx) error {

	// * Check if data exist
	pengurus := []models.Pengurus{}
	if err := config.DB.Preload(clause.Associations).Find(&pengurus).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Data pengurus tidak ditemukan.",
		})
	}

	var pengurusJSON []map[string]interface{}
	for _, p := range pengurus {
		pJSON := map[string]interface{}{
			"pengurus_id":  p.PengurusId,
			"nama":         p.Nama,
			"email":        p.Email,
			"no_handphone": p.NoHandphone,
			"jabatan":      p.Jabatan,
			"admin_id":     p.AdminId,
			"pengurus":     p.Pengurus,
			"created_at":   p.CreatedAt,
			"updated_at":   p.UpdatedAt,
		}
		if p.Foto != nil {
			pJSON["foto"] = strings.Split(*p.Foto, ";")
		} else {
			pJSON["foto"] = []string{}
		}
		pengurusJSON = append(pengurusJSON, pJSON)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": pengurusJSON,
	})
}
