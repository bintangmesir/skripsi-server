package santunan

import (
	"server/config"
	"server/models"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm/clause"
)

func NonAktif(c *fiber.Ctx) error {
	// * Check if data exist
	id := c.Params("id")
	anakYatim := models.AnakYatim{AnakYatimId: id}
	if err := config.DB.Preload(clause.Associations).First(&anakYatim).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Data anak yatim tidak ditemukan.",
		})
	}

	newAnakYatim := models.AnakYatim{
		StatusSantunan: models.NonAktif,
		UpdatedAt:      time.Now(),
	}

	config.DB.Model(&anakYatim).Updates(&newAnakYatim)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Pemberhentian program sementara santunan anak asuh berhasil.",
	})
}
