package postgres

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"ppo/domain"
	"ppo/internal/config"
	"ppo/services/dto"
)

type recipeRepository struct {
	//db *pgxpool.Pool
	db IPool
}

func NewRecipeRepository(db IPool) domain.IRecipeRepository {
	return &recipeRepository{
		db: db,
	}
}

func (r *recipeRepository) Create(ctx context.Context, recipe *domain.Recipe) (uuid.UUID, error) {
	query := `insert into saladRecipes.recipe(saladId, status, numberOfServings, timeToCook)
		values ($1, $2, $3, $4)
	returning id`

	id := uuid.Nil
	err := r.db.QueryRow(
		ctx,
		query,
		recipe.SaladID,
		recipe.Status,
		recipe.NumberOfServings,
		recipe.TimeToCook,
	).Scan(&id)
	if err != nil {
		return id, fmt.Errorf("creating recipe: %w", err)
	}
	return id, nil
}

func (r *recipeRepository) GetById(ctx context.Context, id uuid.UUID) (*domain.Recipe, error) {
	query := `select id, saladId, status, numberOfServings, timeToCook, rating
		from saladRecipes.recipe
		where id = $1`

	recipe := new(domain.Recipe)
	err := r.db.QueryRow(
		ctx,
		query,
		id,
	).Scan(
		&recipe.ID,
		&recipe.SaladID,
		&recipe.Status,
		&recipe.NumberOfServings,
		&recipe.TimeToCook,
		&recipe.Rating)

	if err != nil {
		return nil, fmt.Errorf("getting recipe by id: %w", err)
	}
	return recipe, nil
}

func (r *recipeRepository) GetBySaladId(ctx context.Context, saladId uuid.UUID) (*domain.Recipe, error) {
	query := `select id, saladId, status, numberOfServings, timeToCook, rating
		from saladRecipes.recipe
		where saladId = $1`

	recipe := new(domain.Recipe)
	err := r.db.QueryRow(
		ctx,
		query,
		saladId,
	).Scan(
		&recipe.ID,
		&recipe.SaladID,
		&recipe.Status,
		&recipe.NumberOfServings,
		&recipe.TimeToCook,
		&recipe.Rating)

	if err != nil {
		return nil, fmt.Errorf("getting recipe by salad id: %w", err)
	}
	return recipe, nil
}

func (r *recipeRepository) GetAll(ctx context.Context, filter *dto.RecipeFilter, page int) ([]*domain.Recipe, error) {
	// todo: simplify query

	ingredientUUIDS := ""
	if filter.AvailableIngredients != nil && len(filter.AvailableIngredients) != 0 {
		ingredientUUIDS = uuidsToString(filter.AvailableIngredients)
	} else {
		ingredientUUIDS = "ingredientId"
	}

	saladTypesUUIDS := ""
	if filter.SaladTypes != nil && len(filter.SaladTypes) != 0 {
		saladTypesUUIDS = uuidsToString(filter.SaladTypes)
	} else {
		saladTypesUUIDS = "typeId"
	}

	query := `with matchCounts as (select recipeId, count(*) as matches
		from saladRecipes.recipeIngredient
		where ingredientId in ($1)
		group by recipeId),
	totalIngredients as (select recipeId, count(*) as ingredientsCount
		from saladRecipes.recipeIngredient
		group by recipeId),
	availableRecipes as (select id, saladId, status, numberOfServings, timeToCook, rating
		from saladRecipes.recipe
		where id in (select matchCounts.recipeId
             from matchCounts join totalIngredients on matchCounts.recipeId = totalIngredients.recipeId
             where matches = ingredientsCount)),
	requestedTypes as (select saladId
		from saladRecipes.typesOfSalads
		where typeId in ($2))
		group by saladId)
	select id, saladId, status, numberOfServings, timeToCook, rating
		from availableRecipes
		where
			saladId in (select saladId from requestedTypes) and
			(rating is null or rating > $3)
		order by case
    		when rating is null then rating
    		else 0 end, rating desc 
		offset $4
		limit $5`

	rows, err := r.db.Query(
		ctx,
		query,
		ingredientUUIDS, //uuidsToString(filter.AvailableIngredients),
		saladTypesUUIDS, //uuidsToString(filter.SaladTypes),
		filter.MinRate,
		config.PageSize*(page-1),
		config.PageSize,
	)
	if err != nil {
		return nil, fmt.Errorf("querying recipes: %w", err)
	}

	recipes := make([]*domain.Recipe, 0)
	for rows.Next() {
		tmp := new(domain.Recipe)
		err = rows.Scan(
			&tmp.ID,
			&tmp.SaladID,
			&tmp.Status,
			&tmp.NumberOfServings,
			&tmp.TimeToCook,
			&tmp.Rating,
		)
		recipes = append(recipes, tmp)
		if err != nil {
			return nil, fmt.Errorf("scanning recipes: %w", err)
		}
	}

	return recipes, nil
}

func (r *recipeRepository) Update(ctx context.Context, recipe *domain.Recipe) error {
	query := `update saladRecipes.recipe
		set
			status = $1,
			numberOfServings = $2,
			timeToCook = $3,
			rating = $4
		where id = $5`

	_, err := r.db.Exec(
		ctx,
		query,
		recipe.Status,
		recipe.NumberOfServings,
		recipe.TimeToCook,
		recipe.Rating,
		recipe.ID,
	)
	if err != nil {
		return fmt.Errorf("updating recipe: %w", err)
	}
	return nil
}

func (r *recipeRepository) DeleteById(ctx context.Context, id uuid.UUID) error {
	query := `delete from saladRecipes.recipe
       where id = $1`

	_, err := r.db.Exec(
		ctx,
		query,
		id)
	if err != nil {
		return fmt.Errorf("deleting recipe by id: %w", err)
	}
	return nil
}
