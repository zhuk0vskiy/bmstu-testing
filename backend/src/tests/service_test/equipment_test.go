package service_test

import (
	"backend/src/internal/model/dto"
	"backend/src/internal/repository/interface/mocks"
	"backend/src/internal/service/impl"
	"backend/src/tests/utils"
	"context"
	"fmt"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
)

type EquipmentSuite struct {
	suite.Suite
}

func (suite *EquipmentSuite) TestEquipmentAdd01(t provider.T) {
	t.Title("[Add] success")
	t.Tags("equipment", "add")
	t.Parallel()

	t.WithNewStep("successfully add", func(sCtx provider.StepCtx) {
		ctx := context.Background()
		builder := utils.EquipmentBuilder{}
		equipmentAdd := builder.
			WithName("1").
			WithStudioId(1).
			WithEquipmentType(1).
			ToDto()

		//repo := pool.IUserRepository{t}
		equipmentRepo := new(mocks.IEquipmentRepository)
		reserveRepo := new(mocks.IReserveRepository)
		equipmentRepo.On("Add", ctx, &dto.AddEquipmentRequest{
			Name:     equipmentAdd.Name,
			StudioId: equipmentAdd.StudioId,
			Type:     equipmentAdd.EquipmentType,
		}).Return(
			nil,
		).Once()

		logger := utils.NewMockLogger()

		service := impl.NewEquipmentService(logger, equipmentRepo, reserveRepo)
		sCtx.WithNewParameters("ctx", ctx, "request", equipmentAdd)
		err := service.Add(&dto.AddEquipmentRequest{
			Name:     equipmentAdd.Name,
			Type:     equipmentAdd.EquipmentType,
			StudioId: equipmentAdd.StudioId,
		})

		sCtx.Assert().NoError(err)
	})
}

func (suite *EquipmentSuite) TestEquipmentAdd03(t provider.T) {
	t.Title("[Add] incorrect name")
	t.Tags("equipment", "add")
	t.Parallel()

	t.WithNewStep("fail add", func(sCtx provider.StepCtx) {
		ctx := context.Background()
		builder := utils.EquipmentBuilder{}
		equipmentAdd := builder.
			WithName("1").
			WithStudioId(1).
			WithEquipmentType(1).
			ToDto()

		//repo := pool.IUserRepository{t}
		equipmentRepo := new(mocks.IEquipmentRepository)
		reserveRepo := new(mocks.IReserveRepository)
		equipmentRepo.On("Add", ctx, &dto.AddEquipmentRequest{
			Name:     equipmentAdd.Name,
			StudioId: equipmentAdd.StudioId,
			Type:     equipmentAdd.EquipmentType,
		}).Return(
			fmt.Errorf("incorrect name"),
		).Once()

		logger := utils.NewMockLogger()

		service := impl.NewEquipmentService(logger, equipmentRepo, reserveRepo)
		sCtx.WithNewParameters("ctx", ctx, "request", equipmentAdd)
		err := service.Add(&dto.AddEquipmentRequest{
			Name:     equipmentAdd.Name,
			Type:     equipmentAdd.EquipmentType,
			StudioId: equipmentAdd.StudioId,
		})

		sCtx.Assert().Error(err)
	})
}

//func (suite *EquipmentSuite) TestEquipmentAdd02(t provider.T) {
//	t.Title("[Add] incorrect name")
//	t.Tags("equipment", "add")
//	t.Parallel()
//
//	t.WithNewStep("fail add", func(sCtx provider.StepCtx) {
//		ctx := context.Background()
//		builder := utils.EquipmentBuilder{}
//		equipmentAdd := builder.
//			WithName("1").
//			WithStudioId(1).
//			WithEquipmentType(1).
//			ToDto()
//
//		//repo := pool.IUserRepository{t}
//		//equipmentRepo := new(mocks.IEquipmentRepository)
//
//		//mock, err := pgxmock.NewPool()?
//		pool := new(mocks.IPool)
//
//		reserveRepo := postgresql.NewReserveRepository(pool)
//		equipmentRepo := postgresql.NewEquipmentRepository(pool)
//
//		logger := utils.NewMockLogger()
//
//		service := impl.NewEquipmentService(logger, equipmentRepo, reserveRepo)
//		sCtx.WithNewParameters("ctx", ctx, "request", equipmentAdd)
//		err := service.Add(&dto.AddEquipmentRequest{
//			Name:     equipmentAdd.Name,
//			Type:     equipmentAdd.EquipmentType,
//			StudioId: equipmentAdd.StudioId,
//		})
//
//		sCtx.Assert().Error(err)
//	})
//}
