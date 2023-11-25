package payment

import (
	"errors"
	"fmt"
	"time"

	"github.com/iqrahadian/sma-metro/common"
	"github.com/iqrahadian/sma-metro/src/model"
	"github.com/iqrahadian/sma-metro/src/service/card"
	"github.com/iqrahadian/sma-metro/src/service/route"
)

func NewPaymentHandler(
	rs *route.RouteService,
	cs *card.CardService,
) *PaymentGateway {

	if rs == nil {
		panic("Route service not initiated")
	}

	if cs == nil {
		panic("Card service not initiated")
	}

	return &PaymentGateway{rs, cs}
}

type PaymentGateway struct {
	rs *route.RouteService
	cs *card.CardService
}

func (p *PaymentGateway) Charge(
	card *model.SmartCard,
	travelRoute model.TravelRoute,
) (cost int, err common.Error) {

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

func (p *PaymentGateway) Topup(card *model.SmartCard, amount int) common.Error {

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

func (p *PaymentGateway) getProcessor(cardType model.CardType) (paymentProcessor, common.Error) {

	switch cardType {
	case model.CreditCardType:
		return &creditCardProcessor{p.rs, p.cs}, common.Error{}
	default:
		return nil, common.Error{
			Code: common.CardTypeUnrecognized,
			Error: errors.New(
				fmt.Sprintf("Oops, Cannot recognize card type %s", cardType),
			),
		}
	}

}
