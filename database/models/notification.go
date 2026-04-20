package models

import (
	"time"
)

// Notification represents a user notification
type Notification struct {
	ID         string    `json:"_id"`
	UserID     string    `json:"userId"`
	BusinessID *string   `json:"businessId,omitempty"`
	Type       string    `json:"type"`
	Title      string    `json:"title"`
	Body       string    `json:"body"`
	Data       *string   `json:"data,omitempty"` // Route data for deep linking
	IsRead     bool      `json:"isRead"`
	CreatedAt  int64     `json:"createdAt"`
	CreationTime time.Time `json:"_creationTime"`
}

// NotificationSettings represents user notification preferences
type NotificationSettings struct {
	ID               string    `json:"_id"`
	UserID           string    `json:"userId"`
	PushEnabled      bool      `json:"pushEnabled"`
	PromotionsEnabled bool     `json:"promotionsEnabled"` // Weekly digests, local digests
	SocialEnabled    bool      `json:"socialEnabled"`       // Follows and likes
	BusinessEnabled  bool      `json:"businessEnabled"`   // Alerts for their owned business
	CreationTime     time.Time `json:"_creationTime"`
}
