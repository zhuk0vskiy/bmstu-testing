package model

const (
	Client = 1
	Admin  = 2
)

type User struct {
	Id         int64  `json:"id"`
	Login      string `json:"login"`
	Password   string `json:"password"`
	Role       string `json:"role"`
	FirstName  string `json:"first_name"`
	SecondName string `json:"second_name"`
	ThirdName  string `json:"third_name"`
}
