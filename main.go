package main

import (
	"fmt"
	"time"
)

type PeaktimeHour struct {
	Start time.Time
	End   time.Time
}

type TravelRoute struct {
	From     Route
	To       Route
	TripTime string
}

var travelroute = []TravelRoute{
	{GreenLine, GreenLine, "2021-03-24T07:58:30"},
	{GreenLine, RedLine, "2021-03-24T09:58:30"},
	{RedLine, RedLine, "2021-03-25T11:58:30"},
}

func main() {

	var TravelFaresMap map[string]TravelFaresConfig = parseTravelFaresConfig()
	fmt.Println(TravelFaresMap)

	var PeaktimeMap map[Weekday][]PeaktimeHour = parsePeakTimeConfig()
	fmt.Println(PeaktimeMap)

}

func parseTravelFaresConfig() map[string]TravelFaresConfig {

	faresMap := map[string]TravelFaresConfig{}

	for _, fare := range FaresConfigArr {
		key := fmt.Sprintf("%s%s", fare.Departure, fare.Destination)
		if _, ok := faresMap[key]; !ok {
			faresMap[key] = fare
		}
	}

	return faresMap

}

func parsePeakTimeConfig() map[Weekday][]PeaktimeHour {

	peaktimeMap := map[Weekday][]PeaktimeHour{}

	for _, peaktime := range PeakTimeConfigArr {

		startTimeStr := fmt.Sprintf("%s:00", peaktime.StartHour)
		endTimeStr := fmt.Sprintf("%s:00", peaktime.EndHour)

		startTime, err := time.Parse(TIME_FORMAT, startTimeStr)
		if err != nil {
			fmt.Println("Error parsing time:", peaktime.StartHour)
			panic(err)
		}

		endTime, err := time.Parse(TIME_FORMAT, endTimeStr)
		if err != nil {
			fmt.Println("Error parsing time:", peaktime.EndHour)
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
