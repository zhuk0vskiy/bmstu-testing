package model

type Reserve struct {
	Id                int64
	UserId            int64
	RoomId            int64
	ProducerId        int64
	InstrumentalistId int64
	TimeInterval      *TimeInterval
}
