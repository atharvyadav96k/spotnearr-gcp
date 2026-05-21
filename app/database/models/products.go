package models

import (
	"time"

	"github.com/google/uuid"
)

type Product struct {
	ID         uuid.UUID `db:"id" json:"id"`
	CategoryID uuid.UUID `db:"category_id" json:"category_id"`

	Name        string  `db:"name" json:"name"`
	Description *string `db:"description" json:"description,omitempty"`

	Unit string `db:"unit" json:"unit"`

	IsAvailable bool `db:"is_available" json:"is_available"`
	IsActive    bool `db:"is_active" json:"is_active"`

	Tags []string `db:"tags" json:"tags"`

	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

type ProductInventory struct {
	ID uuid.UUID `db:"id" json:"id"`

	ProductID uuid.UUID `db:"product_id" json:"product_id"`

	LocationID uuid.UUID `db:"location_id" json:"location_id"`

	Price int64 `db:"price" json:"price"`

	DiscountedPrice int64 `db:"discounted_price" json:"discounted_price"`

	Stock int `db:"stock" json:"stock"`

	IsAvailable bool `db:"is_available" json:"is_available"`

	CreatedAt time.Time `db:"created_at" json:"created_at"`

	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

type ProductCategory struct {
	ID uuid.UUID `db:"id" json:"id"`

	Name string `db:"name" json:"name"`

	ParentID *uuid.UUID `db:"parent_id" json:"parent_id,omitempty"`

	IconURL *string `db:"icon_url" json:"icon_url,omitempty"`

	CreatedAt time.Time `db:"created_at" json:"created_at"`
}

type Media struct {
	ID uuid.UUID `db:"id" json:"id"`

	URL string `db:"url" json:"url"`

	Type string `db:"type" json:"type"`

	ThumbnailURL *string `db:"thumbnail_url" json:"thumbnail_url,omitempty"`

	Hash *string `db:"hash" json:"hash,omitempty"`

	CreatedAt time.Time `db:"created_at" json:"created_at"`
}

type ProductMedia struct {
	ProductID uuid.UUID `db:"product_id" json:"product_id"`

	MediaID uuid.UUID `db:"media_id" json:"media_id"`

	IsPrimary bool `db:"is_primary" json:"is_primary"`

	SortOrder int `db:"sort_order" json:"sort_order"`
}

type ProductResponse struct {
	ID          uuid.UUID `json:"id"`
	CategoryID  uuid.UUID `json:"category_id"`
	Name        string    `json:"name"`
	Description *string   `json:"description,omitempty"`
	Unit        string    `json:"unit"`
	IsAvailable bool      `json:"is_available"`
	IsActive    bool      `json:"is_active"`
	Tags        []string  `json:"tags"`
	CreatedAt   time.Time `json:"created_at"`
}

func (p Product) ToResponse() any {
	return ProductResponse{
		ID: p.ID, CategoryID: p.CategoryID, Name: p.Name, Description: p.Description,
		Unit: p.Unit, IsAvailable: p.IsAvailable, IsActive: p.IsActive,
		Tags: p.Tags, CreatedAt: p.CreatedAt,
	}
}

type ProductInventoryResponse struct {
	ID              uuid.UUID `json:"id"`
	ProductID       uuid.UUID `json:"product_id"`
	LocationID      uuid.UUID `json:"location_id"`
	Price           int64     `json:"price"`
	DiscountedPrice int64     `json:"discounted_price"`
	Stock           int       `json:"stock"`
	IsAvailable     bool      `json:"is_available"`
	CreatedAt       time.Time `json:"created_at"`
}

func (p ProductInventory) ToResponse() any {
	return ProductInventoryResponse{
		ID: p.ID, ProductID: p.ProductID, LocationID: p.LocationID,
		Price: p.Price, DiscountedPrice: p.DiscountedPrice, Stock: p.Stock,
		IsAvailable: p.IsAvailable, CreatedAt: p.CreatedAt,
	}
}

type ProductCategoryResponse struct {
	ID       uuid.UUID  `json:"id"`
	Name     string     `json:"name"`
	ParentID *uuid.UUID `json:"parent_id,omitempty"`
	IconURL  *string    `json:"icon_url,omitempty"`
}

func (p ProductCategory) ToResponse() any {
	return ProductCategoryResponse{ID: p.ID, Name: p.Name, ParentID: p.ParentID, IconURL: p.IconURL}
}

type MediaResponse struct {
	ID           uuid.UUID `json:"id"`
	URL          string    `json:"url"`
	Type         string    `json:"type"`
	ThumbnailURL *string   `json:"thumbnail_url,omitempty"`
}

func (m Media) ToResponse() any {
	return MediaResponse{ID: m.ID, URL: m.URL, Type: m.Type, ThumbnailURL: m.ThumbnailURL}
}

type ProductMediaResponse struct {
	ProductID uuid.UUID `json:"product_id"`
	MediaID   uuid.UUID `json:"media_id"`
	IsPrimary bool      `json:"is_primary"`
	SortOrder int       `json:"sort_order"`
}

func (p ProductMedia) ToResponse() any {
	return ProductMediaResponse{ProductID: p.ProductID, MediaID: p.MediaID, IsPrimary: p.IsPrimary, SortOrder: p.SortOrder}
}
