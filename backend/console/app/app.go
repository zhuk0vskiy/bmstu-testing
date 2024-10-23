package app

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rivo/tview"
	"log"
	mail2 "net/mail"
	"ppo/domain"
	"ppo/internal/config"
	"ppo/internal/storage/postgres"
	"ppo/pkg/base"
	"ppo/pkg/logger"
	"ppo/services"
	"ppo/services/dto"
	"strconv"
	"strings"
)

var (
	pages      = tview.NewPages()
	app        = tview.NewApplication()
	form       = tview.NewForm()
	errorForm  = tview.NewForm()
	list       = tview.NewList().ShowSecondaryText(true)
	recipeList = tview.NewList().ShowSecondaryText(true)
)

const (
	GuestPage      = "Menu (guest)"
	AuthorizedPage = "Menu (authorized)"
	AdminPage      = "Menu (admin)"

	RegisterPage = "Register"
	LoginPage    = "Login"
	ErrorPage    = "Error page"

	SaladsRequestPage = "Salads request page"
	SaladsPage        = "Salads page"
	RecipePage        = "Recipe page"
	CreateSaladPage   = "Create salad page"
	CreateRecipeStep  = "Create recipe step"
	IngredientsPage   = "Ingredients page"
	SaladTypesPage    = "Salad types page"
	MeasurementsPage  = "Measurements page"
)

const (
	UnauthorizedUser = "unauthorized"
	AuthorizedUser   = "user"
	AuthorizedAdmin  = "admin"
)

type App struct {
	config *config.Config

	authService domain.IAuthService
	userService domain.IUserService

	commentService   domain.ICommentService
	saladService     domain.ISaladInteractor
	saladTypeService domain.ISaladTypeService

	recipeService      domain.IRecipeService
	recipeStepService  domain.IRecipeStepInteractor
	ingredientService  domain.IIngredientService
	measurementService domain.IMeasurementService
}

type Tui struct {
	app      *App
	userInfo *base.JwtPayload
}

func Run(db *pgxpool.Pool, cfg *config.Config, logger logger.ILogger) *tview.Application {
	var tui Tui
	tui.app = NewApp(db, cfg, logger)
	tui.userInfo = new(base.JwtPayload)
	tui.userInfo.Role = UnauthorizedUser

	pages.AddPage(GuestPage, tui.CreateGuestMenu(form, pages, app), true, true)
	pages.AddPage(RegisterPage, form, true, true).
		AddPage(LoginPage, form, true, true).
		AddPage(ErrorPage, errorForm, true, true).
		AddPage(SaladsRequestPage, form, true, true).
		AddPage(SaladsPage, list, true, true).
		AddPage(RecipePage, recipeList, true, true).
		AddPage(IngredientsPage, recipeList, true, true).
		AddPage(SaladTypesPage, recipeList, true, true).
		AddPage(MeasurementsPage, recipeList, true, true)

	pages.AddPage(AuthorizedPage, tui.CreateAuthorizedMenu(form, pages, app), true, true).
		AddPage(CreateSaladPage, form, true, true).
		AddPage(CreateRecipeStep, form, true, true)

	pages.AddPage(AdminPage, tui.CreateAdminMenu(form, pages, app), true, true)

	pages.SwitchToPage(GuestPage)
	if err := app.SetRoot(pages, true).EnableMouse(true).Run(); err != nil {
		log.Fatalln(err)
	}

	return app
}

func NewApp(db *pgxpool.Pool, cfg *config.Config, logger logger.ILogger) *App {
	authRepo := postgres.NewAuthRepository(db)
	userRepo := postgres.NewUserRepository(db)

	commentRepo := postgres.NewCommentRepository(db)
	saladRepo := postgres.NewSaladRepository(db)
	recipeRepo := postgres.NewRecipeRepository(db)
	recipeStepRepo := postgres.NewRecipeStepRepository(db)
	ingredientRepo := postgres.NewIngredientRepository(db)
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
		config:         cfg,
		authService:    services.NewAuthService(authRepo, logger, crypto, cfg.Jwt.Key),
		userService:    services.NewUserService(userRepo, logger),
		commentService: services.NewCommentService(commentRepo, logger),
		saladService: services.NewSaladInteractor(
			services.NewSaladService(saladRepo, logger),
			validators,
		),
		recipeService: services.NewRecipeService(recipeRepo, logger),
		recipeStepService: services.NewRecipeStepInteractor(
			services.NewRecipeStepService(recipeStepRepo, logger),
			validators,
		),
		ingredientService:  services.NewIngredientService(ingredientRepo, logger),
		saladTypeService:   services.NewSaladTypeService(saladTypeRepo, logger),
		measurementService: services.NewMeasurementService(measurementRepo, logger),
	}
}

func (tui *Tui) ErrorForm(form *tview.Form, pages *tview.Pages, textView *tview.TextView, prevPage string) *tview.Form {
	form.Clear(true)
	form.AddFormItem(textView)

	form.AddButton("OK", func() {
		pages.SwitchToPage(prevPage)
		form.Clear(true)
	})

	return form
}

func (tui *Tui) RegisterForm(form *tview.Form, pages *tview.Pages) *tview.Form {
	user := &domain.User{}
	errorTextView := tview.NewTextView().
		SetDynamicColors(true).
		SetRegions(true).
		SetWrap(true).
		SetTextAlign(tview.AlignCenter).
		SetText("")

	form.AddInputField("Name", "", 20, nil, func(name string) {
		user.Name = name
	})
	form.AddInputField("Email", "", 20, nil, func(email string) {
		mail := mail2.Address{
			Name:    "",
			Address: email,
		}
		user.Email = mail
	})
	form.AddInputField("Username", "", 20, nil, func(username string) {
		user.Username = username
	})
	form.AddPasswordField("Password", "", 20, '*', func(password string) {
		user.Password = password
	})

	form.AddButton("Register", func() {
		_, err := tui.app.authService.Register(context.Background(), user)
		if err != nil {
			errorTextView.SetText(err.Error())
			tui.ErrorForm(errorForm, pages, errorTextView, RegisterPage)
			pages.SwitchToPage(ErrorPage)
			return
		}
		pages.SwitchToPage(AuthorizedPage)

		fullDataUser, _ := tui.app.userService.GetByUsername(context.Background(), user.Username)

		//tui.userInfo.Username = fullDataUser.Username
		tui.userInfo.ID = fullDataUser.ID.String()

		tui.userInfo.Role = fullDataUser.Role
	})
	form.AddButton("Back", func() {
		pages.SwitchToPage(GuestPage)
	})

	return form
}

func (tui *Tui) LoginForm(form *tview.Form, pages *tview.Pages) *tview.Form {
	authInfo := &domain.UserAuth{}
	errorTextView := tview.NewTextView().
		SetDynamicColors(true).
		SetRegions(true).
		SetWrap(true).
		SetTextAlign(tview.AlignCenter).
		SetText("")

	form.AddInputField("Username", "", 20, nil, func(username string) {
		authInfo.Username = username
	})
	form.AddPasswordField("Password", "", 20, '*', func(password string) {
		authInfo.Password = password
	})

	form.AddButton("Login", func() {
		token, err := tui.app.authService.Login(context.Background(), authInfo)
		if err != nil {
			errorTextView.SetText(err.Error())
			tui.ErrorForm(errorForm, pages, errorTextView, LoginPage)
			pages.SwitchToPage(ErrorPage)
			return
		}
		userInfo, err := base.VerifyAuthToken(token, tui.app.config.Jwt.Key)
		if err != nil {
			errorTextView.SetText(err.Error())
			tui.ErrorForm(errorForm, pages, errorTextView, LoginPage)
			pages.SwitchToPage(ErrorPage)
			return
		}
		tui.userInfo = userInfo
		if userInfo.Role == AuthorizedAdmin {
			pages.SwitchToPage(AdminPage)
		} else if userInfo.Role == AuthorizedUser {
			pages.SwitchToPage(AuthorizedPage)
		}
	})
	form.AddButton("Back", func() {
		pages.SwitchToPage(GuestPage)
	})

	return form
}

func stringToUUIDs(uuids string) ([]uuid.UUID, error) {
	tmp := make([]uuid.UUID, 0)
	for _, id := range strings.Fields(uuids) {
		currUuid, err := uuid.Parse(id)
		if err != nil {
			return nil, err
		}
		tmp = append(tmp, currUuid)
	}
	return tmp, nil
}

func (tui *Tui) ShowSalads(form *tview.Form, pages *tview.Pages, list *tview.List) *tview.Form {
	page := 1
	errorTextView := tview.NewTextView().
		SetDynamicColors(true).
		SetRegions(true).
		SetWrap(true).
		SetTextAlign(tview.AlignCenter).
		SetText("")

	filter := new(dto.RecipeFilter)
	filter.Status = dto.PublishedSaladStatus

	ingredientIds := ""
	typeIds := ""

	form.AddInputField("Ingredient IDs", "", 100, nil, func(ingredients string) {
		ingredientIds = ingredients
	})

	form.AddInputField("Type IDs", "", 100, nil, func(types string) {
		typeIds = types
	})

	form.AddButton("Show", func() {
		tmp, err := stringToUUIDs(ingredientIds)
		if err != nil {
			errorTextView.SetText(err.Error())
			tui.ErrorForm(errorForm, pages, errorTextView, SaladsRequestPage)
			pages.SwitchToPage(ErrorPage)
			return
		}
		filter.AvailableIngredients = tmp

		tmp, err = stringToUUIDs(typeIds)
		if err != nil {
			errorTextView.SetText(err.Error())
			tui.ErrorForm(errorForm, pages, errorTextView, SaladsRequestPage)
			pages.SwitchToPage(ErrorPage)
			return
		}
		filter.SaladTypes = tmp

		salads, _, err := tui.app.saladService.GetAll(context.Background(), filter, page)
		if err != nil {
			errorTextView.SetText(err.Error())
			tui.ErrorForm(errorForm, pages, errorTextView, SaladsRequestPage)
			pages.SwitchToPage(ErrorPage)
			return
		}

		tui.appendSaladsToList(list, pages, salads)
		list.AddItem("Back", "", 'b', func() {
			pages.SwitchToPage(SaladsRequestPage)
		})
		pages.SwitchToPage(SaladsPage)
	})

	if tui.userInfo.Role == AuthorizedAdmin || tui.userInfo.Role == AuthorizedUser {
		var tmpId string

		var commentSaladRating int
		var commentSaladText string

		form.AddInputField("Salad ID (to comment)", "", 100, nil, func(saladId string) {
			tmpId = saladId
		})
		form.AddDropDown("Rating", []string{"1", "2", "3", "4", "5"}, 4, func(option string, optionIndex int) {
			commentSaladRating = optionIndex + 1
		})
		form.AddInputField("Text", "", 100, nil, func(text string) {
			commentSaladText = text
		})

		form.AddButton("Add comment", func() {
			commentSaladId, err := uuid.Parse(tmpId)
			if err != nil {
				errorTextView.SetText(err.Error())
				tui.ErrorForm(errorForm, pages, errorTextView, SaladsRequestPage)
				pages.SwitchToPage(ErrorPage)
				return
			}

			//user, _ := tui.app.userService.GetByUsername(context.Background(), tui.userInfo.Username)
			id, _ := uuid.Parse(tui.userInfo.ID)
			user, _ := tui.app.userService.GetById(context.Background(), id)

			comment := domain.Comment{
				AuthorID: user.ID,
				SaladID:  commentSaladId,
				Text:     commentSaladText,
				Rating:   commentSaladRating,
			}
			err = tui.app.commentService.Create(context.Background(), &comment)
			if err != nil {
				errorTextView.SetText(err.Error())
				tui.ErrorForm(errorForm, pages, errorTextView, SaladsRequestPage)
				pages.SwitchToPage(ErrorPage)
				return
			}
		})
	}

	form.AddButton("Show ingredients", func() {
		recipeList = tui.appendIngredientsToList(uuid.Nil, recipeList)
		recipeList.AddItem("Back", "", 'b', func() {
			pages.SwitchToPage(SaladsRequestPage)
		})
		pages.SwitchToPage(IngredientsPage)
	})

	form.AddButton("Show types", func() {
		recipeList = tui.appendSaladTypesToList(uuid.Nil, recipeList)
		recipeList.AddItem("Back", "", 'b', func() {
			pages.SwitchToPage(SaladsRequestPage)
		})
		pages.SwitchToPage(SaladTypesPage)
	})

	form.AddButton("Back", func() {
		form.Clear(true)
		tui.BackToMenu(pages)
	})

	return form
}

func (tui *Tui) appendRecipeToList(recipe *domain.Recipe, pages *tview.Pages) *tview.List {
	recipeList.Clear()
	recipeList.AddItem(recipe.ID.String(),
		fmt.Sprintf("numberOfServings: %d, rating: %f, timeToCook: %d", recipe.NumberOfServings, recipe.Rating, recipe.TimeToCook),
		'*',
		nil)

	types, _ := tui.app.saladTypeService.GetAllBySaladId(context.Background(), recipe.SaladID)
	recipeList.AddItem("TYPES:",
		"",
		'*',
		nil)
	for _, saladType := range types {
		recipeList.AddItem(saladType.ID.String(),
			fmt.Sprintf("Name: %s, description: %s", saladType.Name, saladType.Description),
			'*',
			nil)
	}

	recipeList.AddItem("INGREDIENTS:",
		"",
		'*',
		nil)
	ingredients, _ := tui.app.ingredientService.GetAllByRecipeId(context.Background(), recipe.ID)
	for _, ingredient := range ingredients {
		recipeList.AddItem(ingredient.ID.String(),
			fmt.Sprintf("typeId: %s, name: %s, calories: %d", ingredient.TypeID.String(), ingredient.Name, ingredient.Calories),
			'*',
			nil)
	}

	recipeList.AddItem("STEPS:",
		"",
		'*',
		nil)
	steps, _ := tui.app.recipeStepService.GetAllByRecipeID(context.Background(), recipe.ID)
	for _, step := range steps {
		recipeList.AddItem(step.ID.String(),
			fmt.Sprintf("%d) name: %s, description: %s", step.StepNum, step.Name, step.Description),
			'*',
			nil)
	}

	if tui.userInfo.Role == AuthorizedAdmin && recipe.Status == dto.ModerationSaladStatus {
		recipeList.AddItem("Verify", "", 'v', func() {
			recipe.Status = dto.PublishedSaladStatus
			tui.app.recipeService.Update(context.Background(), recipe)

			form.Clear(true)
			tui.ShowSaladsToModerate(form, pages, list)
			pages.SwitchToPage(SaladsPage)
		})
	}
	if tui.userInfo.Role == AuthorizedAdmin && recipe.Status == dto.PublishedSaladStatus {
		recipeList.AddItem("Delete", "", 'd', func() {
			tui.app.saladService.DeleteById(context.Background(), recipe.SaladID)
			pages.SwitchToPage(SaladsRequestPage)
		})
	}

	return recipeList
}

func (tui *Tui) appendSaladsToList(list *tview.List, pages *tview.Pages, salads []*domain.Salad) {
	list.Clear()

	for _, salad := range salads {
		list.AddItem(
			salad.ID.String(),
			fmt.Sprintf("Author: %s, Name: %s, Description: %s",
				salad.AuthorID.String(),
				salad.Name,
				salad.Description),
			'*',
			nil,
		).AddItem("Go in", "", '*', func() {
			recipe, _ := tui.app.recipeService.GetBySaladId(context.Background(), salad.ID)
			tui.appendRecipeToList(recipe, pages)
			recipeList.AddItem("Back", "", 'b', func() {
				pages.SwitchToPage(SaladsPage)
			})
			pages.SwitchToPage(RecipePage)
		})
	}
}

func (tui *Tui) BackToMenu(pages *tview.Pages) {
	if tui.userInfo.Role == AuthorizedAdmin {
		pages.SwitchToPage(AdminPage)
	} else if tui.userInfo.Role == AuthorizedUser {
		pages.SwitchToPage(AuthorizedPage)
	} else {
		pages.SwitchToPage(GuestPage)
	}
}

func (tui *Tui) CreateGuestMenu(form *tview.Form, pages *tview.Pages, exitFunc *tview.Application) *tview.List {
	return tview.NewList().
		AddItem("Register", "", '1', func() {
			form.Clear(true)
			tui.RegisterForm(form, pages)
			pages.SwitchToPage(RegisterPage)
		}).
		AddItem("Login", "", '2', func() {
			form.Clear(true)
			tui.LoginForm(form, pages)
			pages.SwitchToPage(LoginPage)
		}).
		AddItem("View salads", "", '3', func() {
			form.Clear(true)
			tui.ShowSalads(form, pages, list)
			pages.SwitchToPage(SaladsRequestPage)
		}).
		AddItem("Exit", "", '0', func() {
			exitFunc.Stop()
		})
}

func (tui *Tui) showRated(form *tview.Form, pages *tview.Pages, list *tview.List) *tview.Form {
	page := 1
	errorTextView := tview.NewTextView().
		SetDynamicColors(true).
		SetRegions(true).
		SetWrap(true).
		SetTextAlign(tview.AlignCenter).
		SetText("")

	//user, _ := tui.app.userService.GetByUsername(context.Background(), tui.userInfo.Username)
	id, _ := uuid.Parse(tui.userInfo.ID)
	user, _ := tui.app.userService.GetById(context.Background(), id)

	salads, _, err := tui.app.saladService.GetAllRatedByUser(context.Background(), user.ID, page)
	if err != nil {
		errorTextView.SetText(err.Error())
		tui.ErrorForm(errorForm, pages, errorTextView, SaladsRequestPage)
		pages.SwitchToPage(ErrorPage)
		return form
	}

	tui.appendSaladsToList(list, pages, salads)
	list.AddItem("Back", "", 'b', func() {
		form.Clear(true)
		tui.BackToMenu(pages)
	})

	return form
}

func (tui *Tui) createSaladForm(form *tview.Form, pages *tview.Pages) *tview.Form {
	errorTextView := tview.NewTextView().
		SetDynamicColors(true).
		SetRegions(true).
		SetWrap(true).
		SetTextAlign(tview.AlignCenter).
		SetText("")

	salad := new(domain.Salad)
	tmpName := "Salad"
	tmpDescription := "Description"

	recipe := new(domain.Recipe)
	ttcStr := "10"
	nosStr := "5"

	form.AddInputField("Name", "Salad", 20, nil, func(text string) {
		tmpName = text
	})
	form.AddInputField("Description", "Description", 20, nil, func(text string) {
		tmpDescription = text
	})

	form.AddInputField("Time to cook (minutes)", "10", 20, nil, func(text string) {
		ttcStr = text
	})
	form.AddInputField("Number of servings", "5", 20, nil, func(text string) {
		nosStr = text
	})

	form.AddButton("Create", func() {
		//user, _ := tui.app.userService.GetByUsername(context.Background(), tui.userInfo.Username)
		uid, _ := uuid.Parse(tui.userInfo.ID)
		user, _ := tui.app.userService.GetById(context.Background(), uid)

		salad.Name = tmpName
		salad.Description = tmpDescription
		salad.AuthorID = user.ID

		tmp, err := strconv.Atoi(ttcStr)
		if err != nil {
			errorTextView.SetText(err.Error())
			tui.ErrorForm(errorForm, pages, errorTextView, CreateSaladPage)
			pages.SwitchToPage(ErrorPage)
			return
		} else if tmp <= 0 {
			errorTextView.SetText("Negative or zero time to cook")
			tui.ErrorForm(errorForm, pages, errorTextView, CreateSaladPage)
			pages.SwitchToPage(ErrorPage)
			return
		}
		recipe.TimeToCook = tmp

		tmp, err = strconv.Atoi(nosStr)
		if err != nil {
			errorTextView.SetText(err.Error())
			tui.ErrorForm(errorForm, pages, errorTextView, CreateSaladPage)
			pages.SwitchToPage(ErrorPage)
			return
		} else if tmp <= 0 {
			errorTextView.SetText("Negative or zero number of servings")
			tui.ErrorForm(errorForm, pages, errorTextView, CreateSaladPage)
			pages.SwitchToPage(ErrorPage)
			return
		}
		recipe.NumberOfServings = tmp

		id, err := tui.app.saladService.Create(context.Background(), salad)
		if err != nil {
			errorTextView.SetText(err.Error())
			tui.ErrorForm(errorForm, pages, errorTextView, CreateSaladPage)
			pages.SwitchToPage(ErrorPage)
			return
		}
		salad.ID = id

		recipe.SaladID = salad.ID
		recipe.Status = dto.EditingSaladStatus
		_, err = tui.app.recipeService.Create(context.Background(), recipe)
		if err != nil {
			errorTextView.SetText(err.Error())
			tui.ErrorForm(errorForm, pages, errorTextView, CreateSaladPage)
			pages.SwitchToPage(ErrorPage)
			return
		}

		createdRecipe, _ := tui.app.recipeService.GetBySaladId(context.Background(), salad.ID)
		form.Clear(true)
		tui.createStepsForm(form, pages, createdRecipe)
		pages.SwitchToPage(CreateRecipeStep)
	})

	form.AddButton("Back", func() {
		form.Clear(true)
		tui.BackToMenu(pages)
	})

	return form
}

func (tui *Tui) createStepsForm(form *tview.Form, pages *tview.Pages, recipe *domain.Recipe) *tview.Form {
	errorTextView := tview.NewTextView().
		SetDynamicColors(true).
		SetRegions(true).
		SetWrap(true).
		SetTextAlign(tview.AlignCenter).
		SetText("")

	step := new(domain.RecipeStep)
	step.RecipeID = recipe.ID
	step.StepNum = 1

	form.AddInputField("Step name", "", 20, nil, func(text string) {
		step.Name = text
	})
	form.AddInputField("Step description", "", 20, nil, func(text string) {
		step.Description = text
	})

	form.AddButton("Add step", func() {
		err := tui.app.recipeStepService.Create(context.Background(), step)
		if err != nil {
			errorTextView.SetText(err.Error())
			tui.ErrorForm(errorForm, pages, errorTextView, CreateRecipeStep)
			pages.SwitchToPage(ErrorPage)
			return
		}
	})

	form.AddButton("Preview", func() {
		tui.appendRecipeToList(recipe, pages)
		recipeList.AddItem("Back", "", 'b', func() {
			pages.SwitchToPage(CreateRecipeStep)
		})
		pages.SwitchToPage(RecipePage)
	})

	form.AddButton("Send on moderation", func() {
		recipe.Status = dto.ModerationSaladStatus
		err := tui.app.recipeService.Update(context.Background(), recipe)
		if err != nil {
			errorTextView.SetText(err.Error())
			tui.ErrorForm(errorForm, pages, errorTextView, CreateRecipeStep)
			pages.SwitchToPage(ErrorPage)
			return
		}

		form.Clear(true)
		tui.BackToMenu(pages)
	})

	typeIds := ""
	form.AddInputField("Salad type ids", "", 100, nil, func(text string) {
		typeIds = text
	})
	form.AddButton("Link types", func() {
		saladTypes, err := stringToUUIDs(typeIds)
		if err != nil {
			errorTextView.SetText(err.Error())
			tui.ErrorForm(errorForm, pages, errorTextView, CreateRecipeStep)
			pages.SwitchToPage(ErrorPage)
			return
		}

		for _, saladType := range saladTypes {
			err = tui.app.saladTypeService.Link(context.Background(), recipe.SaladID, saladType)
			if err != nil {
				errorTextView.SetText(err.Error())
				tui.ErrorForm(errorForm, pages, errorTextView, CreateRecipeStep)
				pages.SwitchToPage(ErrorPage)
				return
			}
		}
	})
	form.AddButton("Unlink types", func() {
		saladTypes, err := stringToUUIDs(typeIds)
		if err != nil {
			errorTextView.SetText(err.Error())
			tui.ErrorForm(errorForm, pages, errorTextView, CreateRecipeStep)
			pages.SwitchToPage(ErrorPage)
			return
		}

		for _, saladType := range saladTypes {
			err = tui.app.saladTypeService.Unlink(context.Background(), recipe.SaladID, saladType)
			if err != nil {
				errorTextView.SetText(err.Error())
				tui.ErrorForm(errorForm, pages, errorTextView, CreateRecipeStep)
				pages.SwitchToPage(ErrorPage)
				return
			}
		}
	})

	ingredientId := ""
	measurementId := ""
	amountStr := ""
	form.AddInputField("Ingredient id", "", 100, nil, func(text string) {
		ingredientId = text
	})
	form.AddInputField("Measurement id", "", 100, nil, func(text string) {
		measurementId = text
	})
	form.AddInputField("Amount", "", 100, nil, func(text string) {
		amountStr = text
	})
	form.AddButton("Add ingredient", func() {
		ingrId, err := uuid.Parse(ingredientId)
		if err != nil {
			errorTextView.SetText(err.Error())
			tui.ErrorForm(errorForm, pages, errorTextView, CreateRecipeStep)
			pages.SwitchToPage(ErrorPage)
			return
		}
		mId, err := uuid.Parse(measurementId)
		if err != nil {
			errorTextView.SetText(err.Error())
			tui.ErrorForm(errorForm, pages, errorTextView, CreateRecipeStep)
			pages.SwitchToPage(ErrorPage)
			return
		}
		amount, err := strconv.Atoi(amountStr)
		if err != nil {
			errorTextView.SetText(err.Error())
			tui.ErrorForm(errorForm, pages, errorTextView, CreateRecipeStep)
			pages.SwitchToPage(ErrorPage)
			return
		}
		linkId, err := tui.app.ingredientService.Link(context.Background(), recipe.ID, ingrId)
		if err != nil {
			errorTextView.SetText(err.Error())
			tui.ErrorForm(errorForm, pages, errorTextView, CreateRecipeStep)
			pages.SwitchToPage(ErrorPage)
			return
		}
		err = tui.app.measurementService.UpdateLink(context.Background(), linkId, mId, amount)
		if err != nil {
			errorTextView.SetText(err.Error())
			tui.ErrorForm(errorForm, pages, errorTextView, CreateRecipeStep)
			pages.SwitchToPage(ErrorPage)
			return
		}
	})
	form.AddButton("Remove ingredient", func() {
		ingrId, err := uuid.Parse(ingredientId)
		if err != nil {
			errorTextView.SetText(err.Error())
			tui.ErrorForm(errorForm, pages, errorTextView, CreateRecipeStep)
			pages.SwitchToPage(ErrorPage)
			return
		}

		err = tui.app.ingredientService.Unlink(context.Background(), recipe.ID, ingrId)
		if err != nil {
			errorTextView.SetText(err.Error())
			tui.ErrorForm(errorForm, pages, errorTextView, CreateRecipeStep)
			pages.SwitchToPage(ErrorPage)
			return
		}
	})

	form.AddButton("Measurements", func() {
		recipeList = tui.appendMeasurementsToList(recipeList)
		recipeList.AddItem("Back", "", 'b', func() {
			pages.SwitchToPage(CreateRecipeStep)
		})
		pages.SwitchToPage(MeasurementsPage)
	})

	form.AddButton("Ingredients", func() {
		recipeList = tui.appendIngredientsToList(uuid.Nil, recipeList)
		recipeList.AddItem("Back", "", 'b', func() {
			pages.SwitchToPage(CreateRecipeStep)
		})
		pages.SwitchToPage(IngredientsPage)
	})

	form.AddButton("Types", func() {
		recipeList = tui.appendSaladTypesToList(uuid.Nil, recipeList)
		recipeList.AddItem("Back", "", 'b', func() {
			pages.SwitchToPage(CreateRecipeStep)
		})
		pages.SwitchToPage(SaladTypesPage)
	})

	form.AddButton("Back", func() {
		form.Clear(true)
		tui.BackToMenu(pages)
	})

	return form
}

func (tui *Tui) appendIngredientsToList(recipeId uuid.UUID, list *tview.List) *tview.List {
	list.Clear()
	var ingredients []*domain.Ingredient
	page := 1

	if recipeId == uuid.Nil {
		ingredients, _, _ = tui.app.ingredientService.GetAll(context.Background(), page)
	} else {
		ingredients, _ = tui.app.ingredientService.GetAllByRecipeId(context.Background(), recipeId)
	}
	for _, ingredient := range ingredients {
		list.AddItem(ingredient.ID.String(),
			fmt.Sprintf("name: %s, calories: %d", ingredient.Name, ingredient.Calories),
			'*',
			nil)
	}
	return list
}

func (tui *Tui) appendSaladTypesToList(saladId uuid.UUID, list *tview.List) *tview.List {
	list.Clear()
	var saladTypes []*domain.SaladType
	page := 1

	if saladId == uuid.Nil {
		saladTypes, _, _ = tui.app.saladTypeService.GetAll(context.Background(), page)
	} else {
		saladTypes, _ = tui.app.saladTypeService.GetAllBySaladId(context.Background(), saladId)
	}
	for _, saladType := range saladTypes {
		list.AddItem(saladType.ID.String(),
			fmt.Sprintf("Name: %s, description: %s", saladType.Name, saladType.Description),
			'*',
			nil)
	}
	return list
}

func (tui *Tui) appendMeasurementsToList(list *tview.List) *tview.List {
	list.Clear()
	measurements, _ := tui.app.measurementService.GetAll(context.Background())
	for _, unit := range measurements {
		list.AddItem(unit.ID.String(),
			fmt.Sprintf("Name: %s, grams: %d", unit.Name, unit.Grams),
			'*',
			nil)
	}
	return list
}

func (tui *Tui) ShowSaladsToModerate(form *tview.Form, pages *tview.Pages, list *tview.List) *tview.Form {
	page := 1
	errorTextView := tview.NewTextView().
		SetDynamicColors(true).
		SetRegions(true).
		SetWrap(true).
		SetTextAlign(tview.AlignCenter).
		SetText("")

	filter := new(dto.RecipeFilter)
	filter.Status = dto.ModerationSaladStatus

	salads, _, err := tui.app.saladService.GetAll(context.Background(), filter, page)
	if err != nil {
		errorTextView.SetText(err.Error())
		tui.ErrorForm(errorForm, pages, errorTextView, SaladsRequestPage)
		pages.SwitchToPage(ErrorPage)
		return form
	}

	tui.appendSaladsToList(list, pages, salads)
	list.AddItem("Back", "", 'b', func() {
		form.Clear(true)
		tui.BackToMenu(pages)
	})
	pages.SwitchToPage(SaladsPage)

	return form
}

func (tui *Tui) CreateAuthorizedMenu(form *tview.Form, pages *tview.Pages, exitFunc *tview.Application) *tview.List {
	return tview.NewList().
		AddItem("LogOut", "", '1', func() {
			form.Clear(true)
			//tui.userInfo.Username = ""
			tui.userInfo.ID = ""
			tui.userInfo.Role = UnauthorizedUser
			pages.SwitchToPage(GuestPage)
		}).
		AddItem("View salads", "", '2', func() {
			form.Clear(true)
			form.Clear(true)
			tui.ShowSalads(form, pages, list)
			pages.SwitchToPage(SaladsRequestPage)
		}).
		AddItem("Create salad", "", '4', func() {
			form.Clear(true)
			tui.createSaladForm(form, pages)
			pages.SwitchToPage(CreateSaladPage)
		}).
		AddItem("Show rated", "", '5', func() {
			form.Clear(true)
			tui.showRated(form, pages, list)
			pages.SwitchToPage(SaladsPage)
		}).
		AddItem("Exit", "", '0', func() {
			exitFunc.Stop()
		})
}

func (tui *Tui) CreateAdminMenu(form *tview.Form, pages *tview.Pages, exitFunc *tview.Application) *tview.List {
	return tview.NewList().
		AddItem("LogOut", "", '1', func() {
			form.Clear(true)
			//tui.userInfo.Username = ""
			tui.userInfo.ID = ""
			tui.userInfo.Role = UnauthorizedUser
			pages.SwitchToPage(GuestPage)
		}).
		AddItem("View salads", "", '2', func() {
			form.Clear(true)
			form.Clear(true)
			tui.ShowSalads(form, pages, list)
			pages.SwitchToPage(SaladsRequestPage)
		}).
		AddItem("Create salad", "", '3', func() {
			form.Clear(true)
			tui.createSaladForm(form, pages)
			pages.SwitchToPage(CreateSaladPage)
		}).
		AddItem("Show rated", "", '4', func() {
			form.Clear(true)
			tui.showRated(form, pages, list)
			pages.SwitchToPage(SaladsPage)
		}).
		AddItem("Show salads to moderate", "", '5', func() {
			form.Clear(true)
			tui.ShowSaladsToModerate(form, pages, list)
			pages.SwitchToPage(SaladsPage)
		}).
		AddItem("Exit", "", '0', func() {
			exitFunc.Stop()
		})
}

// TODO: automatically calculate rating of the recipe after adding comment (trigger?)
