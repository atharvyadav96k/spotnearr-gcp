package models

import (
	"time"

	"github.com/atharvyadav96k/spotnearr-gcp/app/utils"
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

type UserFollowUser struct {
	FollowerID uuid.UUID `db:"follower_id" json:"follower_id"`

	FollowingID uuid.UUID `db:"following_id" json:"following_id"`

	CreatedAt time.Time `db:"created_at" json:"created_at"`
}

type UserSavedProduct struct {
	UserID uuid.UUID `db:"user_id" json:"user_id"`

	ProductID uuid.UUID `db:"product_id" json:"product_id"`

	CreatedAt time.Time `db:"created_at" json:"created_at"`
}

type UserLikedProduct struct {
	UserID    uuid.UUID `db:"user_id" json:"user_id"`
	ProductID uuid.UUID `db:"product_id" json:"product_id"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
}

type UserLikedBusiness struct {
	UserID     uuid.UUID `db:"user_id" json:"user_id"`
	BusinessID uuid.UUID `db:"business_id" json:"business_id"`
	CreatedAt  time.Time `db:"created_at" json:"created_at"`
}

type LikedBusinessResponse struct {
	BusinessID uuid.UUID `json:"business_id"`
	CreatedAt  time.Time `json:"created_at"`
}

func (l UserLikedBusiness) Validate() error {
	ve := &utils.ValidationErrors{}
	if l.UserID == (uuid.UUID{}) {
		ve.Add("user_id", "user_id is required")
	}
	if l.BusinessID == (uuid.UUID{}) {
		ve.Add("business_id", "business_id is required")
	}
	if ve.HasErrors() {
		return ve
	}
	return nil
}

func (l UserLikedBusiness) ToResponse() any {
	return LikedBusinessResponse{BusinessID: l.BusinessID, CreatedAt: l.CreatedAt}
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

type FollowBusinessResponse struct {
	BusinessID uuid.UUID `json:"business_id"`
	CreatedAt  time.Time `json:"created_at"`
}

func (f UserFollowBusiness) Validate() error {
	ve := &utils.ValidationErrors{}
	if f.UserID == (uuid.UUID{}) {
		ve.Add("user_id", "user_id is required")
	}
	if f.BusinessID == (uuid.UUID{}) {
		ve.Add("business_id", "business_id is required")
	}
	if ve.HasErrors() {
		return ve
	}
	return nil
}

func (f UserFollowBusiness) ToResponse() any {
	return FollowBusinessResponse{BusinessID: f.BusinessID, CreatedAt: f.CreatedAt}
}

type UserFollowResponse struct {
	FollowerID  uuid.UUID `json:"follower_id"`
	FollowingID uuid.UUID `json:"following_id"`
	CreatedAt   time.Time `json:"created_at"`
}

func (f UserFollowUser) Validate() error {
	ve := &utils.ValidationErrors{}
	if f.FollowerID == (uuid.UUID{}) {
		ve.Add("follower_id", "follower_id is required")
	}
	if f.FollowingID == (uuid.UUID{}) {
		ve.Add("following_id", "following_id is required")
	}
	if ve.HasErrors() {
		return ve
	}
	return nil
}

func (f UserFollowUser) ToResponse() any {
	return UserFollowResponse{FollowerID: f.FollowerID, FollowingID: f.FollowingID, CreatedAt: f.CreatedAt}
}

type SavedProductResponse struct {
	ProductID uuid.UUID `json:"product_id"`
	CreatedAt time.Time `json:"created_at"`
}

func (s UserSavedProduct) Validate() error {
	ve := &utils.ValidationErrors{}
	if s.UserID == (uuid.UUID{}) {
		ve.Add("user_id", "user_id is required")
	}
	if s.ProductID == (uuid.UUID{}) {
		ve.Add("product_id", "product_id is required")
	}
	if ve.HasErrors() {
		return ve
	}
	return nil
}

func (s UserSavedProduct) ToResponse() any {
	return SavedProductResponse{ProductID: s.ProductID, CreatedAt: s.CreatedAt}
}

type LikedProductResponse struct {
	ProductID uuid.UUID `json:"product_id"`
	CreatedAt time.Time `json:"created_at"`
}

func (l UserLikedProduct) Validate() error {
	ve := &utils.ValidationErrors{}
	if l.UserID == (uuid.UUID{}) {
		ve.Add("user_id", "user_id is required")
	}
	if l.ProductID == (uuid.UUID{}) {
		ve.Add("product_id", "product_id is required")
	}
	if ve.HasErrors() {
		return ve
	}
	return nil
}

func (l UserLikedProduct) ToResponse() any {
	return LikedProductResponse{ProductID: l.ProductID, CreatedAt: l.CreatedAt}
}

type ReviewResponse struct {
	ID         uuid.UUID `json:"id"`
	BusinessID uuid.UUID `json:"business_id"`
	UserID     uuid.UUID `json:"user_id"`
	Rating     int       `json:"rating"`
	Comment    *string   `json:"comment,omitempty"`
	IsVerified bool      `json:"is_verified"`
	CreatedAt  time.Time `json:"created_at"`
}

func (r BusinessReview) Validate() error {
	ve := &utils.ValidationErrors{}
	if r.BusinessID == (uuid.UUID{}) {
		ve.Add("business_id", "business_id is required")
	}
	if r.UserID == (uuid.UUID{}) {
		ve.Add("user_id", "user_id is required")
	}
	if r.Rating < 1 || r.Rating > 5 {
		ve.Add("rating", "rating must be between 1 and 5")
	}
	if ve.HasErrors() {
		return ve
	}
	return nil
}

func (r BusinessReview) ToResponse() any {
	return ReviewResponse{
		ID: r.ID, BusinessID: r.BusinessID, UserID: r.UserID,
		Rating: r.Rating, Comment: r.Comment, IsVerified: r.IsVerified, CreatedAt: r.CreatedAt,
	}
}

type NotificationResponse struct {
	ID        uuid.UUID            `json:"id"`
	Type      NotificationType     `json:"type"`
	Title     string               `json:"title"`
	Body      *string              `json:"body,omitempty"`
	RefID     *uuid.UUID           `json:"ref_id,omitempty"`
	RefType   *NotificationRefType `json:"ref_type,omitempty"`
	IsRead    bool                 `json:"is_read"`
	CreatedAt time.Time            `json:"created_at"`
}

func (n Notification) ToResponse() any {
	return NotificationResponse{
		ID: n.ID, Type: n.Type, Title: n.Title, Body: n.Body,
		RefID: n.RefID, RefType: n.RefType, IsRead: n.IsRead, CreatedAt: n.CreatedAt,
	}
}
