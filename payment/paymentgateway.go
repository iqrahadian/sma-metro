package payment

import (
	"errors"
	"fmt"
	"time"

	"github.com/iqrahadian/sma-metro/card"
	"github.com/iqrahadian/sma-metro/route"
)

type PaymentGateway struct{}

func (p *PaymentGateway) Charge(card *card.SmartCard, travelRoute route.TravelRoute) error {

	time.Sleep(500 * time.Millisecond)
	fmt.Println("Processing payment for trip from", travelRoute.From, "to", travelRoute.To, "on", travelRoute.TripTime)

	processor, err := p.getProcessor(card.Type)
	if err != nil {
		return err
	}

	cost, balance, err := processor.Charge(card, travelRoute)
	if err != nil {
		return err
	}

	time.Sleep(500 * time.Millisecond)
	fmt.Println(fmt.Sprintf("This trip cost $%d, your balance is $%d", cost, balance))
	fmt.Println("--------------------------------------------------------->")

	return nil
}

func (p *PaymentGateway) Topup(card *card.SmartCard, amount int) error {

	time.Sleep(500 * time.Millisecond)
	fmt.Println(fmt.Sprintf("Processing Card topup for $%d", amount))
	processor, err := p.getProcessor(card.Type)
	if err != nil {
		return err
	}

	err = processor.Topup(card, amount)
	if err != nil {
		return err
	}

	time.Sleep(500 * time.Millisecond)
	fmt.Println(fmt.Sprintf("Success processing Card topup your new balance : $%d", card.Balance))
	fmt.Println("--------------------------------------------------------->")

	return nil
}

func (p *PaymentGateway) getProcessor(cardType card.CardType) (paymentProcessor, error) {

	switch cardType {
	case card.CreditCardType:
		return &creditCardProcessor{}, nil
	default:
		return nil, errors.New(
			fmt.Sprintf("Oops, Cannot recognize card type %s", cardType),
		)
	}

}
