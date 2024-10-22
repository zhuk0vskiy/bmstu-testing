package time_parser

import (
	"backend/src/internal/model"
	serviceImpl "backend/src/internal/service/impl"
	"fmt"
	"strconv"
	"strings"
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

func ToTime(date string, startHour string, endHour string) (time *model.TimeInterval, err error) {
	startHourInt, err := strconv.Atoi(startHour)
	if err != nil {
		return nil, fmt.Errorf("failed to parse start hour: %v", err)
	}

	endHourInt, err := strconv.Atoi(endHour)
	if err != nil {
		return nil, fmt.Errorf("failed to parse end hour: %v", err)
	}

	if startHourInt >= endHourInt {
		return nil, fmt.Errorf("start hour must be less than end hour")
	}

	startTimeString := strings.TrimSpace(date) + " " +
		strings.TrimSpace(startHour) + ":00:00"

	endTimeString := strings.TrimSpace(date) + " " +
		strings.TrimSpace(endHour) + ":00:00"

	//_, err = time_parser.StringToDate(startTimeString)
	startTime, err := StringToDate(startTimeString)
	if err != nil {
		return nil, err
	}
	//_, err = time_parser.StringToDate(endTimeString)
	endTime, err := StringToDate(endTimeString)
	if err != nil {
		return nil, err
	}

	return serviceImpl.NewTimeInterval(startTime, endTime), nil
}
