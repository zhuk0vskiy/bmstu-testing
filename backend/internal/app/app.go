package app

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"ppo/domain"
	"ppo/internal/config"
	"ppo/internal/storage/postgres"
	"ppo/pkg/base"
	"ppo/pkg/logger"
	"ppo/services"
)

type App struct {
	Config                *config.Config
	Logger                logger.ILogger
	AuthService           domain.IAuthService
	UserService           domain.IUserService
	CommentService        domain.ICommentService
	SaladService          domain.ISaladInteractor
	SaladTypeService      domain.ISaladTypeService
	RecipeService         domain.IRecipeService
	RecipeStepService     domain.IRecipeStepInteractor
	IngredientService     domain.IIngredientService
	IngredientTypeService domain.IIngredientTypeService
	MeasurementService    domain.IMeasurementService
}

func NewApp(db *pgxpool.Pool, cfg *config.Config, logger logger.ILogger) *App {
	authRepo := postgres.NewAuthRepository(db)
	userRepo := postgres.NewUserRepository(db)

	commentRepo := postgres.NewCommentRepository(db)
	saladRepo := postgres.NewSaladRepository(db)
	recipeRepo := postgres.NewRecipeRepository(db)
	recipeStepRepo := postgres.NewRecipeStepRepository(db)
	ingredientRepo := postgres.NewIngredientRepository(db)
	ingredientTypeRepo := postgres.NewIngredientTypeRepository(db)
	saladTypeRepo := postgres.NewSaladTypeRepository(db)
	measurementRepo := postgres.NewMeasrementRepository(db)

	validatorRepo := postgres.NewKeywordValidatorRepository(db)

	crypto := base.NewHashCrypto()

	keyWordValidator, err := services.NewKeywordValidatorService(context.Background(), validatorRepo, logger)
	if err != nil {
		return nil
	}
	urlValidator := services.NewUrlValidatorService(logger)
	validators := []domain.IValidatorService{keyWordValidator, urlValidator}

	return &App{
		Config:         cfg,
		Logger:         logger,
		AuthService:    services.NewAuthService(authRepo, logger, crypto, cfg.Jwt.Key),
		UserService:    services.NewUserService(userRepo, logger),
		CommentService: services.NewCommentService(commentRepo, logger),
		SaladService: services.NewSaladInteractor(
			services.NewSaladService(saladRepo, logger),
			validators,
		),
		RecipeService: services.NewRecipeService(recipeRepo, logger),
		RecipeStepService: services.NewRecipeStepInteractor(
			services.NewRecipeStepService(recipeStepRepo, logger),
			validators,
		),
		IngredientService:     services.NewIngredientService(ingredientRepo, logger),
		IngredientTypeService: services.NewIngredientTypeService(ingredientTypeRepo, logger),
		SaladTypeService:      services.NewSaladTypeService(saladTypeRepo, logger),
		MeasurementService:    services.NewMeasurementService(measurementRepo, logger),
	}
}
