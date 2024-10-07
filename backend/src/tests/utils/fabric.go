package utils

import (
	"backend/src/internal/model"
	"backend/src/internal/model/dto"
	"time"
)

type SignInFabric struct {
	Id int64
}

func (f SignInFabric) CorrectUserSignIn() (*dto.AddUserRequest, *dto.SignInRequest) {
	return &dto.AddUserRequest{
			Login:      "1",
			Password:   "hashedPass123",
			Role:       "client",
			FirstName:  "3",
			SecondName: "4",
			ThirdName:  "5",
		},
		&dto.SignInRequest{
			Login:      "1",
			Password:   "2",
			FirstName:  "3",
			SecondName: "4",
			ThirdName:  "5",
		}
}

func (f SignInFabric) IncorrectUserSignIn() (*dto.AddUserRequest, *dto.SignInRequest) {
	return &dto.AddUserRequest{
			Login:      "",
			Password:   "hashedPass123",
			Role:       "client",
			FirstName:  "3",
			SecondName: "4",
			ThirdName:  "5",
		},
		&dto.SignInRequest{
			Login:      "",
			Password:   "2",
			FirstName:  "3",
			SecondName: "4",
			ThirdName:  "5",
		}
}

type LogInFabric struct {
	Id int64
}

func (f LogInFabric) CorrectUserLogIn() *dto.LogInRequest {
	return &dto.LogInRequest{
		Login:    "1",
		Password: "2",
	}
}

type InstrumentalistFabric struct {
	Id int64
}

func (f InstrumentalistFabric) CorrectInstrumentalistAdd() *dto.AddInstrumentalistRequest {
	return &dto.AddInstrumentalistRequest{
		Name:      "1",
		StudioId:  1,
		StartHour: 1,
		EndHour:   2,
	}
}

func (f InstrumentalistFabric) IncorrectInstrumentalistAdd() *dto.AddInstrumentalistRequest {
	return &dto.AddInstrumentalistRequest{
		Name:      "",
		StudioId:  1,
		StartHour: 1,
		EndHour:   2,
	}
}

func (f InstrumentalistFabric) CorrectInstrumentalistDelete() *dto.DeleteInstrumentalistRequest {
	return &dto.DeleteInstrumentalistRequest{
		Id: f.Id,
	}
}

func (f InstrumentalistFabric) IncorrectInstrumentalistDelete() *dto.DeleteInstrumentalistRequest {
	return &dto.DeleteInstrumentalistRequest{
		Id: f.Id,
	}
}

func (f InstrumentalistFabric) CorrectInstrumentalistGet() *dto.GetInstrumentalistRequest {
	return &dto.GetInstrumentalistRequest{
		Id: f.Id,
	}
}

func (f InstrumentalistFabric) IncorrectInstrumentalistGet() *dto.GetInstrumentalistRequest {
	return &dto.GetInstrumentalistRequest{
		Id: f.Id,
	}
}

type RoomFabric struct {
	Id int64
}

func (f RoomFabric) CorrectRoomAdd() *dto.AddRoomRequest {
	return &dto.AddRoomRequest{
		Name:      "1",
		StudioId:  1,
		StartHour: 1,
		EndHour:   2,
	}
}

func (f RoomFabric) IncorrectRoomAdd() *dto.AddRoomRequest {
	return &dto.AddRoomRequest{
		Name:      "",
		StudioId:  1,
		StartHour: 1,
		EndHour:   2,
	}
}

func (f RoomFabric) CorrectRoomDelete() *dto.DeleteRoomRequest {
	return &dto.DeleteRoomRequest{
		Id: f.Id,
	}
}

func (f RoomFabric) IncorrectRoomDelete() *dto.DeleteRoomRequest {
	return &dto.DeleteRoomRequest{
		Id: f.Id,
	}
}

type ReserveFabric struct {
	Id int64
}

func (f ReserveFabric) CorrectReserveAdd() *dto.AddReserveRequest {
	return &dto.AddReserveRequest{
		UserId:            1,
		RoomId:            1,
		ProducerId:        1,
		InstrumentalistId: 1,
		TimeInterval: &model.TimeInterval{
			StartTime: time.Date(2024, 4, 14, 12, 0, 0, 0, time.UTC),
			EndTime:   time.Date(2024, 4, 14, 13, 0, 0, 0, time.UTC),
		},
		EquipmentId: nil,
	}
}

func (f ReserveFabric) IncorrectReserveAdd() *dto.AddReserveRequest {
	return &dto.AddReserveRequest{
		UserId:            0,
		RoomId:            1,
		ProducerId:        1,
		InstrumentalistId: 1,
		TimeInterval: &model.TimeInterval{
			StartTime: time.Date(2024, 4, 14, 12, 0, 0, 0, time.UTC),
			EndTime:   time.Date(2024, 4, 14, 13, 0, 0, 0, time.UTC),
		},
		EquipmentId: nil,
	}
}

func (f ReserveFabric) CorrectReserveGetAll() *dto.GetAllReserveRequest {
	return &dto.GetAllReserveRequest{}
}

type StudioFabric struct {
	Id int64
}

func (f StudioFabric) CorrectStudioAdd() *dto.AddStudioRequest {
	return &dto.AddStudioRequest{
		Name: "1",
	}
}

func (f StudioFabric) IncorrectStudioAdd() *dto.AddStudioRequest {
	return &dto.AddStudioRequest{
		Name: "",
	}
}

func (f StudioFabric) CorrectStudioDelete() *dto.DeleteStudioRequest {
	return &dto.DeleteStudioRequest{
		Id: f.Id,
	}
}

func (f StudioFabric) IncorrectStudioDelete() *dto.DeleteStudioRequest {
	return &dto.DeleteStudioRequest{
		Id: f.Id,
	}
}

type ProducerFabric struct {
	Id int64
}

func (f ProducerFabric) CorrectProducerGet() *dto.GetProducerRequest {
	return &dto.GetProducerRequest{
		Id: f.Id,
	}
}

func (f ProducerFabric) IncorrectProducerGet() *dto.GetProducerRequest {
	return &dto.GetProducerRequest{
		Id: f.Id,
	}
}

func (f ProducerFabric) CorrectProducerAdd() *dto.AddProducerRequest {
	return &dto.AddProducerRequest{
		Name:      "1",
		StudioId:  1,
		StartHour: 1,
		EndHour:   2,
	}
}

func (f ProducerFabric) IncorrectProducerAdd() *dto.AddProducerRequest {
	return &dto.AddProducerRequest{
		Name:      "",
		StudioId:  1,
		StartHour: 1,
		EndHour:   2,
	}
}

func (f ProducerFabric) CorrectProducerDelete() *dto.DeleteProducerRequest {
	return &dto.DeleteProducerRequest{
		Id: f.Id,
	}
}

func (f ProducerFabric) IncorrectProducerDelete() *dto.DeleteProducerRequest {
	return &dto.DeleteProducerRequest{
		Id: f.Id,
	}
}

type UserFabric struct {
	Id int64
}

func (f UserFabric) UserDelete() *dto.DeleteUserRequest {
	return &dto.DeleteUserRequest{
		Id: f.Id,
	}
}

func (f UserFabric) CorrectUserAdd() *dto.AddUserRequest {
	return &dto.AddUserRequest{
		Login:      "1",
		Password:   "hashedPass123",
		Role:       "client",
		FirstName:  "3",
		SecondName: "4",
		ThirdName:  "5",
	}
}

func (f UserFabric) IncorrectUserAdd() *dto.AddUserRequest {
	return &dto.AddUserRequest{
		Login:      "",
		Password:   "hashedPass123",
		Role:       "client",
		FirstName:  "3",
		SecondName: "4",
		ThirdName:  "5",
	}
}
