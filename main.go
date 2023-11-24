package main

import (
	"fmt"

	"github.com/iqrahadian/sma-metro/card"
	"github.com/iqrahadian/sma-metro/payment"
	"github.com/iqrahadian/sma-metro/route"
)

var travelRoutes = []route.TravelRoute{
	{route.GreenLine, route.GreenLine, "2021-03-01T07:58:30"},
	{route.GreenLine, route.GreenLine, "2021-03-01T12:58:30"},
	{route.GreenLine, route.GreenLine, "2021-03-01T07:58:30"},
	{route.GreenLine, route.GreenLine, "2021-03-01T12:58:30"},
	{route.GreenLine, route.GreenLine, "2021-03-01T07:58:30"},
	{route.GreenLine, route.RedLine, "2021-03-12T09:58:30"},
	{route.GreenLine, route.RedLine, "2021-03-12T09:58:30"},
	{route.RedLine, route.RedLine, "2021-03-30T11:58:30"},
}

func main() {

	paymentGateway := payment.PaymentGateway{}

	smartCard := card.InitCard(card.CreditCardType)

	paymentGateway.Topup(&smartCard, 100)

	for _, travelRoute := range travelRoutes {

		err := paymentGateway.Charge(&smartCard, travelRoute)
		if err != nil {
			panic(fmt.Errorf("Failed to charge card %v", err))
		}
	}

	// for key, value := range *&smartCard.Transactions {
	// 	fmt.Println("Key:", key, "Value:", value)
	// }

	fmt.Println("Final Card Balance : ", smartCard.Balance)

}
