package dto

type GetUserRequest struct {
	Id int64
}

type UpdateUserRequest struct {
	Id         int64
	Login      string
	Password   string
	FirstName  string
	SecondName string
	ThirdName  string
}

type DeleteUserRequest struct {
	Id int64
}

type GetUserReservesRequest struct {
	Id int64
}

type AddUserRequest struct {
	Login      string
	Password   string
	Role       string
	FirstName  string
	SecondName string
	ThirdName  string
}
type GetUserByLoginRequest struct {
	Login string
}
