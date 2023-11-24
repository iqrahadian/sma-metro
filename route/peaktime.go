package route

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/iqrahadian/sma-metro/common"
	"github.com/iqrahadian/sma-metro/util"
)

type PeakTimeConfig struct {
	FromDay   time.Weekday
	ToDay     time.Weekday
	StartHour string
	EndHour   string
}

type PeaktimeHour struct {
	Start time.Time
	End   time.Time
}

var PeaktimeMap map[time.Weekday][]PeaktimeHour

func init() {
	PeaktimeMap = parsePeakTimeConfig()
}

func IsPeaktimePrice(route TravelRoute) (bool, common.Error) {

	travelTime, _ := time.Parse(util.DATE_TIME_FORMAT, route.TripTime)

	peakTimes, _ := PeaktimeMap[travelTime.Weekday()]

	for _, peakTime := range peakTimes {

		isPeak, err := util.IsTimeBetween(travelTime, peakTime.Start, peakTime.End)
		if err.Error != nil {
			return isPeak, err
		} else if isPeak {
			return isPeak, err
		}

	}

	return false, common.Error{}

}

func parsePeakTimeConfig() map[time.Weekday][]PeaktimeHour {

	var daysOfWeek = map[string]time.Weekday{
		"sunday":    time.Sunday,
		"monday":    time.Monday,
		"tuesday":   time.Tuesday,
		"wednesday": time.Wednesday,
		"thursday":  time.Thursday,
		"friday":    time.Friday,
		"saturday":  time.Saturday,
	}

	peaktimeMap := map[time.Weekday][]PeaktimeHour{}
	PeakTimeConfigArr := []PeakTimeConfig{}

	f, err := os.Open("./input/peaktime.csv")
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
			tmpConfig := PeakTimeConfig{
				FromDay:   daysOfWeek[line[0]],
				ToDay:     daysOfWeek[line[1]],
				StartHour: line[2],
				EndHour:   line[3],
			}

			PeakTimeConfigArr = append(PeakTimeConfigArr, tmpConfig)
		}
	}

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
