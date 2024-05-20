package donatur

import (
	"server/config"
	"server/models"
	"server/pkg"

	"github.com/gofiber/fiber/v2"
)

func DonaturDelete(c *fiber.Ctx) error {
	// * Check if data exist
	id := c.Params("id")
	donatur := models.Donatur{DonaturId: id}
	if err := config.DB.First(&donatur).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Data donatur tidak ditemukan.",
		})
	}

	// * Handle if fotoProfile	Doesn't Exist
	if donatur.Foto != nil {
		if err := pkg.DeleteFile(donatur.Foto, pkg.DIR_IMG_DONATUR); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Mohon maaf terjadi kesalahan pada server.",
			})
		}
	}

	config.DB.Delete(&donatur)
	return c.SendStatus(fiber.StatusNoContent)
}
