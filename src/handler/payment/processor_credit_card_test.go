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
		DailyCap:  12,
		WeeklyCap: 39,
	}

	testArr := []testData{
		{9, 3, 10},  // validate use daily as base calculation
		{0, 12, 12}, // validate use daily as base calculation
		{0, 5, 39},  // validate use daily as base calculation
		{1, 3, 38},  // validate use weekly as base calculation
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

func TestCreditCardProcessorgetMaxDeductionNocap(t *testing.T) {

	type testData struct {
		result      int
		dailySpend  int
		weeklySpend int
		routeFare   model.TravelFaresConfig
	}

	ccProcessor := &creditCardProcessor{nil, nil}

	testArr := []testData{
		{10, 50, 90, model.TravelFaresConfig{DailyCap: 0, WeeklyCap: 100}},
		{70, 20, 30, model.TravelFaresConfig{DailyCap: 0, WeeklyCap: 100}},
		{17, 3, 38, model.TravelFaresConfig{DailyCap: 20, WeeklyCap: 0}},
		{999, 3, 38, model.TravelFaresConfig{DailyCap: 0, WeeklyCap: 0}},
	}

	for _, singleTest := range testArr {

		maxDeduction := ccProcessor.getMaxDeduction(
			singleTest.routeFare,
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
