package main

type TransactionType string

const (
	TransactionTopup TransactionType = "topup"
	TransactionPay                   = "pay"
)

type Transaction struct {
	Type    TransactionType
	Route   TravelRoute
	Amount  int
	Remarks string
}
