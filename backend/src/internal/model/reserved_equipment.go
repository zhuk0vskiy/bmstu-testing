package model

import "time"

type ReservedEquipment struct {
	ReserveId   int64
	EquipmentId int64
	Type        int64
	StudioId    int64
	StartTime   time.Time
	EndTime     time.Time
}
