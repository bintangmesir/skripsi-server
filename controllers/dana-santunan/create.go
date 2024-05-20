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

func DanaSantunanCreate(c *fiber.Ctx) error {

	form, err := c.MultipartForm()
	if err != nil {
		return err
	}

	nama := form.Value["nama"][0]
	tanggal := form.Value["tanggal"][0]
	validasi := form.Value["validasi"][0]
	keterangan := form.Value["keterangan"][0]
	nominal := form.Value["nominal"][0]
	file := form.File["file"]

	//* Handle dana santunan id
	danaSantunanId, err := HandleDanaSantunanId()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Mohon maaf terjadi kesalahan pada server.",
		})
	}

	//* Handle tanggal
	tanggalParsedDate, err := time.Parse("Mon Jan 2 2006 15:04:05 GMT+0700 (Western Indonesia Time)", tanggal)
	if err != nil {
		log.Fatal(err)
	}

	//* Handle nominal
	nominalConverted, err := strconv.Atoi(nominal)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Mohon maaf terjadi kesalahan pada server.",
		})
	}

	// * Handle File
	uploadedFileNames, err := pkg.UploadFile(file, pkg.DIR_FILE_DANA_SANTUNAN)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Mohon maaf terjadi kesalahan pada server.",
		})
	}

	id, err := pkg.GetUserActive(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Mohon login terlebih dahulu.",
		})
	}

	newDanaSantunan := models.DanaSantunan{
		DanaSantunanId: danaSantunanId,
		Tanggal:        tanggalParsedDate,
		Nama:           nama,
		Nominal:        nominalConverted,
		Keterangan:     &keterangan,
		Validasi:       models.ValidationEnum(validasi),
		File:           &uploadedFileNames,
		PengurusId:     &id,
	}

	config.DB.Create(&newDanaSantunan)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Data dana zis berhasil di tambah",
	})
}
