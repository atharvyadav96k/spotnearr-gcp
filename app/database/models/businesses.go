package models

import (
	"strings"
	"time"

	"github.com/atharvyadav96k/spotnearr-gcp/app/utils"
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

type BusinessResponse struct {
	ID           uuid.UUID      `json:"id"`
	OwnerID      uuid.UUID      `json:"owner_id"`
	Name         string         `json:"name"`
	Description  *string        `json:"description,omitempty"`
	CategoryID   uuid.UUID      `json:"category_id"`
	Email        *string        `json:"email,omitempty"`
	Phone        *string        `json:"phone,omitempty"`
	Website      *string        `json:"website,omitempty"`
	LogoURL      *string        `json:"logo_url,omitempty"`
	CoverURL     *string        `json:"cover_url,omitempty"`
	Status       BusinessStatus `json:"status"`
	IsVerified   bool           `json:"is_verified"`
	Rating       float64        `json:"rating"`
	TotalReviews int            `json:"total_reviews"`
	CreatedAt    time.Time      `json:"created_at"`
}

func (b Business) ToResponse() any {
	return BusinessResponse{
		ID: b.ID, OwnerID: b.OwnerID, Name: b.Name, Description: b.Description,
		CategoryID: b.CategoryID, Email: b.Email, Phone: b.Phone, Website: b.Website,
		LogoURL: b.LogoURL, CoverURL: b.CoverURL, Status: b.Status, IsVerified: b.IsVerified,
		Rating: b.Rating, TotalReviews: b.TotalReviews, CreatedAt: b.CreatedAt,
	}
}

type BusinessCategoryResponse struct {
	ID      uuid.UUID `json:"id"`
	Name    string    `json:"name"`
	IconURL *string   `json:"icon_url,omitempty"`
}

func (b BusinessCategory) ToResponse() any {
	return BusinessCategoryResponse{ID: b.ID, Name: b.Name, IconURL: b.IconURL}
}

type BusinessLocationResponse struct {
	ID           uuid.UUID `json:"id"`
	BusinessID   uuid.UUID `json:"business_id"`
	BranchName   string    `json:"branch_name"`
	AddressLine1 string    `json:"address_line1"`
	AddressLine2 *string   `json:"address_line2,omitempty"`
	City         string    `json:"city"`
	State        string    `json:"state"`
	PinCode      string    `json:"pin_code"`
	Latitude     float64   `json:"latitude"`
	Longitude    float64   `json:"longitude"`
	IsMain       bool      `json:"is_main"`
	OpeningTime  string    `json:"opening_time"`
	ClosingTime  string    `json:"closing_time"`
	WorkingDays  []string  `json:"working_days"`
	CreatedAt    time.Time `json:"created_at"`
}

func (b Business) Validate() error {
	ve := &utils.ValidationErrors{}
	if strings.TrimSpace(b.Name) == "" {
		ve.Add("name", "name is required")
	}
	if b.CategoryID == (uuid.UUID{}) {
		ve.Add("category_id", "category_id is required")
	}
	if ve.HasErrors() {
		return ve
	}
	return nil
}

func (l BusinessLocation) Validate() error {
	ve := &utils.ValidationErrors{}
	if l.BusinessID == (uuid.UUID{}) {
		ve.Add("business_id", "business_id is required")
	}
	if strings.TrimSpace(l.BranchName) == "" {
		ve.Add("branch_name", "branch_name is required")
	}
	if strings.TrimSpace(l.AddressLine1) == "" {
		ve.Add("address_line1", "address_line1 is required")
	}
	if strings.TrimSpace(l.City) == "" {
		ve.Add("city", "city is required")
	}
	if strings.TrimSpace(l.State) == "" {
		ve.Add("state", "state is required")
	}
	if strings.TrimSpace(l.PinCode) == "" {
		ve.Add("pin_code", "pin_code is required")
	}
	if strings.TrimSpace(l.OpeningTime) == "" {
		ve.Add("opening_time", "opening_time is required")
	}
	if strings.TrimSpace(l.ClosingTime) == "" {
		ve.Add("closing_time", "closing_time is required")
	}
	if len(l.WorkingDays) == 0 {
		ve.Add("working_days", "at least one working day is required")
	}
	if ve.HasErrors() {
		return ve
	}
	return nil
}

func (b BusinessLocation) ToResponse() any {
	return BusinessLocationResponse{
		ID: b.ID, BusinessID: b.BusinessID, BranchName: b.BranchName,
		AddressLine1: b.AddressLine1, AddressLine2: b.AddressLine2,
		City: b.City, State: b.State, PinCode: b.PinCode,
		Latitude: b.Latitude, Longitude: b.Longitude,
		IsMain: b.IsMain, OpeningTime: b.OpeningTime, ClosingTime: b.ClosingTime,
		WorkingDays: b.WorkingDays, CreatedAt: b.CreatedAt,
	}
}
