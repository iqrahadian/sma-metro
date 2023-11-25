package payment

import (
	"testing"

	"github.com/iqrahadian/sma-metro/src/model"
)

func TestCreditCardProcessorgetMaxDeduction(t *testing.T) {

	type testData struct {
		result      int
		dailySpend  int
		weeklySpend int
	}

	ccProcessor := &creditCardProcessor{nil, nil}
	routeFares := model.TravelFaresConfig{
		DailyCap:  10,
		WeeklyCap: 50,
	}

	testArr := []testData{
		{5, 5, 10},
	}

	for _, singleTest := range testArr {

		maxDeduction := ccProcessor.getMaxDeduction(
			routeFares,
			singleTest.dailySpend,
			singleTest.weeklySpend,
		)

		if maxDeduction != singleTest.result {
			t.Errorf("Max Deduction Result : %v, Want %v", maxDeduction, singleTest.result)
			return
		}
	}

	return

}
