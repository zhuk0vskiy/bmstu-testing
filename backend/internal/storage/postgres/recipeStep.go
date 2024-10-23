package postgres

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"ppo/domain"
)

type recipeStepRepository struct {
	//db *pgxpool.Pool
	db IPool
}

func NewRecipeStepRepository(db IPool) domain.IRecipeStepRepository {
	return &recipeStepRepository{
		db: db,
	}
}

func (r *recipeStepRepository) Create(ctx context.Context, recipeStep *domain.RecipeStep) error {
	query := `with freeRecipeStep as (select case
    		when max(stepNum) is null then 1
    		else max(stepNum) + 1
			end as freeStep
		from saladRecipes.recipeStep
		where recipeId = $3)
	insert into saladRecipes.recipeStep(name, description, recipeId, stepNum)
		values ($1, $2, $3, (select freeStep from freeRecipeStep))`

	_, err := r.db.Exec(
		ctx,
		query,
		recipeStep.Name,
		recipeStep.Description,
		recipeStep.RecipeID)
	if err != nil {
		return fmt.Errorf("creating recipe step: %w", err)
	}

	return nil
}

func (r *recipeStepRepository) GetById(ctx context.Context, id uuid.UUID) (*domain.RecipeStep, error) {
	query := `select id, name, description, recipeId, stepNum
		from saladRecipes.recipeStep
		where id = $1`

	step := new(domain.RecipeStep)
	err := r.db.QueryRow(
		ctx,
		query,
		id,
	).Scan(
		&step.ID,
		&step.Name,
		&step.Description,
		&step.RecipeID,
		&step.StepNum)

	if err != nil {
		return nil, fmt.Errorf("getting recipe step by id: %w", err)
	}
	return step, nil
}

func (r *recipeStepRepository) GetAllByRecipeID(ctx context.Context, recipeId uuid.UUID) ([]*domain.RecipeStep, error) {
	query := `select id, name, description, recipeId, stepNum
		from saladRecipes.recipeStep
		where recipeId = $1
		order by stepNum`
	rows, err := r.db.Query(
		ctx,
		query,
		recipeId,
	)
	if err != nil {
		return nil, fmt.Errorf("getting all recipe steps: %w", err)
	}

	steps := make([]*domain.RecipeStep, 0)
	for rows.Next() {
		tmp := new(domain.RecipeStep)
		err = rows.Scan(
			&tmp.ID,
			&tmp.Name,
			&tmp.Description,
			&tmp.RecipeID,
			&tmp.StepNum,
		)
		steps = append(steps, tmp)
		if err != nil {
			return nil, fmt.Errorf("scanning recipe steps: %w", err)
		}
	}

	return steps, nil
}

func (r *recipeStepRepository) Update(ctx context.Context, recipeStep *domain.RecipeStep) error {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return fmt.Errorf("create transaction: %w", err)
	}
	defer func() {
		if err != nil {
			rollbackErr := tx.Rollback(ctx)
			if rollbackErr != nil {
				err = fmt.Errorf("%v: %w", rollbackErr, err)
			}
		}
	}()

	maxStepNum := 0
	prevStepNum := 0
	selectMaxStepNumQuery := `with maxStepNum as (select case
    	when max(stepNum) is null then 0
    	else max(stepNum)
	end as maxNum
	from saladRecipes.recipeStep
	where recipeId = $1)
	select (select maxNum from maxStepNum), stepNum
	from saladRecipes.recipeStep
	where id = $2`
	err = tx.QueryRow(
		ctx,
		selectMaxStepNumQuery,
		recipeStep.RecipeID,
		recipeStep.ID,
	).Scan(
		&maxStepNum,
		&prevStepNum)
	if err != nil {
		return fmt.Errorf("updating recipe step (checking max step num): %w", err)
	}
	if recipeStep.StepNum > maxStepNum {
		return fmt.Errorf("updating recipe step: step num out of range")
	}

	changeStepNumsQuery := `update saladRecipes.recipeStep
	set
    	stepNum = stepNum + $4
	where recipeId = $1 and stepNum between $2+1 and $3`
	change := -1
	if recipeStep.StepNum < prevStepNum {
		change = 1
		changeStepNumsQuery = `update saladRecipes.recipeStep
		set
    		stepNum = stepNum + $4
		where recipeId = $1 and stepNum between $3 and $2-1`
	}
	_, err = tx.Exec(
		ctx,
		changeStepNumsQuery,
		recipeStep.RecipeID,
		prevStepNum,
		recipeStep.StepNum,
		change,
	)
	if err != nil {
		return fmt.Errorf("updating recipe step (moving other steps): %w", err)
	}

	query := `update saladRecipes.recipeStep
		set
			name = $1,
			description = $2,
			stepNum = $3
		where id = $4`
	_, err = tx.Exec(
		ctx,
		query,
		recipeStep.Name,
		recipeStep.Description,
		recipeStep.StepNum,
		recipeStep.ID,
	)
	if err != nil {
		return fmt.Errorf("updating recipe step: %w", err)
	}

	err = tx.Commit(ctx)
	if err != nil {
		return fmt.Errorf("updating recipe step (commiting transaction): %w", err)
	}
	return nil
}

func (r *recipeStepRepository) DeleteById(ctx context.Context, id uuid.UUID) error {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return fmt.Errorf("create transaction: %w", err)
	}
	defer func() {
		if err != nil {
			rollbackErr := tx.Rollback(ctx)
			if rollbackErr != nil {
				err = fmt.Errorf("%v: %w", rollbackErr, err)
			}
		}
	}()

	var recipeId uuid.UUID
	stepNum := 0
	selectRecipeIdQuery := `select recipeId, stepNum
	from saladRecipes.recipeStep
	where id = $1`
	err = tx.QueryRow(
		ctx,
		selectRecipeIdQuery,
		id,
	).Scan(
		&recipeId,
		&stepNum,
	)
	if err != nil {
		return fmt.Errorf("deleting recipe step by id (getting recipe ID): %w", err)
	}

	changeStepNumsQuery := `update saladRecipes.recipeStep
	set
    	stepNum = stepNum - 1
	where recipeId = $1 and stepNum > $2`
	_, err = tx.Exec(
		ctx,
		changeStepNumsQuery,
		recipeId,
		stepNum,
	)
	if err != nil {
		return fmt.Errorf("updating recipe step (moving other steps): %w", err)
	}

	query := `delete from saladRecipes.recipeStep
       where id = $1`
	_, err = tx.Exec(
		ctx,
		query,
		id)
	if err != nil {
		return fmt.Errorf("deleting recipe step by id: %w", err)
	}

	err = tx.Commit(ctx)
	if err != nil {
		return fmt.Errorf("deleting recipe step (commiting transaction): %w", err)
	}
	return nil
}

func (r *recipeStepRepository) DeleteAllByRecipeID(ctx context.Context, recipeId uuid.UUID) error {
	query := `delete from saladRecipes.recipeStep
       where recipeId = $1`

	_, err := r.db.Exec(
		ctx,
		query,
		recipeId)
	if err != nil {
		return fmt.Errorf("deleting all recipe steps by id: %w", err)
	}
	return nil
}
