package model

type Reserve struct {
	Id                int64
	UserId            int64
	RoomId            int64
	ProducerId        int64
	InstrumentalistId int64
	TimeInterval      *TimeInterval
}

type ReserveExt struct {
	Id                int64         `json:"id"`
	UserId            int64         `json:"user_id"`
	RoomId            int64         `json:"room_id"`
	ProducerId        int64         `json:"producer_id"`
	InstrumentalistId int64         `json:"instrumentalist_id"`
	EquipmentsId      []int64       `json:"equipments_id"`
	TimeInterval      *TimeInterval `json:"time_interval"`
}
