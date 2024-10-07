package impl

import (
	"backend/src/internal/model"
	"backend/src/internal/model/dto"
	_interface "backend/src/internal/repository/interface"
	"backend/src/internal/repository/interface/mocks"
	"backend/src/pkg/base"
	"context"
	"github.com/stretchr/testify/mock"
	"reflect"
	"testing"
	"time"
)

func TestUserService_Delete(t *testing.T) {
	type fields struct {
		userRepo    _interface.IUserRepository
		reserveRepo _interface.IReserveRepository
		crypto      base.IHashCrypto
	}
	type args struct {
		request *dto.DeleteUserRequest
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
				request: &dto.DeleteUserRequest{
					Id: 1,
				},
			},
			wantErr: false,
		},
		{
			name: "test_neg_01",
			args: args{
				request: &dto.DeleteUserRequest{
					Id: 0,
				},
			},
			wantErr: true,
		},
	}
	userRepo := new(mocks.IUserRepository)

	for _, tt := range tests {

		userRepo.On("Delete", context.Background(), tt.args.request).Return(nil)
		t.Run(tt.name, func(t *testing.T) {
			s := UserService{
				userRepo:    userRepo,
				reserveRepo: tt.fields.reserveRepo,
				crypto:      tt.fields.crypto,
			}
			if err := s.Delete(tt.args.request); (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUserService_Get(t *testing.T) {
	type fields struct {
		userRepo    _interface.IUserRepository
		reserveRepo _interface.IReserveRepository
		crypto      base.IHashCrypto
	}
	type args struct {
		request *dto.GetUserRequest
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantUser *model.User
		wantErr  bool
	}{
		{
			name: "test_pos_01",
			args: args{
				request: &dto.GetUserRequest{
					Id: 1,
				},
			},
			wantUser: &model.User{
				Id:         1,
				Login:      "1",
				Password:   "1",
				FirstName:  "1",
				SecondName: "1",
				ThirdName:  "1",
			},
		},
		{
			name: "test_neg_01",
			args: args{
				request: &dto.GetUserRequest{
					Id: 0,
				},
			},
			wantUser: nil,
			wantErr:  true,
		},
	}
	userRepo := new(mocks.IUserRepository)
	for _, tt := range tests {
		userRepo.On("Get", context.Background(), tt.args.request).Return(tt.wantUser, nil)
		t.Run(tt.name, func(t *testing.T) {
			s := UserService{
				userRepo:    userRepo,
				reserveRepo: tt.fields.reserveRepo,
				crypto:      tt.fields.crypto,
			}
			gotUser, err := s.Get(tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotUser, tt.wantUser) {
				t.Errorf("Get() gotUser = %v, want %v", gotUser, tt.wantUser)
			}
		})
	}
}

func TestUserService_GetReserves(t *testing.T) {
	type fields struct {
		userRepo    _interface.IUserRepository
		reserveRepo _interface.IReserveRepository
		crypto      base.IHashCrypto
	}
	type args struct {
		request *dto.GetUserReservesRequest
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
				request: &dto.GetUserReservesRequest{
					Id: 1,
				},
			},
			wantReserves: []*model.Reserve{
				&model.Reserve{
					Id:                1,
					UserId:            1,
					RoomId:            1,
					ProducerId:        1,
					InstrumentalistId: 1,
					TimeInterval: &model.TimeInterval{
						StartTime: time.Date(2024, 4, 14, 12, 00, 00, 00, time.UTC),
						EndTime:   time.Date(2024, 4, 14, 12, 00, 00, 00, time.UTC),
					},
				},
			},
			wantErr: false,
		},
		{
			name: "test_neg_01",
			args: args{
				request: &dto.GetUserReservesRequest{
					Id: 0,
				},
			},
			wantReserves: nil,
			wantErr:      true,
		},
	}

	reserveRepo := new(mocks.IReserveRepository)

	for _, tt := range tests {
		reserveRepo.On("GetUserReserves", context.Background(), tt.args.request).Return(tt.wantReserves, nil)

		t.Run(tt.name, func(t *testing.T) {
			s := UserService{
				userRepo:    tt.fields.userRepo,
				reserveRepo: reserveRepo,
				crypto:      tt.fields.crypto,
			}
			gotReserves, err := s.GetReserves(tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetReserves() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotReserves, tt.wantReserves) {
				t.Errorf("GetReserves() gotReserves = %v, want %v", gotReserves, tt.wantReserves)
			}
		})
	}
}

func TestUserService_Update(t *testing.T) {
	type fields struct {
		userRepo    _interface.IUserRepository
		reserveRepo _interface.IReserveRepository
		crypto      base.IHashCrypto
	}
	type args struct {
		request *dto.UpdateUserRequest
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
				request: &dto.UpdateUserRequest{
					Id:         1,
					Login:      "1",
					Password:   "Hash1234",
					FirstName:  "1",
					SecondName: "1",
					ThirdName:  "1",
				},
			},
		},
		{
			name: "test_neg_01",
			args: args{
				request: &dto.UpdateUserRequest{
					Id:         0,
					Login:      "1",
					Password:   "1",
					FirstName:  "1",
					SecondName: "1",
					ThirdName:  "1",
				},
			},
			wantErr: true,
		},
		{
			name: "test_neg_02",
			args: args{
				request: &dto.UpdateUserRequest{
					Id:         1,
					Login:      "",
					Password:   "1",
					FirstName:  "1",
					SecondName: "1",
					ThirdName:  "1",
				},
			},
			wantErr: true,
		},
		{
			name: "test_neg_03",
			args: args{
				request: &dto.UpdateUserRequest{
					Id:         1,
					Login:      "1",
					Password:   "",
					FirstName:  "1",
					SecondName: "1",
					ThirdName:  "1",
				},
			},
			wantErr: true,
		},
		{
			name: "test_neg_04",
			args: args{
				request: &dto.UpdateUserRequest{
					Id:         1,
					Login:      "1",
					Password:   "1",
					FirstName:  "",
					SecondName: "1",
					ThirdName:  "1",
				},
			},
			wantErr: true,
		},
		{
			name: "test_neg_05",
			args: args{
				request: &dto.UpdateUserRequest{
					Id:         1,
					Login:      "1",
					Password:   "1",
					FirstName:  "1",
					SecondName: "",
					ThirdName:  "1",
				},
			},
			wantErr: true,
		},
		{
			name: "test_neg_06",
			args: args{
				request: &dto.UpdateUserRequest{
					Id:         1,
					Login:      "1",
					Password:   "1",
					FirstName:  "1",
					SecondName: "1",
					ThirdName:  "",
				},
			},
			wantErr: true,
		},
	}

	userRepo := new(mocks.IUserRepository)
	cryptoRepo := new(mocks.IHashCrypto)

	for _, tt := range tests {
		cryptoRepo.On("GenerateHashPass", mock.Anything, mock.Anything).Return("Hash1234", nil)
		//tt.args.request.Password = "Hash1234"
		userRepo.On("Update", context.Background(), tt.args.request).Return(nil)
		t.Run(tt.name, func(t *testing.T) {
			s := UserService{
				userRepo:    userRepo,
				reserveRepo: tt.fields.reserveRepo,
				crypto:      cryptoRepo,
			}
			if err := s.Update(tt.args.request); (err != nil) != tt.wantErr {
				t.Errorf("Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
