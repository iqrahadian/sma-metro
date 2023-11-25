package route

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/iqrahadian/sma-metro/src/model"
)

func parseTravelFaresConfig() map[string]model.TravelFaresConfig {

	faresMap := map[string]model.TravelFaresConfig{}

	f, err := os.Open("./data/fares.csv")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	for i, line := range data {
		if i > 0 { // omit header line
			var tmpFare model.TravelFaresConfig
			tmpFare.Departure = strings.ToLower(line[0])
			tmpFare.Destination = strings.ToLower(line[1])

			peakCost, _ := strconv.Atoi(line[2])
			tmpFare.PeakCost = peakCost

			standardCost, _ := strconv.Atoi(line[3])
			tmpFare.StandardCost = standardCost

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
