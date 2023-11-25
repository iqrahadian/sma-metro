package route

import (
	"testing"

	"github.com/iqrahadian/sma-metro/src/model"
)

func TestIsPeakTime(t *testing.T) {

	peakTimeMap := parsePeakTimeConfig("../../../data/peaktime_test.csv")
	travelFaresMap := parseTravelFaresConfig("../../../data/fares_test.csv")
	rs := RouteService{
		peakTimeMap,
		travelFaresMap,
	}

	type testStruct struct {
		result      bool
		travelRoute model.TravelRoute
	}

	testData := []testStruct{
		{false, model.TravelRoute{TripTime: "2021-03-01T07:58:30"}},
		{true, model.TravelRoute{TripTime: "2021-03-01T08:58:30"}},
		{false, model.TravelRoute{TripTime: "2021-03-01T12:58:30"}},
		{true, model.TravelRoute{TripTime: "2021-03-01T18:58:30"}},
		{false, model.TravelRoute{TripTime: "2021-03-06T08:58:30"}},
		{true, model.TravelRoute{TripTime: "2021-03-06T10:58:30"}},
	}

	for _, singleTest := range testData {

		isPeaktime, _ := rs.isPeaktime(singleTest.travelRoute)
		if singleTest.result != isPeaktime {
			t.Errorf("Peaktime Result : %v, Want %v", isPeaktime, singleTest.result)
			return
		}
	}

}
