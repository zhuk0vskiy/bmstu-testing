//go:build e2e_test

package e2e

import (
	"github.com/Mx1q/ppo_services/tests/mocks"
	"github.com/gavv/httpexpect/v2"
	"github.com/golang/mock/gomock"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
	"net/http"
	"os"
	"ppo/domain"
)

type E2ESuite struct {
	suite.Suite
	ctrl *gomock.Controller

	logger      *mocks.MockILogger
	crypto      *mocks.MockIHashCrypto
	saladS      domain.ISaladInteractor
	authS       domain.IAuthService
	recipeStepS domain.IRecipeStepInteractor
	recipeS     domain.IRecipeService

	e httpexpect.Expect
}

func (s *E2ESuite) BeforeAll(t provider.T) {
	s.ctrl = gomock.NewController(t)

	t.Title("[e2e] init test repository")
	//aRepo := postgres.NewActivityFieldRepository(TestDbInstance)
	//uRepo := postgres.NewUserRepository(TestDbInstance)
	//auRepo := postgres.NewAuthRepository(TestDbInstance)
	//cRepo := postgres.NewCompanyRepository(TestDbInstance)
	//fRepo := postgres.NewFinReportRepository(TestDbInstance)
	//
	s.logger = mocks.NewMockILogger(s.ctrl)
	s.crypto = mocks.NewMockIHashCrypto(s.ctrl)
	//s.aSvc = activity_field.NewService(aRepo, cRepo, s.logger)
	//s.uSvc = user.NewService(uRepo, cRepo, aRepo, s.logger)
	//s.auSvc = auth.NewService(auRepo, s.crypto, "jwt123", s.logger)
	//s.cSvc = company.NewService(cRepo, aRepo, s.logger)
	//s.fSvc = fin_report.NewService(fRepo, s.logger)

	s.e = *httpexpect.WithConfig(httpexpect.Config{
		Client:   &http.Client{},
		BaseURL:  "http://localhost:8083",
		Reporter: httpexpect.NewAssertReporter(t),
		Printers: []httpexpect.Printer{
			httpexpect.NewDebugPrinter(t, true),
		},
	})

	t.Tags("fixture", "e2e")
	done := make(chan os.Signal, 1)
	ok := make(chan struct{}, 2)
	go RunTheApp(testDbInstance, done, ok)
	for {
		select {
		case <-ok:
			return
		}
	}
}

func (s *E2ESuite) Test1(t provider.T) {
	t.Title("[e2e] Test2")
	t.Tags("e2e", "postgres")
	t.Parallel()
	t.WithNewStep("Success", func(sCtx provider.StepCtx) {
		type SignUpReq struct {
			Name     string `json:"name"`
			Username string `json:"username"`
			Password string `json:"password"`
			Email    string `json:"email"`
		}
		sReq := SignUpReq{
			Name:     "name",
			Username: "user",
			Password: "pass",
			Email:    "user@example.com",
		}

		type LoginReq struct {
			Login    string `json:"login"`
			Password string `json:"password"`
		}
		lReq := LoginReq{"user", "pass"}

		s.e.POST("/signup").
			WithJSON(sReq).
			Expect().
			Status(http.StatusOK).
			JSON().
			Object().
			NotEmpty().
			HasValue("status", "success")

		s.e.POST("/login").
			WithJSON(lReq).
			Expect().
			Status(http.StatusOK).
			JSON().
			Object().
			NotEmpty().
			HasValue("status", "success")
	})
}
