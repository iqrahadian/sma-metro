package main

import (
	"fmt"
	"time"
)

type TravelRoute struct {
	From     Route
	To       Route
	TripTime string
}

var travelroute = []TravelRoute{
	{GreenLine, GreenLine, "2021-03-01T07:58:30"},
	{GreenLine, GreenLine, "2021-03-01T12:58:30"},
	{GreenLine, GreenLine, "2021-03-01T07:58:30"},
	{GreenLine, GreenLine, "2021-03-01T12:58:30"},
	{GreenLine, GreenLine, "2021-03-01T07:58:30"},
	{GreenLine, RedLine, "2021-03-12T09:58:30"},
	{GreenLine, RedLine, "2021-03-12T09:58:30"},
	{RedLine, RedLine, "2021-03-30T11:58:30"},
}

var (
	TravelFaresMap map[string]TravelFaresConfig    = parseTravelFaresConfig()
	PeaktimeMap    map[time.Weekday][]PeaktimeHour = parsePeakTimeConfig()
)

func main() {

	// var card = CreditCard{CreditCardType, 0, FareUsage{}}
	var card = InitCard(CreditCardType)

	card.Topup(100)
	// err := TopupCard(card, 100)
	// if err != nil {
	// 	fmt.Println("Failed to topup card", err)
	// }

	for _, route := range travelroute {
		err := ChargeCard(card, route)
		if err != nil {
			panic(fmt.Errorf("Failed to charge card %v", err))
		}
	}

	fmt.Println(card)

	for key, value := range *card.GetUsages() {
		fmt.Println("Key:", key, "Value:", value)
	}

	fmt.Println("Final Card Balance : ", card.GetBalance())

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

func parsePeakTimeConfig() map[time.Weekday][]PeaktimeHour {

	peaktimeMap := map[time.Weekday][]PeaktimeHour{}

	for _, peaktime := range PeakTimeConfigArr {

		if peaktime.FromDay > peaktime.ToDay {
			panic("Oops, something wrong with peaktime config, Start & End day config is not correct")
		}

		startTimeStr := fmt.Sprintf("%s:00", peaktime.StartHour)
		endTimeStr := fmt.Sprintf("%s:00", peaktime.EndHour)

		startTime, err := time.Parse(TIME_FORMAT, startTimeStr)
		if err != nil {
			fmt.Println("Oops, something wrong with peaktime config, Failed to parse time:", peaktime.StartHour)
			panic(err)
		}

		endTime, err := time.Parse(TIME_FORMAT, endTimeStr)
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
