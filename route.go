package main

type Route string

const (
	GreenLine Route = "Green"
	RedLine         = "Red"
)

type TravelFares struct {
	Departure    Route
	Destination  Route
	StandardCost int
	PeakCost     int
	DailyCap     int
	WeeklyCap    int
}

var FaresConfig = []TravelFares{
	{GreenLine, GreenLine, 2, 1, 8, 55},
	{RedLine, RedLine, 3, 2, 12, 70},
	{GreenLine, RedLine, 4, 3, 15, 90},
	{RedLine, GreenLine, 3, 2, 15, 90},
}
