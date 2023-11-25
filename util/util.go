package util

import (
	"time"

	"github.com/iqrahadian/sma-metro/common"
)

const (
	TIME_FORMAT      string = "15:04:05"
	DATE_TIME_FORMAT string = "2006-01-02T15:04:05"
)

func IsTimeBetween(checkTime, startTime, endTime time.Time) (bool, common.Error) {

	timeStr := checkTime.Format(TIME_FORMAT)
	newTime, err := time.Parse(TIME_FORMAT, timeStr)
	if err != nil {
		return false, common.Error{Error: err, Code: common.InternalParseTriptime}
	}

	return !newTime.Before(startTime) && !newTime.After(endTime), common.Error{}
}
