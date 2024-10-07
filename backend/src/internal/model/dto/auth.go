package dto

type SignInRequest struct {
	Login      string
	Password   string
	FirstName  string
	SecondName string
	ThirdName  string
}

type LogInRequest struct {
	Login    string
	Password string
	//Role     string
}
