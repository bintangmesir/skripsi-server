package pengurus

import (
	"server/config"
	"server/models"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func PengurusGetById(c *fiber.Ctx) error {
	id := c.Params("id")
	pengurus := models.Pengurus{}

	// * Check if data exist
	if err := config.DB.First(&pengurus, "pengurus_id = ?", id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Data pengurus tidak ditemukan.",
		})
	}

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
