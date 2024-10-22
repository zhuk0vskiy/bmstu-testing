package model

type Producer struct {
	Id        int64  `json:"id"`
	Name      string `json:"name"`
	StudioId  int64  `json:"studio_id"`
	StartHour int64  `json:"start_hour"`
	EndHour   int64  `json:"end_hour"`
}
