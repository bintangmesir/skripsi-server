package danasantunan

import (
	"server/config"
	"server/models"
	"server/pkg"

	"github.com/gofiber/fiber/v2"
)

func DanaSantunanDelete(c *fiber.Ctx) error {

	// * Check if data exist
	id := c.Params("id")
	danaSantunan := models.DanaSantunan{DanaSantunanId: id}
	if err := config.DB.First(&danaSantunan).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Data dana santunan tidak ditemukan.",
		})
	}

	// * Handle If File	Doesn't Exist
	if err := pkg.DeleteFile(danaSantunan.File, pkg.DIR_FILE_DANA_SANTUNAN); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Mohon maaf terjadi kesalahan pada server.",
		})
	}

	config.DB.Delete(&danaSantunan)
	return c.SendStatus(fiber.StatusNoContent)
}
