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
