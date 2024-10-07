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

//func TestNewRoomService(t *testing.T) {
//	type args struct {
//		roomRepo _interface.IRoomRepository
//		reserveRepo  _interface.IReserveRepository
//	}
//	test_data := []struct {
//		name string
//		args args
//		want _interface.IRoomService
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range test_data {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := NewRoomService(tt.args.roomRepo, tt.args.reserveRepo); !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("NewRoomService() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}

func TestRoomService_Add(t *testing.T) {
	type fields struct {
		roomRepo    _interface.IRoomRepository
		reserveRepo _interface.IReserveRepository
	}
	type args struct {
		request *dto.AddRoomRequest
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
				request: &dto.AddRoomRequest{
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
				request: &dto.AddRoomRequest{
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
				request: &dto.AddRoomRequest{
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
		prodRepo := new(mocks.IRoomRepository)
		//ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		//defer cancel()
		prodRepo.On("Add", context.Background(), tt.args.request).Return(nil)
		s := RoomService{
			roomRepo: prodRepo,
			//reserveRepo:  tt.fields.reserveRepo,
		}
		t.Run(tt.name, func(t *testing.T) {

			if err := s.Add(tt.args.request); (err != nil) != tt.wantErr {
				t.Errorf("Add() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRoomService_Delete(t *testing.T) {
	type fields struct {
		roomRepo    _interface.IRoomRepository
		reserveRepo _interface.IReserveRepository
	}
	type args struct {
		request *dto.DeleteRoomRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "test_pos_01",
			args: args{&dto.DeleteRoomRequest{
				Id: 1,
			}},
			wantErr: false,
		},
		{
			name: "test_neg_01",
			args: args{
				&dto.DeleteRoomRequest{
					Id: -1,
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		//new(pool.IRoomRepository).On("Delete", context.Background(), tt.args.request).Return(nil)
		prodRepo := new(mocks.IRoomRepository)
		prodRepo.On("Delete", context.Background(), tt.args.request).Return(nil)
		t.Run(tt.name, func(t *testing.T) {
			s := RoomService{
				roomRepo:    prodRepo,
				reserveRepo: tt.fields.reserveRepo,
			}
			if err := s.Delete(tt.args.request); (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRoomService_Get(t *testing.T) {
	type fields struct {
		roomRepo    _interface.IRoomRepository
		reserveRepo _interface.IReserveRepository
	}
	type args struct {
		request *dto.GetRoomRequest
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantRoom *model.Room
		wantErr  bool
	}{
		{
			name: "test_pos_01",
			args: args{
				&dto.GetRoomRequest{
					Id: 1,
				},
			},
			wantErr: false,
			wantRoom: &model.Room{
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
				&dto.GetRoomRequest{
					Id: 0,
				},
			},
			wantErr:  true,
			wantRoom: nil,
		},
	}
	for _, tt := range tests {
		prodRepo := new(mocks.IRoomRepository)
		prodRepo.On("Get", context.Background(), tt.args.request).Return(&model.Room{
			Id:        1,
			Name:      "1",
			StudioId:  1,
			StartHour: 1,
			EndHour:   2,
		}, nil)
		t.Run(tt.name, func(t *testing.T) {
			s := RoomService{
				roomRepo: prodRepo,
				//reserveRepo:  tt.fields.reserveRepo,
			}
			gotRoom, err := s.Get(tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRoom, tt.wantRoom) {
				t.Errorf("Get() gotRoom = %v, want %v", gotRoom, tt.wantRoom)
			}
		})
	}
}

func TestRoomService_GetByStudio(t *testing.T) {
	type fields struct {
		roomRepo    _interface.IRoomRepository
		reserveRepo _interface.IReserveRepository
	}
	type args struct {
		request *dto.GetRoomByStudioRequest
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		wantRooms []*model.Room
		wantErr   bool
	}{
		{
			name: "test_pos_01",
			args: args{
				&dto.GetRoomByStudioRequest{
					StudioId: 1,
				},
			},
			wantErr: false,
			wantRooms: []*model.Room{
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
				&dto.GetRoomByStudioRequest{
					StudioId: 2,
				},
			},
			wantErr: false,
			wantRooms: []*model.Room{
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
				&dto.GetRoomByStudioRequest{
					StudioId: 0,
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		prodRepo := new(mocks.IRoomRepository)

		prodRepo.On("GetByStudio", context.Background(), &dto.GetRoomByStudioRequest{
			StudioId: 1,
		}).Return([]*model.Room{
			{
				Id:        1,
				Name:      "1",
				StudioId:  1,
				StartHour: 1,
				EndHour:   2,
			},
		}, nil)

		prodRepo.On("GetByStudio", context.Background(), &dto.GetRoomByStudioRequest{
			StudioId: 2,
		}).Return([]*model.Room{
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
			s := RoomService{
				roomRepo: prodRepo,
				//reserveRepo:  tt.fields.reserveRepo,
			}
			gotRooms, err := s.GetByStudio(tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetByStudio() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRooms, tt.wantRooms) {
				t.Errorf("GetByStudio() gotRooms = %v, want %v", gotRooms, tt.wantRooms)
			}
		})
	}
}

func TestRoomService_Update(t *testing.T) {
	type fields struct {
		roomRepo    _interface.IRoomRepository
		reserveRepo _interface.IReserveRepository
	}
	type args struct {
		request *dto.UpdateRoomRequest
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
				&dto.UpdateRoomRequest{
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
				&dto.UpdateRoomRequest{
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
		prodRepo := new(mocks.IRoomRepository)
		reserveRepo := new(mocks.IReserveRepository)

		reserveRepo.On("IsRoomReserve", context.Background(), &dto.IsRoomReserveRequest{
			RoomId: 1,
		}).Return(false, nil)

		prodRepo.On("Update", context.Background(),
			tt.args.request,
		).Return(nil)
		t.Run(tt.name, func(t *testing.T) {
			s := RoomService{
				roomRepo:    prodRepo,
				reserveRepo: reserveRepo,
			}
			if err := s.Update(tt.args.request); (err != nil) != tt.wantErr {
				t.Errorf("Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
