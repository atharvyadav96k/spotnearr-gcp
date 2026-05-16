package models

import (
	"time"

	"github.com/google/uuid"
)

type SpotlightType string
type SpotlightStatus string
type SpotlightMediaType string
type SpotlightInteractionType string

const (
	SpotlightProduct SpotlightType = "product"
	SpotlightOffer   SpotlightType = "offer"
	SpotlightGeneral SpotlightType = "general"

	SpotlightDraft     SpotlightStatus = "draft"
	SpotlightPublished SpotlightStatus = "published"
	SpotlightExpired   SpotlightStatus = "expired"

	MediaImage SpotlightMediaType = "image"
	MediaVideo SpotlightMediaType = "video"

	InteractionLike  SpotlightInteractionType = "like"
	InteractionShare SpotlightInteractionType = "share"
	InteractionSave  SpotlightInteractionType = "save"
)

type Spotlight struct {
	ID uuid.UUID `db:"id" json:"id"`

	BusinessID uuid.UUID `db:"business_id" json:"business_id"`

	Type   SpotlightType   `db:"type" json:"type"`
	Status SpotlightStatus `db:"status" json:"status"`

	Title       string  `db:"title" json:"title"`
	Description *string `db:"description" json:"description,omitempty"`

	MediaURL     string             `db:"media_url" json:"media_url"`
	MediaType    SpotlightMediaType `db:"media_type" json:"media_type"`
	ThumbnailURL *string            `db:"thumbnail_url" json:"thumbnail_url,omitempty"`

	ProductID *uuid.UUID `db:"product_id" json:"product_id,omitempty"`
	OfferID   *uuid.UUID `db:"offer_id" json:"offer_id,omitempty"`

	ViewCount    int `db:"view_count" json:"view_count"`
	LikeCount    int `db:"like_count" json:"like_count"`
	ShareCount   int `db:"share_count" json:"share_count"`
	CommentCount int `db:"comment_count" json:"comment_count"`

	ExpiresAt *time.Time `db:"expires_at" json:"expires_at,omitempty"`

	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

type SpotlightFeedItem struct {
	ID uuid.UUID `db:"id" json:"id"`

	UserID      uuid.UUID `db:"user_id" json:"user_id"`
	SpotlightID uuid.UUID `db:"spotlight_id" json:"spotlight_id"`

	IsSeen bool `db:"is_seen" json:"is_seen"`

	SeenAt *time.Time `db:"seen_at" json:"seen_at,omitempty"`

	CreatedAt time.Time `db:"created_at" json:"created_at"`
}

type SpotlightInteraction struct {
	ID uuid.UUID `db:"id" json:"id"`

	UserID      uuid.UUID `db:"user_id" json:"user_id"`
	SpotlightID uuid.UUID `db:"spotlight_id" json:"spotlight_id"`

	Type SpotlightInteractionType `db:"type" json:"type"`

	CreatedAt time.Time `db:"created_at" json:"created_at"`
}

type SpotlightComment struct {
	ID uuid.UUID `db:"id" json:"id"`

	SpotlightID uuid.UUID `db:"spotlight_id" json:"spotlight_id"`
	UserID      uuid.UUID `db:"user_id" json:"user_id"`

	ParentID *uuid.UUID `db:"parent_id" json:"parent_id,omitempty"`

	Content string `db:"content" json:"content"`

	IsDeleted bool `db:"is_deleted" json:"is_deleted"`

	CreatedAt time.Time `db:"created_at" json:"created_at"`
}
