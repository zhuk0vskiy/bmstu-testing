package postgresql

import (
	"backend/src/internal/model"
	"backend/src/internal/model/dto"
	repositoryInterface "backend/src/internal/repository/interface"
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
)

type InstrumentalistRepository struct {
	db *pgxpool.Pool
}

func NewInstrumentalistRepository(db *pgxpool.Pool) repositoryInterface.IInstrumentalistRepository {
	return &InstrumentalistRepository{
		db: db,
	}
}

func (r InstrumentalistRepository) Get(ctx context.Context, request *dto.GetInstrumentalistRequest) (instrumentalist *model.Instrumentalist, err error) {

	query := `    select id, 
				       name,
				       studio_id,
				       start_hour,
				       end_hour
				from instrumentalist
				where id = $1`

	instrumentalist = new(model.Instrumentalist)
	//var startTime string = ""
	//var endTime string = ""

	err = r.db.QueryRow(
		ctx,
		query,
		request.Id,
	).Scan(
		&instrumentalist.Id,
		&instrumentalist.Name,
		&instrumentalist.StudioId,
		&instrumentalist.StartHour,
		&instrumentalist.EndHour,
	)

	//&instrumentalist.StartHour = time_parser.StringToDate(startTime)

	if err != nil {
		return nil, fmt.Errorf("запрос не выполнен: %w", err)
	}

	return instrumentalist, err
}

func (r InstrumentalistRepository) GetByStudio(ctx context.Context, request *dto.GetInstrumentalistByStudioRequest) (instrumentalists []*model.Instrumentalist, err error) {
	query := `
				select id, 
				       name,
				       studio_id,
				       start_hour,
				       end_hour
				from instrumentalist 
				where studio_id = $1`

	rows, err := r.db.Query(
		ctx,
		query,
		request.StudioId,
	)

	if err != nil {
		return nil, fmt.Errorf("ошибка при запросе: %w", err)
	}

	instrumentalists = make([]*model.Instrumentalist, 0)

	for rows.Next() {
		tmp := new(model.Instrumentalist)

		err = rows.Scan(
			&tmp.Id,
			&tmp.Name,
			&tmp.StudioId,
			&tmp.StartHour,
			&tmp.EndHour,
		)

		instrumentalists = append(instrumentalists, tmp)

		if err != nil {
			return nil, fmt.Errorf("сканирование полученных строк: %w", err)
		}
	}

	return instrumentalists, err
}

func (r InstrumentalistRepository) Add(ctx context.Context, request *dto.AddInstrumentalistRequest) (err error) {
	query := `insert into instrumentalist(name, studio_id, start_hour, end_hour) values ($1, $2, $3, $4)`

	_, err = r.db.Exec(
		ctx,
		query,
		request.Name,
		request.StudioId,
		request.StartHour,
		request.EndHour,
	)
	if err != nil {
		return fmt.Errorf("добавление продюсера: %w", err)
	}

	return err
}

func (r InstrumentalistRepository) Update(ctx context.Context, request *dto.UpdateInstrumentalistRequest) (err error) {
	query := `
			update instrumentalist
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
		return fmt.Errorf("обновление информации о продюсере: %w", err)
	}

	return err
}

func (r InstrumentalistRepository) Delete(ctx context.Context, request *dto.DeleteInstrumentalistRequest) (err error) {
	query := `delete from instrumentalist where id = $1`

	_, err = r.db.Exec(
		ctx,
		query,
		request.Id,
	)
	if err != nil {
		return fmt.Errorf("удаление продюсера по id: %w", err)
	}

	return nil
}
