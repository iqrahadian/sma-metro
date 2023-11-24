package route

import (
	"fmt"
	"time"

	"github.com/iqrahadian/sma-metro/util"
)

type PeaktimeHour struct {
	Start time.Time
	End   time.Time
}

type PeakTimeConfig struct {
	FromDay   time.Weekday
	ToDay     time.Weekday
	StartHour string
	EndHour   string
}

var PeakTimeConfigArr = []PeakTimeConfig{
	{time.Monday, time.Friday, "08:00", "10:00"},
	{time.Monday, time.Friday, "16:30", "19:00"},
	{time.Saturday, time.Saturday, "10:00", "14:00"},
	{time.Saturday, time.Saturday, "18:00", "23:00"},
	{time.Sunday, time.Sunday, "18:00", "23:00"},
}

var PeaktimeMap map[time.Weekday][]PeaktimeHour

func init() {
	PeaktimeMap = parsePeakTimeConfig()
}

func IsPeaktimePrice(route TravelRoute) bool {

	travelTime, _ := time.Parse(util.DATE_TIME_FORMAT, route.TripTime)

	peakTimes, _ := PeaktimeMap[travelTime.Weekday()]

	for _, peakTime := range peakTimes {

		if util.IsTimeBetween(travelTime, peakTime.Start, peakTime.End) {
			return true
		}

	}

	return false

}

func parsePeakTimeConfig() map[time.Weekday][]PeaktimeHour {

	peaktimeMap := map[time.Weekday][]PeaktimeHour{}

	for _, peaktime := range PeakTimeConfigArr {

		if peaktime.FromDay > peaktime.ToDay {
			panic("Oops, something wrong with peaktime config, Start & End day config is not correct")
		}

		startTimeStr := fmt.Sprintf("%s:00", peaktime.StartHour)
		endTimeStr := fmt.Sprintf("%s:00", peaktime.EndHour)

		startTime, err := time.Parse(util.TIME_FORMAT, startTimeStr)
		if err != nil {
			fmt.Println("Oops, something wrong with peaktime config, Failed to parse time:", peaktime.StartHour)
			panic(err)
		}

		endTime, err := time.Parse(util.TIME_FORMAT, endTimeStr)
		if err != nil {
			fmt.Println("Oops, something wrong with peaktime config, Failed to parse time:", peaktime.EndHour)
			panic(err)
		}

		for i := peaktime.FromDay; i <= peaktime.ToDay; i++ {

			if val, ok := peaktimeMap[i]; !ok {
				peaktimeMap[i] = []PeaktimeHour{{startTime, endTime}}
			} else {
				peaktimeMap[i] = append(val, PeaktimeHour{startTime, endTime})
			}

		}

	}

	return peaktimeMap
}
