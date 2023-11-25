package route

import (
	"errors"
	"fmt"
	"time"

	"github.com/iqrahadian/sma-metro/common"
	"github.com/iqrahadian/sma-metro/src/model"
	"github.com/iqrahadian/sma-metro/util"
)

func NewRouteService() *RouteService {

	peakTimeMap := parsePeakTimeConfig()
	travelFaresMap := parseTravelFaresConfig()

	return &RouteService{
		peakTimeMap,
		travelFaresMap,
	}
}

type RouteService struct {
	peakTimeMap    map[time.Weekday][]model.PeaktimeHour
	travelFaresMap map[string]model.TravelFaresConfig
}

func (r *RouteService) GetTravelCost(travelRoute model.TravelRoute) (cost int, error common.Error) {

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

func (r *RouteService) GetRouteFare(stasion string) (routeFare model.TravelFaresConfig, err common.Error) {

	routeFare, ok := r.travelFaresMap[stasion]
	if !ok {
		return routeFare, common.Error{Error: errors.New("Unkown Route"), Code: common.FaresUnknown}
	}

	return routeFare, err

}

func (r *RouteService) isPeaktime(route model.TravelRoute) (bool, common.Error) {

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
