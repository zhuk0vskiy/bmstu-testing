package postgresql

import (
	"backend/src/internal/model"
	"backend/src/internal/model/dto"
	repositoryInterface "backend/src/internal/repository/interface"
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ProducerRepository struct {
	db *pgxpool.Pool
}

func NewProducerRepository(db *pgxpool.Pool) repositoryInterface.IProducerRepository {
	return &ProducerRepository{
		db: db,
	}
}

func (r ProducerRepository) Get(ctx context.Context, request *dto.GetProducerRequest) (producer *model.Producer, err error) {

	query := `    select id, 
				       name,
				       studio_id,
				       start_hour,
				       end_hour
				from producer
				where id = $1`

	producer = new(model.Producer)
	//var startTime string = ""
	//var endTime string = ""

	err = r.db.QueryRow(
		ctx,
		query,
		request.Id,
	).Scan(
		&producer.Id,
		&producer.Name,
		&producer.StudioId,
		&producer.StartHour,
		&producer.EndHour,
	)

	//&producer.StartHour = time_parser.StringToDate(startTime)

	if err != nil {
		return nil, fmt.Errorf("запрос не выполнен: %w", err)
	}

	return producer, err
}

func (r ProducerRepository) GetByStudio(ctx context.Context, request *dto.GetProducerByStudioRequest) (producers []*model.Producer, err error) {
	query := `
				select id, 
				       name,
				       studio_id,
				       start_hour,
				       end_hour
				from producer 
				where studio_id = $1`

	rows, err := r.db.Query(
		ctx,
		query,
		request.StudioId,
	)

	if err != nil {
		return nil, fmt.Errorf("ошибка при запросе: %w", err)
	}

	producers = make([]*model.Producer, 0)

	for rows.Next() {
		tmp := new(model.Producer)

		err = rows.Scan(
			&tmp.Id,
			&tmp.Name,
			&tmp.StudioId,
			&tmp.StartHour,
			&tmp.EndHour,
		)

		producers = append(producers, tmp)

		if err != nil {
			return nil, fmt.Errorf("сканирование полученных строк: %w", err)
		}
	}

	return producers, err
}

func (r ProducerRepository) Add(ctx context.Context, request *dto.AddProducerRequest) (err error) {
	query := `insert into producer(name, studio_id, start_hour, end_hour) values ($1, $2, $3, $4)`

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

func (r ProducerRepository) Update(ctx context.Context, request *dto.UpdateProducerRequest) (err error) {
	query := `
			update producer
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

func (r ProducerRepository) Delete(ctx context.Context, request *dto.DeleteProducerRequest) (err error) {
	query := `delete from producer where id = $1`

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
