package impl

import (
	"backend/src/internal/model"
	"backend/src/internal/model/dto"
	_interface "backend/src/internal/repository/interface"
	"backend/src/internal/repository/interface/mocks"
	"backend/src/pkg/base"
	"context"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestAuthService_LogIn(t *testing.T) {
	type fields struct {
		userRepo _interface.IUserRepository
		crypto   base.IHashCrypto
		jwtKey   string
	}
	type args struct {
		request *dto.LogInRequest
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		wantToken string
		wantErr   bool
	}{
		{
			name: "test_pos_01",
			args: args{
				request: &dto.LogInRequest{
					Login:    "log1234",
					Password: "1234",
				},
			},
			wantErr: true,
		},
	}

	userRepo := new(mocks.IUserRepository)
	crypto := new(mocks.IHashCrypto)

	userRepo.On("GetByLogin", context.Background(), &dto.GetUserByLoginRequest{
		Login: "log1234",
	}).Return(&model.User{
		Id:         1,
		Login:      "log1234",
		Password:   "hashed1234",
		Role:       "test",
		FirstName:  "1",
		SecondName: "2",
		ThirdName:  "3",
	}, nil)

	crypto.On("CheckPasswordHash", "1234", "hashed1234").Return(true)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := AuthService{
				userRepo: userRepo,
				crypto:   crypto,
				jwtKey:   tt.fields.jwtKey,
			}
			gotToken, err := s.LogIn(tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("LogIn() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotToken != tt.wantToken {
				t.Errorf("LogIn() gotToken = %v, want %v", gotToken, tt.wantToken)
			}
		})
	}
}

func TestAuthService_SignIn(t *testing.T) {
	type fields struct {
		userRepo _interface.IUserRepository
		crypto   base.IHashCrypto
		jwtKey   string
	}
	type args struct {
		request *dto.SignInRequest
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
				request: &dto.SignInRequest{
					Login:      "log1234",
					Password:   "1234",
					FirstName:  "1",
					SecondName: "2",
					ThirdName:  "3",
				},
			},
			wantErr: false,
		},
		{
			name: "test_neg_01",
			args: args{
				request: &dto.SignInRequest{
					Login:      "",
					Password:   "1234",
					FirstName:  "1",
					SecondName: "2",
					ThirdName:  "3",
				},
			},
			wantErr: true,
		},
		{
			name: "test_neg_02",
			args: args{
				request: &dto.SignInRequest{
					Login:      "",
					Password:   "1234",
					FirstName:  "1",
					SecondName: "2",
					ThirdName:  "",
				},
			},
			wantErr: true,
		},
	}

	cryptoRepo := new(mocks.IHashCrypto)
	userRepo := new(mocks.IUserRepository)

	cryptoRepo.On("GenerateHashPass", mock.Anything, mock.Anything).Return("Hash1234", nil)

	for _, tt := range tests {

		userRepo.On("Add", context.Background(), &dto.AddUserRequest{
			Login:      tt.args.request.Login,
			Password:   "Hash1234",
			FirstName:  tt.args.request.FirstName,
			SecondName: tt.args.request.SecondName,
			ThirdName:  tt.args.request.ThirdName,
		}).Return(nil)

		t.Run(tt.name, func(t *testing.T) {
			s := AuthService{
				userRepo: userRepo,
				crypto:   cryptoRepo,
				jwtKey:   tt.fields.jwtKey,
			}
			if err := s.SignIn(tt.args.request); (err != nil) != tt.wantErr {
				t.Errorf("SignIn() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
