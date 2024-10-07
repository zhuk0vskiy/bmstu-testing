package model

const (
	Client = 1
	Admin  = 2
)

type User struct {
	Id         int64
	Login      string
	Password   string
	Role       string
	FirstName  string
	SecondName string
	ThirdName  string
}
