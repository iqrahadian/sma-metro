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
) (totalCost int, newBalance int, error common.Error) {

	tripTime, err := time.Parse(util.DATE_TIME_FORMAT, travelRoute.TripTime)
	if err != nil {
		return totalCost, newBalance, common.Error{err, common.InternalParseTriptime}
	}
	_, currentWeek := tripTime.ISOWeek()

	stasion := fmt.Sprintf("%s%s", travelRoute.From, travelRoute.To)
	routeFare, ok := route.TravelFaresMap[stasion]
	if !ok {
		return totalCost, newBalance, common.Error{Error: errors.New("Unkown Route"), Code: common.FaresUnknown}
	}

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

	totalCost = routeFare.StandardCost
	isPeak, error := route.IsPeaktimePrice(travelRoute)
	if error.Error != nil {
		return totalCost, newBalance, error
	} else if isPeak {
		totalCost = routeFare.PeakCost
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

	if totalCost > maxDeduction {
		totalCost = maxDeduction
	}

	if smartCard.Balance < totalCost {
		return totalCost, newBalance, common.Error{errors.New("Not enough balance"), common.CardInsufficientBalance}
	}

	smartCard.Balance -= totalCost
	newBalance = smartCard.Balance

	fareUsages.DailySpending += totalCost
	fareUsages.WeeklySpending += totalCost
	fareUsages.LastWeekUsed = currentWeek
	fareUsages.LastDayUsed = int(tripTime.Weekday())

	return totalCost, newBalance, error

}

func (c *creditCardProcessor) Topup(card *card.SmartCard, amount int) common.Error {
	card.Balance += amount
	return common.Error{}
}
