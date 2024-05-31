package laporandanasantunananakasuh

import (
	"server/config"
	"server/models"
	"strings"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm/clause"
)

// ! Get Data Dana Zis
func LaporanDanaSantunanAnakAsuhGet(c *fiber.Ctx) error {

	// * Check if data exist
	laporanDanaSantunanAnakAsuh := []models.LaporanDanaSantunanAnakAsuh{}
	if err := config.DB.Preload(clause.Associations).Order("created_at DESC").Find(&laporanDanaSantunanAnakAsuh).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Data laporan dana santunan tidak ditemukan.",
		})
	}

	var laporanDanaSantunanAnakAsuhJSON []map[string]interface{}
	for _, ds := range laporanDanaSantunanAnakAsuh {

		dsJSON := map[string]interface{}{
			"laporan_dana_santunan_anak_asuh_id": ds.LaporanDanaSantunanAnakAsuhId,
			"judul":                              ds.Judul,
			"saldo_awal":                         ds.SaldoAwal,
			"saldo_sisa":                         ds.SaldoSisa,
			"tanggal_tanda_tangan":               ds.TanggalTandaTangan,
			"nama_tanda_tangan":                  ds.NamaTandaTangan,
			"tanda_tangan":                       ds.TandaTangan,
			"keterangan":                         ds.Keterangan,
			"validasi":                           ds.Validasi,
			"dana_santunan_anak_asuh":            ds.DanaSantunanAnakAsuh,
			"pengurus_id":                        ds.PengurusId,
			"donatur_id":                         ds.DonaturId,
			"created_at":                         ds.CreatedAt,
			"updated_at":                         ds.UpdatedAt,
		}

		if ds.File != nil {
			dsJSON["file"] = strings.Split(*ds.File, ";")
		} else {
			dsJSON["file"] = []string{}
		}

		if ds.TandaTangan != nil {
			dsJSON["tanda_tangan"] = strings.Split(*ds.TandaTangan, ";")
		} else {
			dsJSON["tanda_tangan"] = []string{}
		}
		laporanDanaSantunanAnakAsuhJSON = append(laporanDanaSantunanAnakAsuhJSON, dsJSON)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": laporanDanaSantunanAnakAsuhJSON,
	})
}
