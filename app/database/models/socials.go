package models

import (
	"time"

	"github.com/google/uuid"
)

type NotificationType string
type NotificationRefType string

const (
	NotificationSpotlightPosted NotificationType = "spotlight_posted"
	NotificationOfferLive       NotificationType = "offer_live"
	NotificationNewProduct      NotificationType = "new_product"
	NotificationReview          NotificationType = "review"

	RefSpotlight NotificationRefType = "spotlight"
	RefOffer     NotificationRefType = "offer"
	RefProduct   NotificationRefType = "product"
)

type UserFollowBusiness struct {
	UserID uuid.UUID `db:"user_id" json:"user_id"`

	BusinessID uuid.UUID `db:"business_id" json:"business_id"`

	CreatedAt time.Time `db:"created_at" json:"created_at"`
}

type UserSavedProduct struct {
	UserID uuid.UUID `db:"user_id" json:"user_id"`

	ProductID uuid.UUID `db:"product_id" json:"product_id"`

	CreatedAt time.Time `db:"created_at" json:"created_at"`
}

type BusinessReview struct {
	ID uuid.UUID `db:"id" json:"id"`

	BusinessID uuid.UUID `db:"business_id" json:"business_id"`

	UserID uuid.UUID `db:"user_id" json:"user_id"`

	Rating int `db:"rating" json:"rating"`

	Comment *string `db:"comment" json:"comment,omitempty"`

	IsVerified bool `db:"is_verified" json:"is_verified"`

	CreatedAt time.Time `db:"created_at" json:"created_at"`
}

type Notification struct {
	ID uuid.UUID `db:"id" json:"id"`

	UserID uuid.UUID `db:"user_id" json:"user_id"`

	Type NotificationType `db:"type" json:"type"`

	Title string `db:"title" json:"title"`

	Body *string `db:"body" json:"body,omitempty"`

	RefID *uuid.UUID `db:"ref_id" json:"ref_id,omitempty"`

	RefType *NotificationRefType `db:"ref_type" json:"ref_type,omitempty"`

	IsRead bool `db:"is_read" json:"is_read"`

	CreatedAt time.Time `db:"created_at" json:"created_at"`
}
