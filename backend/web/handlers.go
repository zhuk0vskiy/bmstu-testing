package web

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"net/http"
	"net/mail"
	"ppo/domain"
	"ppo/internal/app"
	"ppo/pkg/base"
	"ppo/services/dto"
	"strconv"
	"time"
)

func LoginHandler(app *app.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		prompt := "аутентификация"

		type Req struct {
			Login    string `json:"login"`
			Password string `json:"password"`
		}
		var req Req

		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			errorResponse(w, fmt.Errorf("%s: %w", prompt, err).Error(), http.StatusBadRequest)
			return
		}

		ua := &domain.UserAuth{Username: req.Login, Password: req.Password}
		token, err := app.AuthService.Login(r.Context(), ua)
		if err != nil {
			errorResponse(w, fmt.Errorf("%s: %w", prompt, err).Error(), http.StatusUnauthorized)
			return
		}

		_, err = base.VerifyAuthToken(token, app.Config.Jwt.Key)
		if err != nil {
			errorResponse(w, fmt.Errorf("%s: проверка JWT-токена: %w", prompt, err).Error(), http.StatusInternalServerError)
			return
		}

		cookie := http.Cookie{
			Name:    "access_token",
			Value:   token,
			Path:    "/",
			Secure:  true,
			Expires: time.Now().Add(3600 * 24 * time.Second),
		}

		http.SetCookie(w, &cookie)
		successResponse(w, http.StatusOK, map[string]string{"token": token})
	}
}

func RegisterHandler(app *app.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		prompt := "регистрация"

		type Req struct {
			Name     string `json:"name"`
			Username string `json:"username"`
			Password string `json:"password"`
			Email    string `json:"email"`
		}
		var req Req

		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			errorResponse(w, fmt.Errorf("%s: %w", prompt, err).Error(), http.StatusBadRequest)
			return
		}

		var mailAddr mail.Address
		mailAddr.Address = req.Email

		ua := &domain.User{
			Name:     req.Name,
			Username: req.Username,
			Password: req.Password,
			Email:    mailAddr,
		}
		token, err := app.AuthService.Register(r.Context(), ua)
		if err != nil {
			errorResponse(w, fmt.Errorf("%s: %w", prompt, err).Error(), http.StatusBadRequest)
			return
		}

		_, err = base.VerifyAuthToken(token, app.Config.Jwt.Key)
		if err != nil {
			errorResponse(w, fmt.Errorf("%s: проверка JWT-токена: %w", prompt, err).Error(), http.StatusInternalServerError)
			return
		}

		cookie := http.Cookie{
			Name:    "access_token",
			Value:   token,
			Path:    "/",
			Secure:  true,
			Expires: time.Now().Add(3600 * 24 * time.Second),
		}

		http.SetCookie(w, &cookie)

		successResponse(w, http.StatusOK, map[string]string{"token": token})
	}
}

func GetUser(app *app.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		prompt := "получение информации о пользователе"

		id := chi.URLParam(r, "id")
		if id == "" {
			errorResponse(w, fmt.Errorf("%s: пустой id", prompt).Error(), http.StatusBadRequest)
			return
		}
		idUuid, err := uuid.Parse(id)
		if err != nil {
			errorResponse(w, fmt.Errorf("%s: преобразование id к uuid: %w", prompt, err).Error(), http.StatusBadRequest)
			return
		}

		user, err := app.UserService.GetById(r.Context(), idUuid)
		if err != nil {
			errorResponse(w, fmt.Errorf("%s: %w", prompt, err).Error(), http.StatusInternalServerError)
			return
		}

		successResponse(w, http.StatusOK, map[string]interface{}{"user": toUserTransport(user)})
	}
}

func GetSalads(app *app.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		prompt := "получение списка салатов"
		r.ParseForm()

		page := r.URL.Query().Get("page")
		if page == "" {
			errorResponse(w, fmt.Errorf("%s: пустой номер страницы", prompt).Error(), http.StatusBadRequest)
			return
		}
		pageInt, err := strconv.Atoi(page)
		if err != nil {
			errorResponse(w, fmt.Errorf("%s: преобразование номера страницы к int: %w", prompt, err).Error(), http.StatusBadRequest)
			return
		}

		minRate := r.URL.Query().Get("minRate")
		if minRate == "" {
			minRate = "0.0"
		}
		minRateFloat, err := strconv.ParseFloat(minRate, 64)
		if err != nil {
			errorResponse(w, fmt.Errorf("%s: преобразование минимального рейтинга к float: %w", prompt, err).Error(), http.StatusBadRequest)
			return
		}

		ingredients := r.Form["ingredients"]
		if len(ingredients[0]) == 0 {
			ingredients = make([]string, 0)
		}
		ingredientUuids := make([]uuid.UUID, len(ingredients))
		for i := 0; i < len(ingredients); i++ {
			ingredientUuids[i], err = uuid.Parse(ingredients[i])
			if err != nil {
				errorResponse(w, fmt.Errorf("%s: преобразование id ингредиента к uuid: %w", prompt, err).Error(), http.StatusBadRequest)
				return
			}
		}

		saladTypes := r.Form["types"]
		if len(saladTypes[0]) == 0 {
			saladTypes = make([]string, 0)
		}
		saladUuids := make([]uuid.UUID, len(saladTypes))
		for i := 0; i < len(saladTypes); i++ {
			saladUuids[i], err = uuid.Parse(saladTypes[i])
			if err != nil {
				errorResponse(w, fmt.Errorf("%s: преобразование id типа салата к uuid: %w", prompt, err).Error(), http.StatusBadRequest)
				return
			}
		}

		filter := new(dto.RecipeFilter)

		filter.MinRate = minRateFloat
		filter.Status = dto.PublishedSaladStatus
		filter.SaladTypes = saladUuids
		filter.AvailableIngredients = ingredientUuids

		salads, numPages, err := app.SaladService.GetAll(r.Context(), filter, pageInt)
		if err != nil {
			errorResponse(w, fmt.Errorf("%s: %w", prompt, err).Error(), http.StatusInternalServerError)
			return
		}

		saladsTransport := make([]Salad, len(salads))
		for i, salad := range salads {
			saladsTransport[i] = toSaladTransport(salad)
		}

		successResponse(w, http.StatusOK, map[string]interface{}{"num_pages": numPages, "salads": saladsTransport})
	}
}

func GetSaladsWithStatus(app *app.App) http.HandlerFunc { // FIXME
	return func(w http.ResponseWriter, r *http.Request) {
		prompt := "получение списка салатов (по статусу)"
		r.ParseForm()

		page := r.URL.Query().Get("page")
		if page == "" {
			errorResponse(w, fmt.Errorf("%s: пустой номер страницы", prompt).Error(), http.StatusBadRequest)
			return
		}
		pageInt, err := strconv.Atoi(page)
		if err != nil {
			errorResponse(w, fmt.Errorf("%s: преобразование номера страницы к int: %w", prompt, err).Error(), http.StatusBadRequest)
			return
		}

		minRate := r.URL.Query().Get("minRate")
		if minRate == "" {
			minRate = "0.0"
		}
		minRateFloat, err := strconv.ParseFloat(minRate, 64)
		if err != nil {
			errorResponse(w, fmt.Errorf("%s: преобразование минимального рейтинга к float: %w", prompt, err).Error(), http.StatusBadRequest)
			return
		}

		ingredients := r.Form["ingredients"]
		if len(ingredients[0]) == 0 {
			ingredients = make([]string, 0)
		}
		ingredientUuids := make([]uuid.UUID, len(ingredients))
		for i := 0; i < len(ingredients); i++ {
			ingredientUuids[i], err = uuid.Parse(ingredients[i])
			if err != nil {
				errorResponse(w, fmt.Errorf("%s: преобразование id ингредиента к uuid: %w", prompt, err).Error(), http.StatusBadRequest)
				return
			}
		}

		saladTypes := r.Form["types"]
		if len(saladTypes[0]) == 0 {
			saladTypes = make([]string, 0)
		}
		saladUuids := make([]uuid.UUID, len(saladTypes))
		for i := 0; i < len(saladTypes); i++ {
			saladUuids[i], err = uuid.Parse(saladTypes[i])
			if err != nil {
				errorResponse(w, fmt.Errorf("%s: преобразование id типа салата к uuid: %w", prompt, err).Error(), http.StatusBadRequest)
				return
			}
		}

		filter := new(dto.RecipeFilter)

		status := r.URL.Query().Get("status")
		if page == "" {
			errorResponse(w, fmt.Errorf("%s: пустой номер статуса", prompt).Error(), http.StatusBadRequest)
			return
		}
		statusInt, err := strconv.Atoi(status)
		if err != nil {
			errorResponse(w, fmt.Errorf("%s: преобразование номера статуса к int: %w", prompt, err).Error(), http.StatusBadRequest)
			return
		}
		filter.Status = statusInt
		filter.MinRate = minRateFloat

		//filter.Status = dto.PublishedSaladStatus
		filter.SaladTypes = saladUuids
		filter.AvailableIngredients = ingredientUuids

		salads, numPages, err := app.SaladService.GetAll(r.Context(), filter, pageInt)
		if err != nil {
			errorResponse(w, fmt.Errorf("%s: %w", prompt, err).Error(), http.StatusInternalServerError)
			return
		}

		saladsTransport := make([]Salad, len(salads))
		for i, salad := range salads {
			saladsTransport[i] = toSaladTransport(salad)
		}

		successResponse(w, http.StatusOK, map[string]interface{}{"num_pages": numPages, "salads": saladsTransport})
	}
}

func GetUserSalads(app *app.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		prompt := "получение списка салатов"
		r.ParseForm()

		userId, err := getStringClaimFromJWT(r.Context(), "user_id")
		if err != nil {
			errorResponse(w, fmt.Errorf("получение id авторизованного пользователя: %w", err).Error(), http.StatusBadRequest)
			return
		}
		userUuid, err := uuid.Parse(userId)
		if err != nil {
			errorResponse(w, fmt.Errorf("преобразование id пользователя к uuid: %w", err).Error(), http.StatusInternalServerError)
			return
		}

		page := r.URL.Query().Get("page")
		if page == "" {
			errorResponse(w, fmt.Errorf("%s: пустой номер страницы", prompt).Error(), http.StatusBadRequest)
			return
		}
		pageInt, err := strconv.Atoi(page)
		if err != nil {
			errorResponse(w, fmt.Errorf("%s: преобразование номера страницы к int: %w", prompt, err).Error(), http.StatusBadRequest)
			return
		}

		ingredients := r.Form["ingredients"]
		if len(ingredients[0]) == 0 {
			ingredients = make([]string, 0)
		}
		ingredientUuids := make([]uuid.UUID, len(ingredients))
		for i := 0; i < len(ingredients); i++ {
			ingredientUuids[i], err = uuid.Parse(ingredients[i])
			if err != nil {
				errorResponse(w, fmt.Errorf("%s: преобразование id ингредиента к uuid: %w", prompt, err).Error(), http.StatusBadRequest)
				return
			}
		}

		saladTypes := r.Form["types"]
		if len(saladTypes[0]) == 0 {
			saladTypes = make([]string, 0)
		}
		saladUuids := make([]uuid.UUID, len(saladTypes))
		for i := 0; i < len(saladTypes); i++ {
			saladUuids[i], err = uuid.Parse(saladTypes[i])
			if err != nil {
				errorResponse(w, fmt.Errorf("%s: преобразование id типа салата к uuid: %w", prompt, err).Error(), http.StatusBadRequest)
				return
			}
		}

		salads, err := app.SaladService.GetAllByUserId(r.Context(), userUuid)
		if err != nil {
			errorResponse(w, fmt.Errorf("%s: %w", prompt, err).Error(), http.StatusInternalServerError)
			return
		}

		saladsTransport := make([]Salad, len(salads))
		for i, salad := range salads {
			saladsTransport[i] = toSaladTransport(salad)
		}

		successResponse(w, http.StatusOK, map[string]interface{}{"num_pages": pageInt, "salads": saladsTransport})
	}
}

func GetUserRatedSalads(app *app.App) http.HandlerFunc { // FIXME
	return func(w http.ResponseWriter, r *http.Request) {
		prompt := "получение списка салатов, оцененных пользователем"

		userId, err := getStringClaimFromJWT(r.Context(), "user_id")
		if err != nil {
			errorResponse(w, fmt.Errorf("получение id авторизованного пользователя: %w", err).Error(), http.StatusBadRequest)
			return
		}
		userUuid, err := uuid.Parse(userId)
		if err != nil {
			errorResponse(w, fmt.Errorf("преобразование id пользователя к uuid: %w", err).Error(), http.StatusInternalServerError)
			return
		}

		page := r.URL.Query().Get("page")
		if page == "" {
			errorResponse(w, fmt.Errorf("%s: пустой номер страницы", prompt).Error(), http.StatusBadRequest)
			return
		}
		pageInt, err := strconv.Atoi(page)
		if err != nil {
			errorResponse(w, fmt.Errorf("%s: преобразование номера страницы к int: %w", prompt, err).Error(), http.StatusBadRequest)
			return
		}

		salads, numPages, err := app.SaladService.GetAllRatedByUser(r.Context(), userUuid, pageInt)
		if err != nil {
			errorResponse(w, fmt.Errorf("%s: %w", prompt, err).Error(), http.StatusInternalServerError)
			return
		}

		saladsTransport := make([]Salad, len(salads))
		for i, salad := range salads {
			saladsTransport[i] = toSaladTransport(salad)
		}

		successResponse(w, http.StatusOK, map[string]interface{}{"num_pages": numPages, "salads": saladsTransport})
	}
}

func GetSaladById(app *app.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		prompt := "получение салата по id"

		id := chi.URLParam(r, "id")
		if id == "" {
			errorResponse(w, fmt.Errorf("%s: пустой id", prompt).Error(), http.StatusBadRequest)
			return
		}
		idUuid, err := uuid.Parse(id)
		if err != nil {
			errorResponse(w, fmt.Errorf("%s: %w", prompt, err).Error(), http.StatusBadRequest)
			return
		}

		salad, err := app.SaladService.GetById(r.Context(), idUuid)
		if err != nil {
			errorResponse(w, fmt.Errorf("%s: %w", prompt, err).Error(), http.StatusInternalServerError)
			return
		}

		successResponse(w, http.StatusOK, map[string]interface{}{"salad": toSaladTransport(salad)})
	}
}

func GetSaladRating(app *app.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		prompt := "получение рейтинга салата"

		id := chi.URLParam(r, "id")
		if id == "" {
			errorResponse(w, fmt.Errorf("%s: пустой id", prompt).Error(), http.StatusBadRequest)
			return
		}
		idUuid, err := uuid.Parse(id)
		if err != nil {
			errorResponse(w, fmt.Errorf("%s: %w", prompt, err).Error(), http.StatusBadRequest)
			return
		}

		recipe, err := app.RecipeService.GetBySaladId(r.Context(), idUuid)
		if err != nil {
			errorResponse(w, fmt.Errorf("%s: %w", prompt, err).Error(), http.StatusInternalServerError)
			return
		}

		successResponse(w, http.StatusOK, map[string]float32{"rating": recipe.Rating})
	}
}

func GetSaladRecipe(app *app.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		prompt := "получение рецепта салата"

		id := chi.URLParam(r, "id")
		if id == "" {
			errorResponse(w, fmt.Errorf("%s: пустой id", prompt).Error(), http.StatusBadRequest)
			return
		}
		idUuid, err := uuid.Parse(id)
		if err != nil {
			errorResponse(w, fmt.Errorf("%s: %w", prompt, err).Error(), http.StatusBadRequest)
			return
		}

		recipe, err := app.RecipeService.GetBySaladId(r.Context(), idUuid)
		if err != nil {
			errorResponse(w, fmt.Errorf("%s: %w", prompt, err).Error(), http.StatusInternalServerError)
			return
		}

		successResponse(w, http.StatusOK, map[string]interface{}{
			"recipe": toRecipeTransport(recipe),
		})
	}
}

func GetRecipeSteps(app *app.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		prompt := "получение рецепта салата"

		id := chi.URLParam(r, "id")
		if id == "" {
			errorResponse(w, fmt.Errorf("%s: пустой id", prompt).Error(), http.StatusBadRequest)
			return
		}
		idUuid, err := uuid.Parse(id)
		if err != nil {
			errorResponse(w, fmt.Errorf("%s: %w", prompt, err).Error(), http.StatusBadRequest)
			return
		}

		steps, err := app.RecipeStepService.GetAllByRecipeID(r.Context(), idUuid)
		if err != nil {
			errorResponse(w, fmt.Errorf("%s: %w", prompt, err).Error(), http.StatusInternalServerError)
			return
		}
		stepsTransport := make([]RecipeStep, len(steps))
		for i, step := range steps {
			stepsTransport[i] = toRecipeStepTransport(step)
		}

		successResponse(w, http.StatusOK, map[string]interface{}{
			"steps": stepsTransport,
		})
	}
}

func GetRecipeIngredients(app *app.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		prompt := "получение ингредиентов рецепта"

		id := chi.URLParam(r, "id")
		if id == "" {
			errorResponse(w, fmt.Errorf("%s: пустой id", prompt).Error(), http.StatusBadRequest)
			return
		}
		idUuid, err := uuid.Parse(id)
		if err != nil {
			errorResponse(w, fmt.Errorf("%s: %w", prompt, err).Error(), http.StatusBadRequest)
			return
		}

		ingredients, err := app.IngredientService.GetAllByRecipeId(r.Context(), idUuid)
		if err != nil {
			errorResponse(w, fmt.Errorf("%s: %w", prompt, err).Error(), http.StatusInternalServerError)
			return
		}
		ingredientsTransport := make([]Ingredient, len(ingredients))
		for i, ingredient := range ingredients {
			ingredientsTransport[i] = toIngredientTransport(ingredient)
		}

		successResponse(w, http.StatusOK, map[string]interface{}{
			"ingredients": ingredientsTransport,
		})
	}
}

func GetSaladTypes(app *app.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		prompt := "получение типов рецепта"

		id := chi.URLParam(r, "id")
		if id == "" {
			errorResponse(w, fmt.Errorf("%s: пустой id", prompt).Error(), http.StatusBadRequest)
			return
		}
		idUuid, err := uuid.Parse(id)
		if err != nil {
			errorResponse(w, fmt.Errorf("%s: %w", prompt, err).Error(), http.StatusBadRequest)
			return
		}

		types, err := app.SaladTypeService.GetAllBySaladId(r.Context(), idUuid)
		if err != nil {
			errorResponse(w, fmt.Errorf("%s: %w", prompt, err).Error(), http.StatusInternalServerError)
			return
		}
		typesTransport := make([]SaladType, len(types))
		for i, saladType := range types {
			typesTransport[i] = toSaladTypeTransport(saladType)
		}

		successResponse(w, http.StatusOK, map[string]interface{}{
			"types": typesTransport,
		})
	}
}

func GetIngredientType(app *app.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		prompt := "получение типа ингредиента"

		id := chi.URLParam(r, "id")
		if id == "" {
			errorResponse(w, fmt.Errorf("%s: пустой id", prompt).Error(), http.StatusBadRequest)
			return
		}
		idUuid, err := uuid.Parse(id)
		if err != nil {
			errorResponse(w, fmt.Errorf("%s: %w", prompt, err).Error(), http.StatusBadRequest)
			return
		}

		ingredientType, err := app.IngredientTypeService.GetById(r.Context(), idUuid)
		if err != nil {
			errorResponse(w, fmt.Errorf("%s: %w", prompt, err).Error(), http.StatusInternalServerError)
			return
		}
		successResponse(w, http.StatusOK, map[string]interface{}{
			"ingredientType": toIngredientTypeTransport(ingredientType),
		})
	}
}

func GetMeasurementByRecipe(app *app.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		prompt := "получение количества ингредиента"

		ingredientId := r.URL.Query().Get("ingredient")
		if ingredientId == "" {
			errorResponse(w, fmt.Errorf("%s: пустой id ингредиента", prompt).Error(), http.StatusBadRequest)
			return
		}
		ingredientUuid, err := uuid.Parse(ingredientId)
		if err != nil {
			errorResponse(w, fmt.Errorf("%s: %w", prompt, err).Error(), http.StatusBadRequest)
			return
		}

		recipeId := r.URL.Query().Get("recipe")
		if ingredientId == "" {
			errorResponse(w, fmt.Errorf("%s: пустой id рецепта", prompt).Error(), http.StatusBadRequest)
			return
		}
		recipeUuid, err := uuid.Parse(recipeId)
		if err != nil {
			errorResponse(w, fmt.Errorf("%s: %w", prompt, err).Error(), http.StatusBadRequest)
			return
		}

		measurement, count, err := app.MeasurementService.GetByRecipeId(r.Context(), ingredientUuid, recipeUuid)
		if err != nil {
			errorResponse(w, fmt.Errorf("%s: %w", prompt, err).Error(), http.StatusInternalServerError)
			return
		}
		successResponse(w, http.StatusOK, map[string]interface{}{
			"measurement": toMeasurementTransport(measurement),
			"count":       count,
		})
	}
}

func GetIngredientsByPage(app *app.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		prompt := "получение списка ингредиентов"

		page := r.URL.Query().Get("page")
		if page == "" {
			errorResponse(w, fmt.Errorf("%s: пустой номер страницы", prompt).Error(), http.StatusBadRequest)
			return
		}
		pageInt, err := strconv.Atoi(page)
		if err != nil {
			errorResponse(w, fmt.Errorf("%s: преобразование номера страницы к int: %w", prompt, err).Error(), http.StatusBadRequest)
			return
		}

		ingredients, numPages, err := app.IngredientService.GetAll(r.Context(), pageInt)
		if err != nil {
			errorResponse(w, fmt.Errorf("%s: %w", prompt, err).Error(), http.StatusInternalServerError)
			return
		}

		ingredientsTransport := make([]Ingredient, len(ingredients))
		for i, ingredient := range ingredients {
			ingredientsTransport[i] = toIngredientTransport(ingredient)
		}

		successResponse(w, http.StatusOK, map[string]interface{}{"num_pages": numPages, "ingredients": ingredientsTransport})
	}
}

func GetSaladTypesByPage(app *app.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		prompt := "получение списка типов салатов"

		page := r.URL.Query().Get("page")
		if page == "" {
			errorResponse(w, fmt.Errorf("%s: пустой номер страницы", prompt).Error(), http.StatusBadRequest)
			return
		}
		pageInt, err := strconv.Atoi(page)
		if err != nil {
			errorResponse(w, fmt.Errorf("%s: преобразование номера страницы к int: %w", prompt, err).Error(), http.StatusBadRequest)
			return
		}

		saladTypes, numPages, err := app.SaladTypeService.GetAll(r.Context(), pageInt)
		if err != nil {
			errorResponse(w, fmt.Errorf("%s: %w", prompt, err).Error(), http.StatusInternalServerError)
			return
		}

		typesTransport := make([]SaladType, len(saladTypes))
		for i, saladType := range saladTypes {
			typesTransport[i] = toSaladTypeTransport(saladType)
		}

		successResponse(w, http.StatusOK, map[string]interface{}{"num_pages": numPages, "salad_types": typesTransport})
	}
}

func GetCommentsBySalad(app *app.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		prompt := "получение комментариев к салату"

		page := r.URL.Query().Get("page")
		if page == "" {
			errorResponse(w, fmt.Errorf("%s: пустой номер страницы", prompt).Error(), http.StatusBadRequest)
			return
		}
		pageInt, err := strconv.Atoi(page)
		if err != nil {
			errorResponse(w, fmt.Errorf("%s: преобразование номера страницы к int: %w", prompt, err).Error(), http.StatusBadRequest)
			return
		}

		saladId := r.URL.Query().Get("salad")
		if saladId == "" {
			errorResponse(w, fmt.Errorf("%s: пустой id салата", prompt).Error(), http.StatusBadRequest)
			return
		}
		saladUuid, err := uuid.Parse(saladId)
		if err != nil {
			errorResponse(w, fmt.Errorf("%s: преобразование id салата к uuid: %w", prompt, err).Error(), http.StatusBadRequest)
			return
		}

		comments, numPages, err := app.CommentService.GetAllBySaladID(r.Context(), saladUuid, pageInt)
		if err != nil {
			errorResponse(w, fmt.Errorf("%s: %w", prompt, err).Error(), http.StatusInternalServerError)
			return
		}

		commentsTransport := make([]Comment, len(comments))
		for i, comment := range comments {
			commentsTransport[i] = toCommentTransport(comment)
		}

		successResponse(w, http.StatusOK, map[string]interface{}{"num_pages": numPages, "comments": commentsTransport})
	}
}

func GetUserComment(app *app.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		prompt := "получение комментария пользователя к салату"

		userId := r.URL.Query().Get("user")
		if userId == "" {
			errorResponse(w, fmt.Errorf("%s: пустой id пользователя", prompt).Error(), http.StatusBadRequest)
			return
		}
		userUuid, err := uuid.Parse(userId)
		if err != nil {
			errorResponse(w, fmt.Errorf("%s: преобразование id пользователя к uuid: %w", prompt, err).Error(), http.StatusBadRequest)
			return
		}

		saladId := r.URL.Query().Get("salad")
		if saladId == "" {
			errorResponse(w, fmt.Errorf("%s: пустой id салата", prompt).Error(), http.StatusBadRequest)
			return
		}
		saladUuid, err := uuid.Parse(saladId)
		if err != nil {
			errorResponse(w, fmt.Errorf("%s: преобразование id салата к uuid: %w", prompt, err).Error(), http.StatusBadRequest)
			return
		}

		comment, err := app.CommentService.GetBySaladAndUser(r.Context(), saladUuid, userUuid)
		if err != nil {
			errorResponse(w, fmt.Errorf("%s: %w", prompt, err).Error(), http.StatusInternalServerError)
			return
		}

		successResponse(w, http.StatusOK, map[string]interface{}{"comment": toCommentTransport(comment)})
	}
}

func DeleteComment(app *app.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		prompt := "удаление комментария"

		userId, err := getStringClaimFromJWT(r.Context(), "user_id")
		if err != nil {
			errorResponse(w, fmt.Errorf("получение id авторизованного пользователя: %w", err).Error(), http.StatusBadRequest)
			return
		}
		userUuid, err := uuid.Parse(userId)
		if err != nil {
			errorResponse(w, fmt.Errorf("преобразование id пользователя к uuid: %w", err).Error(), http.StatusInternalServerError)
			return
		}

		id := chi.URLParam(r, "id")
		if id == "" {
			app.Logger.Infof("getting id")
			errorResponse(w, fmt.Errorf("%s: пустой id", prompt).Error(), http.StatusBadRequest)
			return
		}
		idUuid, err := uuid.Parse(id)
		if err != nil {
			errorResponse(w, fmt.Errorf("%s: %w", prompt, err).Error(), http.StatusBadRequest)
			return
		}

		comment, err := app.CommentService.GetById(r.Context(), idUuid)
		if err != nil {
			errorResponse(w, fmt.Errorf("%s: %w", prompt, err).Error(), http.StatusInternalServerError)
			return
		}
		if comment.AuthorID != userUuid { // TODO: mb admin can delete every comment
			errorResponse(w, fmt.Errorf("%s: только автор комментария может удалить его", prompt).Error(), http.StatusBadRequest)
			return
		}

		err = app.CommentService.DeleteById(r.Context(), idUuid)
		if err != nil {
			errorResponse(w, fmt.Errorf("%s: %w", prompt, err).Error(), http.StatusInternalServerError)
			return
		}

		successResponse(w, http.StatusOK, nil)
	}
}

func GetCommentById(app *app.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		prompt := "получение комментария по id"

		id := chi.URLParam(r, "id")
		if id == "" {
			errorResponse(w, fmt.Errorf("%s: пустой id", prompt).Error(), http.StatusBadRequest)
			return
		}
		idUuid, err := uuid.Parse(id)
		if err != nil {
			errorResponse(w, fmt.Errorf("%s: %w", prompt, err).Error(), http.StatusBadRequest)
			return
		}

		comment, err := app.CommentService.GetById(r.Context(), idUuid)
		if err != nil {
			errorResponse(w, fmt.Errorf("%s: %w", prompt, err).Error(), http.StatusInternalServerError)
			return
		}
		successResponse(w, http.StatusOK, map[string]interface{}{
			"comment": toCommentTransport(comment),
		})
	}
}

func UpdateComment(app *app.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		prompt := "обновление комментария"

		userId, err := getStringClaimFromJWT(r.Context(), "user_id")
		if err != nil {
			errorResponse(w, fmt.Errorf("получение id авторизованного пользователя: %w", err).Error(), http.StatusBadRequest)
			return
		}
		userUuid, err := uuid.Parse(userId)
		if err != nil {
			errorResponse(w, fmt.Errorf("преобразование id пользователя к uuid: %w", err).Error(), http.StatusInternalServerError)
			return
		}

		id := chi.URLParam(r, "id")
		if id == "" {
			errorResponse(w, fmt.Errorf("%s: пустой id комментария", prompt).Error(), http.StatusBadRequest)
			return
		}
		idUuid, err := uuid.Parse(id)
		if err != nil {
			errorResponse(w, fmt.Errorf("%s: преобразование id комментария к uuid: %w", prompt, err).Error(), http.StatusBadRequest)
			return
		}

		commentDb, err := app.CommentService.GetById(r.Context(), idUuid)
		if err != nil {
			errorResponse(w, fmt.Errorf("%s: %w", prompt, err).Error(), http.StatusInternalServerError)
			return
		}
		if commentDb.AuthorID != userUuid {
			errorResponse(w, fmt.Errorf("%s: только автор комментария может изменить его", prompt).Error(), http.StatusBadRequest)
			return
		}

		var req Comment
		err = json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			errorResponse(w, err.Error(), http.StatusBadRequest)
			return
		}
		if req.Text != "" {
			commentDb.Text = req.Text
		}
		if req.Rating != 0 {
			commentDb.Rating = req.Rating
		}

		err = app.CommentService.Update(r.Context(), commentDb)
		if err != nil {
			errorResponse(w, fmt.Errorf("%s: %w", prompt, err).Error(), http.StatusInternalServerError)
			return
		}

		successResponse(w, http.StatusOK, nil)
	}
}

func CreateComment(app *app.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		prompt := "создание комментария"

		userId, err := getStringClaimFromJWT(r.Context(), "user_id")
		if err != nil {
			errorResponse(w, fmt.Errorf("%s: получение id авторизованного пользователя: %w", prompt, err).Error(), http.StatusBadRequest)
			return
		}
		userUuid, err := uuid.Parse(userId)
		if err != nil {
			errorResponse(w, fmt.Errorf("%s: преобразование id пользователя к uuid: %w", prompt, err).Error(), http.StatusInternalServerError)
			return
		}

		var req Comment
		err = json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			errorResponse(w, err.Error(), http.StatusBadRequest)
			return
		}

		comment := toCommentModel(&req)
		comment.AuthorID = userUuid

		err = app.CommentService.Create(r.Context(), comment)
		if err != nil {
			errorResponse(w, fmt.Errorf("%s: %w", prompt, err).Error(), http.StatusInternalServerError)
			return
		}

		successResponse(w, http.StatusOK, nil)
	}
}

func CreateSalad(app *app.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		prompt := "создание салата"

		userId, err := getStringClaimFromJWT(r.Context(), "user_id")
		if err != nil {
			errorResponse(w, fmt.Errorf("%s: получение id авторизованного пользователя: %w", prompt, err).Error(), http.StatusBadRequest)
			return
		}
		userUuid, err := uuid.Parse(userId)
		if err != nil {
			errorResponse(w, fmt.Errorf("%s: преобразование id пользователя к uuid: %w", prompt, err).Error(), http.StatusInternalServerError)
			return
		}

		var req Salad
		err = json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			errorResponse(w, err.Error(), http.StatusBadRequest)
			return
		}

		salad := toSaladModel(&req)
		salad.AuthorID = userUuid

		saladUuid, err := app.SaladService.Create(r.Context(), salad)
		if err != nil {
			errorResponse(w, fmt.Errorf("%s: %w", prompt, err).Error(), http.StatusInternalServerError)
			return
		}
		salad.ID = saladUuid

		successResponse(w, http.StatusOK, map[string]interface{}{
			"salad": toSaladTransport(salad),
		})
	}
}

func UpdateSalad(app *app.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		prompt := "обновление салата"

		userId, err := getStringClaimFromJWT(r.Context(), "user_id")
		if err != nil {
			errorResponse(w, fmt.Errorf("получение id авторизованного пользователя: %w", err).Error(), http.StatusBadRequest)
			return
		}
		userUuid, err := uuid.Parse(userId)
		if err != nil {
			errorResponse(w, fmt.Errorf("преобразование id пользователя к uuid: %w", err).Error(), http.StatusInternalServerError)
			return
		}

		id := chi.URLParam(r, "id")
		if id == "" {
			errorResponse(w, fmt.Errorf("%s: пустой id комментария", prompt).Error(), http.StatusBadRequest)
			return
		}
		idUuid, err := uuid.Parse(id)
		if err != nil {
			errorResponse(w, fmt.Errorf("%s: преобразование id комментария к uuid: %w", prompt, err).Error(), http.StatusBadRequest)
			return
		}

		saladDb, err := app.SaladService.GetById(r.Context(), idUuid)
		if err != nil {
			errorResponse(w, fmt.Errorf("%s: %w", prompt, err).Error(), http.StatusInternalServerError)
			return
		}
		if saladDb.AuthorID != userUuid { // TODO: mb admin can change every salad?
			errorResponse(w, fmt.Errorf("%s: только автор салата может изменить его", prompt).Error(), http.StatusBadRequest)
			return
		}

		var req Salad
		err = json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			errorResponse(w, err.Error(), http.StatusBadRequest)
			return
		}
		if req.Name != "" {
			saladDb.Name = req.Name
		}
		if req.Description != "" {
			saladDb.Description = req.Description
		}

		err = app.SaladService.Update(r.Context(), saladDb)
		if err != nil {
			errorResponse(w, fmt.Errorf("%s: %w", prompt, err).Error(), http.StatusInternalServerError)
			return
		}

		successResponse(w, http.StatusOK, map[string]interface{}{
			"salad": toSaladTransport(saladDb),
		})
	}
}

func DeleteSalad(app *app.App) http.HandlerFunc { // FIXME
	return func(w http.ResponseWriter, r *http.Request) {
		prompt := "удаление салата"

		userId, err := getStringClaimFromJWT(r.Context(), "user_id")
		if err != nil {
			errorResponse(w, fmt.Errorf("получение id авторизованного пользователя: %w", err).Error(), http.StatusBadRequest)
			return
		}
		userUuid, err := uuid.Parse(userId)
		if err != nil {
			errorResponse(w, fmt.Errorf("преобразование id пользователя к uuid: %w", err).Error(), http.StatusInternalServerError)
			return
		}

		userRole, err := getStringClaimFromJWT(r.Context(), "role")
		if err != nil {
			errorResponse(w, fmt.Errorf("получение роли авторизованного пользователя: %w", err).Error(), http.StatusBadRequest)
			return
		}

		id := chi.URLParam(r, "id")
		if id == "" {
			errorResponse(w, fmt.Errorf("%s: пустой id комментария", prompt).Error(), http.StatusBadRequest)
			return
		}
		idUuid, err := uuid.Parse(id)
		if err != nil {
			errorResponse(w, fmt.Errorf("%s: преобразование id комментария к uuid: %w", prompt, err).Error(), http.StatusBadRequest)
			return
		}

		saladDb, err := app.SaladService.GetById(r.Context(), idUuid)
		if err != nil {
			errorResponse(w, fmt.Errorf("%s: %w", prompt, err).Error(), http.StatusInternalServerError)
			return
		}
		if saladDb.AuthorID != userUuid && userRole != "admin" { // TODO: mb admin can change every salad?
			errorResponse(w, fmt.Errorf("%s: только автор салата может изменить его", prompt).Error(), http.StatusBadRequest)
			return
		}

		err = app.SaladService.DeleteById(r.Context(), idUuid)
		if err != nil {
			errorResponse(w, fmt.Errorf("%s: %w", prompt, err).Error(), http.StatusInternalServerError)
			return
		}

		successResponse(w, http.StatusOK, nil)
	}
}

func CreateRecipe(app *app.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		prompt := "создание рецепта"

		userId, err := getStringClaimFromJWT(r.Context(), "user_id")
		if err != nil {
			errorResponse(w, fmt.Errorf("%s: получение id авторизованного пользователя: %w", prompt, err).Error(), http.StatusBadRequest)
			return
		}
		userUuid, err := uuid.Parse(userId)
		if err != nil {
			errorResponse(w, fmt.Errorf("%s: преобразование id пользователя к uuid: %w", prompt, err).Error(), http.StatusInternalServerError)
			return
		}

		var req Recipe
		err = json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			errorResponse(w, err.Error(), http.StatusBadRequest)
			return
		}

		recipe := toRecipeModel(&req)
		recipe.Status = dto.EditingSaladStatus

		saladDb, err := app.SaladService.GetById(r.Context(), recipe.SaladID)
		if err != nil {
			errorResponse(w, fmt.Errorf("%s: %w", prompt, err).Error(), http.StatusInternalServerError)
			return
		}
		if saladDb.AuthorID != userUuid {
			errorResponse(w, fmt.Errorf("%s: только автор салата может создать рецепт", prompt).Error(), http.StatusBadRequest)
			return
		}

		recipeUuid, err := app.RecipeService.Create(r.Context(), recipe)
		if err != nil {
			errorResponse(w, fmt.Errorf("%s: %w", prompt, err).Error(), http.StatusInternalServerError)
			return
		}
		recipe.ID = recipeUuid

		successResponse(w, http.StatusOK, map[string]interface{}{
			"recipe": toRecipeTransport(recipe),
		})
	}
}

func UpdateRecipe(app *app.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		prompt := "обновление рецепта"

		userId, err := getStringClaimFromJWT(r.Context(), "user_id")
		if err != nil {
			errorResponse(w, fmt.Errorf("получение id авторизованного пользователя: %w", err).Error(), http.StatusBadRequest)
			return
		}
		userUuid, err := uuid.Parse(userId)
		if err != nil {
			errorResponse(w, fmt.Errorf("преобразование id пользователя к uuid: %w", err).Error(), http.StatusInternalServerError)
			return
		}

		userRole, err := getStringClaimFromJWT(r.Context(), "role")
		if err != nil {
			errorResponse(w, fmt.Errorf("получение роли авторизованного пользователя: %w", err).Error(), http.StatusBadRequest)
			return
		}

		app.Logger.Infof("UPDATING ROLE: %s", userRole)

		id := chi.URLParam(r, "id")
		if id == "" {
			errorResponse(w, fmt.Errorf("%s: пустой id комментария", prompt).Error(), http.StatusBadRequest)
			return
		}
		idUuid, err := uuid.Parse(id)
		if err != nil {
			errorResponse(w, fmt.Errorf("%s: преобразование id комментария к uuid: %w", prompt, err).Error(), http.StatusBadRequest)
			return
		}

		recipeDb, err := app.RecipeService.GetById(r.Context(), idUuid)
		if err != nil {
			errorResponse(w, fmt.Errorf("%s: %w", prompt, err).Error(), http.StatusInternalServerError)
			return
		}
		saladDb, err := app.SaladService.GetById(r.Context(), recipeDb.SaladID)
		if err != nil {
			errorResponse(w, fmt.Errorf("%s: %w", prompt, err).Error(), http.StatusInternalServerError)
			return
		}
		if saladDb.AuthorID != userUuid && userRole != "admin" { // TODO: mb admin can change every recipe?
			errorResponse(w, fmt.Errorf("%s: только автор рецепта может изменить его", prompt).Error(), http.StatusBadRequest)
			return
		}

		var req Recipe
		err = json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			errorResponse(w, err.Error(), http.StatusBadRequest)
			return
		}
		if req.NumberOfServings != 0 {
			recipeDb.NumberOfServings = req.NumberOfServings
		}
		if req.TimeToCook != 0 {
			recipeDb.TimeToCook = req.TimeToCook
		}
		if req.Status != 0 {
			if userRole == domain.DefaultRole {
				if req.Status == dto.EditingSaladStatus ||
					req.Status == dto.ModerationSaladStatus ||
					req.Status == dto.StoredSaladStatus {
					recipeDb.Status = req.Status
				}
			} else if userRole == "admin" {
				recipeDb.Status = req.Status
			}
		}

		err = app.RecipeService.Update(r.Context(), recipeDb)
		if err != nil {
			errorResponse(w, fmt.Errorf("%s: %w", prompt, err).Error(), http.StatusInternalServerError)
			return
		}

		successResponse(w, http.StatusOK, map[string]interface{}{
			"recipe": toRecipeTransport(recipeDb),
		})
	}
}

func LinkTypeToSalad(app *app.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		prompt := "добавление типа к салату"

		userId, err := getStringClaimFromJWT(r.Context(), "user_id")
		if err != nil {
			errorResponse(w, fmt.Errorf("%s: получение id авторизованного пользователя: %w", prompt, err).Error(), http.StatusBadRequest)
			return
		}
		userUuid, err := uuid.Parse(userId)
		if err != nil {
			errorResponse(w, fmt.Errorf("%s: преобразование id пользователя к uuid: %w", prompt, err).Error(), http.StatusInternalServerError)
			return
		}

		var req LinkSaladType
		err = json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			errorResponse(w, err.Error(), http.StatusBadRequest)
			return
		}

		saladUuid, typeUuid := toLinkSaladTypeModel(&req)

		saladDb, err := app.SaladService.GetById(r.Context(), saladUuid)
		if err != nil {
			errorResponse(w, fmt.Errorf("%s: %w", prompt, err).Error(), http.StatusInternalServerError)
			return
		}
		if saladDb.AuthorID != userUuid {
			errorResponse(w, fmt.Errorf("%s: только автор салата может добавлять типы к салату", prompt).Error(), http.StatusBadRequest)
			return
		}

		err = app.SaladTypeService.Link(r.Context(), saladUuid, typeUuid)
		if err != nil {
			errorResponse(w, fmt.Errorf("%s: %w", prompt, err).Error(), http.StatusInternalServerError)
			return
		}

		successResponse(w, http.StatusOK, nil)
	}
}

func UnlinkTypeFromSalad(app *app.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		prompt := "открепление типа от салата"

		userId, err := getStringClaimFromJWT(r.Context(), "user_id")
		if err != nil {
			errorResponse(w, fmt.Errorf("%s: получение id авторизованного пользователя: %w", prompt, err).Error(), http.StatusBadRequest)
			return
		}
		userUuid, err := uuid.Parse(userId)
		if err != nil {
			errorResponse(w, fmt.Errorf("%s: преобразование id пользователя к uuid: %w", prompt, err).Error(), http.StatusInternalServerError)
			return
		}

		var req LinkSaladType
		err = json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			errorResponse(w, err.Error(), http.StatusBadRequest)
			return
		}

		saladUuid, typeUuid := toLinkSaladTypeModel(&req)

		saladDb, err := app.SaladService.GetById(r.Context(), saladUuid)
		if err != nil {
			errorResponse(w, fmt.Errorf("%s: %w", prompt, err).Error(), http.StatusInternalServerError)
			return
		}
		if saladDb.AuthorID != userUuid {
			errorResponse(w, fmt.Errorf("%s: только автор салата может изменять типы салата", prompt).Error(), http.StatusBadRequest)
			return
		}

		err = app.SaladTypeService.Unlink(r.Context(), saladUuid, typeUuid)
		if err != nil {
			errorResponse(w, fmt.Errorf("%s: %w", prompt, err).Error(), http.StatusInternalServerError)
			return
		}

		successResponse(w, http.StatusOK, nil)
	}
}

func LinkIngredientToSalad(app *app.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		prompt := "добавление ингредиента к салату"

		userId, err := getStringClaimFromJWT(r.Context(), "user_id")
		if err != nil {
			errorResponse(w, fmt.Errorf("%s: получение id авторизованного пользователя: %w", prompt, err).Error(), http.StatusBadRequest)
			return
		}
		userUuid, err := uuid.Parse(userId)
		if err != nil {
			errorResponse(w, fmt.Errorf("%s: преобразование id пользователя к uuid: %w", prompt, err).Error(), http.StatusInternalServerError)
			return
		}

		var req LinkIngredientSalad
		err = json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			errorResponse(w, err.Error(), http.StatusBadRequest)
			return
		}

		recipeUuid, saladUuid, ingredientUuid, measurementuuid, amount := toLinkIngredientModel(&req)

		saladDb, err := app.SaladService.GetById(r.Context(), saladUuid)
		if err != nil {
			errorResponse(w, fmt.Errorf("%s: %w", prompt, err).Error(), http.StatusInternalServerError)
			return
		}
		if saladDb.AuthorID != userUuid {
			errorResponse(w, fmt.Errorf("%s: только автор салата может добавлять ингредиенты к салату", prompt).Error(), http.StatusBadRequest)
			return
		}

		linkUuid, err := app.IngredientService.Link(r.Context(), recipeUuid, ingredientUuid)
		if err != nil {
			errorResponse(w, fmt.Errorf("%s: %w", prompt, err).Error(), http.StatusInternalServerError)
			return
		}
		err = app.MeasurementService.UpdateLink(r.Context(), linkUuid, measurementuuid, amount)
		if err != nil {
			errorResponse(w, fmt.Errorf("%s: %w", prompt, err).Error(), http.StatusInternalServerError)
			return
		}

		successResponse(w, http.StatusOK, map[string]string{
			"link_id": linkUuid.String(),
		})
	}
}

func UnlinkIngredientFromSalad(app *app.App) http.HandlerFunc { // FIXME
	return func(w http.ResponseWriter, r *http.Request) {
		prompt := "открепление ингредиента от салата"

		userId, err := getStringClaimFromJWT(r.Context(), "user_id")
		if err != nil {
			errorResponse(w, fmt.Errorf("%s: получение id авторизованного пользователя: %w", prompt, err).Error(), http.StatusBadRequest)
			return
		}
		userUuid, err := uuid.Parse(userId)
		if err != nil {
			errorResponse(w, fmt.Errorf("%s: преобразование id пользователя к uuid: %w", prompt, err).Error(), http.StatusInternalServerError)
			return
		}

		var req LinkIngredientSalad
		err = json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			errorResponse(w, err.Error(), http.StatusBadRequest)
			return
		}

		recipeUuid, saladUuid, ingredientUuid, _, _ := toLinkIngredientModel(&req)

		saladDb, err := app.SaladService.GetById(r.Context(), saladUuid)
		if err != nil {
			errorResponse(w, fmt.Errorf("%s: %w", prompt, err).Error(), http.StatusInternalServerError)
			return
		}
		if saladDb.AuthorID != userUuid {
			errorResponse(w, fmt.Errorf("%s: только автор салата может удалять ингредиенты салата", prompt).Error(), http.StatusBadRequest)
			return
		}

		err = app.IngredientService.Unlink(r.Context(), recipeUuid, ingredientUuid)
		if err != nil {
			errorResponse(w, fmt.Errorf("%s: %w", prompt, err).Error(), http.StatusInternalServerError)
			return
		}

		successResponse(w, http.StatusOK, nil)
	}
}

func GetAllMeasurements(app *app.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		prompt := "получение списка единиц измерения"

		measurements, err := app.MeasurementService.GetAll(r.Context())
		if err != nil {
			errorResponse(w, fmt.Errorf("%s: %w", prompt, err).Error(), http.StatusInternalServerError)
			return
		}

		measurementsTransport := make([]Measurement, len(measurements))
		for i, measurement := range measurements {
			measurementsTransport[i] = toMeasurementTransport(measurement)
		}

		successResponse(w, http.StatusOK, map[string]interface{}{"measurements": measurementsTransport})
	}
}

func CreateRecipeStep(app *app.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		prompt := "создание шага рецепта"

		userId, err := getStringClaimFromJWT(r.Context(), "user_id")
		if err != nil {
			errorResponse(w, fmt.Errorf("%s: получение id авторизованного пользователя: %w", prompt, err).Error(), http.StatusBadRequest)
			return
		}
		userUuid, err := uuid.Parse(userId)
		if err != nil {
			errorResponse(w, fmt.Errorf("%s: преобразование id пользователя к uuid: %w", prompt, err).Error(), http.StatusInternalServerError)
			return
		}

		var req RecipeStep
		err = json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			errorResponse(w, err.Error(), http.StatusBadRequest)
			return
		}

		recipeStep := toRecipeStepModel(&req)

		recipeDb, err := app.RecipeService.GetById(r.Context(), recipeStep.RecipeID)
		if err != nil {
			errorResponse(w, fmt.Errorf("%s: %w", prompt, err).Error(), http.StatusInternalServerError)
			return
		}
		saladDb, err := app.SaladService.GetById(r.Context(), recipeDb.SaladID)
		if err != nil {
			errorResponse(w, fmt.Errorf("%s: %w", prompt, err).Error(), http.StatusInternalServerError)
			return
		}
		if saladDb.AuthorID != userUuid {
			errorResponse(w, fmt.Errorf("%s: только автор салата может удалять ингредиенты салата", prompt).Error(), http.StatusBadRequest)
			return
		}

		err = app.RecipeStepService.Create(r.Context(), recipeStep)
		if err != nil {
			errorResponse(w, fmt.Errorf("%s: %w", prompt, err).Error(), http.StatusInternalServerError)
			return
		}

		successResponse(w, http.StatusOK, nil)
	}
}

func DeleteRecipeStep(app *app.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		prompt := "удаление шага рецепта"

		userId, err := getStringClaimFromJWT(r.Context(), "user_id")
		if err != nil {
			errorResponse(w, fmt.Errorf("получение id авторизованного пользователя: %w", err).Error(), http.StatusBadRequest)
			return
		}
		userUuid, err := uuid.Parse(userId)
		if err != nil {
			errorResponse(w, fmt.Errorf("преобразование id пользователя к uuid: %w", err).Error(), http.StatusInternalServerError)
			return
		}

		id := chi.URLParam(r, "id")
		if id == "" {
			app.Logger.Infof("getting id")
			errorResponse(w, fmt.Errorf("%s: пустой id", prompt).Error(), http.StatusBadRequest)
			return
		}
		idUuid, err := uuid.Parse(id)
		if err != nil {
			errorResponse(w, fmt.Errorf("%s: %w", prompt, err).Error(), http.StatusBadRequest)
			return
		}

		recipeStep, err := app.RecipeStepService.GetById(r.Context(), idUuid)
		if err != nil {
			errorResponse(w, fmt.Errorf("%s: %w", prompt, err).Error(), http.StatusInternalServerError)
			return
		}
		recipeDb, err := app.RecipeService.GetById(r.Context(), recipeStep.RecipeID)
		if err != nil {
			errorResponse(w, fmt.Errorf("%s: %w", prompt, err).Error(), http.StatusInternalServerError)
			return
		}
		saladDb, err := app.SaladService.GetById(r.Context(), recipeDb.SaladID)
		if err != nil {
			errorResponse(w, fmt.Errorf("%s: %w", prompt, err).Error(), http.StatusInternalServerError)
			return
		}
		if saladDb.AuthorID != userUuid {
			errorResponse(w, fmt.Errorf("%s: только автор рецепта может удалять шаги", prompt).Error(), http.StatusBadRequest)
			return
		}

		err = app.RecipeStepService.DeleteById(r.Context(), idUuid)
		if err != nil {
			errorResponse(w, fmt.Errorf("%s: %w", prompt, err).Error(), http.StatusInternalServerError)
			return
		}

		successResponse(w, http.StatusOK, nil)
	}
}

func UpdateRecipeStep(app *app.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		prompt := "обновление шага рецепта"

		userId, err := getStringClaimFromJWT(r.Context(), "user_id")
		if err != nil {
			errorResponse(w, fmt.Errorf("получение id авторизованного пользователя: %w", err).Error(), http.StatusBadRequest)
			return
		}
		userUuid, err := uuid.Parse(userId)
		if err != nil {
			errorResponse(w, fmt.Errorf("преобразование id пользователя к uuid: %w", err).Error(), http.StatusInternalServerError)
			return
		}

		id := chi.URLParam(r, "id")
		if id == "" {
			errorResponse(w, fmt.Errorf("%s: пустой id комментария", prompt).Error(), http.StatusBadRequest)
			return
		}
		idUuid, err := uuid.Parse(id)
		if err != nil {
			errorResponse(w, fmt.Errorf("%s: преобразование id комментария к uuid: %w", prompt, err).Error(), http.StatusBadRequest)
			return
		}

		recipeStep, err := app.RecipeStepService.GetById(r.Context(), idUuid)
		if err != nil {
			errorResponse(w, fmt.Errorf("%s: %w", prompt, err).Error(), http.StatusInternalServerError)
			return
		}
		recipeDb, err := app.RecipeService.GetById(r.Context(), recipeStep.RecipeID)
		if err != nil {
			errorResponse(w, fmt.Errorf("%s: %w", prompt, err).Error(), http.StatusInternalServerError)
			return
		}
		saladDb, err := app.SaladService.GetById(r.Context(), recipeDb.SaladID)
		if err != nil {
			errorResponse(w, fmt.Errorf("%s: %w", prompt, err).Error(), http.StatusInternalServerError)
			return
		}
		if saladDb.AuthorID != userUuid {
			errorResponse(w, fmt.Errorf("%s: только автор рецепта может изменять шаги", prompt).Error(), http.StatusBadRequest)
			return
		}

		var req RecipeStep
		err = json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			errorResponse(w, err.Error(), http.StatusBadRequest)
			return
		}
		if req.Name != "" {
			recipeStep.Name = req.Name
		}
		if req.Description != "" {
			recipeStep.Description = req.Description
		}
		// FIXME: should i change recipe step num? btw now only deleting steps

		err = app.RecipeStepService.Update(r.Context(), recipeStep)
		if err != nil {
			errorResponse(w, fmt.Errorf("%s: %w", prompt, err).Error(), http.StatusInternalServerError)
			return
		}

		successResponse(w, http.StatusOK, map[string]interface{}{
			"step": toRecipeStepTransport(recipeStep),
		})
	}
}
