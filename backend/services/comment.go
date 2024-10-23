package services

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"ppo/domain"
	"ppo/pkg/logger"
)

type CommentService struct {
	commentRepo domain.ICommentRepository
	logger      logger.ILogger
}

func NewCommentService(commentRepo domain.ICommentRepository, logger logger.ILogger) domain.ICommentService {
	return &CommentService{
		commentRepo: commentRepo,
		logger:      logger,
	}
}

func (s *CommentService) verify(comment *domain.Comment) error {
	if comment.Rating < domain.MinRate || comment.Rating > domain.MaxRate {
		return fmt.Errorf("rate out of range")
	}
	return nil
}

func (s *CommentService) Create(ctx context.Context, comment *domain.Comment) error {
	s.logger.Infof("creating comment by %s to salad %s", comment.AuthorID.String(), comment.SaladID.String())

	err := s.verify(comment)
	if err != nil {
		s.logger.Warnf("creating comment: %s", err.Error())
		return fmt.Errorf("creating comment: %w", err)
	}

	err = s.commentRepo.Create(ctx, comment)
	if err != nil {
		s.logger.Errorf("creating comment: %s", err.Error())
		return fmt.Errorf("creating comment: %w", err)
	}

	return nil
}

func (s *CommentService) Update(ctx context.Context, comment *domain.Comment) error {
	s.logger.Infof("updating comment with id %s", comment.ID.String())

	err := s.verify(comment)
	if err != nil {
		s.logger.Warnf("updating comment: %s", err.Error())
		return fmt.Errorf("updating comment: %w", err)
	}

	err = s.commentRepo.Update(ctx, comment)
	if err != nil {
		s.logger.Errorf("updating comment: %s", err.Error())
		return fmt.Errorf("updating comment: %w", err)
	}
	return nil
}

func (s *CommentService) GetById(ctx context.Context, id uuid.UUID) (*domain.Comment, error) {
	s.logger.Infof("getting comment by id %s", id.String())

	comment, err := s.commentRepo.GetById(ctx, id)
	if err != nil {
		s.logger.Errorf("getting comment by id error %s", err.Error())
		return nil, fmt.Errorf("getting comment by id: %w", err)
	}
	return comment, nil
}

func (s *CommentService) GetBySaladAndUser(ctx context.Context, saladId uuid.UUID, userId uuid.UUID) (*domain.Comment, error) {
	s.logger.Infof("getting comment by salad id %s and user id %s", saladId.String(), userId.String())
	comment, err := s.commentRepo.GetBySaladAndUser(ctx, saladId, userId)
	if err != nil {
		s.logger.Errorf("getting comment by salad and user IDs error: %s", err.Error())
		return nil, fmt.Errorf("getting comment by salad and user IDs: %w", err)
	}
	return comment, nil
}

func (s *CommentService) GetAllBySaladID(ctx context.Context, saladId uuid.UUID, page int) ([]*domain.Comment, int, error) {
	s.logger.Infof("getting all comments by salad id %s", saladId.String())

	comments, numPages, err := s.commentRepo.GetAllBySaladID(ctx, saladId, page)
	if err != nil {
		s.logger.Errorf("getting all comments by salad id error %s", err.Error())
		return nil, 0, fmt.Errorf("getting all comments by salad id: %w", err)
	}
	return comments, numPages, nil
}

func (s *CommentService) DeleteById(ctx context.Context, id uuid.UUID) error {
	s.logger.Infof("deleting comment by id %s", id.String())

	err := s.commentRepo.DeleteById(ctx, id)
	if err != nil {
		s.logger.Errorf("deleting comment by id error %s", err.Error())
		return fmt.Errorf("deleting comment by id: %w", err)
	}
	return nil
}
