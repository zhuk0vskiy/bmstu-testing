package utils

import (
	"backend/src/internal/model"
)

type UserAuthBuilder struct {
	model.User
}

func NewUserAuthBuilder() *UserAuthBuilder {
	return &UserAuthBuilder{}
}

func (u *UserAuthBuilder) WithId(id int64) *UserAuthBuilder {
	u.Id = id
	return u
}

func (u *UserAuthBuilder) WithLogin(login string) *UserAuthBuilder {
	u.Login = login
	return u
}

func (u *UserAuthBuilder) WithPassword(password string) *UserAuthBuilder {
	u.Password = password
	return u
}

func (u *UserAuthBuilder) WithRole(role string) *UserAuthBuilder {
	u.Role = role
	return u
}

func (u *UserAuthBuilder) WithFirstName(firstName string) *UserAuthBuilder {
	u.FirstName = firstName
	return u
}

func (u *UserAuthBuilder) WithSecondName(secondName string) *UserAuthBuilder {
	u.SecondName = secondName
	return u
}

func (u *UserAuthBuilder) WithThirdName(thirdName string) *UserAuthBuilder {
	u.ThirdName = thirdName
	return u
}

func (u *UserAuthBuilder) ToDto() *model.User {
	return &model.User{
		Id:         u.Id,
		Login:      u.Login,
		Password:   u.Password,
		Role:       u.Role,
		FirstName:  u.FirstName,
		SecondName: u.SecondName,
		ThirdName:  u.ThirdName,
	}
}

type ReserveBuilder struct {
	model.Reserve
}

func NewReserveBuilder() *ReserveBuilder {
	return &ReserveBuilder{}
}

func (u *ReserveBuilder) WithId(id int64) *ReserveBuilder {
	u.Id = id
	return u
}

func (u *ReserveBuilder) WithProducerId(id int64) *ReserveBuilder {
	u.ProducerId = id
	return u
}

func (u *ReserveBuilder) WithUserId(id int64) *ReserveBuilder {
	u.UserId = id
	return u
}

func (u *ReserveBuilder) WithInstrumentalistId(id int64) *ReserveBuilder {
	u.InstrumentalistId = id
	return u
}

func (u *ReserveBuilder) WithTimeInterval(timeInterval *model.TimeInterval) *ReserveBuilder {
	u.TimeInterval = timeInterval
	return u
}

func (u *ReserveBuilder) ToDto() *model.Reserve {
	return &model.Reserve{
		Id:                u.Id,
		RoomId:            u.RoomId,
		UserId:            u.UserId,
		ProducerId:        u.ProducerId,
		InstrumentalistId: u.InstrumentalistId,
		TimeInterval:      u.TimeInterval,
	}
}

type StudioBuilder struct {
	model.Studio
}

func NewStudioBuilder() *StudioBuilder {
	return &StudioBuilder{}
}

func (u *StudioBuilder) WithId(id int64) *StudioBuilder {
	u.Id = id
	return u
}

func (u *StudioBuilder) WithName(name string) *StudioBuilder {
	u.Name = name
	return u
}

func (u *StudioBuilder) ToDto() *model.Studio {
	return &model.Studio{
		Id:   u.Id,
		Name: u.Name,
	}
}

type RoomBuilder struct {
	model.Room
}

func NewRoomBuilder() *RoomBuilder {
	return &RoomBuilder{}
}

func (u *RoomBuilder) WithId(id int64) *RoomBuilder {
	u.Id = id
	return u
}

func (u *RoomBuilder) WithName(name string) *RoomBuilder {
	u.Name = name
	return u
}

func (u *RoomBuilder) WithStudioId(id int64) *RoomBuilder {
	u.StudioId = id
	return u
}

func (u *RoomBuilder) WithStartHour(hour int64) *RoomBuilder {
	u.StartHour = hour
	return u
}

func (u *RoomBuilder) WithEndHour(hour int64) *RoomBuilder {
	u.EndHour = hour
	return u
}
func (u *RoomBuilder) ToDto() *model.Room {
	return &model.Room{
		Id:        u.Id,
		Name:      u.Name,
		StudioId:  u.StudioId,
		StartHour: u.StartHour,
		EndHour:   u.EndHour,
	}
}

type ProducerBuilder struct {
	model.Producer
}

func NewProducerBuilder() *ProducerBuilder {
	return &ProducerBuilder{}
}

func (u *ProducerBuilder) WithId(id int64) *ProducerBuilder {
	u.Id = id
	return u
}

func (u *ProducerBuilder) WithName(name string) *ProducerBuilder {
	u.Name = name
	return u
}

func (u *ProducerBuilder) WithStudioId(id int64) *ProducerBuilder {
	u.Id = id
	return u
}

func (u *ProducerBuilder) WithStartHour(hour int64) *ProducerBuilder {
	u.StartHour = hour
	return u
}

func (u *ProducerBuilder) WithEndHour(hour int64) *ProducerBuilder {
	u.EndHour = hour
	return u
}
func (u *ProducerBuilder) ToDto() *model.Producer {
	return &model.Producer{
		Id:        u.Id,
		Name:      u.Name,
		StudioId:  u.StudioId,
		StartHour: u.StartHour,
		EndHour:   u.EndHour,
	}
}

type InstrumentalistBuilder struct {
	model.Instrumentalist
}

func NewInstrumentalistBuilder() *InstrumentalistBuilder {
	return &InstrumentalistBuilder{}
}

func (u *InstrumentalistBuilder) WithId(id int64) *InstrumentalistBuilder {
	u.Id = id
	return u
}

func (u *InstrumentalistBuilder) WithName(name string) *InstrumentalistBuilder {
	u.Name = name
	return u
}

func (u *InstrumentalistBuilder) WithStudioId(id int64) *InstrumentalistBuilder {
	u.Id = id
	return u
}

func (u *InstrumentalistBuilder) WithStartHour(hour int64) *InstrumentalistBuilder {
	u.StartHour = hour
	return u
}

func (u *InstrumentalistBuilder) WithEndHour(hour int64) *InstrumentalistBuilder {
	u.EndHour = hour
	return u
}
func (u *InstrumentalistBuilder) ToDto() *model.Instrumentalist {
	return &model.Instrumentalist{
		Id:        u.Id,
		Name:      u.Name,
		StudioId:  u.StudioId,
		StartHour: u.StartHour,
		EndHour:   u.EndHour,
	}
}

type EquipmentBuilder struct {
	model.Equipment
}

func NewEquipmentBuilder() *EquipmentBuilder {
	return &EquipmentBuilder{}
}

func (u *EquipmentBuilder) WithId(id int64) *EquipmentBuilder {
	u.Id = id
	return u
}

func (u *EquipmentBuilder) WithName(name string) *EquipmentBuilder {
	u.Name = name
	return u
}

func (u *EquipmentBuilder) WithEquipmentType(equipmentType int64) *EquipmentBuilder {
	u.EquipmentType = equipmentType
	return u
}

func (u *EquipmentBuilder) WithStudioId(id int64) *EquipmentBuilder {
	u.StudioId = id
	return u
}

func (u *EquipmentBuilder) ToDto() *model.Equipment {
	return &model.Equipment{
		Id:            u.Id,
		Name:          u.Name,
		EquipmentType: u.EquipmentType,
		StudioId:      u.StudioId,
	}
}
