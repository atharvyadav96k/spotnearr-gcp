package models

import (
	"strings"
	"time"

	"github.com/atharvyadav96k/spotnearr-gcp/app/utils"
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

type SpotlightResponse struct {
	ID           uuid.UUID          `json:"id"`
	BusinessID   uuid.UUID          `json:"business_id"`
	Type         SpotlightType      `json:"type"`
	Status       SpotlightStatus    `json:"status"`
	Title        string             `json:"title"`
	Description  *string            `json:"description,omitempty"`
	MediaURL     string             `json:"media_url"`
	MediaType    SpotlightMediaType `json:"media_type"`
	ThumbnailURL *string            `json:"thumbnail_url,omitempty"`
	ProductID    *uuid.UUID         `json:"product_id,omitempty"`
	OfferID      *uuid.UUID         `json:"offer_id,omitempty"`
	ViewCount    int                `json:"view_count"`
	LikeCount    int                `json:"like_count"`
	ShareCount   int                `json:"share_count"`
	CommentCount int                `json:"comment_count"`
	ExpiresAt    *time.Time         `json:"expires_at,omitempty"`
	CreatedAt    time.Time          `json:"created_at"`
}

func (s Spotlight) Validate() error {
	ve := &utils.ValidationErrors{}
	if s.BusinessID == (uuid.UUID{}) {
		ve.Add("business_id", "business_id is required")
	}
	if s.Type == "" {
		ve.Add("type", "type is required (product|offer|general)")
	}
	if strings.TrimSpace(s.Title) == "" {
		ve.Add("title", "title is required")
	}
	if strings.TrimSpace(s.MediaURL) == "" {
		ve.Add("media_url", "media_url is required")
	}
	if s.MediaType == "" {
		ve.Add("media_type", "media_type is required (image|video)")
	}
	if ve.HasErrors() {
		return ve
	}
	return nil
}

func (s Spotlight) ToResponse() any {
	return SpotlightResponse{
		ID: s.ID, BusinessID: s.BusinessID, Type: s.Type, Status: s.Status,
		Title: s.Title, Description: s.Description, MediaURL: s.MediaURL, MediaType: s.MediaType,
		ThumbnailURL: s.ThumbnailURL, ProductID: s.ProductID, OfferID: s.OfferID,
		ViewCount: s.ViewCount, LikeCount: s.LikeCount, ShareCount: s.ShareCount,
		CommentCount: s.CommentCount, ExpiresAt: s.ExpiresAt, CreatedAt: s.CreatedAt,
	}
}

type SpotlightFeedItemResponse struct {
	ID          uuid.UUID  `json:"id"`
	SpotlightID uuid.UUID  `json:"spotlight_id"`
	IsSeen      bool       `json:"is_seen"`
	SeenAt      *time.Time `json:"seen_at,omitempty"`
	CreatedAt   time.Time  `json:"created_at"`
}

func (s SpotlightFeedItem) ToResponse() any {
	return SpotlightFeedItemResponse{ID: s.ID, SpotlightID: s.SpotlightID, IsSeen: s.IsSeen, SeenAt: s.SeenAt, CreatedAt: s.CreatedAt}
}

type SpotlightInteractionResponse struct {
	ID          uuid.UUID                `json:"id"`
	SpotlightID uuid.UUID                `json:"spotlight_id"`
	Type        SpotlightInteractionType `json:"type"`
	CreatedAt   time.Time                `json:"created_at"`
}

func (s SpotlightInteraction) ToResponse() any {
	return SpotlightInteractionResponse{ID: s.ID, SpotlightID: s.SpotlightID, Type: s.Type, CreatedAt: s.CreatedAt}
}

type SpotlightCommentResponse struct {
	ID          uuid.UUID  `json:"id"`
	SpotlightID uuid.UUID  `json:"spotlight_id"`
	UserID      uuid.UUID  `json:"user_id"`
	ParentID    *uuid.UUID `json:"parent_id,omitempty"`
	Content     string     `json:"content"`
	CreatedAt   time.Time  `json:"created_at"`
}

func (c SpotlightComment) Validate() error {
	ve := &utils.ValidationErrors{}
	if c.SpotlightID == (uuid.UUID{}) {
		ve.Add("spotlight_id", "spotlight_id is required")
	}
	if c.UserID == (uuid.UUID{}) {
		ve.Add("user_id", "user_id is required")
	}
	if strings.TrimSpace(c.Content) == "" {
		ve.Add("content", "content is required")
	}
	if ve.HasErrors() {
		return ve
	}
	return nil
}

func (c SpotlightComment) ToResponse() any {
	return SpotlightCommentResponse{
		ID: c.ID, SpotlightID: c.SpotlightID, UserID: c.UserID,
		ParentID: c.ParentID, Content: c.Content, CreatedAt: c.CreatedAt,
	}
}
