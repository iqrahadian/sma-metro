package model

type CardType string

const (
	CreditCardType      CardType = "credit"
	NFCDebitType                 = "nfcdebit"
	NFCRechargeableType          = "nfcrechargeable"
)

type FareUsage map[string]*FareSpending

type FareSpending struct {
	LastWeekUsed   int
	LastDayUsed    int
	WeeklySpending int
	DailySpending  int
}

type SmartCard struct {
	Type         CardType
	Balance      int
	Transactions FareUsage
	asd          int
}
