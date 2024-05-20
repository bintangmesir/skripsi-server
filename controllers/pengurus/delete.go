package pengurus

import (
	"server/config"
	"server/models"
	"server/pkg"

	"github.com/gofiber/fiber/v2"
)

func PengurusDelete(c *fiber.Ctx) error {
	// * Check if data exist
	id := c.Params("id")
	pengurus := models.Pengurus{PengurusId: id}
	if err := config.DB.First(&pengurus).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Data pengurus tidak ditemukan.",
		})
	}

	//* Handle active user
	idACtiveUser, err := pkg.GetUserActive(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Mohon login terlebih dahulu.",
		})
	}

	// * Check if pengurusId is creator
	if *pengurus.AdminId != idACtiveUser {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"message": "Mohon maaf anda tidak dilarang untuk melakukan perubahan pada data ini.",
		})
	}

	// * Handle if fotoProfile doesn't exist
	if pengurus.Foto != nil {
		if err := pkg.DeleteFile(pengurus.Foto, pkg.DIR_IMG_PENGURUS); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Mohon maaf terjadi kesalahan pada server.",
			})
		}
	}

	config.DB.Delete(&pengurus)
	return c.SendStatus(fiber.StatusNoContent)
}
