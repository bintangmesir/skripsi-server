package danasantunananakasuh

import (
	"server/config"
	"server/models"
	"strconv"
)

func HandleDanaSantunanAnakAsuhId() (userId string, err error) {
	danaSantunanAnakAsuh := []models.DanaSantunanAnakAsuh{}
	var count int64
	config.DB.Find(&danaSantunanAnakAsuh).Count(&count)

	var danaSantunanAnakAsuhId string
	if count == 0 {
		danaSantunanAnakAsuhId = "DSA-1"
	} else {
		count = 0
		for {
			danaSantunanAnakAsuhId = "DSA-" + strconv.Itoa(int(count+1))
			err := config.DB.Where("dana_santunan_anak_asuh_id = ?", danaSantunanAnakAsuhId).First(&danaSantunanAnakAsuh).Error
			if err != nil {
				break
			}
			count++
		}
	}
	return danaSantunanAnakAsuhId, nil
}
