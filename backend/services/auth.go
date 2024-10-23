package services

import (
	"context"
	"fmt"
	"net/mail"
	"ppo/domain"
	"ppo/pkg/base"
	"ppo/pkg/logger"
)

type AuthService struct {
	logger   logger.ILogger
	authRepo domain.IAuthRepository
	crypto   base.IHashCrypto
	jwtKey   string
}

func NewAuthService(repo domain.IAuthRepository, logger logger.ILogger, crypto base.IHashCrypto, jwtKey string) domain.IAuthService {
	return &AuthService{
		logger:   logger,
		authRepo: repo,
		crypto:   crypto,
		jwtKey:   jwtKey,
	}
}

func (s *AuthService) Register(ctx context.Context, user *domain.User) (string, error) {
	s.logger.Infof("register user with username: %s", user.Username)

	if user.Name == "" {
		s.logger.Warnf("register user: empty name")
		return "", fmt.Errorf("empty name")
	}

	if user.Username == "" {
		s.logger.Warnf("register user: empty username")
		return "", fmt.Errorf("empty username")
	}

	if user.Password == "" {
		s.logger.Warnf("register user: empty password")
		return "", fmt.Errorf("empty password")
	}

	if _, err := mail.ParseAddress(user.Email.Address); err != nil {
		s.logger.Warnf("register user: invalid email (%s)", err.Error())
		return "", fmt.Errorf("invalid email: %w", err)
	}

	hashedPass, err := s.crypto.GenerateHashPass(user.Password)
	if err != nil {
		s.logger.Warnf("register user: generating hash error (%s)", err.Error())
		return "", fmt.Errorf("generating hash: %w", err)
	}

	user.Password = hashedPass

	uid, err := s.authRepo.Register(ctx, user)
	if err != nil {
		s.logger.Errorf("register user: repo error (%s)", err.Error())
		return "", fmt.Errorf("registration user: %w", err)
	}

	token, err := base.GenerateAuthToken(uid.String(), s.jwtKey, domain.DefaultRole)
	if err != nil {
		s.logger.Warnf("login user: geerating auth token error (%s)", err.Error())
		return "", fmt.Errorf("generating token: %w", err)
	}

	return token, nil
}

func (s *AuthService) Login(ctx context.Context, authInfo *domain.UserAuth) (token string, err error) {
	s.logger.Infof("login user with username: %s", authInfo.Username)

	if authInfo.Username == "" {
		s.logger.Warnf("login user: empty username")
		return "", fmt.Errorf("empty username")
	}

	if authInfo.Password == "" {
		s.logger.Warnf("login user: empty password")
		return "", fmt.Errorf("empty password")
	}

	userAuth, err := s.authRepo.GetByUsername(ctx, authInfo.Username)
	if err != nil {
		s.logger.Errorf("login user: getting data from repo error (%s)", err.Error())
		return "", fmt.Errorf("getting user by name: %w", err)
	}

	if !s.crypto.CheckPasswordHash(authInfo.Password, userAuth.HashedPass) {
		s.logger.Warnf("login user: invalid password")
		return "", fmt.Errorf("invalid password")
	}

	token, err = base.GenerateAuthToken(userAuth.ID.String(), s.jwtKey, userAuth.Role)
	if err != nil {
		s.logger.Warnf("login user: geerating auth token error (%s)", err.Error())
		return "", fmt.Errorf("generating token: %w", err)
	}

	return token, nil
}
