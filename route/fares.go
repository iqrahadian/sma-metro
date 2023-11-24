package route

import "fmt"

type TravelFaresConfig struct {
	Departure    string
	Destination  string
	StandardCost int
	PeakCost     int
	DailyCap     int
	WeeklyCap    int
}

var FaresConfigArr = []TravelFaresConfig{
	{"green", "green", 2, 1, 8, 55},
	{"red", "red", 3, 2, 12, 70},
	{"green", "red", 4, 3, 15, 90},
	{"red", "green", 3, 2, 15, 90},
}

var TravelFaresMap map[string]TravelFaresConfig

func init() {
	TravelFaresMap = parseTravelFaresConfig()
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
