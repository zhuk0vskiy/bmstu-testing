package e2e_tests

import (
	v1 "backend/src/web/v1"
	"github.com/gavv/httpexpect/v2"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
	"go.uber.org/mock/gomock"
	"net/http"
	"os"
)

type E2ESuite struct {
	suite.Suite
	ctrl *gomock.Controller

	//logger *mocks.MockILogger
	//crypto *mocks.MockIHashCrypto
	//aSvc  svcInterface.IUserService
	//uSvc  svcInterace.IUserService
	//auSvc svcInterace.IAuthService
	//cSvc  svcInterace.ICompanyService
	//fSvc  svcInterace.IFinancialReportService

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
	//s.logger = mocks.NewMockILogger(s.ctrl)
	//s.crypto = mocks.NewMockIHashCrypto(s.ctrl)
	//s.aSvc = activity_field.NewService(aRepo, cRepo, s.logger)
	//s.uSvc = user.NewService(uRepo, cRepo, aRepo, s.logger)
	//s.auSvc = auth.NewService(auRepo, s.crypto, "jwt123", s.logger)
	//s.cSvc = company.NewService(cRepo, aRepo, s.logger)
	//s.fSvc = fin_report.NewService(fRepo, s.logger)

	s.e = *httpexpect.WithConfig(httpexpect.Config{
		Client:   &http.Client{},
		BaseURL:  "http://localhost:8081/api/v1",
		Reporter: httpexpect.NewAssertReporter(t),
		Printers: []httpexpect.Printer{
			httpexpect.NewDebugPrinter(t, true),
		},
	})

	t.Tags("fixture", "e2e")
	done := make(chan os.Signal, 1)
	ok := make(chan struct{}, 2)
	go RunTheApp(TestDbInstance, done, ok)
	for {
		select {
		case <-ok:
			return
		}
	}
}

func (s *E2ESuite) Test2(t provider.T) {
	t.Title("[e2e] Test2")
	t.Tags("e2e", "postgres")
	t.Parallel()
	t.WithNewStep("Success", func(sCtx provider.StepCtx) {

		regReq := &v1.SignInRequest{
			Login:      "nik",
			Password:   "nik",
			FirstName:  "nik",
			SecondName: "nik",
			ThirdName:  "nik",
		}
		logReq := &v1.LogInRequest{
			Login:    "nik",
			Password: "nik",
		}

		s.e.POST("/signin").
			WithJSON(regReq).
			Expect().
			Status(http.StatusOK).
			JSON().
			Object().
			NotEmpty().
			HasValue("status", "success")

		s.e.POST("/login").
			WithJSON(logReq).
			Expect().
			Status(http.StatusOK).
			JSON().
			Object().
			NotEmpty().
			HasValue("status", "success")
		//
		//addRoomReq := &v1.AddRoomRequest{
		//	Name:      "e",
		//	StudioId:  1,
		//	StartHour: 1,
		//	EndHour:   2,
		//}
		//s.e.POST("/rooms/").
		//	WithJSON(addRoomReq).
		//	Expect().
		//	Status(http.StatusUnauthorized)
		//JSON().
		//Object().
		//NotEmpty().
		//HasValue("status", "success")

		//errorAddRoomReq := &v1.AddRoomRequest{
		//	Name:      "b",
		//	StudioId:  1,
		//	StartHour: 2,
		//	EndHour:   1,
		//}
		//s.e.POST("/rooms").
		//	WithJSON(errorAddRoomReq).
		//	Expect().
		//	Status(http.StatusUnauthorized)
		//JSON().
		//Object().
		//NotEmpty()
		//HasValue("status", "error")
	})
}
