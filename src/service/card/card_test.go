package card

import (
	"testing"

	"github.com/iqrahadian/sma-metro/src/model"
)

func TestCardAddBalance(t *testing.T) {

	smartCard := model.SmartCard{
		Balance:      0,
		Transactions: make(model.FareUsage),
	}

	cs := NewCardService()

	err := cs.UpdateCardBalance(
		&smartCard,
		10,
		10,
		10,
		model.TravelRoute{
			From:     "green",
			To:       "red",
			TripTime: "2021-03-01T07:58:30",
		},
	)

	if err.Error != nil {
		t.Errorf(err.Error.Error())
	}

	if smartCard.Balance != 10 {
		t.Errorf("Card Balance Result : %v, Want %v", smartCard.Balance, 10)
		return
	}

	if fareSpending, ok := smartCard.Transactions["greenred"]; !ok {
		t.Errorf("Card Fare FareSpending does not generated")
		return
	} else {
		if fareSpending.DailySpending != 10 {
			t.Errorf("Card FareSpending Daily not correct Result : %v, Want %v", fareSpending.DailySpending, 10)
			return
		}

		if fareSpending.WeeklySpending != 10 {
			t.Errorf("Card FareSpending Weekly not correct Result : %v, Want %v", fareSpending.WeeklySpending, 10)
			return
		}

		if fareSpending.LastDayUsed != 1 {
			t.Errorf("Card FareSpending LastDayUsed not correct Result : %v, Want %v", fareSpending.LastDayUsed, 1)
			return
		}

		if fareSpending.LastWeekUsed != 9 {
			t.Errorf("Card FareSpending LastWeekUsed not correct Result : %v, Want %v", fareSpending.LastWeekUsed, 9)
			return
		}
	}

}
