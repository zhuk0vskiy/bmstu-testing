package postgresql

import (
	"backend/src/internal/model"
	"backend/src/internal/model/dto"
	repositoryInterface "backend/src/internal/repository/interface"
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
)

type StudioRepository struct {
	db *pgxpool.Pool
}

func NewStudioRepository(db *pgxpool.Pool) repositoryInterface.IStudioRepository {
	return &StudioRepository{
		db: db,
	}
}

func (r StudioRepository) Get(ctx context.Context, request *dto.GetStudioRequest) (studio *model.Studio, err error) {
	query := `    select id, 
				       name
				from studio 
				where id = $1`

	studio = new(model.Studio)

	err = r.db.QueryRow(
		ctx,
		query,
		request.Id,
	).Scan(
		&studio.Id,
		&studio.Name,
	)

	if err != nil {
		return nil, fmt.Errorf("get запрос не выполнен: %w", err)
	}

	return studio, err
}

func (r StudioRepository) GetAll(ctx context.Context, request *dto.GetStudioAllRequest) (studios []*model.Studio, err error) {

	query := `
				select id, 
				       name  
				from studio`

	rows, err := r.db.Query(
		ctx,
		query,
	)

	if err != nil {
		return nil, fmt.Errorf("ошибка при getAll запросе: %w", err)
	}

	studios = make([]*model.Studio, 0)

	for rows.Next() {
		tmp := new(model.Studio)

		err = rows.Scan(
			&tmp.Id,
			&tmp.Name,
		)

		studios = append(studios, tmp)

		if err != nil {
			return nil, fmt.Errorf("сканирование полученных строк: %w", err)
		}
	}

	return studios, err
}

func (r StudioRepository) Update(ctx context.Context, request *dto.UpdateStudioRequest) (err error) {
	query := `
			update studio
			set 
			    id = $1,
			    name = $2
			where id = $1`

	_, err = r.db.Exec(
		ctx,
		query,
		request.Id,
		request.Name,
	)
	if err != nil {
		return fmt.Errorf("обновление информации о студии: %w", err)
	}

	return err
}

func (r StudioRepository) Add(ctx context.Context, request *dto.AddStudioRequest) (err error) {
	query := `insert into studio(name) 
	values ($1)`

	_, err = r.db.Exec(
		ctx,
		query,
		request.Name,
	)
	if err != nil {
		return fmt.Errorf("создание финансового отчета: %w", err)
	}

	return err
}

func (r StudioRepository) Delete(ctx context.Context, request *dto.DeleteStudioRequest) (err error) {
	query := `delete from studio where id = $1`

	_, err = r.db.Exec(
		ctx,
		query,
		request.Id,
	)
	if err != nil {
		return fmt.Errorf("удаление отчета по id: %w", err)
	}

	return nil
}
