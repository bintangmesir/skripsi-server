package laporandanasantunan

import (
	"server/config"
	"server/models"
	"strconv"
)

func HandleLaporanDanaSantunanId() (userId string, err error) {
	laporanDanaSantunan := []models.LaporanDanaSantunan{}
	var count int64
	config.DB.Find(&laporanDanaSantunan).Count(&count)

	var laporanDanaSantunanId string
	if count == 0 {
		laporanDanaSantunanId = "LDSI-1"
	} else {
		count = 0
		for {
			laporanDanaSantunanId = "LDSI" + "-" + strconv.Itoa(int(count+1))
			err := config.DB.Where("laporan_dana_santunan_id = ?", laporanDanaSantunanId).First(&laporanDanaSantunan).Error
			if err != nil {
				break
			}
			count++
		}
	}
	return laporanDanaSantunanId, nil
}
