package time_parser

import (
	"fmt"
	"time"
)

func StringToDate(date string) (parsed time.Time, err error) {
	t, err := time.Parse("2006-01-02 15:04:05", date)
	if err != nil {
		return time.Time{}, fmt.Errorf("failed to parse date: %v", err)
	}
	//fmt.Println(2)
	//parsed = new(time.Time)
	parsed = time.Date(
		t.Year(),
		t.Month(),
		t.Day(),
		t.Hour(),
		t.Minute(),
		t.Second(),
		0,
		time.UTC,
	)
	//fmt.Println(parsed)

	return parsed, err
}
