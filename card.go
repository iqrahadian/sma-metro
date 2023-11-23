package main

import "time"

type SmartCard interface {
	Charge(route TravelRoute) error
	Topup(amount int)
}

type CardType string

const (
	CreditCard      CardType = "credit"
	NFCDebit                 = "nfcdebit"
	NFCRechargeable          = "nfcrechargeable"
)

type Card struct {
	Type               CardType
	Balance            int
	DailyUsage         int
	WeeklyUsage        int
	TransactionHistory map[time.Month]map[int]map[Weekday][]Transaction
}
