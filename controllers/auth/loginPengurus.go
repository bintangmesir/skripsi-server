package auth

import (
	"server/config"
	"server/models"
	"server/pkg"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

func LoginPengurus(c *fiber.Ctx) error {
	form, err := c.MultipartForm()
	if err != nil {
		return err
	}

	email := form.Value["email"][0]
	password := form.Value["password"][0]

	// * Check if email exist
	existingUser := models.Pengurus{}
	if err := config.DB.Where("email = ?", email).First(&existingUser).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Email belum terdaftar.",
		})
	}
	pengurus := models.Pengurus{Email: email}
	config.DB.Find(&pengurus, "email = ?", email)

	// * Decode password
	if err := bcrypt.CompareHashAndPassword([]byte(pengurus.Password), []byte(password)); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "Password yang anda masukkan salah.",
		})
	}

	// * Generate Cookies
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    pengurus.PengurusId,
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	})
	token, err := claims.SignedString([]byte(pkg.SECRET_KEY))
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "Mohon maaf terjadi kesalahan pada server.",
		})
	}

	// * Set Cookies
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
		Secure:   true,
		SameSite: "Strict",
	}
	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "Selamat datang " + pengurus.Nama + " 👋.",
	})
}
