package models

import (
	"time"
)

// InteractionBuffer stores raw interaction events temporarily
type InteractionBuffer struct {
	ID          string    `json:"_id"`
	UserID      *string   `json:"userId,omitempty"`
	BusinessID  string    `json:"businessId"`
	SpotlightID *string   `json:"spotlightId,omitempty"`
	Type        string    `json:"type"` // view, like, share, visit, unlike, comment, ignore
	CreatedAt   int64     `json:"createdAt"`
	CreationTime time.Time `json:"_creationTime"`
}

// Follow represents a user following a business
type Follow struct {
	ID          string    `json:"_id"`
	UserID      string    `json:"userId"`
	BusinessID  string    `json:"businessId"`
	IsCelebrity bool      `json:"isCelebrity"`
	CreationTime time.Time `json:"_creationTime"`
}

// Review represents a user review for a business
type Review struct {
	ID          string    `json:"_id"`
	UserID      string    `json:"userId"`
	BusinessID  string    `json:"businessId"`
	Rating      int64     `json:"rating"` // 1 to 5
	Content     string    `json:"content"`
	ImageURL    *string   `json:"imageUrl,omitempty"`
	CreatedAt   int64     `json:"createdAt"`
	CreationTime time.Time `json:"_creationTime"`
}

// Interaction type constants
const (
	InteractionTypeView     = "view"
	InteractionTypeLike     = "like"
	InteractionTypeShare    = "share"
	InteractionTypeVisit    = "visit"
	InteractionTypeUnlike   = "unlike"
	InteractionTypeComment  = "comment"
	InteractionTypeIgnore   = "ignore"
)
