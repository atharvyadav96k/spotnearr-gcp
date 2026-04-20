package models

import (
	"time"
)

// Notification represents a user notification
type Notification struct {
	ID         string    `json:"_id" bson:"_id"`
	UserID     string    `json:"userId" bson:"userId"`
	BusinessID *string   `json:"businessId,omitempty" bson:"businessId,omitempty"`
	Type       string    `json:"type" bson:"type"`
	Title      string    `json:"title" bson:"title"`
	Body       string    `json:"body" bson:"body"`
	Data       *string   `json:"data,omitempty" bson:"data,omitempty"` // Route data for deep linking
	IsRead     bool      `json:"isRead" bson:"isRead"`
	CreatedAt  int64     `json:"createdAt" bson:"createdAt"`
	CreationTime time.Time `json:"_creationTime" bson:"_creationTime"`
}

// NotificationSettings represents user notification preferences
type NotificationSettings struct {
	ID               string    `json:"_id" bson:"_id"`
	UserID           string    `json:"userId" bson:"userId"`
	PushEnabled      bool      `json:"pushEnabled" bson:"pushEnabled"`
	PromotionsEnabled bool     `json:"promotionsEnabled" bson:"promotionsEnabled"` // Weekly digests, local digests
	SocialEnabled    bool      `json:"socialEnabled" bson:"socialEnabled"`       // Follows and likes
	BusinessEnabled  bool      `json:"businessEnabled" bson:"businessEnabled"`   // Alerts for their owned business
	CreationTime     time.Time `json:"_creationTime" bson:"_creationTime"`
}
