package services

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"net/mail"
	"ppo/domain"
	"ppo/pkg/logger"
)

type UserService struct {
	userRepo domain.IUserRepository
	logger   logger.ILogger
}

func NewUserService(userRepo domain.IUserRepository, logger logger.ILogger) domain.IUserService {
	return &UserService{
		userRepo: userRepo,
		logger:   logger,
	}
}

func (s *UserService) verify(user *domain.User) error {
	if user.Username == "" {
		return fmt.Errorf("empty username")
	}

	if user.Password == "" {
		return fmt.Errorf("empty password")
	}

	if user.Name == "" {
		return fmt.Errorf("empty name")
	}

	if _, err := mail.ParseAddress(user.Email.Address); err != nil {
		return err
	}

	return nil
}

func (s *UserService) Create(ctx context.Context, user *domain.User) error {
	s.logger.Infof("creating user: %v", user)

	err := s.verify(user)
	if err != nil {
		s.logger.Warnf("failed to verify user: %s", err.Error())
		return fmt.Errorf("creating user: %w", err)
	}

	err = s.userRepo.Create(ctx, user)
	if err != nil {
		s.logger.Errorf("creating user error: %s", err.Error())
		return fmt.Errorf("creating user: %w", err)
	}

	return nil
}

func (s *UserService) Update(ctx context.Context, user *domain.User) error {
	s.logger.Warnf("updating user: %v", user)

	err := s.verify(user)
	if err != nil {
		s.logger.Warnf("failed to verify user: %s", err.Error())
		return fmt.Errorf("updating user: %w", err)
	}

	err = s.userRepo.Update(ctx, user)
	if err != nil {
		s.logger.Errorf("updating user error: %s", err.Error())
		return fmt.Errorf("updating user: %w", err)
	}
	return nil
}

func (s *UserService) GetById(ctx context.Context, id uuid.UUID) (*domain.User, error) {
	s.logger.Infof("getting user by id: %s", id.String())

	user, err := s.userRepo.GetById(ctx, id)
	if err != nil {
		s.logger.Errorf("getting user by id error: %s", err.Error())
		return nil, fmt.Errorf("getting user by id: %w", err)
	}
	return user, nil
}

func (s *UserService) GetAll(ctx context.Context, page int) ([]*domain.User, error) {
	s.logger.Infof("getting all users on page %d", page)

	users, err := s.userRepo.GetAll(ctx, page)
	if err != nil {
		s.logger.Errorf("getting all users error: %s", err.Error())
		return nil, fmt.Errorf("getting all users: %w", err)
	}
	return users, nil
}

func (s *UserService) DeleteById(ctx context.Context, id uuid.UUID) error {
	s.logger.Infof("deleting user by id: %s", id.String())

	err := s.userRepo.DeleteById(ctx, id)
	if err != nil {
		s.logger.Errorf("deleting user by id error: %s", err.Error())
		return fmt.Errorf("deleting user by id: %w", err)
	}
	return nil
}

func (s *UserService) GetByUsername(ctx context.Context, username string) (*domain.User, error) {
	s.logger.Infof("getting user by username: %s", username)

	user, err := s.userRepo.GetByUsername(ctx, username)
	if err != nil {
		s.logger.Errorf("getting user by username error: %s", err.Error())
		return nil, fmt.Errorf("getting user by username: %w", err)
	}
	return user, nil
}
