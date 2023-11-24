package payment

import (
	"errors"
	"fmt"
	"time"

	"github.com/iqrahadian/sma-metro/card"
	"github.com/iqrahadian/sma-metro/common"
	"github.com/iqrahadian/sma-metro/route"
)

type PaymentGateway struct{}

func (p *PaymentGateway) Charge(card *card.SmartCard, travelRoute route.TravelRoute) (cost int, err common.Error) {

	time.Sleep(500 * time.Millisecond)
	fmt.Println("Processing payment for trip from", travelRoute.From, "to", travelRoute.To, "on", travelRoute.TripTime)

	processor, err := p.getProcessor(card.Type)
	if err.Error != nil {
		return cost, err
	}

	cost, balance, err := processor.Charge(card, travelRoute)
	if err.Error != nil {
		return cost, err
	}

	time.Sleep(500 * time.Millisecond)
	fmt.Println(fmt.Sprintf("This trip cost $%d, your balance is $%d", cost, balance))
	fmt.Println("--------------------------------------------------------->")

	return cost, err
}

func (p *PaymentGateway) Topup(card *card.SmartCard, amount int) common.Error {

	time.Sleep(500 * time.Millisecond)
	fmt.Println(fmt.Sprintf("Processing Card topup for $%d", amount))
	processor, err := p.getProcessor(card.Type)
	if err.Error != nil {
		return err
	}

	err = processor.Topup(card, amount)
	if err.Error != nil {
		return err
	}

	time.Sleep(500 * time.Millisecond)
	fmt.Println(fmt.Sprintf("Success processing Card topup your new balance : $%d", card.Balance))
	fmt.Println("--------------------------------------------------------->")

	return common.Error{}
}

func (p *PaymentGateway) getProcessor(cardType card.CardType) (paymentProcessor, common.Error) {

	switch cardType {
	case card.CreditCardType:
		return &creditCardProcessor{}, common.Error{}
	default:
		return nil, common.Error{
			Code: common.CardTypeUnrecognized,
			Error: errors.New(
				fmt.Sprintf("Oops, Cannot recognize card type %s", cardType),
			),
		}
	}

}
