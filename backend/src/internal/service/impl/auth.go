package impl

import (
	"backend/src/internal/model/dto"
	repositoryInterface "backend/src/internal/repository/interface"
	"backend/src/pkg/base"
	"backend/src/pkg/logger"
	"context"
	"fmt"
	"strconv"
)

type AuthService struct {
	userRepo repositoryInterface.IUserRepository
	crypto   base.IHashCrypto
	jwtKey   string
	logger   logger.Interface
}

func NewAuthService(logger logger.Interface,
	repo repositoryInterface.IUserRepository,
	crypto base.IHashCrypto,
	jwtKey string) *AuthService {
	return &AuthService{
		logger:   logger,
		userRepo: repo,
		crypto:   crypto,
		jwtKey:   jwtKey,
	}
}

func (s AuthService) SignIn(request *dto.SignInRequest) (err error) {
	s.logger.Infof("регистрация пользователя %s", request.Login)
	if request.Login == "" {
		s.logger.Infof("ошибка при регистрации пользователя: %s", fmt.Errorf("должно быть указано имя пользователя"))
		return fmt.Errorf("должно быть указано имя пользователя")
	}

	if request.Password == "" {
		s.logger.Infof("ошибка при регистрации пользователя: %s", fmt.Errorf("должен быть указан пароль"))
		return fmt.Errorf("должен быть указан пароль")
	}

	if request.FirstName == "" {
		s.logger.Infof("ошибка при регистрации пользователя: %s", fmt.Errorf("должно быть указано имя"))
		return fmt.Errorf("должно быть указано имя")
	}

	if request.SecondName == "" {
		s.logger.Infof("ошибка при регистрации пользователя: %s", fmt.Errorf("должна быть указано фамилия"))
		return fmt.Errorf("должна быть указано фамилия")
	}

	if request.ThirdName == "" {
		s.logger.Infof("ошибка при регистрации пользователя: %s", fmt.Errorf("должно быть указано отчество"))
		return fmt.Errorf("должно быть указано отчество")
	}

	hashedPassword, err := s.crypto.GenerateHashPass(request.Password)

	if err != nil {
		s.logger.Errorf("ошибка при регистрации пользователя: %s", err.Error())
		return fmt.Errorf("генерация хэша: %w", err)
	}

	ctx := context.Background() //, cancel := context.WithTimeout(context.Background(), cmd.TimeOut*time.Second)
	//defer cancel()

	err = s.userRepo.Add(ctx, &dto.AddUserRequest{
		Login:      request.Login,
		Password:   hashedPassword,
		Role:       "client",
		FirstName:  request.FirstName,
		SecondName: request.SecondName,
		ThirdName:  request.ThirdName,
	})
	if err != nil {
		s.logger.Errorf("ошибка при регистрации пользователя: %s", err.Error())
		return fmt.Errorf("регистрация пользователя: %w", err)
	}
	s.logger.Infof("пользователь %s зарегистрировался", request.Login)

	return err
}

func (s AuthService) LogIn(ctx context.Context, request *dto.LogInRequest) (token string, err error) {
	if request.Login == "" {
		s.logger.Infof("ошибка при входе пользователя: %s", fmt.Errorf("должно быть указано имя пользователя"))
		return "", fmt.Errorf("должно быть указано имя пользователя")
	}

	if request.Password == "" {
		s.logger.Infof("ошибка при входе пользователя: %s", fmt.Errorf("должен быть указан пароль"))
		return "", fmt.Errorf("должен быть указан пароль")
	}

	//ctx := context.Background() //, cancel := context.WithTimeout(context.Background(), cmd.TimeOut*time.Second)
	//defer cancel()
	user, err := s.userRepo.GetByLogin(ctx, &dto.GetUserByLoginRequest{
		Login: request.Login,
	})
	if err != nil {
		s.logger.Errorf("ошибка при входе пользователя %s: %s", request.Login, err.Error())
		return "", fmt.Errorf("получение пользователя по login: %w", err) // FIXME: invalid_username
	}

	if !s.crypto.CheckPasswordHash(request.Password, user.Password) {
		s.logger.Warnf("ошибка при регистрации пользователя %d: %s", user.Id, fmt.Errorf("неверный пароль"))
		return "", fmt.Errorf("неверный пароль")
	}

	userId := strconv.Itoa(int(user.Id))

	token, err = base.GenerateAuthToken(userId, request.Login, s.jwtKey, user.Role)
	if err != nil {
		s.logger.Errorf("ошибка при регистрации пользователя %d: %s", user.Id, err.Error())
		return "", fmt.Errorf("генерация токена: %w", err)
	}

	s.logger.Infof("пользователь %s вошел", request.Login)

	return token, err
}
