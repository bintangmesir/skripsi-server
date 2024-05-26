package laporandanasantunananakasuh

import (
	"server/config"
	"server/models"
	"strconv"
)

func HandleLaporanDanaSantunanAnakAsuhId() (userId string, err error) {
	laporanDanaSantunanAnakAsuh := []models.LaporanDanaSantunanAnakAsuh{}
	var count int64
	config.DB.Find(&laporanDanaSantunanAnakAsuh).Count(&count)

	var laporanDanaSantunanAnakAsuhId string
	if count == 0 {
		laporanDanaSantunanAnakAsuhId = "LDSA-1"
	} else {
		count = 0
		for {
			laporanDanaSantunanAnakAsuhId = "LDSA" + "-" + strconv.Itoa(int(count+1))
			err := config.DB.Where("laporan_dana_santunan_anak_asuh_id = ?", laporanDanaSantunanAnakAsuhId).First(&laporanDanaSantunanAnakAsuh).Error
			if err != nil {
				break
			}
			count++
		}
	}
	return laporanDanaSantunanAnakAsuhId, nil
}
