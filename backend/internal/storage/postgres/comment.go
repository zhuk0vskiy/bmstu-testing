package postgres

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"ppo/domain"
	"ppo/internal/config"
)

type commentRepository struct {
	//db *pgxpool.Pool
	db IPool
}

func NewCommentRepository(db IPool) domain.ICommentRepository {
	return &commentRepository{
		db: db,
	}
}

func (r *commentRepository) Create(ctx context.Context, comment *domain.Comment) error {
	query := `insert into saladRecipes.comment(author, salad, text, rating)
		values ($1, $2, $3, $4)`

	_, err := r.db.Exec(
		ctx,
		query,
		comment.AuthorID,
		comment.SaladID,
		comment.Text,
		comment.Rating)
	if err != nil {
		return fmt.Errorf("creating comment: %w", err)
	}
	return nil
}

func (r *commentRepository) GetById(ctx context.Context, id uuid.UUID) (*domain.Comment, error) {
	query := `select id, author, salad, text, rating
		from saladRecipes.comment
		where id = $1`

	comment := new(domain.Comment)
	err := r.db.QueryRow(
		ctx,
		query,
		id,
	).Scan(
		&comment.ID,
		&comment.AuthorID,
		&comment.SaladID,
		&comment.Text,
		&comment.Rating)

	if err != nil {
		return nil, fmt.Errorf("getting comment by id: %w", err)
	}
	return comment, nil
}

func (r *commentRepository) GetBySaladAndUser(ctx context.Context, saladId uuid.UUID, userId uuid.UUID) (*domain.Comment, error) {
	query := `select id, author, salad, text, rating
		from saladRecipes.comment
		where salad = $1 and author = $2`

	comment := new(domain.Comment)
	err := r.db.QueryRow(
		ctx,
		query,
		saladId,
		userId,
	).Scan(
		&comment.ID,
		&comment.AuthorID,
		&comment.SaladID,
		&comment.Text,
		&comment.Rating)

	if err != nil {
		return nil, fmt.Errorf("getting comment by salad and user IDs: %w", err)
	}
	return comment, nil
}

func (r *commentRepository) GetAllBySaladID(ctx context.Context, saladId uuid.UUID, page int) ([]*domain.Comment, int, error) {
	query := `select id, author, salad, text, rating
		from saladRecipes.comment
		where salad = $1
		offset $2
		limit $3`

	rows, err := r.db.Query(
		ctx,
		query,
		saladId,
		config.PageSize*(page-1),
		config.PageSize,
	)
	if err != nil {
		return nil, 0, fmt.Errorf("getting comments by salad id: %w", err)
	}

	comments := make([]*domain.Comment, 0)
	for rows.Next() {
		tmp := new(domain.Comment)
		err = rows.Scan(
			&tmp.ID,
			&tmp.AuthorID,
			&tmp.SaladID,
			&tmp.Text,
			&tmp.Rating,
		)
		comments = append(comments, tmp)
		if err != nil {
			return nil, 0, fmt.Errorf("scanning comments: %w", err)
		}
	}

	rowsQuery := `select count(*) from saladRecipes.comment
	where salad = $1`
	numRows := 0
	err = r.db.QueryRow(
		ctx,
		rowsQuery,
		saladId,
	).Scan(&numRows)
	numPages := numRows / config.PageSize
	if numRows%config.PageSize != 0 {
		numPages++
	}

	return comments, numPages, nil
}

func (r *commentRepository) Update(ctx context.Context, comment *domain.Comment) error {
	query := `update saladRecipes.comment
		set
			rating = $1,
			text = $2
		where id = $3`

	_, err := r.db.Exec(
		ctx,
		query,
		comment.Rating,
		comment.Text,
		comment.ID,
	)
	if err != nil {
		return fmt.Errorf("updating comment: %w", err)
	}
	return nil
}

func (r *commentRepository) DeleteById(ctx context.Context, id uuid.UUID) error {
	query := `delete from saladRecipes.comment
       where id = $1`

	_, err := r.db.Exec(
		ctx,
		query,
		id)
	if err != nil {
		return fmt.Errorf("deleting comment by id: %w", err)
	}
	return nil
}
