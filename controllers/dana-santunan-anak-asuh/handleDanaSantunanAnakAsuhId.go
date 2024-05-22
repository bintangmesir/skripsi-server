package danasantunananakasuh

import (
	"time"
)

func HandleDanaSantunanAnakAsuhId() (userId string, err error) {
	danaSantunanAnakAsuhId := "DSA-" + time.Now().Format("2006-01-02-15-04-05")

	return danaSantunanAnakAsuhId, nil
}
