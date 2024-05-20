package donatur

import (
	"server/config"
	"server/models"
	"server/pkg"
	"time"

	"github.com/gofiber/fiber/v2"
)

func DonaturUpdate(c *fiber.Ctx) error {

	// * Check if data exist
	id := c.Params("id")
	donatur := models.Donatur{DonaturId: id}
	if err := config.DB.First(&donatur).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Data donatur tidak ditemukan.",
		})
	}

	form, err := c.MultipartForm()
	if err != nil {
		return err
	}

	//* Handle user active
	idActiveUser, err := pkg.GetUserActive(c)
	if err != nil {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"message": "Anda tidak di izinkan untuk melakukan akses pada menu ini.",
		})
	}

	validasi := form.Value["validasi"][0]

	newDonatur := models.Donatur{
		Validasi:   models.RoleEnum(validasi),
		PengurusId: &idActiveUser,
		UpdatedAt:  time.Now(),
	}

	config.DB.Model(donatur).Updates(newDonatur)
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Data donatur berhasil di edit",
	})
}
