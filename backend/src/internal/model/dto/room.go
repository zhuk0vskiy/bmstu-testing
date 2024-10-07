package dto

type GetRoomRequest struct {
	Id int64
}

type GetRoomAllRequest struct {
}

type GetRoomByStudioRequest struct {
	StudioId int64
}

type AddRoomRequest struct {
	Name      string
	StudioId  int64
	StartHour int64
	EndHour   int64
}

type UpdateRoomRequest struct {
	Id        int64
	Name      string
	StudioId  int64
	StartHour int64
	EndHour   int64
}

type DeleteRoomRequest struct {
	Id int64
}
