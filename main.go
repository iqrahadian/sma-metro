package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/iqrahadian/sma-metro/common"
	"github.com/iqrahadian/sma-metro/src/handler/payment"
	"github.com/iqrahadian/sma-metro/src/model"
	"github.com/iqrahadian/sma-metro/src/service/card"
	"github.com/iqrahadian/sma-metro/src/service/route"
)

func main() {
	type testData struct {
		result      int
		dailySpend  int
		weeklySpend int
	}

	ccProcessor := payment.Credit{}

	routeFares := model.TravelFaresConfig{
		DailyCap:  10,
		WeeklyCap: 50,
	}

	testArr := []testData{
		{5, 5, 10},
	}

	for _, singleTest := range testArr {

		maxDeduction := ccProcessor.GetMaxDeduction(
			routeFares,
			singleTest.dailySpend,
			singleTest.weeklySpend,
		)

		// fmt.Println("Max Deduction Result : %v, Want %v", maxDeduction, singleTest.result)
		if maxDeduction == singleTest.result {
			// fmt.Println("Max Deduction Result : %v, Want %v", maxDeduction, singleTest.result)
			fmt.Println("sama")
			return
		}
	}
}

func main2() {

	travelRoutes := loadInput()

	rs := route.NewRouteService()
	cs := card.NewCardService()
	paymentGateway := payment.NewPaymentHandler(rs, cs)

	// init new card, in real case we retrieve from storage by id
	smartCard := cs.NewCard(model.CreditCardType)

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

func loadInput() []model.TravelRoute {

	travelRoutes := []model.TravelRoute{}

	f, err := os.Open("./data/input.csv")
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
			var tmpRoute model.TravelRoute
			tmpRoute.From = strings.ToLower(line[0])
			tmpRoute.To = strings.ToLower(line[1])
			tmpRoute.TripTime = line[2]

			travelRoutes = append(travelRoutes, tmpRoute)
		}
	}

	return travelRoutes

}
