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

//func TestNewInstrumentalistService(t *testing.T) {
//	type args struct {
//		instrumentalistRepo _interface.IInstrumentalistRepository
//		reserveRepo  _interface.IReserveRepository
//	}
//	test_data := []struct {
//		name string
//		args args
//		want _interface.IInstrumentalistService
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range test_data {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := NewInstrumentalistService(tt.args.instrumentalistRepo, tt.args.reserveRepo); !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("NewInstrumentalistService() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}

func TestInstrumentalistService_Add(t *testing.T) {
	type fields struct {
		instrumentalistRepo _interface.IInstrumentalistRepository
		reserveRepo         _interface.IReserveRepository
	}
	type args struct {
		request *dto.AddInstrumentalistRequest
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
				request: &dto.AddInstrumentalistRequest{
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
				request: &dto.AddInstrumentalistRequest{
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
				request: &dto.AddInstrumentalistRequest{
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
		prodRepo := new(mocks.IInstrumentalistRepository)
		//ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		//defer cancel()
		prodRepo.On("Add", context.Background(), tt.args.request).Return(nil)
		s := InstrumentalistService{
			instrumentalistRepo: prodRepo,
			//reserveRepo:  tt.fields.reserveRepo,
		}
		t.Run(tt.name, func(t *testing.T) {

			if err := s.Add(tt.args.request); (err != nil) != tt.wantErr {
				t.Errorf("Add() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestInstrumentalistService_Delete(t *testing.T) {
	type fields struct {
		instrumentalistRepo _interface.IInstrumentalistRepository
		reserveRepo         _interface.IReserveRepository
	}
	type args struct {
		request *dto.DeleteInstrumentalistRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "test_pos_01",
			args: args{&dto.DeleteInstrumentalistRequest{
				Id: 1,
			}},
			wantErr: false,
		},
		{
			name: "test_neg_01",
			args: args{
				&dto.DeleteInstrumentalistRequest{
					Id: -1,
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		//new(pool.IInstrumentalistRepository).On("Delete", context.Background(), tt.args.request).Return(nil)
		prodRepo := new(mocks.IInstrumentalistRepository)
		prodRepo.On("Delete", context.Background(), tt.args.request).Return(nil)
		t.Run(tt.name, func(t *testing.T) {
			s := InstrumentalistService{
				instrumentalistRepo: prodRepo,
				reserveRepo:         tt.fields.reserveRepo,
			}
			if err := s.Delete(tt.args.request); (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestInstrumentalistService_Get(t *testing.T) {
	type fields struct {
		instrumentalistRepo _interface.IInstrumentalistRepository
		reserveRepo         _interface.IReserveRepository
	}
	type args struct {
		request *dto.GetInstrumentalistRequest
	}
	tests := []struct {
		name                string
		fields              fields
		args                args
		wantInstrumentalist *model.Instrumentalist
		wantErr             bool
	}{
		{
			name: "test_pos_01",
			args: args{
				&dto.GetInstrumentalistRequest{
					Id: 1,
				},
			},
			wantErr: false,
			wantInstrumentalist: &model.Instrumentalist{
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
				&dto.GetInstrumentalistRequest{
					Id: 0,
				},
			},
			wantErr:             true,
			wantInstrumentalist: nil,
		},
	}
	for _, tt := range tests {
		prodRepo := new(mocks.IInstrumentalistRepository)
		prodRepo.On("Get", context.Background(), tt.args.request).Return(&model.Instrumentalist{
			Id:        1,
			Name:      "1",
			StudioId:  1,
			StartHour: 1,
			EndHour:   2,
		}, nil)
		t.Run(tt.name, func(t *testing.T) {
			s := InstrumentalistService{
				instrumentalistRepo: prodRepo,
				//reserveRepo:  tt.fields.reserveRepo,
			}
			gotInstrumentalist, err := s.Get(tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotInstrumentalist, tt.wantInstrumentalist) {
				t.Errorf("Get() gotInstrumentalist = %v, want %v", gotInstrumentalist, tt.wantInstrumentalist)
			}
		})
	}
}

func TestInstrumentalistService_GetByStudio(t *testing.T) {
	type fields struct {
		instrumentalistRepo _interface.IInstrumentalistRepository
		reserveRepo         _interface.IReserveRepository
	}
	type args struct {
		request *dto.GetInstrumentalistByStudioRequest
	}
	tests := []struct {
		name                 string
		fields               fields
		args                 args
		wantInstrumentalists []*model.Instrumentalist
		wantErr              bool
	}{
		{
			name: "test_pos_01",
			args: args{
				&dto.GetInstrumentalistByStudioRequest{
					StudioId: 1,
				},
			},
			wantErr: false,
			wantInstrumentalists: []*model.Instrumentalist{
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
				&dto.GetInstrumentalistByStudioRequest{
					StudioId: 2,
				},
			},
			wantErr: false,
			wantInstrumentalists: []*model.Instrumentalist{
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
				&dto.GetInstrumentalistByStudioRequest{
					StudioId: 0,
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		prodRepo := new(mocks.IInstrumentalistRepository)

		prodRepo.On("GetByStudio", context.Background(), &dto.GetInstrumentalistByStudioRequest{
			StudioId: 1,
		}).Return([]*model.Instrumentalist{
			{
				Id:        1,
				Name:      "1",
				StudioId:  1,
				StartHour: 1,
				EndHour:   2,
			},
		}, nil)

		prodRepo.On("GetByStudio", context.Background(), &dto.GetInstrumentalistByStudioRequest{
			StudioId: 2,
		}).Return([]*model.Instrumentalist{
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
			s := InstrumentalistService{
				instrumentalistRepo: prodRepo,
				//reserveRepo:  tt.fields.reserveRepo,
			}
			gotInstrumentalists, err := s.GetByStudio(tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetByStudio() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotInstrumentalists, tt.wantInstrumentalists) {
				t.Errorf("GetByStudio() gotInstrumentalists = %v, want %v", gotInstrumentalists, tt.wantInstrumentalists)
			}
		})
	}
}

func TestInstrumentalistService_Update(t *testing.T) {
	type fields struct {
		instrumentalistRepo _interface.IInstrumentalistRepository
		reserveRepo         _interface.IReserveRepository
	}
	type args struct {
		request *dto.UpdateInstrumentalistRequest
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
				&dto.UpdateInstrumentalistRequest{
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
				&dto.UpdateInstrumentalistRequest{
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
		prodRepo := new(mocks.IInstrumentalistRepository)
		reserveRepo := new(mocks.IReserveRepository)

		reserveRepo.On("IsInstrumentalistReserve", context.Background(), &dto.IsInstrumentalistReserveRequest{
			InstrumentalistId: tt.args.request.Id,
		}).Return(false, nil)

		prodRepo.On("Update", context.Background(),
			tt.args.request,
		).Return(nil)
		t.Run(tt.name, func(t *testing.T) {
			s := InstrumentalistService{
				instrumentalistRepo: prodRepo,
				reserveRepo:         reserveRepo,
			}
			if err := s.Update(tt.args.request); (err != nil) != tt.wantErr {
				t.Errorf("Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
