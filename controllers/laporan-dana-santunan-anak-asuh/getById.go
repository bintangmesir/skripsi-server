package laporandanasantunananakasuh

import (
	"server/config"
	"server/models"
	"strings"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm/clause"
)

// ! Get Data Dana Zis
func LaporanDanaSantunanAnakAsuhGetById(c *fiber.Ctx) error {
	id := c.Params("id")
	// * Check if data exist
	laporanDanaSantunanAnakAsuh := models.LaporanDanaSantunanAnakAsuh{LaporanDanaSantunanAnakAsuhId: id}
	if err := config.DB.Preload(clause.Associations).First(&laporanDanaSantunanAnakAsuh).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Data laporan dana santunan anak asuh tidak ditemukan.",
		})
	}

	laporanDanaSantunanAnakAsuhJSON := map[string]interface{}{
		"laporan_dana_santunan_id": laporanDanaSantunanAnakAsuh.LaporanDanaSantunanAnakAsuhId,
		"judul":                    laporanDanaSantunanAnakAsuh.Judul,
		"saldo_awal":               laporanDanaSantunanAnakAsuh.SaldoAwal,
		"saldo_sisa":               laporanDanaSantunanAnakAsuh.SaldoSisa,
		"tanggal_tanda_tangan":     laporanDanaSantunanAnakAsuh.TanggalTandaTangan,
		"nama_tanda_tangan":        laporanDanaSantunanAnakAsuh.NamaTandaTangan,
		"tanda_tangan":             laporanDanaSantunanAnakAsuh.TandaTangan,
		"keterangan":               laporanDanaSantunanAnakAsuh.Keterangan,
		"validasi":                 laporanDanaSantunanAnakAsuh.Validasi,
		"dana_santunan_anak_asuh":  laporanDanaSantunanAnakAsuh.DanaSantunanAnakAsuh,
		"pengurus_id":              laporanDanaSantunanAnakAsuh.PengurusId,
		"donatur_id":               laporanDanaSantunanAnakAsuh.DonaturId,
		"created_at":               laporanDanaSantunanAnakAsuh.CreatedAt,
		"updated_at":               laporanDanaSantunanAnakAsuh.UpdatedAt,
	}

	if laporanDanaSantunanAnakAsuh.File != nil {
		laporanDanaSantunanAnakAsuhJSON["file"] = strings.Split(*laporanDanaSantunanAnakAsuh.File, ";")
	} else {
		laporanDanaSantunanAnakAsuhJSON["file"] = []string{}
	}

	if laporanDanaSantunanAnakAsuh.TandaTangan != nil {
		laporanDanaSantunanAnakAsuhJSON["tanda_tangan"] = strings.Split(*laporanDanaSantunanAnakAsuh.TandaTangan, ";")
	} else {
		laporanDanaSantunanAnakAsuhJSON["tanda_tangan"] = []string{}
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": laporanDanaSantunanAnakAsuhJSON,
	})
}
