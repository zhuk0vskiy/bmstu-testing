package postgresql

import (
	"backend/src/internal/model"
	"backend/src/internal/model/dto"
	repositoryInterface "backend/src/internal/repository/interface"
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
)

type RoomRepository struct {
	db *pgxpool.Pool
}

func NewRoomRepository(db *pgxpool.Pool) repositoryInterface.IRoomRepository {
	return &RoomRepository{
		db: db,
	}
}

func (r RoomRepository) Get(ctx context.Context, request *dto.GetRoomRequest) (room *model.Room, err error) {

	query := `    select id, 
				       name,
				       studio_id,
				       start_hour,
				       end_hour
				from room
				where id = $1`

	room = new(model.Room)
	//var startTime string = ""
	//var endTime string = ""

	err = r.db.QueryRow(
		ctx,
		query,
		request.Id,
	).Scan(
		&room.Id,
		&room.Name,
		&room.StudioId,
		&room.StartHour,
		&room.EndHour,
	)

	//&room.StartHour = time_parser.StringToDate(startTime)

	if err != nil {
		return nil, fmt.Errorf("запрос не выполнен: %w", err)
	}

	return room, err
}

func (r RoomRepository) GetByStudio(ctx context.Context, request *dto.GetRoomByStudioRequest) (rooms []*model.Room, err error) {
	query := `
				select id, 
				       name,
				       studio_id,
				       start_hour,
				       end_hour
				from room 
				where studio_id = $1`

	rows, err := r.db.Query(
		ctx,
		query,
		request.StudioId,
	)

	if err != nil {
		return nil, fmt.Errorf("ошибка при запросе: %w", err)
	}

	rooms = make([]*model.Room, 0)

	for rows.Next() {
		tmp := new(model.Room)

		err = rows.Scan(
			&tmp.Id,
			&tmp.Name,
			&tmp.StudioId,
			&tmp.StartHour,
			&tmp.EndHour,
		)

		rooms = append(rooms, tmp)

		if err != nil {
			return nil, fmt.Errorf("сканирование полученных строк: %w", err)
		}
	}

	return rooms, err
}

func (r RoomRepository) Add(ctx context.Context, request *dto.AddRoomRequest) (err error) {
	query := `insert into room(name, studio_id, start_hour, end_hour) values ($1, $2, $3, $4)`

	_, err = r.db.Exec(
		ctx,
		query,
		request.Name,
		request.StudioId,
		request.StartHour,
		request.EndHour,
	)
	if err != nil {
		return fmt.Errorf("создание финансового отчета: %w", err)
	}

	return err
}

func (r RoomRepository) Update(ctx context.Context, request *dto.UpdateRoomRequest) (err error) {
	query := `
			update room
			set 
			    id = $1,
			    name = $2,
			    studio_id = $3,
			    start_hour = $4,
			    end_hour = $5
			where id = $1`

	_, err = r.db.Exec(
		ctx,
		query,
		request.Id,
		request.Name,
		request.StudioId,
		request.StartHour,
		request.EndHour,
	)
	if err != nil {
		return fmt.Errorf("обновление информации о комнате: %w", err)
	}

	return err
}

func (r RoomRepository) Delete(ctx context.Context, request *dto.DeleteRoomRequest) (err error) {
	query := `delete from room where id = $1`

	_, err = r.db.Exec(
		ctx,
		query,
		request.Id,
	)
	if err != nil {
		return fmt.Errorf("удаление комнаты по id: %w", err)
	}

	return nil
}
