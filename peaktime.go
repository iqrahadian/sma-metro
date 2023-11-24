package main

import "time"

const (
	TIME_FORMAT      string = "15:04:05"
	DATE_TIME_FORMAT string = "2006-01-02T15:04:05"
)

// type Weekday int

// const (
// 	Sunday Weekday = iota
// 	Monday
// 	Tuesday
// 	Wednesday
// 	Thursday
// 	Friday
// 	Saturday
// )

type PeaktimeHour struct {
	Start time.Time
	End   time.Time
}

type PeakTimeConfig struct {
	FromDay   time.Weekday
	ToDay     time.Weekday
	StartHour string
	EndHour   string
}

var PeakTimeConfigArr = []PeakTimeConfig{
	{time.Monday, time.Friday, "08:00", "10:00"},
	{time.Monday, time.Friday, "16:30", "19:00"},
	{time.Saturday, time.Saturday, "10:00", "14:00"},
	{time.Saturday, time.Saturday, "18:00", "23:00"},
	{time.Sunday, time.Sunday, "18:00", "23:00"},
}
