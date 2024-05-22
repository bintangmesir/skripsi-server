package danasantunan

import (
	"time"
)

func HandleDanaSantunanId() (userId string, err error) {
	danaSantunanId := "DSI-" + time.Now().Format("2006-01-02-15-04-05")
	return danaSantunanId, nil
}
