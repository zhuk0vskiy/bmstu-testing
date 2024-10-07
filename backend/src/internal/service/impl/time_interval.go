package impl

import (
	"backend/src/internal/model"
	"backend/src/internal/model/dto"
	"fmt"
	"time"
)

func NewTime(request *dto.NewTimeRequest) (t time.Time, err error) {
	if request.Time.Month > 12 || request.Time.Month < 1 {
		return time.Time{}, fmt.Errorf("неправильное время: %w", err)
	}
	if request.Time.Day < 1 || request.Time.Day > 31 {
		return time.Time{}, fmt.Errorf("неправильное время: %w", err)
	}
	//time.Date(1,2,3,4,5,6,7,time.UTC())
	t = time.Date(
		request.Time.Year,
		time.Month(request.Time.Month),
		request.Time.Day,
		request.Time.Hour,
		0,
		0,
		0,
		time.UTC,
	)
	return t, nil
}

func NewTimeInterval(startTime time.Time, endTime time.Time) *model.TimeInterval {

	if startTime.Unix() >= endTime.Unix() {
		return nil
	}
	return &model.TimeInterval{
		StartTime: startTime,
		EndTime:   endTime,
	}
}

func isIntervalsIntersect(choosenInterval model.TimeInterval, reservedInterval model.TimeInterval) bool {
	choosenStartHour := int64(choosenInterval.StartTime.Hour())
	choosenEndHour := int64(choosenInterval.EndTime.Hour())

	reserveStartHour := int64(reservedInterval.StartTime.Hour())
	reserveEndHour := int64(reservedInterval.EndTime.Hour())
	ans := false
	//fmt.Println("	looking reserve:", reserve.Id)
	if ((choosenStartHour >= reserveStartHour && choosenStartHour < reserveEndHour) ||
		(choosenEndHour <= reserveEndHour && choosenEndHour > reserveStartHour) ||
		(choosenStartHour <= reserveStartHour && choosenEndHour >= reserveEndHour)) && choosenInterval.StartTime.Day() == reservedInterval.StartTime.Day() {
		ans = true
		//break
	}
	//fmt.Println(ans)
	return ans
}
