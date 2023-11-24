package payment

import (
	"errors"
	"fmt"

	"github.com/iqrahadian/sma-metro/card"
	"github.com/iqrahadian/sma-metro/route"
)

type PaymentGateway struct{}

func (p *PaymentGateway) Charge(card *card.SmartCard, travelRoute route.TravelRoute) error {
	processor, err := p.getProcessor(card.Type)
	if err != nil {
		return err
	}

	err = processor.Charge(card, travelRoute)
	if err != nil {
		return err
	}

	return nil
}

func (p *PaymentGateway) Topup(card *card.SmartCard, amount int) error {

	processor, err := p.getProcessor(card.Type)
	if err != nil {
		return err
	}

	err = processor.Topup(card, amount)
	if err != nil {
		return err
	}

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
