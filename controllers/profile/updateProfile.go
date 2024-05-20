package profile

import (
	"server/config"
	"server/models"
	"server/pkg"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

func UpdateProfile(c *fiber.Ctx) error {

	form, err := c.MultipartForm()
	if err != nil {
		return err
	}

	nama := form.Value["nama"][0]
	email := form.Value["email"][0]
	noHandphone := form.Value["noHandphone"][0]
	foto := form.File["foto"]

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

		// * Check if email exist
		var pengurusAll models.Pengurus
		config.DB.Where("email = ? AND pengurus_id != ?", email, pengurusProfile.PengurusId).First(&pengurusAll)
		if pengurusAll.PengurusId != "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Email telah digunakan.",
			})
		}

		newPengurus := models.Pengurus{
			Nama:        nama,
			Email:       email,
			NoHandphone: noHandphone,
			UpdatedAt:   time.Now(),
		}

		// * Check if fotoProfile exist
		if len(foto) > 0 {
			if pengurus.Foto != nil {
				err := pkg.DeleteFile(pengurus.Foto, pkg.DIR_IMG_PENGURUS)
				if err != nil {
					return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
						"message": "Mohon maaf terjadi kesalahan pada server.",
					})
				}
			}

			// * Handle fotoProfile
			uploadedFileNames, err := pkg.UploadFile(foto, pkg.DIR_IMG_PENGURUS)
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"message": "Mohon maaf terjadi kesalahan pada server.",
				})
			}
			newPengurus.Foto = &uploadedFileNames
		}

		config.DB.Model(&pengurus).Updates(&newPengurus)
	}

	if err := config.DB.Where("donatur_id = ?", claims.Issuer).First(&donaturProfile).Error; err == nil {

		jenisKelamin := form.Value["jenisKelamin"][0]
		namaJalan := form.Value["namaJalan"][0]
		rt := form.Value["rt"][0]
		rw := form.Value["rw"][0]
		kelurahan := form.Value["kelurahan"][0]
		kecamatan := form.Value["kecamatan"][0]
		kota := form.Value["kota"][0]
		provinsi := form.Value["provinsi"][0]
		kodePos := form.Value["kodePos"][0]

		// * Check if data exist
		donatur := models.Donatur{DonaturId: donaturProfile.DonaturId}
		if err := config.DB.First(&donatur).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "Data donatur tidak ditemukan.",
			})
		}

		// * Check if email exist
		var donaturAll models.Donatur
		config.DB.Where("email = ? AND donatur_id != ?", email, donaturProfile.DonaturId).First(&donaturAll)
		if donaturAll.DonaturId != "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Email telah digunakan.",
			})
		}

		newDonatur := models.Donatur{
			Nama:         nama,
			Email:        email,
			JenisKelamin: jenisKelamin,
			NamaJalan:    namaJalan,
			Rt:           rt,
			Rw:           rw,
			Kelurahan:    kelurahan,
			Kecamatan:    kecamatan,
			Kota:         kota,
			Provinsi:     provinsi,
			KodePos:      kodePos,
			NoHandphone:  noHandphone,
			UpdatedAt:    time.Now(),
		}

		// * Check if fotoProfile exist
		if len(foto) > 0 {

			// * Handle If File	Doesn't Exist
			if donatur.Foto != nil {
				println("test")
				err := pkg.DeleteFile(donatur.Foto, pkg.DIR_IMG_DONATUR)
				if err != nil {
					return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
						"message": "Mohon maaf terjadi kesalahan pada server.",
					})
				}
			}
			// * Handle File
			uploadedFileNames, err := pkg.UploadFile(foto, pkg.DIR_IMG_DONATUR)
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"message": "Mohon maaf terjadi kesalahan pada server.",
				})
			}
			newDonatur.Foto = &uploadedFileNames
		}

		config.DB.Model(&donatur).Updates(&newDonatur)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Data profile berhasil di edit",
	})
}
