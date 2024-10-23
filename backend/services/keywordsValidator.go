package services

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"ppo/domain"
	"ppo/pkg/logger"
	"strings"
)

type KeywordValidatorService struct {
	validatorRepo domain.IKeywordValidatorRepository
	logger        logger.ILogger
	keywords      map[string]uuid.UUID
}

func NewKeywordValidatorService(ctx context.Context, validatorRepo domain.IKeywordValidatorRepository, logger logger.ILogger) (domain.IKeywordValidatorService, error) {
	keywords, err := validatorRepo.GetAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("creating keywords validator: %w", err)
	}

	return &KeywordValidatorService{
		validatorRepo: validatorRepo,
		keywords:      keywords,
		logger:        logger,
	}, nil
}

func (s *KeywordValidatorService) verifyWord(word *domain.KeyWord) error {
	if word.Word == "" {
		return fmt.Errorf("empty word")
	}
	if words := strings.Fields(word.Word); len(words) > 1 {
		return fmt.Errorf("accepts only 1 word")
	}
	return nil
}

func (s *KeywordValidatorService) Create(ctx context.Context, word *domain.KeyWord) error {
	s.logger.Infof("creating keyword: %s", word.Word)

	err := s.verifyWord(word)
	if err != nil {
		s.logger.Warnf("failed to verify word: %s", err.Error())
		return fmt.Errorf("creating keyword: %w", err)
	}

	err = s.validatorRepo.Create(ctx, word)
	if err != nil {
		s.logger.Errorf("creating keyword error: %s", err.Error())
		return fmt.Errorf("creating keyword: %w", err)
	}
	s.keywords[word.Word] = word.ID
	return nil
}

func (s *KeywordValidatorService) Update(ctx context.Context, word *domain.KeyWord) error {
	s.logger.Infof("updating keyword: %s", word.Word)

	err := s.verifyWord(word)
	if err != nil {
		s.logger.Warnf("failed to verify word: %s", err.Error())
		return fmt.Errorf("updating keyword: %w", err)
	}

	err = s.validatorRepo.Update(ctx, word)
	if err != nil {
		s.logger.Errorf("updating keyword error: %s", err.Error())
		return fmt.Errorf("updating keyword: %w", err)
	}
	s.keywords[word.Word] = word.ID
	return nil
}

func (s *KeywordValidatorService) GetById(ctx context.Context, id uuid.UUID) (*domain.KeyWord, error) {
	s.logger.Infof("getting keyword by id: %s", id.String())

	word, err := s.validatorRepo.GetById(ctx, id)
	if err != nil {
		s.logger.Errorf("getting keyword by id error: %s", err.Error())
		return nil, fmt.Errorf("getting keyword by id: %w", err)
	}
	return word, nil
}

func (s *KeywordValidatorService) GetAll(ctx context.Context) (map[string]uuid.UUID, error) {
	s.logger.Infof("getting all keywords")

	words, err := s.validatorRepo.GetAll(ctx)
	if err != nil {
		s.logger.Errorf("getting all keywords error: %s", err.Error())
		return nil, fmt.Errorf("getting all keywords: %w", err)
	}
	return words, nil
}

func (s *KeywordValidatorService) DeleteById(ctx context.Context, id uuid.UUID) error {
	s.logger.Infof("deleting keyword by id: %s", id.String())

	err := s.validatorRepo.DeleteById(ctx, id)
	if err != nil {
		s.logger.Errorf("deleting keyword by id error: %s", err.Error())
		return fmt.Errorf("deleting keyword by id: %w", err)
	}
	return nil
}

func (s *KeywordValidatorService) Verify(ctx context.Context, word string) error {
	if checkWord := strings.Fields(word); len(checkWord) > 1 {
		s.logger.Warnf("verifying keywords: accepts only 1 word")
		return fmt.Errorf("verifying keywords: accepts only 1 word")
	}

	_, ok := s.keywords[strings.ToLower(word)]
	if ok {
		s.logger.Warnf("verifying keywords: found %s", word)
		return fmt.Errorf("verifying keywords: found %s", word)
	}

	return nil
}
