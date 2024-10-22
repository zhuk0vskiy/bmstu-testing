package v1

type LogInRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type SignInRequest struct {
	Login      string `json:"login"`
	Password   string `json:"password"`
	FirstName  string `json:"first_name"`
	SecondName string `json:"second_name"`
	ThirdName  string `json:"third_name"`
}

type GetUserReservesRequest struct {
}

type CreateReserveRequest struct {
	UserId            string `json:"user_id"`
	RoomID            string `json:"room_id"`
	ProducerId        string `json:"producer_id"`
	InstrumentalistId string `json:"instrumentalist_id"`
	Date              string `json:"date"`
	StartHour         string `json:"start_hour"`
	EndHour           string `json:"end_hour"`
}

type ValidationRequest struct {
	StudioId  int64  `json:"studio_id"`
	Date      string `json:"date"`
	StartHour string `json:"start_hour"`
	EndHour   string `json:"end_hour"`
}

type AddReserveRequest struct {
	//Id          int64  `json:"id"`
	UserId            int64   `json:"user_id"`
	RoomId            int64   `json:"room_id"`
	ProducerId        int64   `json:"producer_id"`
	InstrumentalistId int64   `json:"instrumentalist_id"`
	EquipmentsId      []int64 `json:"equipments_id"`
	Date              string  `json:"date"`
	StartHour         string  `json:"start_hour"`
	EndHour           string  `json:"end_hour"`
	//
}

type AddStudioRequest struct {
	Name string `json:"name"`
}

type AddRoomRequest struct {
	Name      string `json:"name"`
	StudioId  int64  `json:"studio_id"`
	StartHour int64  `json:"start_hour"`
	EndHour   int64  `json:"end_hour"`
}

type AddProducerRequest struct {
	Name      string `json:"name"`
	StudioId  int64  `json:"studio_id"`
	StartHour int64  `json:"start_hour"`
	EndHour   int64  `json:"end_hour"`
}

type AddInstrumentalistRequest struct {
	Name      string `json:"name"`
	StudioId  int64  `json:"studio_id"`
	StartHour int64  `json:"start_hour"`
	EndHour   int64  `json:"end_hour"`
}

type AddEquipmentRequest struct {
	Name     string `json:"name"`
	StudioId int64  `json:"studio_id"`
	TypeId   int64  `json:"type_id"`
}

type UpdateStudioRequest struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

type UpdateRoomRequest struct {
	Id        int64  `json:"id"`
	Name      string `json:"name"`
	StudioId  int64  `json:"studio_id"`
	StartHour int64  `json:"start_hour"`
	EndHour   int64  `json:"end_hour"`
}

type UpdateProducerRequest struct {
	Id        int64  `json:"id"`
	Name      string `json:"name"`
	StudioId  int64  `json:"studio_id"`
	StartHour int64  `json:"start_hour"`
	EndHour   int64  `json:"end_hour"`
}

type UpdateInstrumentalistRequest struct {
	Id        int64  `json:"id"`
	Name      string `json:"name"`
	StudioId  int64  `json:"studio_id"`
	StartHour int64  `json:"start_hour"`
	EndHour   int64  `json:"end_hour"`
}

type UpdateEquipmentRequest struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	StudioId int64  `json:"studio_id"`
	TypeId   int64  `json:"type_id"`
}
