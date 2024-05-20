package pengurus

import (
	"server/config"
	"server/models"
	"server/pkg"
	"time"

	"github.com/gofiber/fiber/v2"
)

func PengurusPatch(c *fiber.Ctx) error {
	form, err := c.MultipartForm()
	if err != nil {
		return err
	}

	nama := form.Value["nama"][0]
	email := form.Value["email"][0]
	noHandphone := form.Value["noHandphone"][0]
	jabatan := form.Value["jabatan"][0]
	foto := form.File["foto"]

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

	// * Check if email exist
	var pengurusAll models.Pengurus
	config.DB.Where("email = ? AND pengurus_id != ?", email, id).First(&pengurusAll)
	if pengurusAll.PengurusId != "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Email telah digunakan.",
		})
	}

	newPengurus := models.Pengurus{
		Nama:        nama,
		Email:       email,
		NoHandphone: noHandphone,
		Jabatan:     models.RoleEnum(jabatan),
		AdminId:     &idACtiveUser,
		UpdatedAt:   time.Now(),
	}

	//* Handle pengurusId
	if pengurus.Jabatan != models.RoleEnum(jabatan) {
		pengurusId, err := HandlePengurusId(jabatan)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Mohon maaf terjadi kesalahan pada server.",
			})
		}
		newPengurus.PengurusId = pengurusId
	}

	// * Check if foto exist
	if len(foto) > 0 {

		// * Handle if foto exist
		if pengurus.Foto != nil {
			err := pkg.DeleteFile(pengurus.Foto, pkg.DIR_IMG_PENGURUS)
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"message": "Mohon maaf terjadi kesalahan pada server.",
				})
			}
		}

		// * Handle foto
		uploadedFileNames, err := pkg.UploadFile(foto, pkg.DIR_IMG_PENGURUS)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Mohon maaf terjadi kesalahan pada server.",
			})
		}
		newPengurus.Foto = &uploadedFileNames
	}

	config.DB.Model(&pengurus).Updates(&newPengurus)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Data pengurus berhasil di edit",
	})
}
