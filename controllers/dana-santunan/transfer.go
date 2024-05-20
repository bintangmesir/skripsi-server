package danasantunan

import (
	"server/config"
	"server/models"
	"server/pkg"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func DanaSantunanCreateTransfer(c *fiber.Ctx) error {

	form, err := c.MultipartForm()
	if err != nil {
		return err
	}

	nomorTransferMidtrans := form.Value["danaSantunanId"][0]
	nama := form.Value["nama"][0]
	nominal := form.Value["nominal"][0]
	file := form.File["file"]

	nominalConverted, err := strconv.Atoi(nominal)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Mohon maaf terjadi kesalahan pada server.",
		})
	}

	newDanaTransfer := models.DanaSantunan{
		Nama:     nama,
		Nominal:  nominalConverted,
		Validasi: models.Pending,
	}

	if nomorTransferMidtrans == "" {
		//* Handle nomor transfer
		nomorTransfer, err := HandleDanaSantunanId()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Mohon maaf terjadi kesalahan pada server.",
			})
		}
		newDanaTransfer.DanaSantunanId = nomorTransfer
	} else {
		newDanaTransfer.DanaSantunanId = nomorTransferMidtrans
		// * Check if data anak yatim exist
		nomorTransfer := models.DanaSantunan{DanaSantunanId: nomorTransferMidtrans}
		if err := config.DB.First(&nomorTransfer).Error; err == nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Data nomor transfer sudah terdaftar pada sistem.",
			})
		}
	}

	// * Handle File
	uploadedFileNames, err := pkg.UploadFile(file, pkg.DIR_FILE_DANA_SANTUNAN)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Mohon maaf terjadi kesalahan pada server.",
		})
	}

	newDanaTransfer.File = &uploadedFileNames

	config.DB.Create(&newDanaTransfer)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Data dana transfer berhasil di tambah",
	})
}
