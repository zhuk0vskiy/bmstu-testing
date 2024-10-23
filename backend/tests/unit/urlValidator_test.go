//go:build unit_test

package unit_tests

import (
	"context"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
	"ppo/services"
	"ppo/tests/utils"
)

type UrlValidatorSuite struct {
	suite.Suite
}

func (suite *SaladTypeSuite) TestUrlValidatorService_Verify1(t provider.T) {
	t.Title("[Url validator verify] verified")
	t.Tags("url_validator", "verify")
	t.Parallel()

	t.WithNewStep("verified", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		str := "not_url"

		logger := utils.NewMockLogger()
		service := services.NewUrlValidatorService(logger)

		sCtx.WithNewParameters("ctx", ctx, "request", str)
		err := service.Verify(ctx, str)

		sCtx.Assert().NoError(err)
	})
}

func (suite *SaladTypeSuite) TestUrlValidatorService_Verify2(t provider.T) {
	t.Title("[Url validator verify] verification failed")
	t.Tags("url_validator", "verify")
	t.Parallel()

	t.WithNewStep("verification failed", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		str := "url.com"

		logger := utils.NewMockLogger()
		service := services.NewUrlValidatorService(logger)

		sCtx.WithNewParameters("ctx", ctx, "request", str)
		err := service.Verify(ctx, str)

		sCtx.Assert().Error(err)
	})
}
