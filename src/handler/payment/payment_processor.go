package payment

import (
	"errors"
	"fmt"
	"time"

	"github.com/iqrahadian/sma-metro/common"
	"github.com/iqrahadian/sma-metro/src/model"
	"github.com/iqrahadian/sma-metro/src/service/card"
	"github.com/iqrahadian/sma-metro/src/service/route"
	"github.com/iqrahadian/sma-metro/util"
)

type paymentProcessor interface {
	Charge(*model.SmartCard, model.TravelRoute) (cost int, balance int, err common.Error)
	Topup(*model.SmartCard, int) common.Error
}

type creditCardProcessor struct {
	rs *route.RouteService
	cs *card.CardService
}

func (c *creditCardProcessor) Charge(
	smartCard *model.SmartCard,
	travelRoute model.TravelRoute,
) (travelCost int, newBalance int, error common.Error) {

	tripTime, err := time.Parse(util.DATE_TIME_FORMAT, travelRoute.TripTime)
	if err != nil {
		return travelCost, newBalance, common.Error{Error: err, Code: common.InternalParseTriptime}
	}
	_, currentWeek := tripTime.ISOWeek()

	stasion := fmt.Sprintf("%s%s", travelRoute.From, travelRoute.To)

	routeFare, error := c.rs.GetRouteFare(stasion)
	if err != nil {
		return travelCost, newBalance, common.Error{Error: err, Code: common.InternalParseTriptime}
	}

	currentFareSpending := c.cs.GetFareSpending(smartCard, stasion)

	weeklySpendTmp := currentFareSpending.WeeklySpending
	dailySpendTmp := currentFareSpending.DailySpending
	if currentFareSpending.LastWeekUsed < currentWeek {
		weeklySpendTmp = 0
		dailySpendTmp = 0
	} else if currentFareSpending.LastDayUsed < int(tripTime.Weekday()) {
		dailySpendTmp = 0
	}

	travelCost, error = c.rs.GetTravelCost(travelRoute)
	if error.Error != nil {
		return travelCost, newBalance, error
	}

	maxDeduction := c.getMaxDeduction(routeFare, dailySpendTmp, weeklySpendTmp)

	if travelCost > maxDeduction {
		travelCost = maxDeduction
	}

	if smartCard.Balance < travelCost {
		return travelCost, newBalance, common.Error{Error: errors.New("Not enough balance"), Code: common.CardInsufficientBalance}
	}

	c.cs.UpdateCardBalance(
		smartCard,
		smartCard.Balance-travelCost,
		dailySpendTmp+travelCost,
		weeklySpendTmp+travelCost,
		travelRoute,
	)

	return travelCost, smartCard.Balance, error

}

func (c *creditCardProcessor) getMaxDeduction(
	routeFare model.TravelFaresConfig,
	dailySpend int,
	weeklySpend int,
) int {

	if routeFare.DailyCap == 0 && routeFare.WeeklyCap == 0 {
		return 999 // assume highest cost
	}

	maxDeduction := 0

	if routeFare.DailyCap > 0 && dailySpend <= routeFare.DailyCap {
		maxDeduction = routeFare.DailyCap - dailySpend

		if maxDeduction == 0 { // return early, daily cap reached
			return 0
		}
	}

	if routeFare.WeeklyCap > 0 && weeklySpend <= routeFare.WeeklyCap {

		maxWeekDeduction := routeFare.WeeklyCap - weeklySpend

		if maxDeduction > maxWeekDeduction || maxDeduction == 0 {
			maxDeduction = maxWeekDeduction
		}
	}

	return maxDeduction
}

func (c *creditCardProcessor) Topup(card *model.SmartCard, amount int) common.Error {
	card.Balance += amount
	return common.Error{}
}
