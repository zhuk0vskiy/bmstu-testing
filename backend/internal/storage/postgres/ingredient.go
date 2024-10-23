package postgres

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"ppo/domain"
	"ppo/internal/config"
)

type ingredientRepository struct {
	//db *pgxpool.Pool
	db IPool
}

func NewIngredientRepository(db IPool) domain.IIngredientRepository {
	return &ingredientRepository{
		db: db,
	}
}

func (r *ingredientRepository) Create(ctx context.Context, ingredient *domain.Ingredient) error {
	query := `insert into saladRecipes.ingredient(name, calories, type) 
	values ($1, $2, $3)`

	_, err := r.db.Exec(
		ctx,
		query,
		ingredient.Name,
		ingredient.Calories,
		ingredient.TypeID)
	if err != nil {
		return fmt.Errorf("creating ingredient: %w", err)
	}
	return nil
}

func (r *ingredientRepository) GetById(ctx context.Context, id uuid.UUID) (*domain.Ingredient, error) {
	query := `select id, name, calories, type
		from saladRecipes.ingredient
		where id = $1`

	ingredient := new(domain.Ingredient)
	err := r.db.QueryRow(
		ctx,
		query,
		id,
	).Scan(
		&ingredient.ID,
		&ingredient.Name,
		&ingredient.Calories,
		&ingredient.TypeID)

	if err != nil {
		return nil, fmt.Errorf("getting ingredient by id: %w", err)
	}
	return ingredient, nil
}

func (r *ingredientRepository) GetAll(ctx context.Context, page int) ([]*domain.Ingredient, int, error) {
	query := `select id, name, calories, type
		from saladRecipes.ingredient
		offset $1
		limit $2`

	rows, err := r.db.Query(
		ctx,
		query,
		config.PageSize*(page-1),
		config.PageSize,
	)

	ingredients := make([]*domain.Ingredient, 0)
	for rows.Next() {
		tmp := new(domain.Ingredient)
		err = rows.Scan(
			&tmp.ID,
			&tmp.Name,
			&tmp.Calories,
			&tmp.TypeID,
		)
		ingredients = append(ingredients, tmp)
		if err != nil {
			return nil, 0, fmt.Errorf("scanning ingredients: %w", err)
		}
	}

	numRows := 0
	err = r.db.QueryRow(
		ctx,
		`select count(*) from saladRecipes.ingredient`,
	).Scan(&numRows)
	numPages := numRows / config.PageSize
	if numRows%config.PageSize != 0 {
		numPages++
	}

	return ingredients, numPages, nil
}

func (r *ingredientRepository) GetAllByRecipeId(ctx context.Context, id uuid.UUID) ([]*domain.Ingredient, error) {
	query := `with needed as (select *
		from saladRecipes.recipeIngredient
		where recipeId = $1)
	select ingredient.id, ingredient.name, ingredient.calories, ingredient.type
	from needed join saladRecipes.ingredient on needed.ingredientId = ingredient.id`

	rows, err := r.db.Query(
		ctx,
		query,
		id,
	)

	ingredients := make([]*domain.Ingredient, 0)
	for rows.Next() {
		tmp := new(domain.Ingredient)
		err = rows.Scan(
			&tmp.ID,
			&tmp.Name,
			&tmp.Calories,
			&tmp.TypeID,
		)
		ingredients = append(ingredients, tmp)
		if err != nil {
			return nil, fmt.Errorf("scanning ingredients: %w", err)
		}
	}
	return ingredients, nil
}

func (r *ingredientRepository) Update(ctx context.Context, ingredient *domain.Ingredient) error {
	query := `update saladRecipes.ingredient
		set
			name = $1, 
			calories = $2,
			type = $3
		where id = $4`

	_, err := r.db.Exec(
		ctx,
		query,
		ingredient.Name,
		ingredient.Calories,
		ingredient.TypeID,
		ingredient.ID,
	)
	if err != nil {
		return fmt.Errorf("updating ingredient: %w", err)
	}
	return nil
}

func (r *ingredientRepository) DeleteById(ctx context.Context, id uuid.UUID) error {
	query := `delete from saladRecipes.ingredient
       where id = $1`

	_, err := r.db.Exec(
		ctx,
		query,
		id,
	)
	if err != nil {
		return fmt.Errorf("deleting ingredient by id: %w", err)
	}
	return nil
}

func (r *ingredientRepository) Link(ctx context.Context, recipeId uuid.UUID, ingredientId uuid.UUID) (uuid.UUID, error) {
	query := `insert into saladRecipes.recipeIngredient(recipeId, ingredientId)
	values ($1, $2)
	returning id`
	var id uuid.UUID

	err := r.db.QueryRow(
		ctx,
		query,
		recipeId,
		ingredientId,
	).Scan(&id)
	if err != nil {
		return uuid.Nil, fmt.Errorf("linking ingredient: %w", err)
	}
	return id, nil
}

func (r *ingredientRepository) Unlink(ctx context.Context, recipeId uuid.UUID, ingredientId uuid.UUID) error {
	query := `delete from saladRecipes.recipeIngredient
	where recipeId = $1 and ingredientId = $2`

	_, err := r.db.Exec(
		ctx,
		query,
		recipeId,
		ingredientId,
	)
	if err != nil {
		return fmt.Errorf("unlinking ingredient by recipe id: %w", err)
	}
	return nil
}
