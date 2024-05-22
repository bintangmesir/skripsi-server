package anakyatim

import (
	"log"
	"server/config"
	"server/models"
	"server/pkg"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm/clause"
)

func AnakYatimUpdate(c *fiber.Ctx) error {
	// * Check if data exist
	id := c.Params("id")
	anakYatim := models.AnakYatim{AnakYatimId: id}
	if err := config.DB.Preload(clause.Associations).First(&anakYatim).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Data anak yatim tidak ditemukan.",
		})
	}

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
	kebutuhan := form.Value["kebutuhan"][0]
	deskripsi := form.Value["deskripsi"][0]
	nominalSantunan := form.Value["nominalSantunan"][0]
	statusSantunan := form.Value["statusSantunan"][0]

	foto := form.File["foto"]

	penghasilanOrangTuaConverted, err := strconv.Atoi(penghasilanOrangTua)
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

	nominalSantunanConverted, err := strconv.Atoi(nominalSantunan)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Mohon maaf terjadi kesalahan pada server.",
		})
	}

	parsedDate, err := time.Parse("Mon Jan 2 2006 15:04:05 GMT+0700 (Western Indonesia Time)", tanggalLahir)
	if err != nil {
		log.Fatal(err)
	}

	newAnakYatim := models.AnakYatim{
		Nama:                nama,
		Status:              status,
		JenisKelamin:        jenisKelamin,
		TanggalLahir:        parsedDate,
		Pendidikan:          pendidikan,
		PekerjaanOrangTua:   pekerjaanOrangTua,
		PenghasilanOrangTua: penghasilanOrangTuaConverted,
		TanggunganOrangTua:  tanggunganOrangTuaConverted,
		Kebutuhan:           kebutuhan,
		Deskripsi:           deskripsi,
		StatusSantunan:      models.StatusEnum(statusSantunan),
		NominalSantunan:     nominalSantunanConverted,
		UpdatedAt:           time.Now(),
	}

	if len(foto) > 0 {

		// * Handle If File	Doesn't Exist
		if anakYatim.Foto != nil {
			err := pkg.DeleteFile(anakYatim.Foto, pkg.DIR_IMG_ANAK_YATIM)
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"message": "Mohon maaf terjadi kesalahan pada server.",
				})
			}
		}

		// * Handle File
		uploadedFileNames, err := pkg.UploadFile(foto, pkg.DIR_IMG_ANAK_YATIM)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Mohon maaf terjadi kesalahan pada server.",
			})
		}
		newAnakYatim.Foto = &uploadedFileNames
	}

	config.DB.Model(&anakYatim).Updates(&newAnakYatim)

	if statusSantunan == "BELUM_MEMILIKI" {
		config.DB.Model(&anakYatim).Where("anak_yatim_id = ?", id).Update("DonaturId", nil)
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Data anak yatim berhasil di edit",
	})
}
