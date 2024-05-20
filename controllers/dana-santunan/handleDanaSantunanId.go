package danasantunan

import (
	"server/config"
	"server/models"
	"strconv"
)

func HandleDanaSantunanId() (userId string, err error) {
	danaSantunan := []models.DanaSantunan{}
	var count int64
	config.DB.Find(&danaSantunan).Count(&count)

	var danaSantunanId string
	if count == 0 {
		danaSantunanId = "DSI-1"
	} else {
		count = 0
		for {
			danaSantunanId = "DSI-" + strconv.Itoa(int(count+1))
			err := config.DB.Where("dana_santunan_id = ?", danaSantunanId).First(&danaSantunan).Error
			if err != nil {
				break
			}
			count++
		}
	}
	return danaSantunanId, nil
}
