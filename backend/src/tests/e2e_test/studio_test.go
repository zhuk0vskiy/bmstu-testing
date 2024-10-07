package tests

import (
	"backend/src/internal/model"
	"backend/src/internal/model/dto"
	"backend/src/internal/repository/impl/postgresql"
	svcImpl "backend/src/internal/service/impl"
	"backend/src/tests/utils"
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"reflect"
	"testing"
)

func TestStudioPostrgresql_Get(t *testing.T) {
	type fields struct {
		db *pgxpool.Pool
	}
	type args struct {
		ctx     context.Context
		request *dto.GetStudioRequest
	}
	tests := []struct {
		name string
		//fields     fields
		args       args
		wantStudio *model.Studio
		wantErr    bool
	}{
		// TODO: Add test cases.
		{
			name: "test_pos_01",
			args: args{

				ctx: context.Background(),
				request: &dto.GetStudioRequest{
					Id: 1,
				},
			},
			wantStudio: &model.Studio{
				Id:   1,
				Name: "first",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := utils.NewMockLogger()
			p := postgresql.NewStudioRepository(testDbInstance)
			s := svcImpl.NewStudioService(l, p)
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

//func TestStudioRepository_GetAll(t *testing.T) {
//	type fields struct {
//		db *pgxpool.Pool
//	}
//	type args struct {
//		ctx     context.Context
//		request *dto.GetStudioAllRequest
//	}
//	tests := []struct {
//		name        string
//		fields      fields
//		args        args
//		wantStudios []*model.Studio
//		wantErr     bool
//	}{
//		{
//			name: "test_pos_01",
//			args: args{
//				ctx:     context.Background(),
//				request: &dto.GetStudioAllRequest{},
//			},
//			wantStudios: []*model.Studio{
//				&model.Studio{
//					Id:   1,
//					Name: "first",
//				},
//				&model.Studio{
//					Id:   2,
//					Name: "second",
//				},
//			},
//			wantErr: false,
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			r := postgresql.NewStudioRepository(testDbInstance)
//			gotStudios, err := r.GetAll(tt.args.ctx, tt.args.request)
//			if (err != nil) != tt.wantErr {
//				t.Errorf("GetAll() error = %v, wantErr %v", err, tt.wantErr)
//				return
//			}
//			if !reflect.DeepEqual(gotStudios, tt.wantStudios) {
//				t.Errorf("GetAll() gotStudios = %v, want %v", gotStudios, tt.wantStudios)
//			}
//		})
//	}
//}

func TestStudioRepository_Update(t *testing.T) {
	type fields struct {
		db *pgxpool.Pool
	}
	type args struct {
		ctx     context.Context
		request *dto.UpdateStudioRequest
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
				ctx: context.Background(),
				request: &dto.UpdateStudioRequest{
					Id:   1,
					Name: "second",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := postgresql.NewStudioRepository(testDbInstance)
			if err := r.Update(tt.args.ctx, tt.args.request); (err != nil) != tt.wantErr {
				t.Errorf("Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestStudioRepository_Add(t *testing.T) {
	type fields struct {
		db *pgxpool.Pool
	}
	type args struct {
		ctx     context.Context
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
				ctx: context.Background(),
				request: &dto.AddStudioRequest{
					Name: "third",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := utils.NewMockLogger()
			r := postgresql.NewStudioRepository(testDbInstance)
			s := svcImpl.NewStudioService(l, r)
			if err := s.Add(tt.args.request); (err != nil) != tt.wantErr {
				t.Errorf("Add() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestStudioRepository_Delete(t *testing.T) {
	type fields struct {
		db *pgxpool.Pool
	}
	type args struct {
		ctx     context.Context
		request *dto.DeleteStudioRequest
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
				ctx: context.Background(),
				request: &dto.DeleteStudioRequest{
					Id: 1,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := utils.NewMockLogger()
			r := postgresql.NewStudioRepository(testDbInstance)
			s := svcImpl.NewStudioService(l, r)
			if err := s.Delete(tt.args.request); (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
