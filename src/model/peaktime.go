package model

import "time"

type PeakTimeConfig struct {
	FromDay   time.Weekday
	ToDay     time.Weekday
	StartHour string
	EndHour   string
}

type PeaktimeHour struct {
	Start time.Time
	End   time.Time
}
