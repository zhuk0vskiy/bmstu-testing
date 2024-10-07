package dto

type GetProducerRequest struct {
	Id int64
}

type GetProducerByStudioRequest struct {
	StudioId int64
}

type AddProducerRequest struct {
	Name      string
	StudioId  int64
	StartHour int64
	EndHour   int64
}

type UpdateProducerRequest struct {
	Id        int64
	Name      string
	StudioId  int64
	StartHour int64
	EndHour   int64
}

type DeleteProducerRequest struct {
	Id int64
}
