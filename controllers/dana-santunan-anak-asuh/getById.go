package danasantunananakasuh

import (
	"server/config"
	"server/models"
	"strings"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm/clause"
)

func DanaSantunanAnakAsuhGetById(c *fiber.Ctx) error {

	id := c.Params("id")
	DanaSantunanAnakAsuh := models.DanaSantunanAnakAsuh{}

	// * Check if data exist
	if err := config.DB.Preload(clause.Associations).First(&DanaSantunanAnakAsuh, "dana_santunan_anak_asuh_id = ?", id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Data santunan anak asuh tidak ditemukan.",
		})
	}

	DanaSantunanAnakAsuhJSON := map[string]interface{}{
		"dana_santunan_anak_asuh_id": DanaSantunanAnakAsuh.DanaSantunanAnakAsuhId,
		"tanggal":                    DanaSantunanAnakAsuh.Tanggal,
		"nominal":                    DanaSantunanAnakAsuh.Nominal,
		"keterangan":                 DanaSantunanAnakAsuh.Keterangan,
		"validasi":                   DanaSantunanAnakAsuh.Validasi,
		"pengurus_id":                DanaSantunanAnakAsuh.PengurusId,
		"pengurus":                   DanaSantunanAnakAsuh.Pengurus,
		"donatur_id":                 DanaSantunanAnakAsuh.DonaturId,
		"donatur":                    DanaSantunanAnakAsuh.Donatur,
		"anak_yatim_id":              DanaSantunanAnakAsuh.AnakYatimId,
		"anak_yatim":                 DanaSantunanAnakAsuh.AnakYatim,
		"created_at":                 DanaSantunanAnakAsuh.CreatedAt,
		"updated_at":                 DanaSantunanAnakAsuh.UpdatedAt,
	}

	if DanaSantunanAnakAsuh.File != nil {
		DanaSantunanAnakAsuhJSON["file"] = strings.Split(*DanaSantunanAnakAsuh.File, ";")
	} else {
		DanaSantunanAnakAsuhJSON["file"] = []string{}
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": DanaSantunanAnakAsuhJSON,
	})
}
