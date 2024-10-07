package tests

import (
	"backend/src/internal/model"
	"backend/src/internal/model/dto"
	"backend/src/internal/repository/impl/postgresql"
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"reflect"
	"testing"
	"time"
)

func TestReserveRepository_Add(t *testing.T) {
	type fields struct {
		db *pgxpool.Pool
	}
	type args struct {
		ctx     context.Context
		request *dto.AddReserveRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		//{
		//	name: "test_pos_01",
		//	args: args{
		//		ctx: context.Background(),
		//		request: &dto.AddReserveRequest{
		//			UserId:            33,
		//			RoomId:            33,
		//			ProducerId:        33,
		//			InstrumentalistId: 33,
		//			TimeInterval: &model.TimeInterval{
		//				StartTime: time.Date(2000, 9, 10, 13, 0, 0, 0, time.UTC),
		//				EndTime:   time.Date(2000, 9, 10, 14, 0, 0, 0, time.UTC),
		//			},
		//			EquipmentId: []int64{1, 2, 3},
		//		},
		//	},
		//	wantErr: false,
		//},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := postgresql.NewReserveRepository(testDbInstance)
			if err := r.Add(tt.args.ctx, tt.args.request); (err != nil) != tt.wantErr {
				t.Errorf("Add() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
	//time.Sleep(20000 * time.Second)
}

func TestReserveRepository_GetByRoomId(t *testing.T) {
	type fields struct {
		db *pgxpool.Pool
	}
	type args struct {
		ctx     context.Context
		request *dto.GetReserveByRoomIdRequest
	}
	tests := []struct {
		name         string
		fields       fields
		args         args
		wantReserves []*model.Reserve
		wantErr      bool
	}{
		{
			name: "test_pos_01",
			args: args{
				ctx: context.Background(),
				request: &dto.GetReserveByRoomIdRequest{
					RoomId: 1,
				},
			},
			wantErr: false,
			wantReserves: []*model.Reserve{
				&model.Reserve{
					Id:                1,
					UserId:            1,
					RoomId:            1,
					ProducerId:        1,
					InstrumentalistId: 1,
					TimeInterval: &model.TimeInterval{
						StartTime: time.Date(2022, 06, 30, 15, 0, 0, 0, time.UTC),
						EndTime:   time.Date(2022, 06, 30, 17, 0, 0, 0, time.UTC),
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := postgresql.NewReserveRepository(testDbInstance)

			gotReserves, err := r.GetByRoomId(tt.args.ctx, tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetByRoomId() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotReserves, tt.wantReserves) {
				t.Errorf("GetByRoomId() gotReserves = %v, want %v", gotReserves, tt.wantReserves)
			}
		})
	}
}

func TestReserveRepository_IsRoomReserve(t *testing.T) {
	type fields struct {
		db *pgxpool.Pool
	}
	type args struct {
		ctx     context.Context
		request *dto.IsRoomReserveRequest
	}
	tests := []struct {
		name          string
		fields        fields
		args          args
		wantIsReserve bool
		wantErr       bool
	}{
		{
			name: "test_pos_01",
			args: args{
				ctx: context.Background(),
				request: &dto.IsRoomReserveRequest{
					RoomId: 1,
				},
			},
			wantErr:       false,
			wantIsReserve: true,
		},
		{
			name: "test_pod_02",
			args: args{
				ctx: context.Background(),
				request: &dto.IsRoomReserveRequest{
					RoomId: 2,
				},
			},
			wantErr:       false,
			wantIsReserve: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := postgresql.NewReserveRepository(testDbInstance)
			gotIsReserve, err := r.IsRoomReserve(tt.args.ctx, tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("IsRoomReserve() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotIsReserve != tt.wantIsReserve {
				t.Errorf("IsRoomReserve() gotIsReserve = %v, want %v", gotIsReserve, tt.wantIsReserve)
			}
		})
	}
}

func TestReserveRepository_IsEquipmentReserve(t *testing.T) {
	type fields struct {
		db *pgxpool.Pool
	}
	type args struct {
		ctx     context.Context
		request *dto.IsEquipmentReserveRequest
	}
	tests := []struct {
		name          string
		fields        fields
		args          args
		wantIsReserve bool
		wantErr       bool
	}{
		{
			name: "test_pos_01",
			args: args{
				ctx: context.Background(),
				request: &dto.IsEquipmentReserveRequest{
					EquipmentId: 3,
				},
			},
			wantErr:       false,
			wantIsReserve: true,
		},
		//{
		//	name: "test_pos_02",
		//	args: args{
		//		ctx: context.Background(),
		//		request: &dto.IsEquipmentReserveRequest{
		//			EquipmentId: 1,
		//		},
		//	},
		//	wantErr:       false,
		//	wantIsReserve: false,
		//},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := postgresql.NewReserveRepository(testDbInstance)
			gotIsReserve, err := r.IsEquipmentReserve(tt.args.ctx, tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("IsEquipmentReserve() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotIsReserve != tt.wantIsReserve {
				t.Errorf("IsEquipmentReserve() gotIsReserve = %v, want %v", gotIsReserve, tt.wantIsReserve)
			}
		})
	}
}
