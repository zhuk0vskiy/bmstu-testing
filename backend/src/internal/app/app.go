package app

import (
	"backend/src/config"
	"backend/src/internal/repository/impl/postgresql"
	serviceImpl "backend/src/internal/service/impl"
	serviceInterface "backend/src/internal/service/interface"
	"backend/src/pkg/base"
	"backend/src/pkg/logger"
	"github.com/jackc/pgx/v5/pgxpool"
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
	reserveSvc := serviceImpl.NewReserveService(logger, reserveRepo, roomRepo, producerRepo, instrumentalistRepo)
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
