package model

type Equipment struct {
	Id            int64
	Name          string
	EquipmentType int64
	StudioId      int64
}

// TODO: изменить на что-то типо enum
const (
	OutOfFirstEquipment = 0 // для удобства проверки в логике

	Microphones = 1
	Instruments = 2
	Headphones  = 3
	Monitors    = 4
	Cabels      = 5
	Stations    = 6

	OutOfLastEquipment = 7 // для удобства проверки в логике
)
