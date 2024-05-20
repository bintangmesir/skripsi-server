package anakyatim

import (
	"server/config"
	"server/models"
	"server/pkg"

	"github.com/gofiber/fiber/v2"
)

func AnakYatimDelete(c *fiber.Ctx) error {
	// * Check if data exist
	id := c.Params("id")
	anakYatim := models.AnakYatim{AnakYatimId: id}
	if err := config.DB.First(&anakYatim).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Data anak yatim tidak ditemukan.",
		})
	}

	// * Handle If File	Doesn't Exist
	if anakYatim.Foto != nil {
		if err := pkg.DeleteFile(anakYatim.Foto, pkg.DIR_IMG_ANAK_YATIM); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Mohon maaf terjadi kesalahan pada server.",
			})
		}
	}
	config.DB.Delete(&anakYatim)
	return c.SendStatus(fiber.StatusNoContent)
}
