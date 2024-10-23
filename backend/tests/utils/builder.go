package utils

import (
	"github.com/google/uuid"
	"net/mail"
	"ppo/domain"
)

type UserAuthBuilder struct {
	domain.UserAuth
}

func NewUserAuthBuilder() *UserAuthBuilder {
	return &UserAuthBuilder{}
}

func (u *UserAuthBuilder) WithId(id uuid.UUID) *UserAuthBuilder {
	u.ID = id
	return u
}

func (u *UserAuthBuilder) WithUsername(username string) *UserAuthBuilder {
	u.Username = username
	return u
}

func (u *UserAuthBuilder) WithPassword(password string) *UserAuthBuilder {
	u.Password = password
	return u
}

func (u *UserAuthBuilder) WithHashedPass(password string) *UserAuthBuilder {
	u.HashedPass = password
	return u
}

func (u *UserAuthBuilder) ToDto() *domain.UserAuth {
	return &domain.UserAuth{
		ID:         u.ID,
		Username:   u.Username,
		Password:   u.Password,
		HashedPass: u.HashedPass,
		Role:       u.Role,
	}
}

type UserBuilder struct {
	domain.User
}

func NewUserBuilder() *UserBuilder {
	return &UserBuilder{}
}

func (u *UserBuilder) WithId(id uuid.UUID) *UserBuilder {
	u.ID = id
	return u
}

func (u *UserBuilder) WithName(name string) *UserBuilder {
	u.Name = name
	return u
}

func (u *UserBuilder) WithUsername(username string) *UserBuilder {
	u.Username = username
	return u
}

func (u *UserBuilder) WithPassword(password string) *UserBuilder {
	u.Password = password
	return u
}

func (u *UserBuilder) WithEmail(email string) *UserBuilder {
	u.Email = mail.Address{
		Address: email,
	}
	return u
}

func (u *UserBuilder) WithRole(role string) *UserBuilder {
	u.Role = role
	return u
}

func (u *UserBuilder) ToDto() *domain.User {
	return &domain.User{
		ID:       u.ID,
		Name:     u.Name,
		Username: u.Username,
		Password: u.Password,
		Email:    u.Email,
		Role:     u.Role,
	}
}

type CommentBuilder struct {
	domain.Comment
}

func NewCommentBuilder() *CommentBuilder {
	return &CommentBuilder{}
}

func (u *CommentBuilder) WithId(id uuid.UUID) *CommentBuilder {
	u.ID = id
	return u
}

func (u *CommentBuilder) WithAuthorId(id uuid.UUID) *CommentBuilder {
	u.AuthorID = id
	return u
}

func (u *CommentBuilder) WithSaladId(id uuid.UUID) *CommentBuilder {
	u.SaladID = id
	return u
}

func (u *CommentBuilder) WithText(text string) *CommentBuilder {
	u.Text = text
	return u
}

func (u *CommentBuilder) WithRating(rating int) *CommentBuilder {
	u.Rating = rating
	return u
}

func (u *CommentBuilder) ToDto() *domain.Comment {
	return &domain.Comment{
		ID:       u.ID,
		AuthorID: u.AuthorID,
		SaladID:  u.SaladID,
		Text:     u.Text,
		Rating:   u.Rating,
	}
}

type IngredientTypeBuilder struct {
	domain.IngredientType
}

func NewIngredientTypeBuilder() *IngredientTypeBuilder {
	return &IngredientTypeBuilder{}
}

func (u *IngredientTypeBuilder) WithId(id uuid.UUID) *IngredientTypeBuilder {
	u.ID = id
	return u
}

func (u *IngredientTypeBuilder) WithName(name string) *IngredientTypeBuilder {
	u.Name = name
	return u
}

func (u *IngredientTypeBuilder) WithDescription(description string) *IngredientTypeBuilder {
	u.Description = description
	return u
}

func (u *IngredientTypeBuilder) ToDto() *domain.IngredientType {
	return &domain.IngredientType{
		ID:          u.ID,
		Name:        u.Name,
		Description: u.Description,
	}
}

type IngredientBuilder struct {
	domain.Ingredient
}

func NewIngredientBuilder() *IngredientBuilder {
	return &IngredientBuilder{}
}

func (u *IngredientBuilder) WithId(id uuid.UUID) *IngredientBuilder {
	u.ID = id
	return u
}

func (u *IngredientBuilder) WithName(name string) *IngredientBuilder {
	u.Name = name
	return u
}

func (u *IngredientBuilder) WithCalories(calories int) *IngredientBuilder {
	u.Calories = calories
	return u
}

func (u *IngredientBuilder) WithTypeId(id uuid.UUID) *IngredientBuilder {
	u.TypeID = id
	return u
}

func (u *IngredientBuilder) ToDto() *domain.Ingredient {
	return &domain.Ingredient{
		ID:       u.ID,
		Name:     u.Name,
		Calories: u.Calories,
		TypeID:   u.TypeID,
	}
}

type KeywordBuilder struct {
	domain.KeyWord
}

func NewKeywordBuilder() *KeywordBuilder {
	return &KeywordBuilder{}
}

func (u *KeywordBuilder) WithId(id uuid.UUID) *KeywordBuilder {
	u.ID = id
	return u
}

func (u *KeywordBuilder) WithWord(word string) *KeywordBuilder {
	u.Word = word
	return u
}

func (u *KeywordBuilder) ToDto() *domain.KeyWord {
	return &domain.KeyWord{
		ID:   u.ID,
		Word: u.Word,
	}
}

type MeasurementBuilder struct {
	domain.Measurement
}

func NewMeasurementBuilder() *MeasurementBuilder {
	return &MeasurementBuilder{}
}

func (u *MeasurementBuilder) WithId(id uuid.UUID) *MeasurementBuilder {
	u.ID = id
	return u
}

func (u *MeasurementBuilder) WithName(name string) *MeasurementBuilder {
	u.Name = name
	return u
}

func (u *MeasurementBuilder) WithGrams(grams int) *MeasurementBuilder {
	u.Grams = grams
	return u
}

func (u *MeasurementBuilder) ToDto() *domain.Measurement {
	return &domain.Measurement{
		ID:    u.ID,
		Name:  u.Name,
		Grams: u.Grams,
	}
}

type RecipeStepBuilder struct {
	domain.RecipeStep
}

func NewRecipeStepBuilder() *RecipeStepBuilder {
	return &RecipeStepBuilder{}
}

func (u *RecipeStepBuilder) WithId(id uuid.UUID) *RecipeStepBuilder {
	u.ID = id
	return u
}

func (u *RecipeStepBuilder) WithRecipeId(id uuid.UUID) *RecipeStepBuilder {
	u.RecipeID = id
	return u
}

func (u *RecipeStepBuilder) WithName(name string) *RecipeStepBuilder {
	u.Name = name
	return u
}

func (u *RecipeStepBuilder) WithDescription(description string) *RecipeStepBuilder {
	u.Description = description
	return u
}

func (u *RecipeStepBuilder) WithStepNum(stepNum int) *RecipeStepBuilder {
	u.StepNum = stepNum
	return u
}

func (u *RecipeStepBuilder) ToDto() *domain.RecipeStep {
	return &domain.RecipeStep{
		ID:          u.ID,
		RecipeID:    u.RecipeID,
		Name:        u.Name,
		Description: u.Description,
		StepNum:     u.StepNum,
	}
}

type RecipeBuilder struct {
	domain.Recipe
}

func NewRecipeBuilder() *RecipeBuilder {
	return &RecipeBuilder{}
}

func (u *RecipeBuilder) WithId(id uuid.UUID) *RecipeBuilder {
	u.ID = id
	return u
}

func (u *RecipeBuilder) WithSaladId(id uuid.UUID) *RecipeBuilder {
	u.SaladID = id
	return u
}

func (u *RecipeBuilder) WithStatus(status int) *RecipeBuilder {
	u.Status = status
	return u
}

func (u *RecipeBuilder) WithNumberOfServings(num int) *RecipeBuilder {
	u.NumberOfServings = num
	return u
}

func (u *RecipeBuilder) WithTimeToCook(time int) *RecipeBuilder {
	u.TimeToCook = time
	return u
}

func (u *RecipeBuilder) WithRating(rating float32) *RecipeBuilder {
	u.Rating = rating
	return u
}

func (u *RecipeBuilder) ToDto() *domain.Recipe {
	return &domain.Recipe{
		ID:               u.ID,
		SaladID:          u.SaladID,
		Status:           u.Status,
		NumberOfServings: u.NumberOfServings,
		TimeToCook:       u.TimeToCook,
		Rating:           u.Rating,
	}
}

type SaladTypeBuilder struct {
	domain.SaladType
}

func NewSaladTypeBuilder() *SaladTypeBuilder {
	return &SaladTypeBuilder{}
}

func (u *SaladTypeBuilder) WithId(id uuid.UUID) *SaladTypeBuilder {
	u.ID = id
	return u
}

func (u *SaladTypeBuilder) WithName(name string) *SaladTypeBuilder {
	u.Name = name
	return u
}

func (u *SaladTypeBuilder) WithDescription(description string) *SaladTypeBuilder {
	u.Description = description
	return u
}

func (u *SaladTypeBuilder) ToDto() *domain.SaladType {
	return &domain.SaladType{
		ID:          u.ID,
		Name:        u.Name,
		Description: u.Description,
	}
}

type SaladBuilder struct {
	domain.Salad
}

func NewSaladBuilder() *SaladBuilder {
	return &SaladBuilder{}
}

func (u *SaladBuilder) WithId(id uuid.UUID) *SaladBuilder {
	u.ID = id
	return u
}

func (u *SaladBuilder) WithAuthorId(id uuid.UUID) *SaladBuilder {
	u.AuthorID = id
	return u
}

func (u *SaladBuilder) WithName(name string) *SaladBuilder {
	u.Name = name
	return u
}

func (u *SaladBuilder) WithDescription(description string) *SaladBuilder {
	u.Description = description
	return u
}

func (u *SaladBuilder) ToDto() *domain.Salad {
	return &domain.Salad{
		ID:          u.ID,
		AuthorID:    u.AuthorID,
		Name:        u.Name,
		Description: u.Description,
	}
}
