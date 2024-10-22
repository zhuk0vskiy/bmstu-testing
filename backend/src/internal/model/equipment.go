package model

type Equipment struct {
	Id            int64  `json:"id"`
	Name          string `json:"name"`
	EquipmentType int64  `json:"equipment_type"`
	StudioId      int64  `json:"studio_id"`
}

// TODO: изменить на что-то типо enum
const (
	OutOfFirstEquipment = 0 // для удобства проверки в логике

	Microphones = 1
	Instruments = 2
	//Headphones  = 3
	//Monitors    = 4
	//Cabels      = 5
	//Stations    = 6

	OutOfLastEquipment = 3 // для удобства проверки в логике
)
