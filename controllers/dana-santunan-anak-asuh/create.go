package danasantunananakasuh

import (
	"log"
	"server/config"
	"server/models"
	"server/pkg"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

func DanaSantunanAnakAsuhCreate(c *fiber.Ctx) error {

	form, err := c.MultipartForm()
	if err != nil {
		return err
	}

	tanggal := form.Value["tanggal"][0]
	keterangan := form.Value["keterangan"][0]
	nominal := form.Value["nominal"][0]
	validasi := form.Value["validasi"][0]
	file := form.File["file"]
	anakYatimId := form.Value["anakYatimId"][0]
	donaturId := form.Value["donaturId"][0]

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

	//* Handle dana santunan id
	danaSantunanAnakAsuhId, err := HandleDanaSantunanAnakAsuhId()
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
	uploadedFileNames, err := pkg.UploadFile(file, pkg.DIR_FILE_DANA_SANTUNAN_ANAK_ASUH)
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

	newDanaSantunanAnakAsuh := models.DanaSantunanAnakAsuh{
		DanaSantunanAnakAsuhId: danaSantunanAnakAsuhId,
		Tanggal:                tanggalParsedDate,
		Nominal:                nominalConverted,
		Keterangan:             &keterangan,
		Validasi:               models.ValidationEnum(validasi),
		File:                   &uploadedFileNames,
		PengurusId:             &id,
		DonaturId:              &donaturId,
		AnakYatimId:            &anakYatimId,
	}

	config.DB.Create(&newDanaSantunanAnakAsuh)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Data dana santunan anak asuh berhasil di tambah",
	})
}
