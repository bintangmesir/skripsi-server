package auth

import (
	"server/config"
	"server/models"
	"strconv"
)

func HandleRegisterId() (userId string, err error) {
	donatur := []models.Donatur{}
	var count int64
	config.DB.Find(&donatur).Count(&count)

	var donaturId string

	if count == 0 {
		donaturId = "DNTR-1"
	} else {
		count = 0
		for {
			donaturId = "DNTR-" + strconv.Itoa(int(count+1))
			err := config.DB.Where("donatur_id = ?", donaturId).First(&donatur).Error
			if err != nil {
				break
			}
			count++
		}
	}
	return donaturId, nil
}
