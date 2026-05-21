package models

import (
	"strings"
	"time"

	"github.com/atharvyadav96k/spotnearr-gcp/app/utils"
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

type OfferResponse struct {
	ID            uuid.UUID         `json:"id"`
	BusinessID    uuid.UUID         `json:"business_id"`
	Title         string            `json:"title"`
	Description   *string           `json:"description,omitempty"`
	DiscountType  OfferDiscountType `json:"discount_type"`
	DiscountValue int64             `json:"discount_value"`
	MinOrderValue int64             `json:"min_order_value"`
	CouponCode    *string           `json:"coupon_code,omitempty"`
	BannerURL     *string           `json:"banner_url,omitempty"`
	IsActive      bool              `json:"is_active"`
	StartsAt      time.Time         `json:"starts_at"`
	ExpiresAt     time.Time         `json:"expires_at"`
	CreatedAt     time.Time         `json:"created_at"`
}

func (o Offer) Validate() error {
	ve := &utils.ValidationErrors{}
	if o.BusinessID == (uuid.UUID{}) {
		ve.Add("business_id", "business_id is required")
	}
	if strings.TrimSpace(o.Title) == "" {
		ve.Add("title", "title is required")
	}
	if o.DiscountType == "" {
		ve.Add("discount_type", "discount_type is required (percentage|flat)")
	}
	if o.DiscountValue <= 0 {
		ve.Add("discount_value", "discount_value must be greater than 0")
	}
	if o.StartsAt.IsZero() {
		ve.Add("starts_at", "starts_at is required")
	}
	if o.ExpiresAt.IsZero() {
		ve.Add("expires_at", "expires_at is required")
	}
	if ve.HasErrors() {
		return ve
	}
	return nil
}

func (o Offer) ToResponse() any {
	return OfferResponse{
		ID: o.ID, BusinessID: o.BusinessID, Title: o.Title, Description: o.Description,
		DiscountType: o.DiscountType, DiscountValue: o.DiscountValue, MinOrderValue: o.MinOrderValue,
		CouponCode: o.CouponCode, BannerURL: o.BannerURL, IsActive: o.IsActive,
		StartsAt: o.StartsAt, ExpiresAt: o.ExpiresAt, CreatedAt: o.CreatedAt,
	}
}

type OfferProductResponse struct {
	OfferID   uuid.UUID `json:"offer_id"`
	ProductID uuid.UUID `json:"product_id"`
}

func (o OfferProduct) ToResponse() any {
	return OfferProductResponse{OfferID: o.OfferID, ProductID: o.ProductID}
}
