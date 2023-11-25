package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/iqrahadian/sma-metro/card"
	"github.com/iqrahadian/sma-metro/common"
	"github.com/iqrahadian/sma-metro/payment"
	"github.com/iqrahadian/sma-metro/route"
)

func main() {

	travelRoutes := loadInput()

	paymentGateway := payment.NewPaymentGateway()

	// init new card, in real case we retrieve from storage by id
	smartCard := card.InitCard(card.CreditCardType)

	paymentGateway.Topup(&smartCard, 100)

	totalFareApplied := 0

	for _, travelRoute := range travelRoutes {

		cost, err := paymentGateway.Charge(&smartCard, travelRoute)
		if err.Error != nil {
			time.Sleep(500 * time.Millisecond)
			fmt.Println(common.GetErrorMessage(common.ErrorCode(err.Code)))
			fmt.Println("--------------------------------------------------------->")
		} else {
			totalFareApplied += cost
		}
	}

	fmt.Println("Final Card Balance : ", smartCard.Balance)
	fmt.Println("Total Fare Applied : ", totalFareApplied)

}

func loadInput() []route.TravelRoute {

	travelRoutes := []route.TravelRoute{}

	f, err := os.Open("./input/input.csv")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// read csv values using csv.Reader
	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	for i, line := range data {
		if i > 0 { // omit header line
			var tmpRoute route.TravelRoute
			tmpRoute.From = strings.ToLower(line[0])
			tmpRoute.To = strings.ToLower(line[1])
			tmpRoute.TripTime = line[2]

			travelRoutes = append(travelRoutes, tmpRoute)
		}
	}

	return travelRoutes

}
