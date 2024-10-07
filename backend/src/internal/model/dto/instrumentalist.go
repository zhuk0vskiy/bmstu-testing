package dto

type GetInstrumentalistRequest struct {
	Id int64
}

type GetInstrumentalistByStudioRequest struct {
	StudioId int64
}

type AddInstrumentalistRequest struct {
	Name      string
	StudioId  int64
	StartHour int64
	EndHour   int64
}

type UpdateInstrumentalistRequest struct {
	Id        int64
	Name      string
	StudioId  int64
	StartHour int64
	EndHour   int64
}

type DeleteInstrumentalistRequest struct {
	Id int64
}
