package santunan

import (
	"server/config"
	"server/models"
	"time"

	"github.com/gofiber/fiber/v2"
)

func Berhenti(c *fiber.Ctx) error {
	// * Check if data exist
	id := c.Params("id")
	anakYatim := models.AnakYatim{AnakYatimId: id}
	if err := config.DB.First(&anakYatim).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Data anak yatim tidak ditemukan.",
		})
	}

	newAnakYatim := models.AnakYatim{
		StatusSantunan: models.BelumMemiliki,
		UpdatedAt:      time.Now(),
	}

	config.DB.Model(&anakYatim).Updates(&newAnakYatim)
	config.DB.Model(&anakYatim).Where("anak_yatim_id = ?", id).Update("donatur_id", nil)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Pemberhentian program santunan asuh berhasil.",
	})
}
