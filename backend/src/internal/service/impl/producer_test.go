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

//func TestNewProducerService(t *testing.T) {
//	type args struct {
//		producerRepo _interface.IProducerRepository
//		reserveRepo  _interface.IReserveRepository
//	}
//	test_data := []struct {
//		name string
//		args args
//		want _interface.IProducerService
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range test_data {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := NewProducerService(tt.args.producerRepo, tt.args.reserveRepo); !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("NewProducerService() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}

func TestProducerService_Add(t *testing.T) {
	type fields struct {
		producerRepo _interface.IProducerRepository
		reserveRepo  _interface.IReserveRepository
	}
	type args struct {
		request *dto.AddProducerRequest
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
				request: &dto.AddProducerRequest{
					Name:      "1",
					StudioId:  1,
					StartHour: 1,
					EndHour:   2,
				},
			},
			wantErr: false,
		},
		{
			name: "test_neg_01",
			args: args{
				request: &dto.AddProducerRequest{
					Name:      "1",
					StudioId:  1,
					StartHour: 1,
					EndHour:   1,
				},
			},
			wantErr: true,
		},
		{
			name: "test_neg_02",
			args: args{
				request: &dto.AddProducerRequest{
					Name:      "1",
					StudioId:  0,
					StartHour: 1,
					EndHour:   2,
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		prodRepo := new(mocks.IProducerRepository)
		//ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		//defer cancel()
		prodRepo.On("Add", context.Background(), tt.args.request).Return(nil)
		s := ProducerService{
			producerRepo: prodRepo,
			//reserveRepo:  tt.fields.reserveRepo,
		}
		t.Run(tt.name, func(t *testing.T) {

			if err := s.Add(tt.args.request); (err != nil) != tt.wantErr {
				t.Errorf("Add() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestProducerService_Delete(t *testing.T) {
	type fields struct {
		producerRepo _interface.IProducerRepository
		reserveRepo  _interface.IReserveRepository
	}
	type args struct {
		request *dto.DeleteProducerRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "test_pos_01",
			args: args{&dto.DeleteProducerRequest{
				Id: 1,
			}},
			wantErr: false,
		},
		{
			name: "test_neg_01",
			args: args{
				&dto.DeleteProducerRequest{
					Id: -1,
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		//new(pool.IProducerRepository).On("Delete", context.Background(), tt.args.request).Return(nil)
		prodRepo := new(mocks.IProducerRepository)
		prodRepo.On("Delete", context.Background(), tt.args.request).Return(nil)
		t.Run(tt.name, func(t *testing.T) {
			s := ProducerService{
				producerRepo: prodRepo,
				reserveRepo:  tt.fields.reserveRepo,
			}
			if err := s.Delete(tt.args.request); (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestProducerService_Get(t *testing.T) {
	type fields struct {
		producerRepo _interface.IProducerRepository
		reserveRepo  _interface.IReserveRepository
	}
	type args struct {
		request *dto.GetProducerRequest
	}
	tests := []struct {
		name         string
		fields       fields
		args         args
		wantProducer *model.Producer
		wantErr      bool
	}{
		{
			name: "test_pos_01",
			args: args{
				&dto.GetProducerRequest{
					Id: 1,
				},
			},
			wantErr: false,
			wantProducer: &model.Producer{
				Id:        1,
				Name:      "1",
				StudioId:  1,
				StartHour: 1,
				EndHour:   2,
			},
		},
		{
			name: "test_neg_01",
			args: args{
				&dto.GetProducerRequest{
					Id: 0,
				},
			},
			wantErr:      true,
			wantProducer: nil,
		},
	}
	for _, tt := range tests {
		prodRepo := new(mocks.IProducerRepository)
		prodRepo.On("Get", context.Background(), tt.args.request).Return(&model.Producer{
			Id:        1,
			Name:      "1",
			StudioId:  1,
			StartHour: 1,
			EndHour:   2,
		}, nil)
		t.Run(tt.name, func(t *testing.T) {
			s := ProducerService{
				producerRepo: prodRepo,
				//reserveRepo:  tt.fields.reserveRepo,
			}
			gotProducer, err := s.Get(tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotProducer, tt.wantProducer) {
				t.Errorf("Get() gotProducer = %v, want %v", gotProducer, tt.wantProducer)
			}
		})
	}
}

func TestProducerService_GetByStudio(t *testing.T) {
	type fields struct {
		producerRepo _interface.IProducerRepository
		reserveRepo  _interface.IReserveRepository
	}
	type args struct {
		request *dto.GetProducerByStudioRequest
	}
	tests := []struct {
		name          string
		fields        fields
		args          args
		wantProducers []*model.Producer
		wantErr       bool
	}{
		{
			name: "test_pos_01",
			args: args{
				&dto.GetProducerByStudioRequest{
					StudioId: 1,
				},
			},
			wantErr: false,
			wantProducers: []*model.Producer{
				{
					Id:        1,
					Name:      "1",
					StudioId:  1,
					StartHour: 1,
					EndHour:   2,
				},
			},
		},
		{
			name: "test_pos_02",
			args: args{
				&dto.GetProducerByStudioRequest{
					StudioId: 2,
				},
			},
			wantErr: false,
			wantProducers: []*model.Producer{
				{
					Id:        1,
					Name:      "1",
					StudioId:  2,
					StartHour: 1,
					EndHour:   2,
				},
				{
					Id:        2,
					Name:      "2",
					StudioId:  2,
					StartHour: 1,
					EndHour:   2,
				},
			},
		},
		{
			name: "test_neg_01",
			args: args{
				&dto.GetProducerByStudioRequest{
					StudioId: 0,
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		prodRepo := new(mocks.IProducerRepository)

		prodRepo.On("GetByStudio", context.Background(), &dto.GetProducerByStudioRequest{
			StudioId: 1,
		}).Return([]*model.Producer{
			{
				Id:        1,
				Name:      "1",
				StudioId:  1,
				StartHour: 1,
				EndHour:   2,
			},
		}, nil)

		prodRepo.On("GetByStudio", context.Background(), &dto.GetProducerByStudioRequest{
			StudioId: 2,
		}).Return([]*model.Producer{
			{
				Id:        1,
				Name:      "1",
				StudioId:  2,
				StartHour: 1,
				EndHour:   2,
			},
			{
				Id:        2,
				Name:      "2",
				StudioId:  2,
				StartHour: 1,
				EndHour:   2,
			},
		}, nil)

		t.Run(tt.name, func(t *testing.T) {
			s := ProducerService{
				producerRepo: prodRepo,
				//reserveRepo:  tt.fields.reserveRepo,
			}
			gotProducers, err := s.GetByStudio(tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetByStudio() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotProducers, tt.wantProducers) {
				t.Errorf("GetByStudio() gotProducers = %v, want %v", gotProducers, tt.wantProducers)
			}
		})
	}
}

func TestProducerService_Update(t *testing.T) {
	type fields struct {
		producerRepo _interface.IProducerRepository
		reserveRepo  _interface.IReserveRepository
	}
	type args struct {
		request *dto.UpdateProducerRequest
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
				&dto.UpdateProducerRequest{
					Id:        1,
					Name:      "1",
					StudioId:  1,
					StartHour: 1,
					EndHour:   2,
				},
			},
			wantErr: false,
		},
		{
			name: "test_neg_01",
			args: args{
				&dto.UpdateProducerRequest{
					Id:        1,
					Name:      "1",
					StudioId:  1,
					StartHour: 2,
					EndHour:   2,
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		prodRepo := new(mocks.IProducerRepository)
		reserveRepo := new(mocks.IReserveRepository)

		reserveRepo.On("IsProducerReserve", context.Background(), &dto.IsProducerReserveRequest{
			ProducerId: 1,
		}).Return(false, nil)

		prodRepo.On("Update", context.Background(),
			tt.args.request,
		).Return(nil)
		t.Run(tt.name, func(t *testing.T) {
			s := ProducerService{
				producerRepo: prodRepo,
				reserveRepo:  reserveRepo,
			}
			if err := s.Update(tt.args.request); (err != nil) != tt.wantErr {
				t.Errorf("Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
