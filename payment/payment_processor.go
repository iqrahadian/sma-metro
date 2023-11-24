package payment

import (
	"errors"
	"fmt"
	"time"

	"github.com/iqrahadian/sma-metro/card"
	"github.com/iqrahadian/sma-metro/common"
	"github.com/iqrahadian/sma-metro/route"
	"github.com/iqrahadian/sma-metro/util"
)

type paymentProcessor interface {
	Charge(*card.SmartCard, route.TravelRoute) (cost int, balance int, err common.Error)
	Topup(*card.SmartCard, int) common.Error
}

type creditCardProcessor struct{}

func (c *creditCardProcessor) Charge(
	smartCard *card.SmartCard,
	travelRoute route.TravelRoute,
) (cost int, balance int, error common.Error) {

	stasion := fmt.Sprintf("%s%s", travelRoute.From, travelRoute.To)
	routeFare, ok := route.TravelFaresMap[stasion]
	if !ok {
		return cost, balance, common.Error{Error: errors.New("Unkown Route"), Code: common.FaresUnknown}
	}

	tripTime, err := time.Parse(util.DATE_TIME_FORMAT, travelRoute.TripTime)
	if err != nil {
		return cost, balance, common.Error{err, common.InternalParseTriptime}
	}
	_, currentWeek := tripTime.ISOWeek()

	cardUsages := &smartCard.Transactions
	fareUsages, ok := (*cardUsages)[stasion]
	if !ok {
		fareUsages = new(card.FareSpending)
		smartCard.Transactions[stasion] = fareUsages
	} else {

		if fareUsages.LastWeekUsed < currentWeek {
			fareUsages.WeeklySpending = 0
			fareUsages.DailySpending = 0
		} else if fareUsages.LastDayUsed < int(tripTime.Weekday()) {
			fareUsages.DailySpending = 0
		}

	}

	maxDeduction := 0
	if fareUsages.DailySpending < routeFare.DailyCap && routeFare.DailyCap > 0 {
		maxDeduction = routeFare.DailyCap - fareUsages.DailySpending
	}

	if maxDeduction > 0 && fareUsages.WeeklySpending < routeFare.WeeklyCap && routeFare.WeeklyCap > 0 {
		maxWeekDeduction := routeFare.WeeklyCap - fareUsages.WeeklySpending

		if maxDeduction > maxWeekDeduction {
			maxDeduction = maxWeekDeduction
		}
	}

	cost = routeFare.StandardCost
	isPeak, error := route.IsPeaktimePrice(travelRoute)
	if error.Error != nil {
		return cost, balance, error
	} else if isPeak {
		cost = routeFare.PeakCost
	}

	if cost > maxDeduction {
		cost = maxDeduction
	}

	if smartCard.Balance < cost {
		return cost, balance, common.Error{errors.New("Not enough balance"), common.CardInsufficientBalance}
	}

	smartCard.Balance -= cost
	balance = smartCard.Balance

	fareUsages.DailySpending += cost
	fareUsages.WeeklySpending += cost
	fareUsages.LastWeekUsed = currentWeek
	fareUsages.LastDayUsed = int(tripTime.Weekday())

	return cost, balance, error

}

func (c *creditCardProcessor) Topup(card *card.SmartCard, amount int) common.Error {
	card.Balance += amount
	return common.Error{}
}
