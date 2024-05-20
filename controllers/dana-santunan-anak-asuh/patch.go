package danasantunananakasuh

import (
	"server/config"
	"server/models"
	"server/pkg"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

func DanaSantunanAnakAsuhUpdate(c *fiber.Ctx) error {

	// * Check if data exist
	id := c.Params("id")
	danaSantunanAnakAsuh := models.DanaSantunanAnakAsuh{DanaSantunanAnakAsuhId: id}
	if err := config.DB.First(&danaSantunanAnakAsuh).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Data dana santunan anak asuh tidak ditemukan.",
		})
	}

	anakYatim := models.AnakYatim{}
	if err := config.DB.First(&anakYatim, "anak_yatim_id = ?", danaSantunanAnakAsuh.AnakYatimId).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Data anak yatim tidak ditemukan.",
		})
	}

	form, err := c.MultipartForm()
	if err != nil {
		return err
	}

	nominal := form.Value["nominal"][0]
	keterangan := form.Value["keterangan"][0]
	validasi := form.Value["validasi"][0]
	file := form.File["file"]

	nominalConverted, err := strconv.Atoi(nominal)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Mohon maaf terjadi kesalahan pada server.",
		})
	}

	newDanaSantunanAnakAsuh := models.DanaSantunanAnakAsuh{
		Keterangan: &keterangan,
		Nominal:    nominalConverted,
		Validasi:   models.ValidationEnum(validasi),
		UpdatedAt:  time.Now(),
	}

	if len(file) > 0 {

		// * Handle File
		pkg.DeleteFile(danaSantunanAnakAsuh.File, pkg.DIR_FILE_DANA_SANTUNAN_ANAK_ASUH)
		uploadedFileNames, err := pkg.UploadFile(file, pkg.DIR_FILE_DANA_SANTUNAN_ANAK_ASUH)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Mohon maaf terjadi kesalahan pada server.",
			})
		}
		newDanaSantunanAnakAsuh.File = &uploadedFileNames
	}

	config.DB.Model(danaSantunanAnakAsuh).Updates(newDanaSantunanAnakAsuh)

	if validasi == "DIVERIFIKASI" {
		newAnakYatim := models.AnakYatim{
			StatusSantunan:  models.SudahMemiliki,
			NominalSantunan: anakYatim.NominalSantunan + nominalConverted,
		}

		config.DB.Model(&anakYatim).Updates(&newAnakYatim)
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Data dana santunan anak asuh berhasil di edit",
	})
}
