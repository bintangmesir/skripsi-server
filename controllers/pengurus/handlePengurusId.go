package pengurus

import (
	"errors"
	"server/config"
	"server/models"
	"strconv"
)

func HandlePengurusId(jabatan string) (userId string, err error) {
	pengurus := []models.Pengurus{}
	var count int64
	config.DB.Find(&pengurus).Where("jabatan = ?", jabatan).Count(&count)

	var jabatanCode string
	switch jabatan {
	case "ADMIN":
		jabatanCode = "A"
	case "KETUA_DKM":
		jabatanCode = "KD"
	case "BENDAHARA":
		jabatanCode = "B"
	case "HUMAS":
		jabatanCode = "H"
	default:
		return "", errors.New("jabatan tidak diketahui")
	}

	var pengurusId string
	if count == 0 {
		pengurusId = "P" + jabatanCode + "-1"
	} else {
		count = 0
		for {
			pengurusId = "P" + jabatanCode + "-" + strconv.Itoa(int(count+1))
			err := config.DB.Where("pengurus_id = ?", pengurusId).First(&pengurus).Error
			if err != nil {
				break
			}
			count++
		}
	}
	return pengurusId, nil
}
