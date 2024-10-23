package dto

import "github.com/google/uuid"

const (
	EditingSaladStatus    = 1
	ModerationSaladStatus = 2
	RejectedSaladStatus   = 3
	PublishedSaladStatus  = 4
	StoredSaladStatus     = 5
)

type RecipeFilter struct {
	AvailableIngredients []uuid.UUID
	MinRate              float64
	SaladTypes           []uuid.UUID
	Status               int
}
