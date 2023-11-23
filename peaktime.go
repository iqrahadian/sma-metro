package main

const TIME_FORMAT string = "15:04:05"

type Weekday int

const (
	Sunday Weekday = iota + 1
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

type PeakTimeConfig struct {
	FromDay   Weekday
	ToDay     Weekday
	StartHour string
	EndHour   string
}

var PeakTimeConfigArr = []PeakTimeConfig{
	{Monday, Friday, "08:00", "10:00"},
	{Monday, Friday, "16:30", "19:00"},
	{Saturday, Saturday, "10:00", "14:00"},
	{Saturday, Saturday, "18:00", "23:00"},
	{Sunday, Sunday, "18:00", "23:00"},
}
