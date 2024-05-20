package pengurus

import (
	"server/config"
	"server/models"
	"server/pkg"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func PengurusPost(c *fiber.Ctx) error {
	form, err := c.MultipartForm()
	if err != nil {
		return err
	}

	nama := form.Value["nama"][0]
	email := form.Value["email"][0]
	noHandphone := form.Value["noHandphone"][0]
	password := form.Value["password"][0]
	jabatan := form.Value["jabatan"][0]
	foto := form.File["foto"]

	// * Check if email exist
	existingUser := models.Pengurus{}
	if err := config.DB.Where("email = ?", email).First(&existingUser).Error; err == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Email telah digunakan.",
		})
	}

	//* Handle user active
	idActiveUser, err := pkg.GetUserActive(c)
	if err != nil {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"message": "Anda tidak di izinkan untuk melakukan akses pada menu ini.",
		})
	}

	//* Handle pengurusId
	pengurusId, err := HandlePengurusId(jabatan)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Mohon maaf terjadi kesalahan pada server.",
		})
	}

	// * Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	newPengurus := models.Pengurus{
		PengurusId:  pengurusId,
		Nama:        nama,
		Email:       email,
		NoHandphone: noHandphone,
		Password:    string(hashedPassword),
		Jabatan:     models.RoleEnum(jabatan),
		AdminId:     &idActiveUser,
	}

	// * Check if foto exist
	if len(foto) > 0 {

		// * Handle foto
		uploadedFileNames, err := pkg.UploadFile(foto, pkg.DIR_IMG_PENGURUS)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Mohon maaf terjadi kesalahan pada server.",
			})
		}
		newPengurus.Foto = &uploadedFileNames
	}

	config.DB.Create(&newPengurus)
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Data pengurus berhasil di tambah",
	})
}
