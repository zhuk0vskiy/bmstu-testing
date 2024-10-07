package model

import "time"

type Time struct {
	Year  int
	Month int
	Day   int
	Hour  int
}

type TimeInterval struct {
	StartTime time.Time
	EndTime   time.Time
}
