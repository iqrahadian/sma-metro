package card

import (
	"errors"
	"fmt"
	"time"

	"github.com/iqrahadian/sma-metro/common"
	"github.com/iqrahadian/sma-metro/src/model"
	"github.com/iqrahadian/sma-metro/util"
)

func NewCardService() *CardService {
	return &CardService{}
}

type CardService struct{}

func (c *CardService) UpdateCardBalance(
	card *model.SmartCard,
	totalCost int,
	dailySpend int,
	weeklySpend int,
	travelRoute model.TravelRoute,
) (error common.Error) {

	tripTime, err := time.Parse(util.DATE_TIME_FORMAT, travelRoute.TripTime)
	if err != nil {
		return common.Error{Error: err, Code: common.InternalParseTriptime}
	}
	_, currentWeek := tripTime.ISOWeek()

	stasion := fmt.Sprintf("%s%s", travelRoute.From, travelRoute.To)

	fareSpending, ok := card.Transactions[stasion]
	if !ok {
		return common.Error{Error: errors.New("Unkown Route"), Code: common.FaresUnknown}
	}

	card.Balance -= totalCost

	fareSpending.LastWeekUsed = currentWeek
	fareSpending.LastDayUsed = int(tripTime.Weekday())
	fareSpending.WeeklySpending = weeklySpend
	fareSpending.DailySpending = dailySpend

	return error
}

func (c *CardService) GetFareSpending(card *model.SmartCard, stasion string) *model.FareSpending {

	fareSpending, ok := card.Transactions[stasion]
	if !ok {
		fareSpending = new(model.FareSpending)
		card.Transactions[stasion] = fareSpending
	}

	return fareSpending
}

func (c *CardService) NewCard(cardType model.CardType) model.SmartCard {

	switch cardType {
	case model.CreditCardType:
	case model.NFCDebitType:
	case model.NFCRechargeableType:
	default:
		panic(fmt.Sprintf("Oops, Cannot recognize card type %s", cardType))
	}

	return model.SmartCard{
		Type:         cardType,
		Balance:      0,
		Transactions: model.FareUsage{},
	}

}
