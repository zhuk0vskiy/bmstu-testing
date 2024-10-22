package v1

import "backend/src/internal/model"

type ValidateResponse struct {
	Rooms            []*model.Room            `json:"rooms"`
	Producers        []*model.Producer        `json:"producers"`
	Instrumentalists []*model.Instrumentalist `json:"instrumentalists"`
	Equipments       *ValidateEquipment       `json:"equipments"`
}

type ValidateEquipment struct {
	Microphones []*model.Equipment `json:"microphones"`
	Guitars     []*model.Equipment `json:"guitars"`
}

type GetStudioResponse struct {
	Studio *model.Studio `json:"studio"`
}

type GetRoomResponse struct {
	Room *model.Room `json:"room"`
}

type GetProducerResponse struct {
	Producer *model.Producer `json:"producer"`
}

type GetInstrumentalistResponse struct {
	Instrumentalist *model.Instrumentalist `json:"instrumentalist"`
}

type GetEquipmentResponse struct {
	Equipment *model.Equipment `json:"equipment"`
}

type UserReservesResponse struct {
	Reserves []*model.ReserveExt `json:"reserves"`
}

type GetRoomsByStudioResponse struct {
	Rooms []*model.Room `json:"rooms"`
}

type GetProducersByStudioResponse struct {
	Producers []*model.Producer `json:"producers"`
}

type GetInstrumentalistsByStudioResponse struct {
	Instrumentalists []*model.Instrumentalist `json:"instrumentalists"`
}

type GetEquipmentsByStudioResponse struct {
	Equipments []*model.Equipment `json:"equipments"`
}
