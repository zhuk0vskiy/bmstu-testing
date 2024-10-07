package repository_test

import (
	"backend/src/internal/model/dto"
	"backend/src/internal/repository/interface/mocks"
	"backend/src/tests/utils"
	"context"
	"fmt"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
)

type EquipmentRepositorySuite struct {
	suite.Suite

	mockRepo mocks.IEquipmentRepository
}

func (suite *EquipmentRepositorySuite) TestEquipmentAdd01(t provider.T) {
	t.Title("[Add] success")
	t.Tags("equipment", "add")
	t.Parallel()

	t.WithNewStep("successfully add", func(sCtx provider.StepCtx) {
		ctx := context.Background()
		builder := utils.EquipmentBuilder{}
		equipment := builder.
			WithName("1").
			WithStudioId(1).
			WithEquipmentType(1).
			ToDto()

		//repo := pool.IUserRepository{t}

		suite.mockRepo.On("Add", ctx, &dto.AddEquipmentRequest{
			Name:     equipment.Name,
			StudioId: equipment.StudioId,
			Type:     equipment.EquipmentType,
		}).Return(
			nil,
		).Once()

		sCtx.WithNewParameters("ctx", ctx, "request", equipment)
		err := suite.mockRepo.Add(ctx, &dto.AddEquipmentRequest{
			Name:     equipment.Name,
			Type:     equipment.EquipmentType,
			StudioId: equipment.StudioId,
		})

		sCtx.Assert().NoError(err)
	})
}

func (suite *EquipmentRepositorySuite) TestEquipmentAdd03(t provider.T) {
	t.Title("[Add] incorrect name")
	t.Tags("equipment", "add")
	t.Parallel()

	t.WithNewStep("fail add", func(sCtx provider.StepCtx) {
		ctx := context.Background()
		builder := utils.EquipmentBuilder{}
		equipmentAdd := builder.
			WithName("").
			WithStudioId(1).
			WithEquipmentType(1).
			ToDto()

		suite.mockRepo.On("Add", ctx, &dto.AddEquipmentRequest{
			Name:     equipmentAdd.Name,
			StudioId: equipmentAdd.StudioId,
			Type:     equipmentAdd.EquipmentType,
		}).Return(
			fmt.Errorf("incorrect name"),
		).Once()

		sCtx.WithNewParameters("ctx", ctx, "request", equipmentAdd)
		err := suite.mockRepo.Add(ctx, &dto.AddEquipmentRequest{
			Name:     equipmentAdd.Name,
			Type:     equipmentAdd.EquipmentType,
			StudioId: equipmentAdd.StudioId,
		})

		sCtx.Assert().Error(err)
	})
}

//func (suite *EquipmentRepositorySuite) TestEquipmentAdd02(t provider.T) {
//	t.Title("[Add] success")
//	t.Tags("equipment", "add")
//	t.Parallel()
//
//	t.WithNewStep("successfully add", func(sCtx provider.StepCtx) {
//		ctx := context.Background()
//		builder := utils.EquipmentBuilder{}
//		equipmentAdd := builder.
//			WithName("1").
//			WithStudioId(1).
//			WithEquipmentType(1).
//			ToDto()
//
//		//mock, err := pgxmock.NewPool()
//		//if err != nil {
//		//	t.Fatal(err)
//		//}
//		//repo := pool.IUserRepository{t}
//		//equipmentRepo := new(mocks.IEquipmentRepository)
//		//reserveRepo := new(mocks.IReserveRepository)
//		//equipmentRepo.On("Add", ctx, &dto.AddEquipmentRequest{
//		//	Name:     equipmentAdd.Name,
//		//	StudioId: equipmentAdd.StudioId,
//		//	Type:     equipmentAdd.EquipmentType,
//		//}).Return(
//		//	nil,
//		//).Once()
//
//		pool := new(mocks.IPool)
//		//logger := utils.NewMockLogger()
//
//		service := postgresql.NewEquipmentRepository(pool)
//		sCtx.WithNewParameters("ctx", ctx, "request", equipmentAdd)
//		err := service.Add(ctx, &dto.AddEquipmentRequest{
//			Name:     equipmentAdd.Name,
//			Type:     equipmentAdd.EquipmentType,
//			StudioId: equipmentAdd.StudioId,
//		})
//
//		sCtx.Assert().NoError(err)
//	})
//}
