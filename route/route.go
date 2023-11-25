package route

import (
	"errors"
	"fmt"
	"time"

	"github.com/iqrahadian/sma-metro/common"
	"github.com/iqrahadian/sma-metro/util"
)

type Route string

const (
	GreenLine Route = "green"
	RedLine         = "red"
)

type TravelRoute struct {
	From     string
	To       string
	TripTime string
}

func NewRouteService() *RouteService {

	peakTimeMap := parsePeakTimeConfig()
	travelFaresMap := parseTravelFaresConfig()

	return &RouteService{
		peakTimeMap,
		travelFaresMap,
	}
}

type RouteService struct {
	peakTimeMap    map[time.Weekday][]PeaktimeHour
	travelFaresMap map[string]TravelFaresConfig
}

func (r *RouteService) GetTravelCost(travelRoute TravelRoute) (cost int, error common.Error) {

	stasion := fmt.Sprintf("%s%s", travelRoute.From, travelRoute.To)

	routeFare, ok := r.travelFaresMap[stasion]
	if !ok {
		return cost, common.Error{Error: errors.New("Unkown Route"), Code: common.FaresUnknown}
	}

	cost = routeFare.StandardCost
	isPeak, error := r.isPeaktime(travelRoute)
	if error.Error != nil {
		return cost, error
	} else if isPeak {
		cost = routeFare.PeakCost
	}

	return cost, error
}

func (r *RouteService) GetRouteFare(stasion string) (routeFare TravelFaresConfig, err common.Error) {

	routeFare, ok := r.travelFaresMap[stasion]
	if !ok {
		return routeFare, common.Error{Error: errors.New("Unkown Route"), Code: common.FaresUnknown}
	}

	return routeFare, err

}

func (r *RouteService) isPeaktime(route TravelRoute) (bool, common.Error) {

	travelTime, _ := time.Parse(util.DATE_TIME_FORMAT, route.TripTime)

	peakTimes, _ := r.peakTimeMap[travelTime.Weekday()]

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
