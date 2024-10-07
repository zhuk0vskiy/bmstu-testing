package impl

import (
	"backend/src/internal/model"
	"backend/src/internal/model/dto"
	_interface "backend/src/internal/repository/interface"
	"backend/src/internal/repository/interface/mocks"
	//_interface "backend/src/internal/service_test/interface"
	"context"
	"testing"
	"time"
)

//func TestNewReserveService(t *testing.T) {
//	type args struct {
//		reserveRepo _interface.IReserveRepository
//	}
//	test_data := []struct {
//		name string
//		args args
//		want _interface.IReserveService
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range test_data {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := NewReserveService(tt.args.reserveRepo); !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("NewReserveService() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}

func TestReserveService_Add(t *testing.T) {
	type fields struct {
		reserveRepo _interface.IReserveRepository
	}
	type args struct {
		request *dto.AddReserveRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "test_pos_01",
			args: args{

				request: &dto.AddReserveRequest{
					UserId:            1,
					RoomId:            1,
					EquipmentId:       []int64{1, 2, 3},
					ProducerId:        1,
					InstrumentalistId: 1,
					TimeInterval: &model.TimeInterval{
						StartTime: time.Date(2024, 4, 14, 12, 0, 0, 0, time.UTC),
						EndTime:   time.Date(2024, 4, 14, 13, 0, 0, 0, time.UTC),
					},
				},
			},
			wantErr: false,
		},
		{
			name: "test_neg_01",
			args: args{

				request: &dto.AddReserveRequest{
					UserId:            0,
					RoomId:            1,
					EquipmentId:       []int64{1, 2, 3},
					ProducerId:        1,
					InstrumentalistId: 1,
					TimeInterval: &model.TimeInterval{
						StartTime: time.Date(2024, 4, 14, 12, 0, 0, 0, time.UTC),
						EndTime:   time.Date(2024, 4, 14, 12, 0, 0, 0, time.UTC),
					},
				},
			},
			wantErr: true,
		},
		{
			name: "test_neg_02",
			args: args{

				request: &dto.AddReserveRequest{
					UserId:            1,
					RoomId:            0,
					EquipmentId:       []int64{1, 2, 3},
					ProducerId:        1,
					InstrumentalistId: 1,
					TimeInterval: &model.TimeInterval{
						StartTime: time.Date(2024, 4, 14, 12, 0, 0, 0, time.UTC),
						EndTime:   time.Date(2024, 4, 14, 12, 0, 0, 0, time.UTC),
					},
				},
			},
			wantErr: true,
		},
		{
			name: "test_neg_03",
			args: args{

				request: &dto.AddReserveRequest{
					UserId:            1,
					RoomId:            1,
					EquipmentId:       nil,
					ProducerId:        1,
					InstrumentalistId: 1,
					TimeInterval: &model.TimeInterval{
						StartTime: time.Date(2024, 4, 14, 12, 0, 0, 0, time.UTC),
						EndTime:   time.Date(2024, 4, 14, 12, 0, 0, 0, time.UTC),
					},
				},
			},
			wantErr: true,
		},
		{
			name: "test_neg_04",
			args: args{

				request: &dto.AddReserveRequest{
					UserId:            1,
					RoomId:            1,
					EquipmentId:       []int64{1, 2, 3},
					ProducerId:        0,
					InstrumentalistId: 1,
					TimeInterval: &model.TimeInterval{
						StartTime: time.Date(2024, 4, 14, 12, 0, 0, 0, time.UTC),
						EndTime:   time.Date(2024, 4, 14, 12, 0, 0, 0, time.UTC),
					},
				},
			},
			wantErr: true,
		},
		{
			name: "test_neg_05",
			args: args{

				request: &dto.AddReserveRequest{
					UserId:            1,
					RoomId:            1,
					EquipmentId:       []int64{1, 2, 3},
					ProducerId:        1,
					InstrumentalistId: 0,
					TimeInterval: &model.TimeInterval{
						StartTime: time.Date(2024, 4, 14, 12, 0, 0, 0, time.UTC),
						EndTime:   time.Date(2024, 4, 14, 12, 0, 0, 0, time.UTC),
					},
				},
			},
			wantErr: true,
		},
		{
			name: "test_neg_06",
			args: args{
				request: &dto.AddReserveRequest{
					UserId:            1,
					RoomId:            1,
					EquipmentId:       []int64{1, 2, 3},
					ProducerId:        1,
					InstrumentalistId: 1,
					TimeInterval: &model.TimeInterval{
						StartTime: time.Date(2024, 4, 14, 12, 0, 0, 0, time.UTC),
						EndTime:   time.Date(2024, 4, 14, 12, 0, 0, 0, time.UTC),
					},
				},
			},
			wantErr: true,
		},
	}
	reserveRepo := new(mocks.IReserveRepository)
	for _, tt := range tests {
		reserveRepo.On("Add", context.Background(), tt.args.request).Return(nil)

		t.Run(tt.name, func(t *testing.T) {
			s := ReserveService{
				reserveRepo: reserveRepo,
			}
			if err := s.Add(tt.args.request); (err != nil) != tt.wantErr {
				t.Errorf("Add() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestReserveService_Delete(t *testing.T) {
	type fields struct {
		reserveRepo _interface.IReserveRepository
	}
	type args struct {
		request *dto.DeleteReserveRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "test_pos_01",
			args: args{
				request: &dto.DeleteReserveRequest{
					Id: 1,
				},
			},
			wantErr: false,
		},
		{
			name: "test_neg_01",
			args: args{
				request: &dto.DeleteReserveRequest{
					Id: 0,
				},
			},
			wantErr: true,
		},
	}
	reserveRepo := new(mocks.IReserveRepository)
	for _, tt := range tests {
		reserveRepo.On("Delete", context.Background(), tt.args.request).Return(nil)
		t.Run(tt.name, func(t *testing.T) {
			s := ReserveService{
				reserveRepo: reserveRepo,
			}
			if err := s.Delete(tt.args.request); (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
