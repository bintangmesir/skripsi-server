package anakyatim

import (
	"log"
	"server/config"
	"server/models"
	"server/pkg"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

func AnakYatimCreate(c *fiber.Ctx) error {
	form, err := c.MultipartForm()
	if err != nil {
		return err
	}

	nama := form.Value["nama"][0]
	status := form.Value["status"][0]
	jenisKelamin := form.Value["jenisKelamin"][0]
	tanggalLahir := form.Value["tanggalLahir"][0]
	pendidikan := form.Value["pendidikan"][0]
	pekerjaanOrangTua := form.Value["pekerjaanOrangTua"][0]
	penghasilanOrangTua := form.Value["penghasilanOrangTua"][0]
	tanggunganOrangTua := form.Value["tanggunganOrangTua"][0]
	nominalSantunan := form.Value["nominalSantunan"][0]
	kebutuhan := form.Value["kebutuhan"][0]
	deskripsi := form.Value["deskripsi"][0]
	foto := form.File["foto"]

	penghasilanOrangTuaConverted, err := strconv.Atoi(penghasilanOrangTua)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Mohon maaf terjadi kesalahan pada server.",
		})
	}

	nominalSantunanConverted, err := strconv.Atoi(nominalSantunan)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Mohon maaf terjadi kesalahan pada server.",
		})
	}

	tanggunganOrangTuaConverted, err := strconv.Atoi(tanggunganOrangTua)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Mohon maaf terjadi kesalahan pada server.",
		})
	}

	tanggalLahirParsedDate, err := time.Parse("Mon Jan 2 2006 15:04:05 GMT+0700 (Western Indonesia Time)", tanggalLahir)
	if err != nil {
		log.Fatal(err)
	}

	// * Handle Anak Yatim Id
	anakYatimId, err := HandleAnakYatimId()
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

	newAnakYatim := models.AnakYatim{
		AnakYatimId:         anakYatimId,
		Nama:                nama,
		Status:              status,
		JenisKelamin:        jenisKelamin,
		TanggalLahir:        tanggalLahirParsedDate,
		Pendidikan:          pendidikan,
		PekerjaanOrangTua:   pekerjaanOrangTua,
		PenghasilanOrangTua: penghasilanOrangTuaConverted,
		TanggunganOrangTua:  tanggunganOrangTuaConverted,
		NominalSantunan:     nominalSantunanConverted,
		Kebutuhan:           kebutuhan,
		Deskripsi:           deskripsi,
		StatusSantunan:      models.BelumMemiliki,
		PengurusId:          &id,
	}

	if len(foto) > 0 {
		// * Handle File
		uploadedFileNames, err := pkg.UploadFile(foto, pkg.DIR_IMG_ANAK_YATIM)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Mohon maaf terjadi kesalahan pada server.",
			})
		}
		newAnakYatim.Foto = &uploadedFileNames
	}

	config.DB.Create(&newAnakYatim)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Data anak yatim berhasil di tambah",
	})
}
