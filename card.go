package main

import (
	"fmt"
)

type SmartCard interface {
	GetBalance() int
	Topup(amount int) error
	GetUsages() *FareUsage
	SetUsages(stasion string, fareSpending *FareSpending)
}

func InitCard(cardType CardType) SmartCard {

	switch cardType {
	case CreditCardType:
		return &CreditCard{CreditCardType, 0, &FareUsage{}}
	default:
		panic(fmt.Sprintf("%s card type is not recognized", cardType))
	}

}

type CardType string

const (
	CreditCardType      CardType = "credit"
	NFCDebitType                 = "nfcdebit"
	NFCRechargeableType          = "nfcrechargeable"
)

type CreditCard struct {
	Type    CardType
	Balance int
	// Transactions FareUsage
	Transactions *FareUsage
}

func (c *CreditCard) GetBalance() int {
	return c.Balance
}

func (c *CreditCard) Topup(amount int) error {
	c.Balance += amount
	return nil
}

func (c *CreditCard) GetUsages() *FareUsage {
	return c.Transactions
}

func (c *CreditCard) SetUsages(stasion string, fareSpending *FareSpending) {
	(*c.Transactions)[stasion] = fareSpending
}

// init another card type

type FareUsage map[string]*FareSpending

type FareSpending struct {
	LastWeekUsed   int
	LastDayUsed    int
	WeeklySpending int
	DailySpending  int
}
