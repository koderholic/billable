package utils

import (
	"time"
	// "fmt"
)

func GetDifferenceInSeconds(cTimeCur string, cTimePast string) (seconds int) {

	seconds = int(0)
	cFormat := "Mon, 02 Jan 2006 15:04:05 WAT"

	if cTimeCur == "" {
		cTimeCur = time.Now().Format(cFormat)
	}
	tTimeCur, errCur := time.Parse(cFormat, cTimeCur)
	if errCur != nil {
		return
	}

	if cTimePast == "" {
		cTimePast = time.Now().Format(cFormat)
	}
	tTimePast, errPast := time.Parse(cFormat, cTimePast)
	if errPast != nil {
		return
	}

	if errCur == nil && errPast == nil {

		secondsCur := int(tTimeCur.Unix()) 
		secondsPast := int(tTimePast.Unix()) 

		seconds = int(secondsCur - secondsPast) / 3600
	}
	return
}