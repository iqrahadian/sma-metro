package route

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type TravelFaresConfig struct {
	Departure    string
	Destination  string
	StandardCost int
	PeakCost     int
	DailyCap     int
	WeeklyCap    int
}

var TravelFaresMap map[string]TravelFaresConfig

func init() {
	TravelFaresMap = parseTravelFaresConfig()
}

func parseTravelFaresConfig() map[string]TravelFaresConfig {

	faresMap := map[string]TravelFaresConfig{}

	f, err := os.Open("./input/fares.csv")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// read csv values using csv.Reader
	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	for i, line := range data {
		if i > 0 { // omit header line
			var tmpFare TravelFaresConfig
			tmpFare.Departure = strings.ToLower(line[0])
			tmpFare.Destination = strings.ToLower(line[1])

			standardCost, _ := strconv.Atoi(line[2])
			tmpFare.StandardCost = standardCost

			peakCost, _ := strconv.Atoi(line[3])
			tmpFare.PeakCost = peakCost

			dailyCap, _ := strconv.Atoi(line[4])
			tmpFare.DailyCap = dailyCap

			weeklyCap, _ := strconv.Atoi(line[2])
			tmpFare.WeeklyCap = weeklyCap

			key := fmt.Sprintf("%s%s", tmpFare.Departure, tmpFare.Destination)

			faresMap[key] = tmpFare
		}
	}

	return faresMap
}
