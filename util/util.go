package util

import (
	"time"
)

const (
	TIME_FORMAT      string = "15:04:05"
	DATE_TIME_FORMAT string = "2006-01-02T15:04:05"
)

func IsTimeBetween(checkTime, startTime, endTime time.Time) bool {

	timeStr := checkTime.Format(TIME_FORMAT)
	newTime, err := time.Parse(TIME_FORMAT, timeStr)
	if err != nil {
		panic(err)
	}

	return !newTime.Before(startTime) && !newTime.After(endTime)
}
