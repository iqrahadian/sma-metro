# sma-metro

- to run simply just `go run main.go`
- to change input, all config stored under /input folder

## code structure/ideas
1. A card is an object that stores balance and transaction data. However, it shouldn't alter its own information.
2. The PaymentGateway class manages all data changes and functions as the representation of payment services.
3. I used a three-digit error code to identify errors. This code will be displayed when errors occur, similar to a card tap machine that can only show limited information.

## object/class description
### Card
A card is an object that stores balance and transaction data. However, it shouldn't alter its own information.
```
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
```

**FareUsage** stores information for each route-line combination, where combinations are denoted as such: Green->Red = GreenRed.  
This feature is utilized to retain data regarding the routes a cardholder has traveled through.  

**FareSpending**, storing usage information of single line combination.  
All four attribute values aim to reflect real-time data as closely as possible.  

**LastWeekUsed** will follow ISO Week, allowing the values ranging from 52 to 53 in a year  
&emsp;if LastWeekUsed < Current ISO Week  
&emsp;**WeeklySpending** & **DailySpending** will reset to 0 before trip calculations.  

**LastDayUsed** will follow numeric weekday system, Sunday as 0 and Saturday as 7  
&emsp;if LastDayUsed < Current Weekday  
&emsp;**DailySpending** will be reset to 0 before trip calculation.  

### Payment Gateway
paymentGateway class will serve payment interface & decide on how to process card based on the card type.  
it has 2 function, **Charge** & **Topup** acting as the interface for external interactions.  

Payment gateway class will create a new **PaymentProcessor** based on processed Card type,  
This structure ensures that each card type's logic is contained within its respective PaymentProcessor.  

```
paymentProcessor interface {
	Charge(*card.SmartCard, route.TravelRoute)
	Topup(*card.SmartCard, int)
}
```

Topup can be as simple as incresing the balance, or can be complex depend on the card type.  

Charge flow typically follows these steps:
1. Check the route fare configuration from origin to destination.
2. Check peak-time or non-peak-time costs.
3. Check card's total spending to determine the maxDeduction for the trip,  
&emsp;&emsp;check 2 things, daily spending and weekly spending, compared to route fare caps.
4. Compare the fare cost to the maxDeduction. If the cost exceeds the maxDeduction,  
&emsp;&emsp;max deduciton used as cost to ensure the user is not overcharged
5. Deduct the trip's cost from card's balance.
6. Update the Card's FareSpending information.
