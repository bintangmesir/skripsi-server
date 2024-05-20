package auth

import (
	"server/config"
	"server/models"
	"time"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *fiber.Ctx) error {
	form, err := c.MultipartForm()
	if err != nil {
		return err
	}

	anakYatimId := form.Value["anakYatimId"][0]
	nama := form.Value["nama"][0]
	email := form.Value["email"][0]
	jenisKelamin := form.Value["jenisKelamin"][0]
	noHandphone := form.Value["noHandphone"][0]
	password := form.Value["password"][0]

	// * Check if anak yatim exist
	anakYatim := models.AnakYatim{AnakYatimId: anakYatimId}
	if err := config.DB.First(&anakYatim).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Data anak yatim tidak ditemukan.",
		})
	}
	// * Check if email exist
	existingUser := models.Donatur{}
	if err := config.DB.Limit(1).Where("email = ?", email).First(&existingUser).Error; err == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Email telah digunakan.",
		})
	}

	// * Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// * Handle register id
	registerId, err := HandleRegisterId()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Mohon maaf terjadi kesalahan pada server.",
		})
	}

	newDonatur := models.Donatur{
		DonaturId:    registerId,
		Nama:         nama,
		Email:        email,
		JenisKelamin: jenisKelamin,
		NoHandphone:  noHandphone,

		Password: string(hashedPassword),
		Foto:     nil,
	}

	newAnakYatim := models.AnakYatim{
		DonaturId: &registerId,
		UpdatedAt: time.Now(),
	}

	config.DB.Create(&newDonatur)
	config.DB.Model(&anakYatim).Updates(&newAnakYatim)
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Registrasi berhasil.",
	})
}
