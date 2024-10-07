package dto

type GetStudioRequest struct {
	Id int64
}

type GetStudioAllRequest struct {
}

type AddStudioRequest struct {
	Name string
}

type UpdateStudioRequest struct {
	Id   int64
	Name string
}

type DeleteStudioRequest struct {
	Id int64
}
