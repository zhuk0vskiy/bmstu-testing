package postgres

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"ppo/domain"
	"ppo/internal/config"
	"ppo/services/dto"
)

type saladRepository struct {
	//db *pgxpool.Pool
	db IPool
}

func NewSaladRepository(db IPool) domain.ISaladRepository {
	return &saladRepository{
		db: db,
	}
}

func (r *saladRepository) Create(ctx context.Context, salad *domain.Salad) (uuid.UUID, error) {
	query := `insert into saladRecipes.salad(name, authorId, description)
		values ($1, $2, $3)
	returning id`

	id := uuid.Nil
	err := r.db.QueryRow(
		ctx,
		query,
		salad.Name,
		salad.AuthorID,
		salad.Description,
	).Scan(&id)

	if err != nil {
		return id, fmt.Errorf("creating salad: %w", err)
	}
	return id, nil
}

func (r *saladRepository) GetById(ctx context.Context, id uuid.UUID) (*domain.Salad, error) {
	query := `select id, name, authorId, description
		from saladRecipes.salad
		where id = $1`

	salad := new(domain.Salad)
	err := r.db.QueryRow(
		ctx,
		query,
		id,
	).Scan(
		&salad.ID,
		&salad.Name,
		&salad.AuthorID,
		&salad.Description)

	if err != nil {
		return nil, fmt.Errorf("getting salad by id: %w", err)
	}
	return salad, nil
}

func (r *saladRepository) GetAll(ctx context.Context, filter *dto.RecipeFilter, page int) ([]*domain.Salad, int, error) {
	query := `with matchCounts as (select recipeId, count(*) as matches
		from saladRecipes.recipeIngredient
		where ingredientId = any ($1) or ($7 = 1)
		group by recipeId),
	totalIngredients as (select recipeId, count(*) as ingredientsCount
		from saladRecipes.recipeIngredient
		group by recipeId),
	availableRecipes as (select id, saladId, status, numberOfServings, timeToCook, rating
		from saladRecipes.recipe
		where id in (select matchCounts.recipeId
			from matchCounts join totalIngredients on matchCounts.recipeId = totalIngredients.recipeId
			where matches = ingredientsCount) or
				(($7 = 1) and id not in (select saladRecipes.recipeIngredient.recipeId from saladRecipes.recipeIngredient)) ),
	requestedTypes as (select saladId
		from saladRecipes.typesOfSalads
		where typeId = any ($2) or ($8 = 1)
		group by saladId)

	select salad.id, salad.name, salad.description, salad.authorId
	from availableRecipes join saladRecipes.salad on availableRecipes.saladId = salad.id
	where
    	(($8 = 1) or saladId in (select saladId from requestedTypes)) and
    	(rating >= $3) and (status = $6)
	order by case
		when rating is null then rating
		else 0 end, rating desc
	offset $4
    limit $5`

	allIngredients := 0
	if filter.AvailableIngredients == nil || len(filter.AvailableIngredients) == 0 {
		allIngredients = 1
	}
	allTypes := 0
	if filter.SaladTypes == nil || len(filter.SaladTypes) == 0 {
		allTypes = 1
	}

	rows, err := r.db.Query(
		ctx,
		query,
		filter.AvailableIngredients,
		filter.SaladTypes,
		filter.MinRate,
		config.PageSize*(page-1),
		config.PageSize,
		filter.Status,
		allIngredients,
		allTypes,
	)
	if err != nil {
		return nil, 0, fmt.Errorf("querying salads: %w", err)
	}

	salads := make([]*domain.Salad, 0)
	for rows.Next() {
		tmp := new(domain.Salad)
		err = rows.Scan(
			&tmp.ID,
			&tmp.Name,
			&tmp.Description,
			&tmp.AuthorID,
		)
		salads = append(salads, tmp)
		if err != nil {
			return nil, 0, fmt.Errorf("scanning salads: %w", err)
		}
	}

	numRows := 0
	err = r.db.QueryRow(
		ctx,
		`select count(*) from saladRecipes.salad`,
	).Scan(&numRows)
	numPages := numRows / config.PageSize
	if numRows%config.PageSize != 0 {
		numPages++
	}

	return salads, numPages, nil
}

func (r *saladRepository) GetAllByUserId(ctx context.Context, id uuid.UUID) ([]*domain.Salad, error) {
	query := `select id, name, authorId, description
		from saladRecipes.salad
		where authorId = $1`

	rows, err := r.db.Query(
		ctx,
		query,
		id,
	)

	salads := make([]*domain.Salad, 0)
	for rows.Next() {
		tmp := new(domain.Salad)
		err = rows.Scan(
			&tmp.ID,
			&tmp.Name,
			&tmp.AuthorID,
			&tmp.Description,
		)
		salads = append(salads, tmp)
		if err != nil {
			return nil, fmt.Errorf("scanning salads: %w", err)
		}
	}

	return salads, nil
}

func (r *saladRepository) GetAllRatedByUser(ctx context.Context, userId uuid.UUID, page int) ([]*domain.Salad, int, error) {
	query := `with rates as (select salad
	from saladRecipes.comment
	where author = $1)
	select id, name, authorId, description
	from saladRecipes.salad
	where id in (select salad from rates)
	offset $2
    limit $3`

	rows, err := r.db.Query(
		ctx,
		query,
		userId,
		config.PageSize*(page-1),
		config.PageSize,
	)

	salads := make([]*domain.Salad, 0)
	for rows.Next() {
		tmp := new(domain.Salad)
		err = rows.Scan(
			&tmp.ID,
			&tmp.Name,
			&tmp.AuthorID,
			&tmp.Description,
		)
		salads = append(salads, tmp)
		if err != nil {
			return nil, 0, fmt.Errorf("scanning salads: %w", err)
		}
	}

	numRows := 0
	err = r.db.QueryRow(
		ctx,
		`with rates as (select salad
	from saladRecipes.comment
	where author = $1)
	select count(*)
	from saladRecipes.salad
	where id in (select salad from rates)`,
	).Scan(&numRows)
	numPages := numRows / config.PageSize
	if numRows%config.PageSize != 0 {
		numPages++
	}

	return salads, numPages, nil
}

func (r *saladRepository) Update(ctx context.Context, salad *domain.Salad) error {
	query := `update saladRecipes.salad
		set
			name = $1,
			authorId = $2,
			description = $3 
		where id = $4`

	_, err := r.db.Exec(
		ctx,
		query,
		salad.Name,
		salad.AuthorID,
		salad.Description,
		salad.ID,
	)
	if err != nil {
		return fmt.Errorf("updating salad: %w", err)
	}
	return nil
}

func (r *saladRepository) DeleteById(ctx context.Context, id uuid.UUID) error {
	query := `delete from saladRecipes.salad
       where id = $1`

	_, err := r.db.Exec(
		ctx,
		query,
		id)
	if err != nil {
		return fmt.Errorf("deleting salad by id: %w", err)
	}
	return nil
}
