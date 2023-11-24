package main

import (
	"errors"
	"fmt"
	"time"
)

// func ChargeCard(card *CreditCard, route TravelRoute) error {
func ChargeCard(card SmartCard, route TravelRoute) error {

	stasion := fmt.Sprintf("%s%s", route.From, route.To)

	routeFare, ok := TravelFaresMap[stasion]
	if !ok {
		return errors.New("Failed to retrieve route fares")
	}

	travelTime, err := time.Parse(DATE_TIME_FORMAT, route.TripTime)
	if err != nil {
		panic(fmt.Errorf("Failed to parse travel time, err : %v", err))
	}
	_, currentWeek := travelTime.ISOWeek()

	// fareUsages, ok := card.Transactions[stasion]
	cardUsages := card.GetUsages()
	fareUsages, ok := (*cardUsages)[stasion]
	if !ok {
		fareUsages = new(FareSpending)
		card.SetUsages(stasion, fareUsages)
		// card.Transactions[stasion] = fareUsages
	} else {

		if fareUsages.LastWeekUsed < currentWeek {
			fareUsages.WeeklySpending = 0
			fareUsages.DailySpending = 0
		} else if fareUsages.LastDayUsed < int(travelTime.Weekday()) {
			fareUsages.DailySpending = 0
		}

	}

	maxDeduction := 0
	if fareUsages.DailySpending < routeFare.DailyCap {
		maxDeduction = routeFare.DailyCap - fareUsages.DailySpending
	}

	if maxDeduction > 0 && fareUsages.WeeklySpending < routeFare.WeeklyCap {
		maxWeekDeduction := routeFare.WeeklyCap - fareUsages.WeeklySpending

		if maxDeduction > maxWeekDeduction {
			maxDeduction = maxWeekDeduction
		}
	}

	cost := routeFare.StandardCost
	if IsPeaktimePrice(route) {
		cost = routeFare.PeakCost
	}

	if cost > maxDeduction {
		cost = maxDeduction
	}

	if card.GetBalance() < cost {
		// if card.Balance < cost {
		return errors.New("Not enough balance")
	}

	card.Topup(cost * -1)

	fareUsages.DailySpending += cost
	fareUsages.WeeklySpending += cost
	fareUsages.LastWeekUsed = currentWeek
	fareUsages.LastDayUsed = int(travelTime.Weekday())

	return nil
}

func IsPeaktimePrice(route TravelRoute) bool {

	travelTime, _ := time.Parse(DATE_TIME_FORMAT, route.TripTime)

	peakTimes, _ := PeaktimeMap[travelTime.Weekday()]

	for _, peakTime := range peakTimes {

		if IsTimeBetween(travelTime, peakTime.Start, peakTime.End) {
			return true
		}

	}

	return false

}

func IsTimeBetween(checkTime, startTime, endTime time.Time) bool {

	timeStr := checkTime.Format(TIME_FORMAT)
	newTime, err := time.Parse(TIME_FORMAT, timeStr)
	if err != nil {
		fmt.Println("HOHO")
		panic(err)
	}

	return !newTime.Before(startTime) && !newTime.After(endTime)
}

func TopupCard(card *CreditCard, amount int) error {

	card.Balance += amount

	return nil
}
