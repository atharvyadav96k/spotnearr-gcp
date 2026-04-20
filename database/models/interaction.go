package models

import (
	"time"
)

// InteractionBuffer stores raw interaction events temporarily
type InteractionBuffer struct {
	ID          string    `json:"_id" bson:"_id"`
	UserID      *string   `json:"userId,omitempty" bson:"userId,omitempty"`
	BusinessID  string    `json:"businessId" bson:"businessId"`
	SpotlightID *string   `json:"spotlightId,omitempty" bson:"spotlightId,omitempty"`
	Type        string    `json:"type" bson:"type"` // view, like, share, visit, unlike, comment, ignore
	CreatedAt   int64     `json:"createdAt" bson:"createdAt"`
	CreationTime time.Time `json:"_creationTime" bson:"_creationTime"`
}

// Follow represents a user following a business
type Follow struct {
	ID          string    `json:"_id" bson:"_id"`
	UserID      string    `json:"userId" bson:"userId"`
	BusinessID  string    `json:"businessId" bson:"businessId"`
	IsCelebrity bool      `json:"isCelebrity" bson:"isCelebrity"`
	CreationTime time.Time `json:"_creationTime" bson:"_creationTime"`
}

// Review represents a user review for a business
type Review struct {
	ID          string    `json:"_id" bson:"_id"`
	UserID      string    `json:"userId" bson:"userId"`
	BusinessID  string    `json:"businessId" bson:"businessId"`
	Rating      int64     `json:"rating" bson:"rating"` // 1 to 5
	Content     string    `json:"content" bson:"content"`
	ImageURL    *string   `json:"imageUrl,omitempty" bson:"imageUrl,omitempty"`
	CreatedAt   int64     `json:"createdAt" bson:"createdAt"`
	CreationTime time.Time `json:"_creationTime" bson:"_creationTime"`
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
