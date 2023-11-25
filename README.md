# sma-metro

- to run simply just `go run main.go`
- to change input, all config stored under /input folder

## code structure/ideas
1. card is an object holding balance & transaction data, but should not modify it's own information
2. all data changes executed by PaymentGateway class, as it's represent a payment services
3. I add 3 digit number for error code identifier, will be printed when error arised, representing a card tap machine that can only show limited information

## object/class description
### Card
card object will hold information about the card, such as CardType, Balance & FareUsages

Card {
    CardType string
    Balance int
    Transactions FareUsages
}

FareUsages map[string]FareSpending

FareSpending{
    LastWeekUsed   int
	LastDayUsed    int
	WeeklySpending int
	DailySpending  int
}

FareUsage, storing information of each route line combination, combination mean : Green->Red = GreenRed
FareUsage will be used as an information on each route card holder has going through

FareSpending, storing usage information of single line combination
All 4 attributes value will be close to realtime as possible
LastWeekUsed will follow ISO Week, so it can be 52-53 in a year
    if LastWeekUsed < Current ISO Week
        WeeklySpending & DailySpending will be reset to 0 before trip calculation
LastDayUsed will follow numeric weekday system, sunday as 0 and saturday as 7
    if LastDayUsed < Current Weekday
        DailySpending will be reset to 0 before trip calculation

### Payment Gateway
paymentGateway class : payment interface & decide on how to process card based on the card type
have 2 function, Charge & Topup as interface to outside world

Payment gateway class will create a new payment processor based on Card type submitted for each function, so to handle each card type the logic will be isolated on the payment processor

paymentProcessor interface {
	Charge(*card.SmartCard, route.TravelRoute)
	Topup(*card.SmartCard, int)
}

Topup can be as simple as incresing the balance, or can be complex depend on the card type
Charge flow generally will looks like :
    1. check route fare config (from->to)
    2. check peaktime/non peaktime cost
    3. check card total spending, deciding on max deduction for this trip
       check 2 things, daily spending and weekly spending, compared to route fare caps
    4. compare fare cost to max deduction, if cost higher than max deduction,
       max deduciton used as cost, making sure the user does not overcharged
    5. deduct balance
    6. update Card FareSpending information