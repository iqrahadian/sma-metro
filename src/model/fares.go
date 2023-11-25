package model

type TravelFaresConfig struct {
	Departure    string
	Destination  string
	StandardCost int
	PeakCost     int
	DailyCap     int
	WeeklyCap    int
}
