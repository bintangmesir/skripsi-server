package anakyatim

import (
	"server/config"
	"server/models"
	"strconv"
)

func HandleAnakYatimId() (userId string, err error) {
	anakYatim := []models.AnakYatim{}
	var count int64
	config.DB.Find(&anakYatim).Count(&count)

	var anakYatimId string
	if count == 0 {
		anakYatimId = "ANKYTM-1"
	} else {
		count = 0
		for {
			anakYatimId = "ANKYTM" + "-" + strconv.Itoa(int(count+1))
			err := config.DB.Where("anak_yatim_id = ?", anakYatimId).First(&anakYatim).Error
			if err != nil {
				break
			}
			count++
		}
	}
	return anakYatimId, nil
}
