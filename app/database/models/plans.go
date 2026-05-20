package models

import (
	"time"

	"github.com/atharvyadav96k/spotnearr-gcp/app/database/enums"
	"github.com/google/uuid"
)

type Plan struct {
	ID          uuid.UUID  `db:"id" json:"id"`
	Name        string     `db:"name" json:"name"`
	Tier        enums.Plan `db:"tier" json:"tier"`
	Price       float64    `db:"price" json:"price"`
	Description *string    `db:"description" json:"description,omitempty"`
	IsActive    bool       `db:"is_active" json:"is_active"`
	CreatedAt   time.Time  `db:"created_at" json:"created_at"`
}

type BusinessSubscription struct {
	ID         uuid.UUID  `db:"id" json:"id"`
	BusinessID uuid.UUID  `db:"business_id" json:"business_id"`
	PlanID     *uuid.UUID `db:"plan_id" json:"plan_id,omitempty"`
	Tier       enums.Plan `db:"tier" json:"tier"`
	PlanName   *string    `db:"plan_name" json:"plan_name,omitempty"`
	Price      *float64   `db:"price" json:"price,omitempty"`
	StartedAt  time.Time  `db:"started_at" json:"started_at"`
	ExpiresAt  *time.Time `db:"expires_at" json:"expires_at,omitempty"`
	CreatedAt  time.Time  `db:"created_at" json:"created_at"`
}
