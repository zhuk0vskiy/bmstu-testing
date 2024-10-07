package tests

import (
	"backend/src/internal/model"
	"backend/src/internal/model/dto"
	"backend/src/internal/repository/impl/postgresql"
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"reflect"
	"testing"
)

func TestRoomRepository_GetByStudio(t *testing.T) {
	type fields struct {
		db *pgxpool.Pool
	}
	type args struct {
		ctx     context.Context
		request *dto.GetRoomByStudioRequest
	}
	tests := []struct {
		name string
		//fields    fields
		args      args
		wantRooms []*model.Room
		wantErr   bool
	}{
		{
			name: "test_pos_01",
			args: args{
				ctx: context.Background(),
				request: &dto.GetRoomByStudioRequest{
					StudioId: 1,
				},
			},
			wantRooms: []*model.Room{
				&model.Room{
					Id:        1,
					Name:      "first",
					StudioId:  1,
					StartHour: 13,
					EndHour:   15,
				},
				&model.Room{
					Id:        2,
					Name:      "second",
					StudioId:  1,
					StartHour: 13,
					EndHour:   15,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := postgresql.NewRoomRepository(testDbInstance)
			gotRooms, err := r.GetByStudio(tt.args.ctx, tt.args.request)
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
