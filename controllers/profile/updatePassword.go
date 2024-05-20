package profile

import (
	"server/config"
	"server/models"
	"server/pkg"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

func UpdatePassword(c *fiber.Ctx) error {

	form, err := c.MultipartForm()
	if err != nil {
		return err
	}

	password := form.Value["password"][0]
	oldPassword := form.Value["oldPassword"][0]

	// * Check if cookies valid
	cookie := c.Cookies("jwt")
	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(pkg.SECRET_KEY), nil
	})
	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "Tolong login terlebih dahulu.",
		})
	}

	// * Get values of Cookies
	claims := token.Claims.(*jwt.StandardClaims)
	var pengurusProfile models.Pengurus
	var donaturProfile models.Donatur

	if err := config.DB.Where("pengurus_id = ?", claims.Issuer).First(&pengurusProfile).Error; err == nil {
		// * Check if data exist
		pengurus := models.Pengurus{PengurusId: pengurusProfile.PengurusId}
		if err := config.DB.First(&pengurus).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "Data pengurus tidak ditemukan.",
			})
		}

		//* Compare old password
		err = bcrypt.CompareHashAndPassword([]byte(pengurus.Password), []byte(oldPassword))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Password lama tidak sesuai.",
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
	}

	if err := config.DB.Where("donatur_id = ?", claims.Issuer).First(&donaturProfile).Error; err == nil {
		// * Check if data exist
		donatur := models.Donatur{DonaturId: donaturProfile.DonaturId}
		if err := config.DB.First(&donatur).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "Data donatur tidak ditemukan.",
			})
		}
		//* Compare old password
		err = bcrypt.CompareHashAndPassword([]byte(donatur.Password), []byte(oldPassword))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Password lama tidak sesuai.",
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
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Password berhasil di edit",
	})
}
