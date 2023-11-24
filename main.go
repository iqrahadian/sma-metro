package main

import (
	"fmt"
	"time"

	"github.com/iqrahadian/sma-metro/card"
	"github.com/iqrahadian/sma-metro/common"
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
	{route.GreenLine, route.GreenLine, "2021-03-02T07:58:30"},
}

func main() {

	// init payment gateway service
	paymentGateway := payment.PaymentGateway{}

	// init new card, in real case we retrieve from storage by id
	smartCard := card.InitCard(card.CreditCardType)

	paymentGateway.Topup(&smartCard, 100)

	for _, travelRoute := range travelRoutes {

		err := paymentGateway.Charge(&smartCard, travelRoute)
		if err.Error != nil {
			time.Sleep(500 * time.Millisecond)
			fmt.Println(common.GetErrorMessage(common.ErrorCode(err.Code)))
			fmt.Println("--------------------------------------------------------->")
		}
	}

	fmt.Println("Final Card Balance : ", smartCard.Balance)

}
