package postgres

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"ppo/domain"
	"ppo/internal/config"
)

type saladTypeRepository struct {
	//db *pgxpool.Pool
	db IPool
}

func NewSaladTypeRepository(db IPool) domain.ISaladTypeRepository {
	return &saladTypeRepository{
		db: db,
	}
}

func (r *saladTypeRepository) Create(ctx context.Context, saladType *domain.SaladType) error {
	query := `insert into saladRecipes.saladType(name, description) 
	values ($1, $2)`

	_, err := r.db.Exec(
		ctx,
		query,
		saladType.Name,
		saladType.Description)
	if err != nil {
		return fmt.Errorf("creating salad type: %w", err)
	}
	return nil
}

func (r *saladTypeRepository) GetById(ctx context.Context, id uuid.UUID) (*domain.SaladType, error) {
	query := `select id, name, description
		from saladRecipes.saladType
		where id = $1`

	saladType := new(domain.SaladType)
	err := r.db.QueryRow(
		ctx,
		query,
		id,
	).Scan(
		&saladType.ID,
		&saladType.Name,
		&saladType.Description)

	if err != nil {
		return nil, fmt.Errorf("getting salad type by id: %w", err)
	}
	return saladType, nil
}

func (r *saladTypeRepository) GetAll(ctx context.Context, page int) ([]*domain.SaladType, int, error) {
	query := `select id, name, description
		from saladRecipes.saladType
		offset $1
		limit $2`

	rows, err := r.db.Query(
		ctx,
		query,
		config.PageSize*(page-1),
		config.PageSize,
	)
	if err != nil {
		return nil, 0, fmt.Errorf("getting all salad types: %w", err)
	}

	saladTypes := make([]*domain.SaladType, 0)
	for rows.Next() {
		tmp := new(domain.SaladType)
		err = rows.Scan(
			&tmp.ID,
			&tmp.Name,
			&tmp.Description,
		)
		saladTypes = append(saladTypes, tmp)
		if err != nil {
			return nil, 0, fmt.Errorf("scanning salad types: %w", err)
		}
	}

	numRows := 0
	err = r.db.QueryRow(
		ctx,
		`select count(*) from saladRecipes.typesOfSalads`,
	).Scan(&numRows)
	numPages := numRows / config.PageSize
	if numRows%config.PageSize != 0 {
		numPages++
	}

	return saladTypes, numPages, nil
}

func (r *saladTypeRepository) GetAllBySaladId(ctx context.Context, saladId uuid.UUID) ([]*domain.SaladType, error) {
	query := `select id, name, description
		from saladRecipes.saladType
		where id in (
		    select typeId
		    from saladRecipes.typesOfSalads
		    where saladId = $1)`

	rows, err := r.db.Query(
		ctx,
		query,
		saladId,
	)
	if err != nil {
		return nil, fmt.Errorf("getting salad types by saladId: %w", err)
	}

	saladTypes := make([]*domain.SaladType, 0)
	for rows.Next() {
		tmp := new(domain.SaladType)
		err = rows.Scan(
			&tmp.ID,
			&tmp.Name,
			&tmp.Description,
		)
		saladTypes = append(saladTypes, tmp)
		if err != nil {
			return nil, fmt.Errorf("scanning salad types: %w", err)
		}
	}

	return saladTypes, nil
}

func (r *saladTypeRepository) Update(ctx context.Context, saladType *domain.SaladType) error {
	query := `update saladRecipes.saladType
		set
			name = $1, 
			description = $2 
		where id = $3`

	_, err := r.db.Exec(
		ctx,
		query,
		saladType.Name,
		saladType.Description,
		saladType.ID,
	)
	if err != nil {
		return fmt.Errorf("updating salad type: %w", err)
	}
	return nil
}

func (r *saladTypeRepository) DeleteById(ctx context.Context, id uuid.UUID) error {
	query := `delete from saladRecipes.saladType
       where id = $1`

	_, err := r.db.Exec(
		ctx,
		query,
		id)
	if err != nil {
		return fmt.Errorf("deleting salad type by id: %w", err)
	}
	return nil
}

func (r *saladTypeRepository) Link(ctx context.Context, saladId uuid.UUID, saladTypeId uuid.UUID) error {
	query := `insert into saladRecipes.typesOfSalads(saladId, typeId)
	values ($1, $2)`

	_, err := r.db.Exec(
		ctx,
		query,
		saladId,
		saladTypeId)
	if err != nil {
		return fmt.Errorf("linking types from salad: %w", err)
	}
	return nil
}

func (r *saladTypeRepository) Unlink(ctx context.Context, saladId uuid.UUID, saladTypeId uuid.UUID) error {
	query := `delete from saladRecipes.typesOfSalads
	where saladId = $1 and typeId = $2`

	_, err := r.db.Exec(
		ctx,
		query,
		saladId,
		saladTypeId)
	if err != nil {
		return fmt.Errorf("unlinking types from salad: %w", err)
	}
	return nil
}
