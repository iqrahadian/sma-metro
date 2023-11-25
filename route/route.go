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

type RouteService struct{}

func (r *RouteService) GetTravelCost(travelRoute TravelRoute) (cost int, error common.Error) {

	stasion := fmt.Sprintf("%s%s", travelRoute.From, travelRoute.To)

	routeFare, ok := TravelFaresMap[stasion]
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

func (r *RouteService) isPeaktime(route TravelRoute) (bool, common.Error) {

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
