package models

import (
	"time"

	"github.com/google/uuid"
)

type BusinessStatus string

const (
	BusinessPending   BusinessStatus = "pending"
	BusinessActive    BusinessStatus = "active"
	BusinessSuspended BusinessStatus = "suspended"
)

type Business struct {
	ID           uuid.UUID      `db:"id" json:"id"`
	OwnerID      uuid.UUID      `db:"owner_id" json:"owner_id"`
	Name         string         `db:"name" json:"name"`
	Description  *string        `db:"description" json:"description,omitempty"`
	CategoryID   uuid.UUID      `db:"category_id" json:"category_id"`
	Email        *string        `db:"email" json:"email,omitempty"`
	Phone        *string        `db:"phone" json:"phone,omitempty"`
	Website      *string        `db:"website" json:"website,omitempty"`
	LogoURL      *string        `db:"logo_url" json:"logo_url,omitempty"`
	CoverURL     *string        `db:"cover_url" json:"cover_url,omitempty"`
	Status       BusinessStatus `db:"status" json:"status"`
	IsVerified   bool           `db:"is_verified" json:"is_verified"`
	Rating       float64        `db:"rating" json:"rating"`
	TotalReviews int            `db:"total_reviews" json:"total_reviews"`
	CreatedAt    time.Time      `db:"created_at" json:"created_at"`
	UpdatedAt    time.Time      `db:"updated_at" json:"updated_at"`
}

type BusinessCategory struct {
	ID        uuid.UUID `db:"id" json:"id"`
	Name      string    `db:"name" json:"name"`
	IconURL   *string   `db:"icon_url" json:"icon_url,omitempty"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
}

type BusinessLocation struct {
	ID         uuid.UUID `db:"id" json:"id"`
	BusinessID uuid.UUID `db:"business_id" json:"business_id"`

	BranchName   string  `db:"branch_name" json:"branch_name"`
	AddressLine1 string  `db:"address_line1" json:"address_line1"`
	AddressLine2 *string `db:"address_line2" json:"address_line2,omitempty"`

	City    string `db:"city" json:"city"`
	State   string `db:"state" json:"state"`
	PinCode string `db:"pin_code" json:"pin_code"`

	Latitude  float64 `db:"latitude" json:"latitude"`
	Longitude float64 `db:"longitude" json:"longitude"`

	Geohash string `db:"geohash" json:"geohash"`

	IsMain bool `db:"is_main" json:"is_main"`

	OpeningTime string    `db:"opening_time" json:"opening_time"`
	ClosingTime string    `db:"closing_time" json:"closing_time"`
	WorkingDays []string  `db:"working_days" json:"working_days"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`

	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}
