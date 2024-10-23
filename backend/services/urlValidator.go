package services

import (
	"context"
	"fmt"
	"github.com/asaskevich/govalidator"
	"ppo/domain"
	"ppo/pkg/logger"
	"strings"
)

type UrlValidatorService struct {
	logger logger.ILogger
}

func NewUrlValidatorService(logger logger.ILogger) domain.IValidatorService {
	return &UrlValidatorService{
		logger: logger,
	}
}

func (s UrlValidatorService) Verify(ctx context.Context, word string) error {
	if checkWord := strings.Fields(word); len(checkWord) > 1 {
		s.logger.Warnf("verifying url: accepts only 1 word")
		return fmt.Errorf("verifying url: accepts only 1 word")
	}
	if govalidator.IsURL(word) {
		s.logger.Warnf("verifying url: found %s", word)
		return fmt.Errorf("verifying url: found %s", word)
	}

	return nil
}
