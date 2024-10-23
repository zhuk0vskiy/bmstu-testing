package postgres

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"ppo/domain"
)

type ingredientTypeRepository struct {
	//db *pgxpool.Pool
	db IPool
}

func NewIngredientTypeRepository(db IPool) domain.IIngredientTypeRepository {
	return &ingredientTypeRepository{
		db: db,
	}
}

func (r *ingredientTypeRepository) Create(ctx context.Context, ingredientType *domain.IngredientType) error {
	query := `insert into saladRecipes.ingredientType(name, description) 
	values ($1, $2)`

	_, err := r.db.Exec(
		ctx,
		query,
		ingredientType.Name,
		ingredientType.Description)
	if err != nil {
		return fmt.Errorf("creating ingredient type: %w", err)
	}
	return nil
}

func (r *ingredientTypeRepository) GetById(ctx context.Context, id uuid.UUID) (*domain.IngredientType, error) {
	query := `select id, name, description
		from saladRecipes.ingredientType
		where id = $1`

	ingredientType := new(domain.IngredientType)
	err := r.db.QueryRow(
		ctx,
		query,
		id, // uuidToString(id),
	).Scan(
		&ingredientType.ID,
		&ingredientType.Name,
		&ingredientType.Description)

	if err != nil {
		return nil, fmt.Errorf("getting ingredient type by id: %w", err)
	}
	return ingredientType, nil
}

func (r *ingredientTypeRepository) GetAll(ctx context.Context) ([]*domain.IngredientType, error) {
	query := `select id, name, description
		from saladRecipes.ingredientType`

	rows, err := r.db.Query(
		ctx,
		query,
	)
	if err != nil {
		return nil, fmt.Errorf("getting all ingredient types: %w", err)
	}

	ingredientTypes := make([]*domain.IngredientType, 0)
	for rows.Next() {
		tmp := new(domain.IngredientType)
		err = rows.Scan(
			&tmp.ID,
			&tmp.Name,
			&tmp.Description,
		)
		ingredientTypes = append(ingredientTypes, tmp)
		if err != nil {
			return nil, fmt.Errorf("scanning ingredients types: %w", err)
		}
	}

	return ingredientTypes, nil
}

func (r *ingredientTypeRepository) Update(ctx context.Context, measurement *domain.IngredientType) error {
	query := `update saladRecipes.ingredientType
		set
			name = $1, 
			description = $2 
		where id = $3`

	_, err := r.db.Exec(
		ctx,
		query,
		measurement.Name,
		measurement.Description,
		measurement.ID,
	)
	if err != nil {
		return fmt.Errorf("updating ingredient type: %w", err)
	}
	return nil
}

func (r *ingredientTypeRepository) DeleteById(ctx context.Context, id uuid.UUID) error {
	query := `delete from saladRecipes.ingredientType
       where id = $1`

	_, err := r.db.Exec(
		ctx,
		query,
		id, //uuidToString(id)
	)
	if err != nil {
		return fmt.Errorf("deleting ingredient type by id: %w", err)
	}
	return nil
}
