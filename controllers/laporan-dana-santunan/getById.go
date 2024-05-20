package laporandanasantunan

import (
	"server/config"
	"server/models"
	"strings"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm/clause"
)

// ! Get Data Dana Zis
func LaporanDanaSantunanGetById(c *fiber.Ctx) error {
	id := c.Params("id")
	// * Check if data exist
	laporanDanaSantunan := models.LaporanDanaSantunan{LaporanDanaSantunanId: id}
	if err := config.DB.Preload(clause.Associations).First(&laporanDanaSantunan).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Data laporan dana santunan tidak ditemukan.",
		})
	}

	laporanDanaSantunanJSON := map[string]interface{}{
		"laporan_dana_santunan_id": laporanDanaSantunan.LaporanDanaSantunanId,
		"judul":                    laporanDanaSantunan.Judul,
		"saldo_awal":               laporanDanaSantunan.SaldoAwal,
		"saldo_sisa":               laporanDanaSantunan.SaldoSisa,
		"tanggal_tanda_tangan":     laporanDanaSantunan.TanggalTandaTangan,
		"nama_tanda_tangan":        laporanDanaSantunan.NamaTandaTangan,
		"tanda_tangan":             laporanDanaSantunan.TandaTangan,
		"keterangan":               laporanDanaSantunan.Keterangan,
		"validasi":                 laporanDanaSantunan.Validasi,
		"dana_santunan":            laporanDanaSantunan.DanaSantunan,
		"pengurus_id":              laporanDanaSantunan.PengurusId,
		"created_at":               laporanDanaSantunan.CreatedAt,
		"updated_at":               laporanDanaSantunan.UpdatedAt,
	}

	if laporanDanaSantunan.File != nil {
		laporanDanaSantunanJSON["file"] = strings.Split(*laporanDanaSantunan.File, ";")
	} else {
		laporanDanaSantunanJSON["file"] = []string{}
	}

	if laporanDanaSantunan.TandaTangan != nil {
		laporanDanaSantunanJSON["tanda_tangan"] = strings.Split(*laporanDanaSantunan.TandaTangan, ";")
	} else {
		laporanDanaSantunanJSON["tanda_tangan"] = []string{}
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": laporanDanaSantunanJSON,
	})
}
