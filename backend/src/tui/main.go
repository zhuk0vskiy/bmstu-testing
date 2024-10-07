package tui

import (
	"backend/src/config"
	"backend/src/internal/model"
	"backend/src/internal/model/dto"
	"backend/src/internal/repository/impl/postgresql"
	serviceImpl "backend/src/internal/service/impl"
	serviceInterface "backend/src/internal/service/interface"
	"backend/src/pkg/base"
	"backend/src/pkg/logger"
	"backend/src/pkg/time_parser"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rivo/tview"
	"log"
	"strconv"
	"strings"
)

type App struct {
	AuthSvc            serviceInterface.IAuthService
	UserSvc            serviceInterface.IUserService
	StudioSvc          serviceInterface.IStudioService
	RoomSvc            serviceInterface.IRoomService
	ProducerSvc        serviceInterface.IProducerService
	InstrumentalistSvc serviceInterface.IInstrumentalistService
	EquipmentSvc       serviceInterface.IEquipmentService
	ReserveSvc         serviceInterface.IReserveService
	ValidateTimeSvc    serviceInterface.IValidateTimeService
	Config             config.Config
}

func NewApp(db *pgxpool.Pool, cfg *config.Config, logger logger.Interface) *App {
	//authRepo := postgresql.NewA
	userRepo := postgresql.NewUserRepository(db)
	studioRepo := postgresql.NewStudioRepository(db)
	roomRepo := postgresql.NewRoomRepository(db)
	producerRepo := postgresql.NewProducerRepository(db)
	instrumentalistRepo := postgresql.NewInstrumentalistRepository(db)
	equipmentRepo := postgresql.NewEquipmentRepository(db)
	reserveRepo := postgresql.NewReserveRepository(db)

	crypto := base.NewHashCrypto()

	authSvc := serviceImpl.NewAuthService(logger, userRepo, crypto, cfg.JwtKey)
	userSvc := serviceImpl.NewUserService(logger, userRepo, reserveRepo, crypto)
	studioSvc := serviceImpl.NewStudioService(logger, studioRepo)
	roomSvc := serviceImpl.NewRoomService(roomRepo, reserveRepo)
	producerSvc := serviceImpl.NewProducerService(logger, producerRepo, reserveRepo)
	instrumentalistSvc := serviceImpl.NewInstrumentalistService(logger, instrumentalistRepo, reserveRepo)
	equipmentSvc := serviceImpl.NewEquipmentService(logger, equipmentRepo, reserveRepo)
	reserveSvc := serviceImpl.NewReserveService(logger, reserveRepo)
	validateTimeSvc := serviceImpl.NewValidateTimeService(logger, roomRepo, equipmentRepo, producerRepo, instrumentalistRepo, reserveRepo)

	return &App{
		AuthSvc:            authSvc,
		UserSvc:            userSvc,
		StudioSvc:          studioSvc,
		RoomSvc:            roomSvc,
		ProducerSvc:        producerSvc,
		InstrumentalistSvc: instrumentalistSvc,
		EquipmentSvc:       equipmentSvc,
		ReserveSvc:         reserveSvc,
		ValidateTimeSvc:    validateTimeSvc,
		Config:             *cfg,
	}
}

var (
	pages     = tview.NewPages()
	app       = tview.NewApplication()
	form      = tview.NewForm()
	errorForm = tview.NewForm()
	list      = tview.NewList().ShowSecondaryText(true)
)

type Tui struct {
	app      *App
	userInfo *base.JwtPayload
}

const (
	GuestPage      = "Menu (guest)"
	AuthorizedPage = "Menu (authorized)"
	AdminPage      = "Menu (admin)"

	GetAll                    = "Посмотреть все брони"
	RegisterPage              = "Регистрация"
	LoginPage                 = "Авторизация"
	ErrorPage                 = "Error page"
	CreateValidatePage        = "Создать время"
	CreateReservePage         = "Создать бронь"
	DeleteReservePage         = "Удалить бронь"
	GetReservesPage           = "Посмотреть брони"
	UpdateUserPage            = "Обновить данные пользователя"
	CreateAddMenu             = "Меню добавления"
	AddRoomPage               = "Добавить комнату"
	AddProducerPage           = "Добавить продюсера"
	AddInstrumentalistPage    = "Добавить инструменталиста"
	AddStudioPage             = "Добавить студию"
	AddEquipmentPage          = "Добавить оборудование"
	CreateDeleteMenu          = "Меню удаления"
	DeleteRoomPage            = "Удалить комнату"
	DeleteProducerPage        = "Уадлить продюсера"
	DeleteInstrumentalistPage = "Удалить инструменталиста"
	DeleteStudioPage          = "Удалить студию"
	DeleteEquipmentPage       = "Удалить оборудование"
	CreateUpdateMenu          = "Меню изменений"
	UpdateRoomPage            = "Изменить комнату"
	UpdateProducerPage        = "Изменить продюсера"
	UpdateInstrumentalistPage = "Изменить инструменталиста"
	UpdateStudioPage          = "Изменить студию"
	UpdateEquipmentPage       = "Изменить оборудование"
)

const (
	UnauthorizedUser = "unauthorized"
	AuthorizedUser   = "client"
	AuthorizedAdmin  = "admin"
)

func Run(db *pgxpool.Pool, cfg *config.Config, logger logger.Interface) *tview.Application {
	var tui Tui
	tui.app = NewApp(db, cfg, logger)
	tui.userInfo = new(base.JwtPayload)
	tui.userInfo.Role = UnauthorizedUser

	pages.AddPage(GuestPage, tui.CreateGuestMenu(form, pages, app), true, true).
		AddPage(RegisterPage, form, true, true).
		AddPage(LoginPage, form, true, true).
		AddPage(GetAll, form, true, true)

	pages.AddPage(AuthorizedPage, tui.CreateAuthorizedMenu(form, pages, app), true, true)

	pages.AddPage(CreateValidatePage, form, true, true).
		AddPage(CreateReservePage, form, true, true).
		AddPage(DeleteReservePage, form, true, true).
		AddPage(GetReservesPage, form, true, true).
		AddPage(UpdateUserPage, form, true, true)

	pages.AddPage(AdminPage, tui.CreateAdminMenu(form, pages, app), true, true).
		AddPage(CreateAddMenu, tui.CreateAddMenu(form, pages, app), true, true).
		AddPage(AddRoomPage, form, true, true).
		AddPage(AddProducerPage, form, true, true).
		AddPage(AddInstrumentalistPage, form, true, true).
		AddPage(AddStudioPage, form, true, true).
		AddPage(AddEquipmentPage, form, true, true).
		AddPage(CreateDeleteMenu, tui.CreateDeleteMenu(form, pages, app), true, true).
		AddPage(DeleteRoomPage, form, true, true).
		AddPage(DeleteProducerPage, form, true, true).
		AddPage(DeleteInstrumentalistPage, form, true, true).
		AddPage(DeleteStudioPage, form, true, true).
		AddPage(DeleteEquipmentPage, form, true, true).
		AddPage(CreateUpdateMenu, tui.CreateUpdateMenu(form, pages, app), true, true).
		AddPage(UpdateRoomPage, form, true, true).
		AddPage(UpdateProducerPage, form, true, true).
		AddPage(UpdateInstrumentalistPage, form, true, true).
		AddPage(UpdateStudioPage, form, true, true).
		AddPage(UpdateEquipmentPage, form, true, true)

	pages.AddPage(ErrorPage, errorForm, true, true)
	//AddPage(CreateRecipeStep, form, true, true)
	// 		tui.CreateReservePage(form, pages, notReservedRooms, notReservedEquipments, notReservedProducers, notReservedInstrumentalists)

	//pages.AddPage(AdminPage, tui.CreateAdminMenu(form, pages, app), true, true)

	pages.SwitchToPage(GuestPage)
	if err := app.SetRoot(pages, true).EnableMouse(true).Run(); err != nil {
		log.Fatalln(err)
	}

	return app
}

func (tui *Tui) ErrorForm(form *tview.Form, pages *tview.Pages, textView *tview.TextView, prevPage string) *tview.Form {
	//form.Clear(true)
	form.AddFormItem(textView)

	form.AddButton("OK", func() {
		pages.SwitchToPage(prevPage)
		form.Clear(true)
		return
	})

	return form
}

func (tui *Tui) CreateGuestMenu(form *tview.Form, pages *tview.Pages, exitFunc *tview.Application) *tview.List {
	return tview.NewList().
		AddItem("Посмотреть брони", "", '1', func() {
			form.Clear(true)
			tui.GetAllReservesPage(form, pages)
			pages.SwitchToPage(GetAll)
		}).
		AddItem("Регистрация", "", '2', func() {
			form.Clear(true)
			tui.RegisterForm(form, pages)
			pages.SwitchToPage(RegisterPage)
		}).
		AddItem("Вход в профиль", "", '3', func() {
			form.Clear(true)
			tui.LoginForm(form, pages)
			pages.SwitchToPage(LoginPage)
		}).
		AddItem("Выход", "", '0', func() {
			exitFunc.Stop()
		})
}

func (tui *Tui) RegisterForm(form *tview.Form, pages *tview.Pages) *tview.Form {
	user := &model.User{}
	errorTextView := tview.NewTextView().
		SetDynamicColors(true).
		SetRegions(true).
		SetWrap(true).
		SetTextAlign(tview.AlignCenter).
		SetText("")

	form.AddInputField("Логин", "", 20, nil, func(username string) {
		user.Login = username
	})
	form.AddPasswordField("Пароль", "", 20, '*', func(password string) {
		user.Password = password
	})
	form.AddInputField("Имя", "", 20, nil, func(name string) {
		user.FirstName = name
	})
	form.AddInputField("Фамилия", "", 20, nil, func(name string) {
		user.SecondName = name
	})
	form.AddInputField("Отчество", "", 20, nil, func(name string) {
		user.ThirdName = name
	})

	form.AddButton("Зарегистрироваться", func() {
		err := tui.app.AuthSvc.SignIn(&dto.SignInRequest{
			Login:      user.Login,
			Password:   user.Password,
			FirstName:  user.FirstName,
			SecondName: user.SecondName,
			ThirdName:  user.ThirdName,
		})
		if err != nil {
			errorTextView.SetText(err.Error())
			tui.ErrorForm(errorForm, pages, errorTextView, RegisterPage)
			pages.SwitchToPage(ErrorPage)
			return
		}
		pages.SwitchToPage(AuthorizedPage)

		fullDataUser, _ := tui.app.UserSvc.GetByLogin(&dto.GetUserByLoginRequest{
			Login: user.Login,
		})
		tui.userInfo.Username = fullDataUser.Login
		tui.userInfo.Role = fullDataUser.Role
	})
	form.AddButton("Назад", func() {
		pages.SwitchToPage(GuestPage)
	})

	return form
}

func (tui *Tui) LoginForm(form *tview.Form, pages *tview.Pages) *tview.Form {
	authInfo := &model.User{}
	errorTextView := tview.NewTextView().
		SetDynamicColors(true).
		SetRegions(true).
		SetWrap(true).
		SetTextAlign(tview.AlignCenter).
		SetText("")

	form.AddInputField("Логин", "", 20, nil, func(username string) {
		authInfo.Login = username
	})
	form.AddPasswordField("Пароль", "", 20, '*', func(password string) {
		authInfo.Password = password
	})

	form.AddButton("Войти", func() {
		token, err := tui.app.AuthSvc.LogIn(&dto.LogInRequest{
			Login:    authInfo.Login,
			Password: authInfo.Password,
		})
		if err != nil {
			errorTextView.SetText(err.Error())
			tui.ErrorForm(errorForm, pages, errorTextView, LoginPage)
			pages.SwitchToPage(ErrorPage)
			return
		}
		userInfo, err := base.VerifyAuthToken(token, tui.app.Config.JwtKey)
		if err != nil {
			errorTextView.SetText("Ошикба")
			tui.ErrorForm(errorForm, pages, errorTextView, LoginPage)
			pages.SwitchToPage(ErrorPage)
			return
		}
		//fmt.Println(userInfo)
		tui.userInfo = userInfo
		if userInfo.Role == AuthorizedAdmin {
			pages.SwitchToPage(AdminPage)
		} else if userInfo.Role == AuthorizedUser {
			pages.SwitchToPage(AuthorizedPage)
		}
	})
	form.AddButton("Назад", func() {
		pages.SwitchToPage(GuestPage)
	})

	return form
}

func (tui *Tui) GetReservesPage(form *tview.Form, pages *tview.Pages) *tview.Form {
	errorTextView := tview.NewTextView().
		SetDynamicColors(true).
		SetRegions(true).
		SetWrap(true).
		SetTextAlign(tview.AlignCenter).
		SetText("")

	user, err := tui.app.UserSvc.GetByLogin(&dto.GetUserByLoginRequest{
		Login: tui.userInfo.Username,
	})
	if err != nil {
		errorTextView.SetText(err.Error())
		tui.ErrorForm(errorForm, pages, errorTextView, AuthorizedPage)
		pages.SwitchToPage(ErrorPage)
		return form
	}

	reserves, err := tui.app.UserSvc.GetReserves(&dto.GetUserReservesRequest{
		Id: user.Id,
	})
	if err != nil {
		errorTextView.SetText(err.Error())
		tui.ErrorForm(errorForm, pages, errorTextView, AuthorizedPage)
		pages.SwitchToPage(ErrorPage)
	}
	//reservesIntervals := make([]string, 0)
	//for _, reserve := range reserves {
	//	reservesIntervals = append(reservesIntervals, reserve.TimeInterval.StartTime.String()+" --- "+
	//		reserve.TimeInterval.EndTime.String())
	//}

	var str string
	for i := 0; i < len(reserves); i++ {

		room, err := tui.app.RoomSvc.Get(&dto.GetRoomRequest{Id: reserves[i].RoomId})
		if err != nil {
			errorTextView.SetText(err.Error())
			tui.ErrorForm(errorForm, pages, errorTextView, AuthorizedPage)
			pages.SwitchToPage(ErrorPage)
		}

		producer, err := tui.app.ProducerSvc.Get(&dto.GetProducerRequest{Id: reserves[i].ProducerId})
		if err != nil {
			errorTextView.SetText(err.Error())
			tui.ErrorForm(errorForm, pages, errorTextView, AuthorizedPage)
			pages.SwitchToPage(ErrorPage)
		}

		instrumentalist, err := tui.app.InstrumentalistSvc.Get(&dto.GetInstrumentalistRequest{Id: reserves[i].InstrumentalistId})
		if err != nil {
			errorTextView.SetText(err.Error())
			tui.ErrorForm(errorForm, pages, errorTextView, AuthorizedPage)
			pages.SwitchToPage(ErrorPage)
		}

		equipments, err := tui.app.EquipmentSvc.GetByReserve(&dto.GetEquipmentByReserveRequest{
			ReserveId: reserves[i].Id,
		})

		var roomName, producerName, instrumentalistName, microphoneName, guitarName string
		if room == nil {
			roomName = "Пусто"
		} else {
			roomName = room.Name
		}
		if producer == nil {
			producerName = "Пусто"
		} else {
			producerName = producer.Name
		}
		if instrumentalist == nil {
			instrumentalistName = "Пусто"
		} else {
			instrumentalistName = instrumentalist.Name
		}
		if equipments == nil {
			microphoneName = "Пусто"
			guitarName = "ПУсто"
		} else if len(equipments) == 1 && equipments[0].EquipmentType == 1 {
			microphoneName = equipments[0].Name
			guitarName = "ПустО"
		} else if len(equipments) == 1 && equipments[0].EquipmentType == 2 {
			microphoneName = "Пусто"
			guitarName = equipments[0].Name
		}

		//eqAndTime, err := postgresql.EquipmentRepository.GetNotFullTimeFreeByStudioAndType(nil, &dto.GetEquipmentNotFullTimeFreeByStudioAndTypeRequest{
		//	StudioId: room.StudioId,
		//	Type: 1,
		//	TimeInterval: reserves[i].TimeInterval,
		//})

		str += fmt.Sprintf("%d. Комната: %s\n   "+
			"Продюсер: %s\n   "+
			"Инструменталист: %s\n   "+
			"Микрофон: %s\n   "+
			"Гитара: %s\n   "+
			"Время: %s\n   ",
			i+1,
			roomName,
			producerName,
			instrumentalistName,
			microphoneName,
			guitarName,
			reserves[i].TimeInterval.StartTime.Format("2006-Jan-02 15:04")+" - "+reserves[i].TimeInterval.EndTime.Format("15:04"),
		)
		str += "\n\n"
	}
	if str == "" {
		str = "Нет броней"
	}
	form.AddTextView("Мои брони: ", str, 50, 20, true, true)
	//fmt.Println(1)
	form.AddButton("Назад", func() {
		pages.SwitchToPage(AuthorizedPage)
	})
	return form

}

func (tui *Tui) GetAllReservesPage(form *tview.Form, pages *tview.Pages) *tview.Form {
	errorTextView := tview.NewTextView().
		SetDynamicColors(true).
		SetRegions(true).
		SetWrap(true).
		SetTextAlign(tview.AlignCenter).
		SetText("")

	reserves, err := tui.app.ReserveSvc.GetAll(&dto.GetAllReserveRequest{})
	if err != nil {
		errorTextView.SetText(err.Error())
		tui.ErrorForm(errorForm, pages, errorTextView, GuestPage)
		pages.SwitchToPage(ErrorPage)
	}
	//reservesIntervals := make([]string, 0)
	//for _, reserve := range reserves {
	//	reservesIntervals = append(reservesIntervals, reserve.TimeInterval.StartTime.String()+" --- "+
	//		reserve.TimeInterval.EndTime.String())
	//}

	var str string
	for i := 0; i < len(reserves); i++ {

		room, err := tui.app.RoomSvc.Get(&dto.GetRoomRequest{Id: reserves[i].RoomId})
		if err != nil {
			errorTextView.SetText(err.Error())
			tui.ErrorForm(errorForm, pages, errorTextView, GuestPage)
			pages.SwitchToPage(ErrorPage)
		}

		producer, err := tui.app.ProducerSvc.Get(&dto.GetProducerRequest{Id: reserves[i].ProducerId})
		if err != nil {
			errorTextView.SetText(err.Error())
			tui.ErrorForm(errorForm, pages, errorTextView, GuestPage)
			pages.SwitchToPage(ErrorPage)
		}

		instrumentalist, err := tui.app.InstrumentalistSvc.Get(&dto.GetInstrumentalistRequest{Id: reserves[i].InstrumentalistId})
		if err != nil {
			errorTextView.SetText(err.Error())
			tui.ErrorForm(errorForm, pages, errorTextView, GuestPage)
			pages.SwitchToPage(ErrorPage)
		}

		equipments, err := tui.app.EquipmentSvc.GetByReserve(&dto.GetEquipmentByReserveRequest{
			ReserveId: reserves[i].Id,
		})

		var roomName, producerName, instrumentalistName, microphoneName, guitarName string
		if room == nil {
			roomName = "Пусто"
		} else {
			roomName = room.Name
		}
		if producer == nil {
			producerName = "Пусто"
		} else {
			producerName = producer.Name
		}
		if instrumentalist == nil {
			instrumentalistName = "Пусто"
		} else {
			instrumentalistName = instrumentalist.Name
		}
		if equipments == nil || len(equipments) == 0 {
			microphoneName = "Пусто"
			guitarName = "Пусто"
		} else if len(equipments) == 1 && equipments[0].EquipmentType == 1 {
			microphoneName = equipments[0].Name
			guitarName = "ПустО"
		} else if len(equipments) == 1 && equipments[0].EquipmentType == 2 {
			microphoneName = "Пусто"
			guitarName = equipments[0].Name
		}

		//eqAndTime, err := postgresql.EquipmentRepository.GetNotFullTimeFreeByStudioAndType(nil, &dto.GetEquipmentNotFullTimeFreeByStudioAndTypeRequest{
		//	StudioId: room.StudioId,
		//	Type: 1,
		//	TimeInterval: reserves[i].TimeInterval,
		//})

		str += fmt.Sprintf("%d. Комната: %s\n   "+
			"Продюсер: %s\n   "+
			"Инструменталист: %s\n   "+
			"Микрофон: %s\n   "+
			"Гитара: %s\n   "+
			"Время: %s\n   ",
			i+1,
			roomName,
			producerName,
			instrumentalistName,
			microphoneName,
			guitarName,
			reserves[i].TimeInterval.StartTime.Format("2006-Jan-02 15:04")+" - "+reserves[i].TimeInterval.EndTime.Format("15:04"),
		)
		str += "\n\n"
	}
	if str == "" {
		str = "Нет броней"
	}
	form.AddTextView("Все брони: ", str, 50, 20, true, true)
	//fmt.Println(1)
	form.AddButton("Назад", func() {
		pages.SwitchToPage(GuestPage)
	})
	return form

}

func (tui *Tui) CreateAuthorizedMenu(form *tview.Form, pages *tview.Pages, exitFunc *tview.Application) *tview.List {
	return tview.NewList().
		AddItem("Посмотреть свои брони", "", '1', func() {
			form.Clear(true)
			form.Clear(true)
			tui.GetReservesPage(form, pages)
			pages.SwitchToPage(GetReservesPage)
		}).
		AddItem("Создать бронь", "", '2', func() {
			form.Clear(true)
			tui.CreateValidatePage(form, pages)
			pages.SwitchToPage(CreateValidatePage)
		}).
		AddItem("Удалить бронь", "", '3', func() {
			form.Clear(true)
			tui.DeleteReservePage(form, pages)
			pages.SwitchToPage(DeleteReservePage)
		}).
		AddItem("Изменить свои данные", "", '4', func() {
			form.Clear(true)
			tui.UpdateUserPage(form, pages)
			pages.SwitchToPage(UpdateUserPage)
		}).
		AddItem("Выйти из профиля", "", '5', func() {
			form.Clear(true)
			tui.userInfo.Username = ""
			tui.userInfo.Role = UnauthorizedUser
			pages.SwitchToPage(GuestPage)
		}).
		AddItem("Выйти из приложения", "", '0', func() {
			exitFunc.Stop()
		})
}

func (tui *Tui) CreateAdminMenu(form *tview.Form, pages *tview.Pages, exitFunc *tview.Application) *tview.List {
	return tview.NewList().
		AddItem("[ АДМИН ] Добавить атрибуты", "", '1', func() {
			tui.CreateAddMenu(form, pages, exitFunc)
			pages.SwitchToPage(CreateAddMenu)
		}).
		AddItem("[ АДМИН ] Удалить атрибуты", "", '2', func() {
			tui.CreateDeleteMenu(form, pages, exitFunc)
			pages.SwitchToPage(CreateDeleteMenu)
		}).
		AddItem("[ АДМИН ] Изменить атрибуты", "", '3', func() {
			tui.CreateUpdateMenu(form, pages, exitFunc)
			pages.SwitchToPage(CreateUpdateMenu)
		}).
		AddItem("Изменить свои данные", "", '4', func() {
			form.Clear(true)
			tui.UpdateUserPage(form, pages)
			pages.SwitchToPage(UpdateUserPage)
		}).
		AddItem("Выйти из профиля", "", '5', func() {
			form.Clear(true)
			tui.userInfo.Username = ""
			tui.userInfo.Role = UnauthorizedUser
			pages.SwitchToPage(GuestPage)
		}).
		AddItem("Выйти из приложения", "", '0', func() {
			exitFunc.Stop()
		})
}

func (tui *Tui) CreateAddMenu(form *tview.Form, pages *tview.Pages, exitFunc *tview.Application) *tview.List {
	return tview.NewList().
		AddItem("Добавить студию", "", '1', func() {
			form.Clear(true)
			tui.AddStudioPage(form, pages)
			pages.SwitchToPage(AddStudioPage)
		}).
		AddItem("Добавить комнату", "", '2', func() {
			form.Clear(true)
			tui.AddRoomPage(form, pages)
			pages.SwitchToPage(AddRoomPage)
		}).
		AddItem("Добавить продюсера", "", '3', func() {
			form.Clear(true)
			tui.AddProducerPage(form, pages)
			pages.SwitchToPage(AddProducerPage)
		}).
		AddItem("Добавить инструменталиста", "", '4', func() {
			form.Clear(true)
			tui.AddInstrumentalistPage(form, pages)
			pages.SwitchToPage(AddInstrumentalistPage)
		}).
		AddItem("Добавить оборудование", "", '5', func() {
			form.Clear(true)
			tui.AddEquipmentPage(form, pages)
			pages.SwitchToPage(AddEquipmentPage)
		}).
		AddItem("Назад", "", '6', func() {
			pages.SwitchToPage(AdminPage)
		})
}

func (tui *Tui) AddRoomPage(form *tview.Form, pages *tview.Pages) *tview.Form {
	room := &model.Room{}
	errorTextView := tview.NewTextView().
		SetDynamicColors(true).
		SetRegions(true).
		SetWrap(true).
		SetTextAlign(tview.AlignCenter).
		SetText("")

	studios, err := tui.app.StudioSvc.GetAll(&dto.GetStudioAllRequest{})
	if err != nil {
		errorTextView.SetText(err.Error())
		tui.ErrorForm(errorForm, pages, errorTextView, AdminPage)
		pages.SwitchToPage(ErrorPage)
	}
	studiosName := make([]string, 0)

	for _, studio := range studios {
		studiosName = append(studiosName, studio.Name)
	}

	form.AddDropDown("Выберите студию: ", studiosName, 0, func(tmp string, tmpInt int) {
		if tmpInt < 0 {
			return
		}
		room.StudioId = studios[tmpInt].Id
	})

	form.AddInputField("Введите название", "", 20, nil, func(tmp string) {
		room.Name = tmp
	})
	form.AddInputField("Введите час начала работы комнаты (с 0 до 23)", "", 20, nil, func(tmpS string) {
		tmp, err := strconv.Atoi(tmpS)
		if err != nil {
			errorTextView.SetText(err.Error())
			tui.ErrorForm(errorForm, pages, errorTextView, AddRoomPage)
			pages.SwitchToPage(ErrorPage)
			return
		}
		room.StartHour = int64(tmp)

	})
	form.AddInputField("Введите час конца работы комнаты (с 0 до 23)", "", 20, nil, func(tmpS string) {
		tmp, err := strconv.Atoi(tmpS)
		if err != nil {
			errorTextView.SetText(err.Error())
			tui.ErrorForm(errorForm, pages, errorTextView, AddRoomPage)
			pages.SwitchToPage(ErrorPage)
			return
		}
		room.EndHour = int64(tmp)
	})

	form.AddButton("Добавить комнату", func() {
		err := tui.app.RoomSvc.Add(&dto.AddRoomRequest{
			Name:      room.Name,
			StudioId:  room.StudioId,
			StartHour: room.StartHour,
			EndHour:   room.EndHour,
		})

		if err != nil {
			errorTextView.SetText(err.Error())
			tui.ErrorForm(errorForm, pages, errorTextView, AddRoomPage)
			pages.SwitchToPage(ErrorPage)
			return
		} else {
			errorTextView.SetText("Комната добавлена")
			tui.ErrorForm(errorForm, pages, errorTextView, AdminPage)
			pages.SwitchToPage(ErrorPage)
		}
		pages.SwitchToPage(AdminPage)
	})
	form.AddButton("Назад", func() {
		pages.SwitchToPage(CreateAddMenu)
	})
	return form
}

func (tui *Tui) AddProducerPage(form *tview.Form, pages *tview.Pages) *tview.Form {
	producer := &model.Producer{}
	errorTextView := tview.NewTextView().
		SetDynamicColors(true).
		SetRegions(true).
		SetWrap(true).
		SetTextAlign(tview.AlignCenter).
		SetText("")

	studios, err := tui.app.StudioSvc.GetAll(&dto.GetStudioAllRequest{})
	if err != nil {
		errorTextView.SetText(err.Error())
		tui.ErrorForm(errorForm, pages, errorTextView, AdminPage)
		pages.SwitchToPage(ErrorPage)
	}
	studiosName := make([]string, 0)
	for _, studio := range studios {
		studiosName = append(studiosName, studio.Name)
	}
	form.AddDropDown("Выберите студию: ", studiosName, 0, func(tmp string, tmpInt int) {
		if tmpInt < 0 {
			return
		}
		producer.StudioId = studios[tmpInt].Id
	})

	form.AddInputField("Введите имя", "", 20, nil, func(tmp string) {
		producer.Name = tmp
	})
	form.AddInputField("Введите час начала работы продюсера (с 0 до 23)", "", 20, nil, func(tmpS string) {
		tmp, err := strconv.Atoi(tmpS)
		if err != nil {
			errorTextView.SetText(err.Error())
			tui.ErrorForm(errorForm, pages, errorTextView, AddProducerPage)
			pages.SwitchToPage(ErrorPage)
			return
		}
		producer.StartHour = int64(tmp)

	})
	form.AddInputField("Введите час конца работы продюсера (с 0 до 23)", "", 20, nil, func(tmpS string) {
		tmp, err := strconv.Atoi(tmpS)
		if err != nil {
			errorTextView.SetText(err.Error())
			tui.ErrorForm(errorForm, pages, errorTextView, AddProducerPage)
			pages.SwitchToPage(ErrorPage)
			return
		}
		producer.EndHour = int64(tmp)
	})

	form.AddButton("Добавить продюсюера", func() {

		err := tui.app.ProducerSvc.Add(&dto.AddProducerRequest{
			Name:      producer.Name,
			StudioId:  producer.StudioId,
			StartHour: producer.StartHour,
			EndHour:   producer.EndHour,
		})

		if err != nil {
			errorTextView.SetText(err.Error())
			tui.ErrorForm(errorForm, pages, errorTextView, AddProducerPage)
			pages.SwitchToPage(ErrorPage)
			return
		}
		pages.SwitchToPage(AdminPage)
	})
	form.AddButton("Назад", func() {
		pages.SwitchToPage(AdminPage)
	})
	return form
}

func (tui *Tui) AddInstrumentalistPage(form *tview.Form, pages *tview.Pages) *tview.Form {
	instrumentalist := &model.Instrumentalist{}
	errorTextView := tview.NewTextView().
		SetDynamicColors(true).
		SetRegions(true).
		SetWrap(true).
		SetTextAlign(tview.AlignCenter).
		SetText("")

	studios, err := tui.app.StudioSvc.GetAll(&dto.GetStudioAllRequest{})
	if err != nil {
		errorTextView.SetText(err.Error())
		tui.ErrorForm(errorForm, pages, errorTextView, AdminPage)
		pages.SwitchToPage(ErrorPage)
	}
	studiosName := make([]string, 0)
	for _, studio := range studios {
		studiosName = append(studiosName, studio.Name)
	}
	form.AddDropDown("Выберите студию: ", studiosName, 0, func(tmp string, tmpInt int) {
		if tmpInt < 0 {
			return
		}
		instrumentalist.StudioId = studios[tmpInt].Id
	})

	form.AddInputField("Введите имя", "", 20, nil, func(tmp string) {
		instrumentalist.Name = tmp
	})
	form.AddInputField("Введите час начала работы инструменталиста (с 0 до 23)", "", 20, nil, func(tmpS string) {
		tmp, err := strconv.Atoi(tmpS)
		if err != nil {
			errorTextView.SetText(err.Error())
			tui.ErrorForm(errorForm, pages, errorTextView, AddInstrumentalistPage)
			pages.SwitchToPage(ErrorPage)
			return
		}
		instrumentalist.StartHour = int64(tmp)

	})
	form.AddInputField("Введите час конца работы инструменталиста (с 0 до 23)", "", 20, nil, func(tmpS string) {
		tmp, err := strconv.Atoi(tmpS)
		if err != nil {
			errorTextView.SetText(err.Error())
			tui.ErrorForm(errorForm, pages, errorTextView, AddInstrumentalistPage)
			pages.SwitchToPage(ErrorPage)
			return
		}
		instrumentalist.EndHour = int64(tmp)
	})

	form.AddButton("Добавить инструменталиста", func() {

		err := tui.app.InstrumentalistSvc.Add(&dto.AddInstrumentalistRequest{
			Name:      instrumentalist.Name,
			StudioId:  instrumentalist.StudioId,
			StartHour: instrumentalist.StartHour,
			EndHour:   instrumentalist.EndHour,
		})

		if err != nil {
			errorTextView.SetText(err.Error())
			tui.ErrorForm(errorForm, pages, errorTextView, AddInstrumentalistPage)
			pages.SwitchToPage(ErrorPage)
			return
		}
		pages.SwitchToPage(AdminPage)
	})
	form.AddButton("Назад", func() {
		pages.SwitchToPage(AdminPage)
	})
	return form
}

func (tui *Tui) AddStudioPage(form *tview.Form, pages *tview.Pages) *tview.Form {
	studio := &model.Studio{}
	errorTextView := tview.NewTextView().
		SetDynamicColors(true).
		SetRegions(true).
		SetWrap(true).
		SetTextAlign(tview.AlignCenter).
		SetText("")

	form.AddInputField("Введите название", "", 20, nil, func(tmp string) {
		studio.Name = tmp
	})

	form.AddButton("Добавить студию", func() {

		err := tui.app.StudioSvc.Add(&dto.AddStudioRequest{
			Name: studio.Name,
		})

		if err != nil {
			errorTextView.SetText(err.Error())
			tui.ErrorForm(errorForm, pages, errorTextView, AddStudioPage)
			pages.SwitchToPage(ErrorPage)
			return
		}
		pages.SwitchToPage(AdminPage)
	})
	form.AddButton("Назад", func() {
		pages.SwitchToPage(AdminPage)
	})
	return form
}

func (tui *Tui) AddEquipmentPage(form *tview.Form, pages *tview.Pages) *tview.Form {
	equipment := &model.Equipment{}
	errorTextView := tview.NewTextView().
		SetDynamicColors(true).
		SetRegions(true).
		SetWrap(true).
		SetTextAlign(tview.AlignCenter).
		SetText("")

	studios, err := tui.app.StudioSvc.GetAll(&dto.GetStudioAllRequest{})
	if err != nil {
		errorTextView.SetText(err.Error())
		tui.ErrorForm(errorForm, pages, errorTextView, AdminPage)
		pages.SwitchToPage(ErrorPage)
	}
	studiosName := make([]string, 0)
	for _, studio := range studios {
		studiosName = append(studiosName, studio.Name)
	}
	form.AddDropDown("Выберите студию: ", studiosName, 0, func(tmp string, tmpInt int) {
		if tmpInt < 0 {
			return
		}
		equipment.StudioId = studios[tmpInt].Id
	})

	form.AddInputField("Введите название: ", "", 20, nil, func(tmp string) {
		equipment.Name = tmp
	})

	form.AddDropDown("Выберите тип: ", []string{"Микрофон", "Гитара"}, 0, func(tmp string, tmpInt int) {
		if tmpInt < 0 {
			return
		}
		equipment.EquipmentType = int64(tmpInt + 1)
	})

	form.AddButton("Добавить оборудование", func() {

		err := tui.app.EquipmentSvc.Add(&dto.AddEquipmentRequest{
			Name:     equipment.Name,
			StudioId: equipment.StudioId,
			Type:     equipment.EquipmentType,
		})

		if err != nil {
			errorTextView.SetText(err.Error())
			tui.ErrorForm(errorForm, pages, errorTextView, AddEquipmentPage)
			pages.SwitchToPage(ErrorPage)
			return
		}
		pages.SwitchToPage(AdminPage)
	})
	form.AddButton("Назад", func() {
		pages.SwitchToPage(AdminPage)
	})
	return form
}

func (tui *Tui) CreateDeleteMenu(form *tview.Form, pages *tview.Pages, exitFunc *tview.Application) *tview.List {
	return tview.NewList().
		AddItem("Удалить комнату", "", '1', func() {
			form.Clear(true)
			tui.DeleteRoomPage(form, pages)
			pages.SwitchToPage(DeleteRoomPage)
		}).
		AddItem("Удалить продюсера", "", '2', func() {
			form.Clear(true)
			tui.DeleteProducerPage(form, pages)
			pages.SwitchToPage(DeleteProducerPage)
		}).
		AddItem("Удалить инструменталиста", "", '3', func() {
			form.Clear(true)
			tui.DeleteInstrumentalistPage(form, pages)
			pages.SwitchToPage(DeleteInstrumentalistPage)
		}).
		AddItem("Удалить студию", "", '4', func() {
			form.Clear(true)
			tui.DeleteStudioPage(form, pages)
			pages.SwitchToPage(DeleteStudioPage)
		}).
		AddItem("Удалить оборудование", "", '5', func() {
			form.Clear(true)
			tui.DeleteEquipmentPage(form, pages)
			pages.SwitchToPage(DeleteEquipmentPage)
		}).
		AddItem("Назад", "", '6', func() {
			pages.SwitchToPage(AdminPage)
		})
}

func (tui *Tui) DeleteRoomPage(form *tview.Form, pages *tview.Pages) *tview.Form {
	//room := &model.Room{}
	errorTextView := tview.NewTextView().
		SetDynamicColors(true).
		SetRegions(true).
		SetWrap(true).
		SetTextAlign(tview.AlignCenter).
		SetText("")

	studios, err := tui.app.StudioSvc.GetAll(&dto.GetStudioAllRequest{})
	if err != nil {
		errorTextView.SetText(err.Error())
		tui.ErrorForm(errorForm, pages, errorTextView, CreateDeleteMenu)
		pages.SwitchToPage(ErrorPage)
	}
	studiosName := make([]string, 0)
	for _, studio := range studios {
		studiosName = append(studiosName, studio.Name)
	}
	var studioId int64
	form.AddDropDown("Выберите студию: ", studiosName, 0, func(tmp string, tmpInt int) {
		if tmpInt < 0 {
			return
		}
		studioId = studios[tmpInt].Id
	})

	form.AddButton("Список комнат", func() {
		rooms, err := tui.app.RoomSvc.GetByStudio(&dto.GetRoomByStudioRequest{
			StudioId: studioId,
		})
		if err != nil {
			errorTextView.SetText(err.Error())
			tui.ErrorForm(errorForm, pages, errorTextView, CreateDeleteMenu)
			pages.SwitchToPage(ErrorPage)
			return
		}
		roomsName := make([]string, 0)
		for _, room := range rooms {
			roomsName = append(roomsName, room.Name)
		}
		var roomId int64
		form.AddDropDown("Выберите комнату: ", roomsName, 0, func(tmp string, tmpInt int) {
			if tmpInt < 0 {
				return
			}
			roomId = rooms[tmpInt].Id
		})
		form.AddButton("Удалить комнату", func() {
			err := tui.app.RoomSvc.Delete(&dto.DeleteRoomRequest{
				Id: roomId,
			})
			if err != nil {
				errorTextView.SetText(err.Error())
				tui.ErrorForm(errorForm, pages, errorTextView, DeleteRoomPage)
				pages.SwitchToPage(ErrorPage)
				return
			}
			pages.SwitchToPage(AdminPage)
		})
		form.AddButton("Назад", func() {
			pages.SwitchToPage(CreateDeleteMenu)
		})

	})

	return form
}

func (tui *Tui) DeleteProducerPage(form *tview.Form, pages *tview.Pages) *tview.Form {
	//room := &model.Room{}
	errorTextView := tview.NewTextView().
		SetDynamicColors(true).
		SetRegions(true).
		SetWrap(true).
		SetTextAlign(tview.AlignCenter).
		SetText("")

	studios, err := tui.app.StudioSvc.GetAll(&dto.GetStudioAllRequest{})
	if err != nil {
		errorTextView.SetText(err.Error())
		tui.ErrorForm(errorForm, pages, errorTextView, CreateDeleteMenu)
		pages.SwitchToPage(ErrorPage)
	}
	studiosName := make([]string, 0)
	for _, studio := range studios {
		studiosName = append(studiosName, studio.Name)
	}
	var studioId int64
	form.AddDropDown("Выберите студию: ", studiosName, 0, func(tmp string, tmpInt int) {
		if tmpInt < 0 {
			return
		}
		studioId = studios[tmpInt].Id
	})

	form.AddButton("Список продюсеров", func() {
		producers, err := tui.app.ProducerSvc.GetByStudio(&dto.GetProducerByStudioRequest{
			StudioId: studioId,
		})
		if err != nil {
			errorTextView.SetText(err.Error())
			tui.ErrorForm(errorForm, pages, errorTextView, CreateDeleteMenu)
			pages.SwitchToPage(ErrorPage)
			return
		}
		producersName := make([]string, 0)
		for _, producer := range producers {
			producersName = append(producersName, producer.Name)
		}
		var producerId int64
		form.AddDropDown("Выберите продюсера: ", producersName, 0, func(tmp string, tmpInt int) {
			if tmpInt < 0 {
				return
			}
			producerId = producers[tmpInt].Id
		})
		form.AddButton("Удалить продюсера", func() {
			err := tui.app.ProducerSvc.Delete(&dto.DeleteProducerRequest{
				Id: producerId,
			})
			if err != nil {
				errorTextView.SetText(err.Error())
				tui.ErrorForm(errorForm, pages, errorTextView, DeleteProducerPage)
				pages.SwitchToPage(ErrorPage)
				return
			}
			pages.SwitchToPage(AdminPage)
		})
		form.AddButton("Назад", func() {
			pages.SwitchToPage(CreateDeleteMenu)
		})

	})

	return form
}

func (tui *Tui) DeleteInstrumentalistPage(form *tview.Form, pages *tview.Pages) *tview.Form {
	//room := &model.Room{}
	errorTextView := tview.NewTextView().
		SetDynamicColors(true).
		SetRegions(true).
		SetWrap(true).
		SetTextAlign(tview.AlignCenter).
		SetText("")

	studios, err := tui.app.StudioSvc.GetAll(&dto.GetStudioAllRequest{})
	if err != nil {
		errorTextView.SetText(err.Error())
		tui.ErrorForm(errorForm, pages, errorTextView, CreateDeleteMenu)
		pages.SwitchToPage(ErrorPage)
	}
	studiosName := make([]string, 0)
	for _, studio := range studios {
		studiosName = append(studiosName, studio.Name)
	}
	var studioId int64
	form.AddDropDown("Выберите студию: ", studiosName, 0, func(tmp string, tmpInt int) {
		if tmpInt < 0 {
			return
		}
		studioId = studios[tmpInt].Id
	})

	form.AddButton("Список комнат", func() {
		instrumentalists, err := tui.app.InstrumentalistSvc.GetByStudio(&dto.GetInstrumentalistByStudioRequest{
			StudioId: studioId,
		})
		if err != nil {
			errorTextView.SetText(err.Error())
			tui.ErrorForm(errorForm, pages, errorTextView, CreateDeleteMenu)
			pages.SwitchToPage(ErrorPage)
			return
		}
		instrumentalistsName := make([]string, 0)
		for _, instrumentalist := range instrumentalists {
			instrumentalistsName = append(instrumentalistsName, instrumentalist.Name)
		}
		var instrumentalistId int64
		form.AddDropDown("Выберите инструменталиста: ", instrumentalistsName, 0, func(tmp string, tmpInt int) {
			if tmpInt < 0 {
				return
			}
			instrumentalistId = instrumentalists[tmpInt].Id
		})
		form.AddButton("Удалить инструменталиста", func() {
			err := tui.app.InstrumentalistSvc.Delete(&dto.DeleteInstrumentalistRequest{
				Id: instrumentalistId,
			})
			if err != nil {
				errorTextView.SetText(err.Error())
				tui.ErrorForm(errorForm, pages, errorTextView, DeleteInstrumentalistPage)
				pages.SwitchToPage(ErrorPage)
				return
			}
			pages.SwitchToPage(AdminPage)
		})
		form.AddButton("Назад", func() {
			pages.SwitchToPage(CreateDeleteMenu)
		})

	})

	return form
}

func (tui *Tui) DeleteStudioPage(form *tview.Form, pages *tview.Pages) *tview.Form {
	errorTextView := tview.NewTextView().
		SetDynamicColors(true).
		SetRegions(true).
		SetWrap(true).
		SetTextAlign(tview.AlignCenter).
		SetText("")

	studios, err := tui.app.StudioSvc.GetAll(&dto.GetStudioAllRequest{})
	if err != nil {
		errorTextView.SetText(err.Error())
		tui.ErrorForm(errorForm, pages, errorTextView, CreateDeleteMenu)
		pages.SwitchToPage(ErrorPage)
	}
	studiosName := make([]string, 0)
	for _, studio := range studios {
		studiosName = append(studiosName, studio.Name)
	}
	var studioId int64
	form.AddDropDown("Выберите студию: ", studiosName, 0, func(tmp string, tmpInt int) {
		if tmpInt < 0 {
			return
		}
		studioId = studios[tmpInt].Id
	})

	form.AddButton("Удалить инструменталиста", func() {
		err := tui.app.InstrumentalistSvc.Delete(&dto.DeleteInstrumentalistRequest{
			Id: studioId,
		})
		if err != nil {
			errorTextView.SetText(err.Error())
			tui.ErrorForm(errorForm, pages, errorTextView, DeleteStudioPage)
			pages.SwitchToPage(ErrorPage)
			return
		}
		pages.SwitchToPage(AdminPage)
	})
	form.AddButton("Назад", func() {
		pages.SwitchToPage(CreateDeleteMenu)
	})

	return form
}

func (tui *Tui) DeleteEquipmentPage(form *tview.Form, pages *tview.Pages) *tview.Form {
	//room := &model.Room{}
	errorTextView := tview.NewTextView().
		SetDynamicColors(true).
		SetRegions(true).
		SetWrap(true).
		SetTextAlign(tview.AlignCenter).
		SetText("")

	studios, err := tui.app.StudioSvc.GetAll(&dto.GetStudioAllRequest{})
	if err != nil {
		errorTextView.SetText(err.Error())
		tui.ErrorForm(errorForm, pages, errorTextView, CreateDeleteMenu)
		pages.SwitchToPage(ErrorPage)
	}
	studiosName := make([]string, 0)
	for _, studio := range studios {
		studiosName = append(studiosName, studio.Name)
	}
	var studioId int64
	form.AddDropDown("Выберите студию: ", studiosName, 0, func(tmp string, tmpInt int) {
		if tmpInt < 0 {
			return
		}
		studioId = studios[tmpInt].Id
	})

	form.AddButton("Список оборудования", func() {
		equipments, err := tui.app.EquipmentSvc.GetByStudio(&dto.GetEquipmentByStudioRequest{
			StudioId: studioId,
		})
		if err != nil {
			errorTextView.SetText(err.Error())
			tui.ErrorForm(errorForm, pages, errorTextView, CreateDeleteMenu)
			pages.SwitchToPage(ErrorPage)
			return
		}
		equipmentsName := make([]string, 0)
		for _, equipment := range equipments {
			equipmentsName = append(equipmentsName, equipment.Name)
		}
		var equipmentId int64
		form.AddDropDown("Выберите оборудование: ", equipmentsName, 0, func(tmp string, tmpInt int) {
			if tmpInt < 0 {
				return
			}
			equipmentId = equipments[tmpInt].Id
		})
		form.AddButton("Удалить оборудование", func() {
			err := tui.app.EquipmentSvc.Delete(&dto.DeleteEquipmentRequest{
				Id: equipmentId,
			})
			if err != nil {
				errorTextView.SetText(err.Error())
				tui.ErrorForm(errorForm, pages, errorTextView, DeleteEquipmentPage)
				pages.SwitchToPage(ErrorPage)
				return
			}
			pages.SwitchToPage(AdminPage)
		})
		form.AddButton("Назад", func() {
			pages.SwitchToPage(CreateDeleteMenu)
		})

	})

	return form
}

func (tui *Tui) CreateUpdateMenu(form *tview.Form, pages *tview.Pages, exitFunc *tview.Application) *tview.List {
	return tview.NewList().
		AddItem("Изменить комнату", "", '1', func() {
			form.Clear(true)
			tui.UpdateRoomPage(form, pages)
			pages.SwitchToPage(UpdateRoomPage)
		}).
		AddItem("Изменить продюсера", "", '2', func() {
			form.Clear(true)
			tui.UpdateProducerPage(form, pages)
			pages.SwitchToPage(UpdateProducerPage)
		}).
		AddItem("Изменить инструменталиста", "", '3', func() {
			form.Clear(true)
			tui.UpdateInstrumentalistPage(form, pages)
			pages.SwitchToPage(UpdateInstrumentalistPage)
		}).
		AddItem("Изменить студию", "", '4', func() {
			form.Clear(true)
			tui.UpdateStudioPage(form, pages)
			pages.SwitchToPage(UpdateStudioPage)
		}).
		AddItem("Изменить оборудование", "", '5', func() {
			form.Clear(true)
			tui.UpdateEquipmentPage(form, pages)
			pages.SwitchToPage(UpdateEquipmentPage)
		}).
		AddItem("Назад", "", '6', func() {
			pages.SwitchToPage(AdminPage)
		})
}

func (tui *Tui) UpdateRoomPage(form *tview.Form, pages *tview.Pages) *tview.Form {
	//room := &model.Room{}
	errorTextView := tview.NewTextView().
		SetDynamicColors(true).
		SetRegions(true).
		SetWrap(true).
		SetTextAlign(tview.AlignCenter).
		SetText("")

	studios, err := tui.app.StudioSvc.GetAll(&dto.GetStudioAllRequest{})
	if err != nil {
		errorTextView.SetText(err.Error())
		tui.ErrorForm(errorForm, pages, errorTextView, CreateUpdateMenu)
		pages.SwitchToPage(ErrorPage)
	}
	studiosName := make([]string, 0)
	for _, studio := range studios {
		studiosName = append(studiosName, studio.Name)
	}
	var studioId int64
	form.AddDropDown("Выберите студию: ", studiosName, 0, func(tmp string, tmpInt int) {
		if tmpInt < 0 {
			return
		}
		studioId = studios[tmpInt].Id
	})

	form.AddButton("Список комнат", func() {
		rooms, err := tui.app.RoomSvc.GetByStudio(&dto.GetRoomByStudioRequest{
			StudioId: studioId,
		})
		if err != nil {
			errorTextView.SetText(err.Error())
			tui.ErrorForm(errorForm, pages, errorTextView, CreateUpdateMenu)
			pages.SwitchToPage(ErrorPage)
			return
		}
		roomsName := make([]string, 0)
		for _, room := range rooms {
			roomsName = append(roomsName, room.Name)
		}
		var roomId int64
		form.AddDropDown("Выберите комнату: ", roomsName, 0, func(tmp string, tmpInt int) {
			if tmpInt < 0 {
				return
			}
			roomId = rooms[tmpInt].Id
		})

		var roomNewName string
		tmpStr := fmt.Sprintf("Введите новое название:")
		form.AddInputField(tmpStr, "", 20, nil, func(tmp string) {
			roomNewName = tmp
		})

		var studioNewId int64
		form.AddDropDown("Выберите новую студию: ", studiosName, 0, func(tmp string, tmpInt int) {
			if tmpInt < 0 {
				return
			}
			studioNewId = studios[tmpInt].Id
		})

		var roomStartHour int64
		form.AddInputField("Введите новый час начала работы комнаты (с 0 до 23)", "", 20, nil, func(tmpS string) {
			tmp, err := strconv.Atoi(tmpS)
			if err != nil {
				errorTextView.SetText(err.Error())
				tui.ErrorForm(errorForm, pages, errorTextView, UpdateRoomPage)
				pages.SwitchToPage(ErrorPage)
				return
			}
			roomStartHour = int64(tmp)

		})
		var roomEndHour int64
		form.AddInputField("Введите новый час конца работы комнаты (с 0 до 23)", "", 20, nil, func(tmpS string) {
			tmp, err := strconv.Atoi(tmpS)
			if err != nil {
				errorTextView.SetText(err.Error())
				tui.ErrorForm(errorForm, pages, errorTextView, UpdateRoomPage)
				pages.SwitchToPage(ErrorPage)
				return
			}
			roomEndHour = int64(tmp)
		})

		form.AddButton("Изменить комнату", func() {
			err := tui.app.RoomSvc.Update(&dto.UpdateRoomRequest{
				Id:        roomId,
				Name:      roomNewName,
				StudioId:  studioNewId,
				StartHour: roomStartHour,
				EndHour:   roomEndHour,
			})
			if err != nil {
				errorTextView.SetText(err.Error())
				tui.ErrorForm(errorForm, pages, errorTextView, UpdateRoomPage)
				pages.SwitchToPage(ErrorPage)
			}

			pages.SwitchToPage(AdminPage)
		})
	})

	form.AddButton("Назад", func() {
		pages.SwitchToPage(CreateUpdateMenu)
	})

	return form
}

func (tui *Tui) UpdateProducerPage(form *tview.Form, pages *tview.Pages) *tview.Form {
	//room := &model.Room{}
	errorTextView := tview.NewTextView().
		SetDynamicColors(true).
		SetRegions(true).
		SetWrap(true).
		SetTextAlign(tview.AlignCenter).
		SetText("")

	studios, err := tui.app.StudioSvc.GetAll(&dto.GetStudioAllRequest{})
	if err != nil {
		errorTextView.SetText(err.Error())
		tui.ErrorForm(errorForm, pages, errorTextView, CreateUpdateMenu)
		pages.SwitchToPage(ErrorPage)
	}
	studiosName := make([]string, 0)
	for _, studio := range studios {
		studiosName = append(studiosName, studio.Name)
	}
	var studioId int64
	form.AddDropDown("Выберите студию: ", studiosName, 0, func(tmp string, tmpInt int) {
		if tmpInt < 0 {
			return
		}
		studioId = studios[tmpInt].Id
	})

	form.AddButton("Список продюсеров", func() {
		producers, err := tui.app.ProducerSvc.GetByStudio(&dto.GetProducerByStudioRequest{
			StudioId: studioId,
		})
		if err != nil {
			errorTextView.SetText(err.Error())
			tui.ErrorForm(errorForm, pages, errorTextView, CreateUpdateMenu)
			pages.SwitchToPage(ErrorPage)
			return
		}
		producersName := make([]string, 0)
		for _, producer := range producers {
			producersName = append(producersName, producer.Name)
		}
		var producerId int64
		form.AddDropDown("Выберите продюсера: ", producersName, 0, func(tmp string, tmpInt int) {
			if tmpInt < 0 {
				return
			}
			producerId = producers[tmpInt].Id
		})

		var producerNewName string
		tmpStr := fmt.Sprintf("Введите новое имя:")
		form.AddInputField(tmpStr, "", 20, nil, func(tmp string) {
			producerNewName = tmp
		})

		var studioNewId int64
		form.AddDropDown("Выберите новую студию: ", studiosName, 0, func(tmp string, tmpInt int) {
			if tmpInt < 0 {
				return
			}
			studioNewId = studios[tmpInt].Id
		})

		var producerStartHour int64
		form.AddInputField("Введите новый час начала работы продюсера (с 0 до 23)", "", 20, nil, func(tmpS string) {
			tmp, err := strconv.Atoi(tmpS)
			if err != nil {
				errorTextView.SetText(err.Error())
				tui.ErrorForm(errorForm, pages, errorTextView, UpdateProducerPage)
				pages.SwitchToPage(ErrorPage)
				return
			}
			producerStartHour = int64(tmp)

		})
		var producerEndHour int64
		form.AddInputField("Введите новый час конца работы продюсера (с 0 до 23)", "", 20, nil, func(tmpS string) {
			tmp, err := strconv.Atoi(tmpS)
			if err != nil {
				errorTextView.SetText(err.Error())
				tui.ErrorForm(errorForm, pages, errorTextView, UpdateProducerPage)
				pages.SwitchToPage(ErrorPage)
				return
			}
			producerEndHour = int64(tmp)
		})

		form.AddButton("Изменить продюсера", func() {
			err := tui.app.ProducerSvc.Update(&dto.UpdateProducerRequest{
				Id:        producerId,
				Name:      producerNewName,
				StudioId:  studioNewId,
				StartHour: producerStartHour,
				EndHour:   producerEndHour,
			})
			if err != nil {
				errorTextView.SetText(err.Error())
				tui.ErrorForm(errorForm, pages, errorTextView, UpdateProducerPage)
				pages.SwitchToPage(ErrorPage)
			}

			pages.SwitchToPage(AdminPage)
		})
	})

	form.AddButton("Назад", func() {
		pages.SwitchToPage(CreateUpdateMenu)
	})

	return form
}

func (tui *Tui) UpdateInstrumentalistPage(form *tview.Form, pages *tview.Pages) *tview.Form {
	//room := &model.Room{}
	errorTextView := tview.NewTextView().
		SetDynamicColors(true).
		SetRegions(true).
		SetWrap(true).
		SetTextAlign(tview.AlignCenter).
		SetText("")

	studios, err := tui.app.StudioSvc.GetAll(&dto.GetStudioAllRequest{})
	if err != nil {
		errorTextView.SetText(err.Error())
		tui.ErrorForm(errorForm, pages, errorTextView, CreateUpdateMenu)
		pages.SwitchToPage(ErrorPage)
	}
	studiosName := make([]string, 0)
	for _, studio := range studios {
		studiosName = append(studiosName, studio.Name)
	}
	var studioId int64
	form.AddDropDown("Выберите студию: ", studiosName, 0, func(tmp string, tmpInt int) {
		if tmpInt < 0 {
			return
		}
		studioId = studios[tmpInt].Id
	})

	form.AddButton("Список инструменталистов", func() {
		instrumentalists, err := tui.app.InstrumentalistSvc.GetByStudio(&dto.GetInstrumentalistByStudioRequest{
			StudioId: studioId,
		})
		if err != nil {
			errorTextView.SetText(err.Error())
			tui.ErrorForm(errorForm, pages, errorTextView, CreateUpdateMenu)
			pages.SwitchToPage(ErrorPage)
			return
		}
		instrumentalistsName := make([]string, 0)
		for _, instrumentalist := range instrumentalists {
			instrumentalistsName = append(instrumentalistsName, instrumentalist.Name)
		}
		var instrumentalistId int64
		form.AddDropDown("Выберите инструменталиста: ", instrumentalistsName, 0, func(tmp string, tmpInt int) {
			if tmpInt < 0 {
				return
			}
			instrumentalistId = instrumentalists[tmpInt].Id
		})

		var instrumentalistNewName string
		tmpStr := fmt.Sprintf("Введите новое имя:")
		form.AddInputField(tmpStr, "", 20, nil, func(tmp string) {
			instrumentalistNewName = tmp
		})

		var studioNewId int64
		form.AddDropDown("Выберите новую студию: ", studiosName, 0, func(tmp string, tmpInt int) {
			if tmpInt < 0 {
				return
			}
			studioNewId = studios[tmpInt].Id
		})

		var instrumentalistStartHour int64
		form.AddInputField("Введите новый час начала работы инструменталиста (с 0 до 23)", "", 20, nil, func(tmpS string) {
			tmp, err := strconv.Atoi(tmpS)
			if err != nil {
				errorTextView.SetText(err.Error())
				tui.ErrorForm(errorForm, pages, errorTextView, UpdateInstrumentalistPage)
				pages.SwitchToPage(ErrorPage)
				return
			}
			instrumentalistStartHour = int64(tmp)

		})
		var instrumentalistEndHour int64
		form.AddInputField("Введите новый час конца работы инструменталиста (с 0 до 23)", "", 20, nil, func(tmpS string) {
			tmp, err := strconv.Atoi(tmpS)
			if err != nil {
				errorTextView.SetText(err.Error())
				tui.ErrorForm(errorForm, pages, errorTextView, UpdateInstrumentalistPage)
				pages.SwitchToPage(ErrorPage)
				return
			}
			instrumentalistEndHour = int64(tmp)
		})
		form.AddButton("Изменить инструменталиста", func() {
			err := tui.app.InstrumentalistSvc.Update(&dto.UpdateInstrumentalistRequest{
				Id:        instrumentalistId,
				Name:      instrumentalistNewName,
				StudioId:  studioNewId,
				StartHour: instrumentalistStartHour,
				EndHour:   instrumentalistEndHour,
			})
			if err != nil {
				errorTextView.SetText(err.Error())
				tui.ErrorForm(errorForm, pages, errorTextView, UpdateInstrumentalistPage)
				pages.SwitchToPage(ErrorPage)
			}
			//form.AddTextView("debug", err.Error(), 10, 10, true, true)
			pages.SwitchToPage(AdminPage)
		})
	})

	form.AddButton("Назад", func() {
		pages.SwitchToPage(CreateUpdateMenu)
	})

	return form
}

func (tui *Tui) UpdateStudioPage(form *tview.Form, pages *tview.Pages) *tview.Form {
	errorTextView := tview.NewTextView().
		SetDynamicColors(true).
		SetRegions(true).
		SetWrap(true).
		SetTextAlign(tview.AlignCenter).
		SetText("")

	studios, err := tui.app.StudioSvc.GetAll(&dto.GetStudioAllRequest{})
	if err != nil {
		errorTextView.SetText(err.Error())
		tui.ErrorForm(errorForm, pages, errorTextView, CreateUpdateMenu)
		pages.SwitchToPage(ErrorPage)
	}
	studiosName := make([]string, 0)
	for _, studio := range studios {
		studiosName = append(studiosName, studio.Name)
	}
	var studioId int64
	form.AddDropDown("Выберите студию: ", studiosName, 0, func(tmp string, tmpInt int) {
		if tmpInt < 0 {
			return
		}
		studioId = studios[tmpInt].Id
	})

	var studioNewName string
	tmpStr := fmt.Sprintf("Введите новое название:")
	form.AddInputField(tmpStr, "", 20, nil, func(tmp string) {
		studioNewName = tmp
	})

	form.AddButton("Изменить студию", func() {
		err := tui.app.StudioSvc.Update(&dto.UpdateStudioRequest{
			Id:   studioId,
			Name: studioNewName,
		})
		if err != nil {
			errorTextView.SetText(err.Error())
			tui.ErrorForm(errorForm, pages, errorTextView, UpdateStudioPage)
			pages.SwitchToPage(ErrorPage)
		}
		//form.AddTextView("debug", err.Error(), 10, 10, true, true)
		pages.SwitchToPage(AdminPage)
	})

	form.AddButton("Назад", func() {
		pages.SwitchToPage(CreateUpdateMenu)
	})

	return form
}

func (tui *Tui) UpdateEquipmentPage(form *tview.Form, pages *tview.Pages) *tview.Form {
	//room := &model.Room{}
	errorTextView := tview.NewTextView().
		SetDynamicColors(true).
		SetRegions(true).
		SetWrap(true).
		SetTextAlign(tview.AlignCenter).
		SetText("")

	studios, err := tui.app.StudioSvc.GetAll(&dto.GetStudioAllRequest{})
	if err != nil {
		errorTextView.SetText(err.Error())
		tui.ErrorForm(errorForm, pages, errorTextView, CreateUpdateMenu)
		pages.SwitchToPage(ErrorPage)
	}
	studiosName := make([]string, 0)
	for _, studio := range studios {
		studiosName = append(studiosName, studio.Name)
	}
	var studioId int64
	form.AddDropDown("Выберите студию: ", studiosName, 0, func(tmp string, tmpInt int) {
		if tmpInt < 0 {
			return
		}
		studioId = studios[tmpInt].Id
	})

	form.AddButton("Список оборудования", func() {
		equipments, err := tui.app.EquipmentSvc.GetByStudio(&dto.GetEquipmentByStudioRequest{
			StudioId: studioId,
		})
		if err != nil {
			errorTextView.SetText(err.Error())
			tui.ErrorForm(errorForm, pages, errorTextView, CreateUpdateMenu)
			pages.SwitchToPage(ErrorPage)
			return
		}
		equipmentsName := make([]string, 0)
		for _, equipment := range equipments {
			equipmentsName = append(equipmentsName, equipment.Name)
		}
		var equipmentId int64
		form.AddDropDown("Выберите продюсера: ", equipmentsName, 0, func(tmp string, tmpInt int) {
			if tmpInt < 0 {
				return
			}
			equipmentId = equipments[tmpInt].Id
		})

		var equipmentNewName string
		tmpStr := fmt.Sprintf("Введите новое название:")
		form.AddInputField(tmpStr, "", 20, nil, func(tmp string) {
			equipmentNewName = tmp
		})

		var studioNewId int64
		form.AddDropDown("Выберите новую студию: ", studiosName, 0, func(tmp string, tmpInt int) {
			if tmpInt < 0 {
				return
			}
			studioNewId = studios[tmpInt].Id
		})

		var equipmentType int64
		form.AddDropDown("Выберите тип: ", []string{"Микрофон", "Гитара"}, 0, func(tmp string, tmpInt int) {
			if tmpInt < 0 {
				return
			}
			equipmentType = int64(tmpInt + 1)
		})

		form.AddButton("Изменить оборудование", func() {
			err := tui.app.EquipmentSvc.Update(&dto.UpdateEquipmentRequest{
				Id:       equipmentId,
				Name:     equipmentNewName,
				StudioId: studioNewId,
				Type:     equipmentType,
			})
			if err != nil {
				errorTextView.SetText(err.Error())
				tui.ErrorForm(errorForm, pages, errorTextView, UpdateEquipmentPage)
				pages.SwitchToPage(ErrorPage)
			}

			pages.SwitchToPage(AdminPage)
		})
	})

	form.AddButton("Назад", func() {
		pages.SwitchToPage(CreateUpdateMenu)
	})

	return form
}

func (tui *Tui) CreateReservePage(form *tview.Form,
	pages *tview.Pages,
	notReservedRooms []*model.Room,
	notReservedEquipments [][]*model.Equipment,
	notReservedProducers []*model.Producer,
	notReservedInstrumentalists []*model.Instrumentalist,
	userInterval *model.TimeInterval,
) *tview.Form {
	errorTextView := tview.NewTextView().
		SetDynamicColors(true).
		SetRegions(true).
		SetWrap(true).
		SetTextAlign(tview.AlignCenter).
		SetText("")

	notReservedRoomsName := make([]string, 0)
	for _, notReservedRoom := range notReservedRooms {
		notReservedRoomsName = append(notReservedRoomsName, notReservedRoom.Name)
	}
	var notReservedRoomId int64
	form.AddDropDown("Выберите комнату: ", notReservedRoomsName, 0, func(tmp string, tmpInt int) {
		if tmpInt < 0 {
			return
		}
		notReservedRoomId = notReservedRooms[tmpInt].Id
	})

	notReservedProducersName := make([]string, 1)
	for _, notReservedProducer := range notReservedProducers {
		notReservedProducersName = append(notReservedProducersName, notReservedProducer.Name)
	}
	var notReservedProducerId int64
	form.AddDropDown("Выберите продюсера: ", notReservedProducersName, 0, func(tmp string, tmpInt int) {
		if tmpInt <= 0 {
			return
		}
		notReservedProducerId = notReservedProducers[tmpInt-1].Id
	})

	notReservedInstrumentalistsName := make([]string, 1)
	for _, notReservedInstrumentalist := range notReservedInstrumentalists {
		notReservedInstrumentalistsName = append(notReservedInstrumentalistsName, notReservedInstrumentalist.Name)
	}
	var notReservedInstrumentalistId int64

	form.AddDropDown("Выберите инструменталиста: ", notReservedInstrumentalistsName, 0, func(tmp string, tmpInt int) {
		if tmpInt <= 0 {
			return
		}
		notReservedInstrumentalistId = notReservedInstrumentalists[tmpInt-1].Id
	})

	var notReservedEquipment []int64
	notReservedMicrophones := make([]string, 1)
	for _, notReservedMicrophone := range notReservedEquipments[0] {
		notReservedMicrophones = append(notReservedMicrophones, notReservedMicrophone.Name)
	}
	var micTmpint int
	form.AddDropDown("Выберите микрофон: ", notReservedMicrophones, 0, func(tmp string, tmpInt int) {
		if tmpInt < 0 || len(notReservedEquipments[0]) == 0 {
			return
		}

		micTmpint = tmpInt
	})

	notReservedGuitars := make([]string, 1)
	for _, notReservedGuitar := range notReservedEquipments[1] {
		notReservedGuitars = append(notReservedGuitars, notReservedGuitar.Name)
	}
	var guitTmpint int
	form.AddDropDown("Выберите гитару: ", notReservedGuitars, 0, func(tmp string, tmpInt int) {
		if tmpInt < 0 {
			return
		}

		guitTmpint = tmpInt
	})

	user, err := tui.app.UserSvc.GetByLogin(&dto.GetUserByLoginRequest{
		Login: tui.userInfo.Username,
	})
	if err != nil {
		errorTextView.SetText(err.Error())
		tui.ErrorForm(errorForm, pages, errorTextView, CreateReservePage)
		pages.SwitchToPage(ErrorPage)
	}

	form.AddButton("Создать бронь", func() {
		if len(notReservedEquipments[0]) != 0 {
			notReservedEquipment = append(notReservedEquipment, notReservedEquipments[0][micTmpint].Id)
		}
		if len(notReservedEquipments[1]) != 0 {
			notReservedEquipment = append(notReservedEquipment, notReservedEquipments[1][guitTmpint].Id)
		}
		err := tui.app.ReserveSvc.Add(&dto.AddReserveRequest{
			UserId:            user.Id,
			RoomId:            notReservedRoomId,
			ProducerId:        notReservedProducerId,
			InstrumentalistId: notReservedInstrumentalistId,
			TimeInterval:      userInterval,
			EquipmentId:       notReservedEquipment,
		})

		if err != nil {
			errorTextView.SetText(err.Error())
			tui.ErrorForm(errorForm, pages, errorTextView, CreateReservePage)
			pages.SwitchToPage(ErrorPage)
		}
		pages.SwitchToPage(AuthorizedPage)
	})
	form.AddButton("Назад", func() {
		form.Clear(true)
		tui.CreateValidatePage(form, pages)
		pages.SwitchToPage(CreateValidatePage)
	})
	return form
}

func (tui *Tui) CreateValidatePage(form *tview.Form, pages *tview.Pages) *tview.Form {
	errorTextView := tview.NewTextView().
		SetDynamicColors(true).
		SetRegions(true).
		SetWrap(true).
		SetTextAlign(tview.AlignCenter).
		SetText("")

	studios, err := tui.app.StudioSvc.GetAll(&dto.GetStudioAllRequest{})
	if err != nil {
		errorTextView.SetText(err.Error())
		tui.ErrorForm(errorForm, pages, errorTextView, CreateValidatePage)
		pages.SwitchToPage(ErrorPage)
	}
	studiosName := make([]string, 0)

	for _, studio := range studios {
		studiosName = append(studiosName, studio.Name)
	}

	var studioId int64
	flag := true
	form.AddDropDown("Выберите студию: ", studiosName, 0, func(tmp string, tmpInt int) {
		if tmpInt < 0 {
			flag = false
			return
		}
		studioId = studios[tmpInt].Id
	})

	var dateString, startHourString, endHourString string

	form.AddInputField("Введите дату начала брони в форме YYYY-MM-DD", "", 20, nil, func(tmp string) {
		dateString = tmp
	})
	form.AddInputField("Введите час начала брони в форме (от 0 до 23):", "", 20, nil, func(tmp string) {
		startHourString = tmp
	})
	form.AddInputField("Введите час конца брони в форме (от 0 до 23):", "", 20, nil, func(name string) {
		endHourString = name
	})
	//form.AddTextView("debug", strconv.FormatInt(studioId, 10), 20, 20, true, true)

	form.AddButton("Продолжить", func() {
		if flag == false {
			errorTextView.SetText("Нет доступных комнат")
			tui.ErrorForm(errorForm, pages, errorTextView, CreateValidatePage)
			pages.SwitchToPage(ErrorPage)
			return
		}
		startTimeString := strings.TrimSpace(dateString) + " " +
			strings.TrimSpace(startHourString) + ":00:00"

		endTimeString := strings.TrimSpace(dateString) + " " +
			strings.TrimSpace(endHourString) + ":00:00"

		//_, err = time_parser.StringToDate(startTimeString)
		startTime, err := time_parser.StringToDate(startTimeString)
		if err != nil {
			errorTextView.SetText(err.Error())
			tui.ErrorForm(errorForm, pages, errorTextView, CreateValidatePage)
			pages.SwitchToPage(ErrorPage)
		}
		//_, err = time_parser.StringToDate(endTimeString)
		endTime, err := time_parser.StringToDate(endTimeString)
		if err != nil {
			errorTextView.SetText(err.Error())
			tui.ErrorForm(errorForm, pages, errorTextView, CreateValidatePage)
			pages.SwitchToPage(ErrorPage)
		}

		userInterval := serviceImpl.NewTimeInterval(startTime, endTime)

		notReservedRooms := make([]*model.Room, 0)
		notReservedEquipments := make([][]*model.Equipment, 0)
		notReservedProducers := make([]*model.Producer, 0)
		notReservedInstrumentalists := make([]*model.Instrumentalist, 0)

		notReservedRooms,
			notReservedEquipments,
			notReservedProducers,
			notReservedInstrumentalists,
			err = tui.app.ValidateTimeSvc.GetSuitableStuff(&dto.GetSuitableStuffRequest{
			TimeInterval: userInterval,
			StudioId:     studioId,
		})
		if err != nil {
			errorTextView.SetText(err.Error())
			tui.ErrorForm(errorForm, pages, errorTextView, CreateValidatePage)
			pages.SwitchToPage(ErrorPage)
		}

		form.Clear(true)
		tui.CreateReservePage(form,
			pages,
			notReservedRooms,
			notReservedEquipments,
			notReservedProducers,
			notReservedInstrumentalists,
			userInterval,
		)
		pages.SwitchToPage(CreateReservePage)
	})
	form.AddButton("Назад", func() {
		pages.SwitchToPage(AuthorizedPage)
	})

	return form
}

//func (tui *Tui) appendToStudiosList(list *tview.List, pages *tview.Pages, studios []*model.Studio) {
//	list.Clear()
//	for _, studio := range studios {
//		list.AddItem(
//			strconv.Itoa(int(studio.Id)),
//			fmt.Sprintf("%s",
//				studio.Name,
//			),
//			'*',
//			nil)
//	}
//	list.AddItem(
//		"Back",
//		"",
//		'b',
//		func() {
//			pages.SwitchToPage("Show info cards")
//		},
//	)
//}

func (tui *Tui) DeleteReservePage(form *tview.Form, pages *tview.Pages) *tview.Form {
	errorTextView := tview.NewTextView().
		SetDynamicColors(true).
		SetRegions(true).
		SetWrap(true).
		SetTextAlign(tview.AlignCenter).
		SetText("")

	user, err := tui.app.UserSvc.GetByLogin(&dto.GetUserByLoginRequest{
		Login: tui.userInfo.Username,
	})
	if err != nil {
		errorTextView.SetText(err.Error())
		tui.ErrorForm(errorForm, pages, errorTextView, DeleteReservePage)
		pages.SwitchToPage(ErrorPage)
	}
	if user == nil {
		pages.SwitchToPage(AuthorizedPage)
	}
	reserves, err := tui.app.UserSvc.GetReserves(&dto.GetUserReservesRequest{
		Id: user.Id,
	})
	if err != nil {
		errorTextView.SetText(err.Error())
		tui.ErrorForm(errorForm, pages, errorTextView, DeleteReservePage)
		pages.SwitchToPage(ErrorPage)
	}

	reservesIntervals := make([]string, 0)
	for _, reserve := range reserves {
		reservesIntervals = append(reservesIntervals, reserve.TimeInterval.StartTime.Format("2006-Jan-02 15:04")+" - "+
			reserve.TimeInterval.EndTime.Format("15:04"))
	}
	var reserveId int64
	form.AddDropDown("Выберите бронь: ", reservesIntervals, 0, func(tmp string, tmpInt int) {
		if tmpInt < 0 {
			return
		}
		reserveId = reserves[tmpInt].Id
	})
	//return form
	form.AddButton("Снять бронь", func() {
		err := tui.app.ReserveSvc.Delete(&dto.DeleteReserveRequest{
			Id: reserveId,
		})
		if err != nil {
			errorTextView.SetText(err.Error())
			tui.ErrorForm(errorForm, pages, errorTextView, DeleteReservePage)
			pages.SwitchToPage(ErrorPage)
		}
		pages.SwitchToPage(AuthorizedPage)

	})
	form.AddButton("Назад", func() {
		pages.SwitchToPage(AuthorizedPage)
	})
	return form
}

func (tui *Tui) UpdateUserPage(form *tview.Form, pages *tview.Pages) *tview.Form {
	errorTextView := tview.NewTextView().
		SetDynamicColors(true).
		SetRegions(true).
		SetWrap(true).
		SetTextAlign(tview.AlignCenter).
		SetText("")
	user, err := tui.app.UserSvc.GetByLogin(&dto.GetUserByLoginRequest{
		Login: tui.userInfo.Username,
	})
	if err != nil {
		errorTextView.SetText(err.Error())
		tui.ErrorForm(errorForm, pages, errorTextView, DeleteReservePage)
		pages.SwitchToPage(ErrorPage)
		return form
	}
	var login string
	tmpStr := fmt.Sprintf("Введите новый логин (%s):", user.Login)
	form.AddInputField(tmpStr, "", 20, nil, func(tmp string) {
		login = tmp
	})

	var password string
	form.AddInputField("Введите новый пароль", "", 20, nil, func(tmp string) {
		password = tmp
	})

	var firstName string
	tmpStr = fmt.Sprintf("Введите новыое имя (%s):", user.FirstName)
	form.AddInputField(tmpStr, "", 20, nil, func(tmp string) {
		firstName = tmp
	})

	var secondName string
	tmpStr = fmt.Sprintf("Введите новую фамилию (%s):", user.SecondName)
	form.AddInputField(tmpStr, "", 20, nil, func(tmp string) {
		secondName = tmp
	})

	var thirdName string
	tmpStr = fmt.Sprintf("Введите новое отчество (%s):", user.Login)
	form.AddInputField(tmpStr, "", 20, nil, func(tmp string) {
		thirdName = tmp
	})

	form.AddButton("Изменить данные", func() {
		err := tui.app.UserSvc.Update(&dto.UpdateUserRequest{
			Id:         user.Id,
			Login:      login,
			Password:   password,
			FirstName:  firstName,
			SecondName: secondName,
			ThirdName:  thirdName,
		})
		if err != nil {
			errorTextView.SetText(err.Error())
			tui.ErrorForm(errorForm, pages, errorTextView, DeleteReservePage)
			pages.SwitchToPage(ErrorPage)
		}
		tui.userInfo.Username = login
		pages.SwitchToPage(AuthorizedPage)
	})
	form.AddButton("Назад", func() {
		pages.SwitchToPage(AuthorizedPage)
	})
	return form
}
