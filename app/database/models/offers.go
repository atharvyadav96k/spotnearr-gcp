package models

import (
	"time"

	"github.com/google/uuid"
)

type OfferDiscountType string

const (
	DiscountPercentage OfferDiscountType = "percentage"
	DiscountFlat       OfferDiscountType = "flat"
)

type Offer struct {
	ID uuid.UUID `db:"id" json:"id"`

	BusinessID uuid.UUID `db:"business_id" json:"business_id"`

	Title string `db:"title" json:"title"`

	Description *string `db:"description" json:"description,omitempty"`

	DiscountType OfferDiscountType `db:"discount_type" json:"discount_type"`

	DiscountValue int64 `db:"discount_value" json:"discount_value"`

	MinOrderValue int64 `db:"min_order_value" json:"min_order_value"`

	CouponCode *string `db:"coupon_code" json:"coupon_code,omitempty"`

	BannerURL *string `db:"banner_url" json:"banner_url,omitempty"`

	IsActive bool `db:"is_active" json:"is_active"`

	StartsAt time.Time `db:"starts_at" json:"starts_at"`

	ExpiresAt time.Time `db:"expires_at" json:"expires_at"`

	CreatedAt time.Time `db:"created_at" json:"created_at"`

	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

type OfferProduct struct {
	OfferID uuid.UUID `db:"offer_id" json:"offer_id"`

	ProductID uuid.UUID `db:"product_id" json:"product_id"`
}
