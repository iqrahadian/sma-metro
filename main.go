package main

import "fmt"

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

	var TravelFaresMap map[string]TravelFares = parseTravelFaresConfig()

	fmt.Println(TravelFaresMap)

}

func parseTravelFaresConfig() map[string]TravelFares {

	faresMap := map[string]TravelFares{}

	for _, fare := range FaresConfig {
		key := fmt.Sprintf("%s%s", fare.Departure, fare.Destination)
		if _, ok := faresMap[key]; !ok {
			faresMap[key] = fare
		}
	}

	return faresMap

}
