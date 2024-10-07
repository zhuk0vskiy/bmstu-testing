package impl

import (
	"backend/src/internal/model"
	"backend/src/internal/model/dto"
	_interface "backend/src/internal/repository/interface"
	"backend/src/internal/repository/interface/mocks"
	//_interface "backend/src/internal/service_test/interface"
	"context"
	"reflect"
	"testing"
)

//func TestNewStudioService(t *testing.T) {
//	type args struct {
//		studioRepo _interface.IStudioRepository
//	}
//	test_data := []struct {
//		name string
//		args args
//		want _interface.IStudioService
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range test_data {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := NewStudioService(tt.args.studioRepo); !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("NewStudioService() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}

func TestStudioService_Add(t *testing.T) {
	type fields struct {
		studioRepo _interface.IStudioRepository
	}
	type args struct {
		request *dto.AddStudioRequest
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
				request: &dto.AddStudioRequest{
					Name: "1",
				},
			},
			wantErr: false,
		},
		{
			name: "test_neg_01",
			args: args{
				request: &dto.AddStudioRequest{
					Name: "",
				},
			},
			wantErr: true,
		},
	}

	studioRepo := new(mocks.IStudioRepository)

	for _, tt := range tests {
		studioRepo.On("Add", context.Background(), tt.args.request).Return(nil)
		t.Run(tt.name, func(t *testing.T) {
			s := StudioService{
				studioRepo: studioRepo,
			}
			if err := s.Add(tt.args.request); (err != nil) != tt.wantErr {
				t.Errorf("Add() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestStudioService_Delete(t *testing.T) {
	type fields struct {
		studioRepo _interface.IStudioRepository
	}
	type args struct {
		request *dto.DeleteStudioRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := StudioService{
				studioRepo: tt.fields.studioRepo,
			}
			if err := s.Delete(tt.args.request); (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestStudioService_Get(t *testing.T) {
	type fields struct {
		studioRepo _interface.IStudioRepository
	}
	type args struct {
		request *dto.GetStudioRequest
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantStudio *model.Studio
		wantErr    bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := StudioService{
				studioRepo: tt.fields.studioRepo,
			}
			gotStudio, err := s.Get(tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotStudio, tt.wantStudio) {
				t.Errorf("Get() gotStudio = %v, want %v", gotStudio, tt.wantStudio)
			}
		})
	}
}

func TestStudioService_GetAll(t *testing.T) {
	type fields struct {
		studioRepo _interface.IStudioRepository
	}
	type args struct {
		request *dto.GetStudioAllRequest
	}
	tests := []struct {
		name        string
		fields      fields
		args        args
		wantStudios []*model.Studio
		wantErr     bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := StudioService{
				studioRepo: tt.fields.studioRepo,
			}
			gotStudios, err := s.GetAll(tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotStudios, tt.wantStudios) {
				t.Errorf("GetAll() gotStudios = %v, want %v", gotStudios, tt.wantStudios)
			}
		})
	}
}

func TestStudioService_Update(t *testing.T) {
	type fields struct {
		studioRepo _interface.IStudioRepository
	}
	type args struct {
		request *dto.UpdateStudioRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := StudioService{
				studioRepo: tt.fields.studioRepo,
			}
			if err := s.Update(tt.args.request); (err != nil) != tt.wantErr {
				t.Errorf("Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
