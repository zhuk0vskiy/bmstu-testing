package postgres

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"ppo/domain"
)

type keywordValidatorRepository struct {
	//db *pgxpool.Pool
	db IPool
}

func NewKeywordValidatorRepository(db IPool) domain.IKeywordValidatorRepository {
	return &keywordValidatorRepository{
		db: db,
	}
}

func (r *keywordValidatorRepository) Create(ctx context.Context, word *domain.KeyWord) error {
	query := `insert into keywords.word(word)
		values ($1)`

	_, err := r.db.Exec(
		ctx,
		query,
		word.Word)
	if err != nil {
		return fmt.Errorf("creating keyword: %w", err)
	}
	return nil
}

func (r *keywordValidatorRepository) GetById(ctx context.Context, id uuid.UUID) (*domain.KeyWord, error) {
	query := `select id, word
		from keywords.word
		where id = $1`

	keyword := new(domain.KeyWord)
	err := r.db.QueryRow(
		ctx,
		query,
		id,
	).Scan(
		&keyword.ID,
		&keyword.Word,
	)

	if err != nil {
		return nil, fmt.Errorf("getting keyword by id: %w", err)
	}
	return keyword, nil
}

func (r *keywordValidatorRepository) GetAll(ctx context.Context) (map[string]uuid.UUID, error) {
	query := `select id, word
		from keywords.word`

	rows, err := r.db.Query(
		ctx,
		query,
	)
	if err != nil {
		return nil, fmt.Errorf("getting all keywords: %w", err)
	}

	keywords := make(map[string]uuid.UUID)
	for rows.Next() {
		tmp := new(domain.KeyWord)
		err = rows.Scan(
			&tmp.ID,
			&tmp.Word,
		)
		if err != nil {
			return nil, fmt.Errorf("scanning keywords: %w", err)
		}
		keywords[tmp.Word] = tmp.ID
	}
	return keywords, nil
}

func (r *keywordValidatorRepository) Update(ctx context.Context, word *domain.KeyWord) error {
	query := `update keywords.word
		set
			word = $1
		where id = $2`

	_, err := r.db.Exec(
		ctx,
		query,
		word.Word,
		word.ID,
	)
	if err != nil {
		return fmt.Errorf("updating keyword: %w", err)
	}
	return nil
}

func (r *keywordValidatorRepository) DeleteById(ctx context.Context, id uuid.UUID) error {
	query := `delete from keywords.word
       where id = $1`

	_, err := r.db.Exec(
		ctx,
		query,
		id)
	if err != nil {
		return fmt.Errorf("deleting keyword by id: %w", err)
	}
	return nil
}
