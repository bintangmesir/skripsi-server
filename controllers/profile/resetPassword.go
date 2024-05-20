package profile

import (
	"server/config"
	"server/models"
	"time"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func ResetPassword(c *fiber.Ctx) error {

	form, err := c.MultipartForm()
	if err != nil {
		return err
	}

	// * Check if data exist
	id := c.Params("id")
	password := form.Value["password"][0]

	var pengurusProfile models.Pengurus
	var donaturProfile models.Donatur

	if err := config.DB.Where("pengurus_id = ?", id).First(&pengurusProfile).Error; err == nil {
		// * Check if data exist
		pengurus := models.Pengurus{PengurusId: pengurusProfile.PengurusId}
		if err := config.DB.First(&pengurus).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "Data pengurus tidak ditemukan.",
			})
		}

		// * Hash password
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}

		newPengurus := models.Pengurus{
			Password:  string(hashedPassword),
			UpdatedAt: time.Now(),
		}

		config.DB.Model(pengurus).Updates(newPengurus)
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Password berhasil di edit",
		})
	}

	if err := config.DB.Where("donatur_id = ?", id).First(&donaturProfile).Error; err == nil {
		// * Check if data exist
		donatur := models.Donatur{DonaturId: donaturProfile.DonaturId}
		if err := config.DB.First(&donatur).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "Data donatur tidak ditemukan.",
			})
		}

		// * Hash password
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}

		newDonatur := models.Donatur{
			Password:  string(hashedPassword),
			UpdatedAt: time.Now(),
		}

		config.DB.Model(&donatur).Updates(&newDonatur)
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Password berhasil di edit",
		})
	}
	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
		"message": "Data akun tidak ditemukan.",
	})
}
