package danasantunananakasuh

import (
	"server/config"
	"server/models"
	"server/pkg"

	"github.com/gofiber/fiber/v2"
)

func DanaSantunanAnakAsuhDelete(c *fiber.Ctx) error {

	// * Check if data exist
	id := c.Params("id")
	danaSantunanAnakAsuh := models.DanaSantunanAnakAsuh{DanaSantunanAnakAsuhId: id}
	if err := config.DB.First(&danaSantunanAnakAsuh).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Data dana santunan anak asuh tidak ditemukan.",
		})
	}

	// * Handle If File	Doesn't Exist
	if err := pkg.DeleteFile(danaSantunanAnakAsuh.File, pkg.DIR_FILE_DANA_SANTUNAN_ANAK_ASUH); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Mohon maaf terjadi kesalahan pada server.",
		})
	}

	config.DB.Delete(&danaSantunanAnakAsuh)
	return c.SendStatus(fiber.StatusNoContent)
}
