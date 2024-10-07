package impl

import (
	"backend/src/internal/model"
	"backend/src/internal/model/dto"
	_interface "backend/src/internal/repository/interface"
	"backend/src/internal/repository/interface/mocks"
	"context"
	//_interface "backend/src/internal/service_test/interface"
	"reflect"
	"testing"
)

//func TestNewEquipmentService(t *testing.T) {
//	type args struct {
//		equipmentRepo _interface.IEquipmentRepository
//		reserveRepo  _interface.IReserveRepository
//	}
//	test_data := []struct {
//		name string
//		args args
//		want _interface.IEquipmentService
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range test_data {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := NewEquipmentService(tt.args.equipmentRepo, tt.args.reserveRepo); !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("NewEquipmentService() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}

func TestEquipmentService_Add(t *testing.T) {
	type fields struct {
		equipmentRepo _interface.IEquipmentRepository
		reserveRepo   _interface.IReserveRepository
	}
	type args struct {
		request *dto.AddEquipmentRequest
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
				request: &dto.AddEquipmentRequest{
					Name:     "1",
					StudioId: 1,
					Type:     1,
				},
			},
			wantErr: false,
		},
		{
			name: "test_neg_01",
			args: args{
				request: &dto.AddEquipmentRequest{
					Name:     "1",
					StudioId: 1,
					Type:     0,
				},
			},
			wantErr: true,
		},
		{
			name: "test_neg_02",
			args: args{
				request: &dto.AddEquipmentRequest{
					Name:     "1",
					StudioId: 0,
					Type:     1,
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		prodRepo := new(mocks.IEquipmentRepository)
		//ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		//defer cancel()
		prodRepo.On("Add", context.Background(), tt.args.request).Return(nil)
		s := EquipmentService{
			equipmentRepo: prodRepo,
			//reserveRepo:  tt.fields.reserveRepo,
		}
		t.Run(tt.name, func(t *testing.T) {

			if err := s.Add(tt.args.request); (err != nil) != tt.wantErr {
				t.Errorf("Add() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestEquipmentService_Delete(t *testing.T) {
	type fields struct {
		equipmentRepo _interface.IEquipmentRepository
		reserveRepo   _interface.IReserveRepository
	}
	type args struct {
		request *dto.DeleteEquipmentRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "test_pos_01",
			args: args{&dto.DeleteEquipmentRequest{
				Id: 1,
			}},
			wantErr: false,
		},
		{
			name: "test_neg_01",
			args: args{
				&dto.DeleteEquipmentRequest{
					Id: -1,
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		//new(pool.IEquipmentRepository).On("Delete", context.Background(), tt.args.request).Return(nil)
		prodRepo := new(mocks.IEquipmentRepository)
		prodRepo.On("Delete", context.Background(), tt.args.request).Return(nil)
		t.Run(tt.name, func(t *testing.T) {
			s := EquipmentService{
				equipmentRepo: prodRepo,
				reserveRepo:   tt.fields.reserveRepo,
			}
			if err := s.Delete(tt.args.request); (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestEquipmentService_Get(t *testing.T) {
	type fields struct {
		equipmentRepo _interface.IEquipmentRepository
		reserveRepo   _interface.IReserveRepository
	}
	type args struct {
		request *dto.GetEquipmentRequest
	}
	tests := []struct {
		name          string
		fields        fields
		args          args
		wantEquipment *model.Equipment
		wantErr       bool
	}{
		{
			name: "test_pos_01",
			args: args{
				&dto.GetEquipmentRequest{
					Id: 1,
				},
			},
			wantErr: false,
			wantEquipment: &model.Equipment{
				Id:            1,
				Name:          "1",
				StudioId:      1,
				EquipmentType: 1,
			},
		},
		{
			name: "test_neg_01",
			args: args{
				&dto.GetEquipmentRequest{
					Id: 0,
				},
			},
			wantErr:       true,
			wantEquipment: nil,
		},
	}
	for _, tt := range tests {
		prodRepo := new(mocks.IEquipmentRepository)
		prodRepo.On("Get", context.Background(), tt.args.request).Return(&model.Equipment{
			Id:            1,
			Name:          "1",
			StudioId:      1,
			EquipmentType: 1,
		}, nil)
		t.Run(tt.name, func(t *testing.T) {
			s := EquipmentService{
				equipmentRepo: prodRepo,
				//reserveRepo:  tt.fields.reserveRepo,
			}
			gotEquipment, err := s.Get(tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotEquipment, tt.wantEquipment) {
				t.Errorf("Get() gotEquipment = %v, want %v", gotEquipment, tt.wantEquipment)
			}
		})
	}
}

func TestEquipmentService_GetByStudio(t *testing.T) {
	type fields struct {
		equipmentRepo _interface.IEquipmentRepository
		reserveRepo   _interface.IReserveRepository
	}
	type args struct {
		request *dto.GetEquipmentByStudioRequest
	}
	tests := []struct {
		name           string
		fields         fields
		args           args
		wantEquipments []*model.Equipment
		wantErr        bool
	}{
		{
			name: "test_pos_01",
			args: args{
				&dto.GetEquipmentByStudioRequest{
					StudioId: 1,
				},
			},
			wantErr: false,
			wantEquipments: []*model.Equipment{
				{
					Id:            1,
					Name:          "1",
					StudioId:      1,
					EquipmentType: 1,
				},
			},
		},
		{
			name: "test_pos_02",
			args: args{
				&dto.GetEquipmentByStudioRequest{
					StudioId: 2,
				},
			},
			wantErr: false,
			wantEquipments: []*model.Equipment{
				{
					Id:            1,
					Name:          "1",
					StudioId:      2,
					EquipmentType: 1,
				},
				{
					Id:            2,
					Name:          "2",
					StudioId:      2,
					EquipmentType: 2,
				},
			},
		},
		{
			name: "test_neg_01",
			args: args{
				&dto.GetEquipmentByStudioRequest{
					StudioId: 0,
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		prodRepo := new(mocks.IEquipmentRepository)

		prodRepo.On("GetByStudio", context.Background(), &dto.GetEquipmentByStudioRequest{
			StudioId: 1,
		}).Return([]*model.Equipment{
			{
				Id:            1,
				Name:          "1",
				StudioId:      1,
				EquipmentType: 1,
			},
		}, nil)

		prodRepo.On("GetByStudio", context.Background(), &dto.GetEquipmentByStudioRequest{
			StudioId: 2,
		}).Return([]*model.Equipment{
			{
				Id:            1,
				Name:          "1",
				StudioId:      2,
				EquipmentType: 1,
			},
			{
				Id:            2,
				Name:          "2",
				StudioId:      2,
				EquipmentType: 2,
			},
		}, nil)

		t.Run(tt.name, func(t *testing.T) {
			s := EquipmentService{
				equipmentRepo: prodRepo,
				//reserveRepo:  tt.fields.reserveRepo,
			}
			gotEquipments, err := s.GetByStudio(tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetByStudio() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotEquipments, tt.wantEquipments) {
				t.Errorf("GetByStudio() gotEquipments = %v, want %v", gotEquipments, tt.wantEquipments)
			}
		})
	}
}

func TestEquipmentService_Update(t *testing.T) {
	type fields struct {
		equipmentRepo _interface.IEquipmentRepository
		reserveRepo   _interface.IReserveRepository
	}
	type args struct {
		request *dto.UpdateEquipmentRequest
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
				&dto.UpdateEquipmentRequest{
					Id:       1,
					Name:     "1",
					StudioId: 1,
					Type:     1,
				},
			},
			wantErr: false,
		},
		{
			name: "test_neg_01",
			args: args{
				&dto.UpdateEquipmentRequest{
					Id:       1,
					Name:     "1",
					StudioId: 1,
					Type:     0,
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		prodRepo := new(mocks.IEquipmentRepository)
		reserveRepo := new(mocks.IReserveRepository)

		reserveRepo.On("IsEquipmentReserve", context.Background(), &dto.IsEquipmentReserveRequest{
			EquipmentId: tt.args.request.Id,
		}).Return(false, nil)

		prodRepo.On("Update", context.Background(),
			tt.args.request,
		).Return(nil)
		t.Run(tt.name, func(t *testing.T) {
			s := EquipmentService{
				equipmentRepo: prodRepo,
				reserveRepo:   reserveRepo,
			}
			if err := s.Update(tt.args.request); (err != nil) != tt.wantErr {
				t.Errorf("Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
