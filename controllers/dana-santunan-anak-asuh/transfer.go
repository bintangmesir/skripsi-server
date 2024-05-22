package danasantunananakasuh

import (
	"server/config"
	"server/models"
	"server/pkg"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func DanaSantunanAnakAsuhCreateTransfer(c *fiber.Ctx) error {

	form, err := c.MultipartForm()
	if err != nil {
		return err
	}

	nomorTransferMidtrans := form.Value["danaSantunanAnakAsuhId"][0]
	anakYatimId := form.Value["anakYatimId"][0]
	donaturId := form.Value["donaturId"][0]
	nominal := form.Value["nominal"][0]
	file := form.File["file"]

	donatur := models.Donatur{}
	anakYatim := models.AnakYatim{}

	// * Check if data exist
	if err := config.DB.First(&donatur, "donatur_id = ?", donaturId).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Data donatur tidak ditemukan.",
		})
	}

	// * Check if data exist
	if err := config.DB.First(&anakYatim, "anak_yatim_id = ?", anakYatimId).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Data anak yatim tidak ditemukan.",
		})
	}

	nominalConverted, err := strconv.Atoi(nominal)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Mohon maaf terjadi kesalahan pada server.",
		})
	}

	newDanaTransfer := models.DanaSantunanAnakAsuh{
		Nominal:     nominalConverted,
		Validasi:    models.Pending,
		DonaturId:   &donaturId,
		AnakYatimId: &anakYatimId,
	}

	newAnakYatim := models.AnakYatim{
		StatusSantunan: models.Aktif,
		DonaturId:      &donaturId,
	}

	if nomorTransferMidtrans == "" {
		//* Handle nomor transfer
		nomorTransfer, err := HandleDanaSantunanAnakAsuhId()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Mohon maaf terjadi kesalahan pada server.",
			})
		}
		newDanaTransfer.DanaSantunanAnakAsuhId = nomorTransfer
	} else {
		newDanaTransfer.DanaSantunanAnakAsuhId = nomorTransferMidtrans
		// * Check if data anak yatim exist
		nomorTransfer := models.DanaSantunanAnakAsuh{DanaSantunanAnakAsuhId: nomorTransferMidtrans}
		if err := config.DB.First(&nomorTransfer).Error; err == nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Data nomor transfer sudah terdaftar pada sistem.",
			})
		}
	}

	// * Handle File
	uploadedFileNames, err := pkg.UploadFile(file, pkg.DIR_FILE_DANA_SANTUNAN_ANAK_ASUH)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Mohon maaf terjadi kesalahan pada server.",
		})
	}

	newDanaTransfer.File = &uploadedFileNames

	config.DB.Create(&newDanaTransfer)
	config.DB.Model(&anakYatim).Updates(&newAnakYatim)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Data dana transfer berhasil di tambah",
	})
}
