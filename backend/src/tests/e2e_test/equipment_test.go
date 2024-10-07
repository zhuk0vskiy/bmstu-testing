package tests

import (
	"backend/src/internal/model"
	"backend/src/internal/model/dto"
	"backend/src/internal/repository/impl/postgresql"
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"reflect"
	"testing"
	"time"
)

//func TestEquipmentRepository_GetFullTimeFreeByStudioAndType(t *testing.T) {
//	type fields struct {
//		db *pgxpool.Pool
//	}
//	type args struct {
//		ctx     context.Context
//		request *dto.GetEquipmentFullTimeFreeByStudioAndTypeRequest
//	}
//	tests := []struct {
//		name string
//		//fields         fields
//		args           args
//		wantEquipments []*model.Equipment
//		wantErr        bool
//	}{
//		{
//			name: "test_pos_01",
//			args: args{
//				ctx: context.Background(),
//				request: &dto.GetEquipmentFullTimeFreeByStudioAndTypeRequest{
//					StudioId: 1,
//					Type:     1,
//				},
//			},
//			wantEquipments: []*model.Equipment{
//				&model.Equipment{
//					Id:            1,
//					Name:          "first",
//					EquipmentType: 1,
//					StudioId:      1,
//				},
//				&model.Equipment{
//					Id:            2,
//					Name:          "second",
//					EquipmentType: 1,
//					StudioId:      1,
//				},
//			},
//			wantErr: false,
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			r := postgresql.NewEquipmentRepository(testDbInstance)
//			var elapsed time.Duration
//			count := 1000
//			gotEquipments, err := r.GetFullTimeFreeByStudioAndType(tt.args.ctx, tt.args.request)
//			for i := 0; i < count; i++ {
//				startTime := time.Now()
//				gotEquipments, err = r.GetFullTimeFreeByStudioAndType(tt.args.ctx, tt.args.request)
//				elapsed = elapsed + time.Since(startTime)
//			}
//			fmt.Println("Время: ", elapsed/1000)
//
//			if (err != nil) != tt.wantErr {
//				t.Errorf("GetFullTimeFreeByStudioAndType() error = %v, wantErr %v", err, tt.wantErr)
//				return
//			}
//			if !reflect.DeepEqual(gotEquipments, tt.wantEquipments) {
//				t.Errorf("GetFullTimeFreeByStudioAndType() gotEquipments = %v, want %v", gotEquipments, tt.wantEquipments)
//			}
//		})
//	}
//}

func TestEquipmentRepository_GetNotFullTimeFreeByStudioAndType(t *testing.T) {
	type fields struct {
		db *pgxpool.Pool
	}
	type args struct {
		ctx     context.Context
		request *dto.GetEquipmentNotFullTimeFreeByStudioAndTypeRequest
	}
	tests := []struct {
		name                  string
		fields                fields
		args                  args
		wantEquipmentsAndTime []*dto.EquipmentAndTime
		wantErr               bool
	}{
		{
			name: "test_pos_01",
			args: args{
				ctx: context.Background(),
				request: &dto.GetEquipmentNotFullTimeFreeByStudioAndTypeRequest{
					StudioId: 1,
					Type:     1,
					TimeInterval: &model.TimeInterval{
						StartTime: time.Date(2024, 4, 22, 12, 0, 0, 0, time.UTC),
						EndTime:   time.Date(2024, 4, 22, 14, 0, 0, 0, time.UTC),
					},
				},
			},
			wantEquipmentsAndTime: []*dto.EquipmentAndTime{
				{
					&model.Equipment{
						Id:            3,
						Name:          "third",
						EquipmentType: 1,
						StudioId:      1,
					},
					&model.TimeInterval{
						StartTime: time.Date(2022, 6, 30, 15, 0, 0, 0, time.UTC),
						EndTime:   time.Date(2022, 6, 30, 17, 0, 0, 0, time.UTC),
					},
				},
				{
					&model.Equipment{
						Id:            4,
						Name:          "fourth",
						EquipmentType: 1,
						StudioId:      1,
					},
					&model.TimeInterval{
						StartTime: time.Date(2022, 6, 30, 15, 0, 0, 0, time.UTC),
						EndTime:   time.Date(2022, 6, 30, 17, 0, 0, 0, time.UTC)},
				},
			},

			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := postgresql.NewEquipmentRepository(testDbInstance)
			var elapsed time.Duration
			count := 1000
			gotEquipmentsAndTime, err := r.GetNotFullTimeFreeByStudioAndType(tt.args.ctx, tt.args.request)
			for i := 0; i < count; i++ {
				startTime := time.Now()
				gotEquipmentsAndTime, err = r.GetNotFullTimeFreeByStudioAndType(tt.args.ctx, tt.args.request)
				elapsed = elapsed + time.Since(startTime)
			}
			fmt.Println("Время: ", elapsed/1000)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetNotFullTimeFreeByStudioAndType() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotEquipmentsAndTime, tt.wantEquipmentsAndTime) {
				t.Errorf("GetNotFullTimeFreeByStudioAndType() gotEquipmentsAndTime = %v, want %v", gotEquipmentsAndTime, tt.wantEquipmentsAndTime)
			}
		})
	}
}
