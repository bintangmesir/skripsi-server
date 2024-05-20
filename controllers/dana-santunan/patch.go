package danasantunan

import (
	"log"
	"server/config"
	"server/models"
	"server/pkg"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

func DanaSantunanUpdate(c *fiber.Ctx) error {

	// * Check if data exist
	id := c.Params("id")
	danaSantunan := models.DanaSantunan{DanaSantunanId: id}
	if err := config.DB.First(&danaSantunan).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Data dana santunan tidak ditemukan.",
		})
	}

	form, err := c.MultipartForm()
	if err != nil {
		return err
	}

	tanggal := form.Value["tanggal"][0]
	nama := form.Value["nama"][0]
	nominal := form.Value["nominal"][0]
	keterangan := form.Value["keterangan"][0]
	validasi := form.Value["validasi"][0]
	file := form.File["file"]

	//* Handle tanggal
	tanggalParsedDate, err := time.Parse("Mon Jan 2 2006 15:04:05 GMT+0700 (Western Indonesia Time)", tanggal)
	if err != nil {
		log.Fatal(err)
	}

	nominalConverted, err := strconv.Atoi(nominal)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Mohon maaf terjadi kesalahan pada server.",
		})
	}

	newDanaSantunan := models.DanaSantunan{
		Tanggal:    tanggalParsedDate,
		Nama:       nama,
		Keterangan: &keterangan,
		Nominal:    nominalConverted,
		Validasi:   models.ValidationEnum(validasi),
		UpdatedAt:  time.Now(),
	}

	if len(file) > 0 {

		// * Handle File
		pkg.DeleteFile(danaSantunan.File, pkg.DIR_FILE_DANA_SANTUNAN)
		uploadedFileNames, err := pkg.UploadFile(file, pkg.DIR_FILE_DANA_SANTUNAN)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Mohon maaf terjadi kesalahan pada server.",
			})
		}
		newDanaSantunan.File = &uploadedFileNames
	}

	config.DB.Model(&danaSantunan).Updates(&newDanaSantunan)
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Data dana santunan anak asuh berhasil di edit",
	})
}
