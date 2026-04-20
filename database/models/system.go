package models

import (
	"time"
)

// SystemConfig represents system configuration
type SystemConfig struct {
	ID          string    `json:"_id" bson:"_id"`
	Key         string    `json:"key" bson:"key"`
	Value       interface{} `json:"value" bson:"value"` // number, string, or boolean
	Description *string   `json:"description,omitempty" bson:"description,omitempty"`
	UpdatedAt   int64     `json:"updatedAt" bson:"updatedAt"`
	CreationTime time.Time `json:"_creationTime" bson:"_creationTime"`
}

// Cache represents distributed cache entries
type Cache struct {
	Key          string      `json:"key" bson:"key"`
	Data         interface{} `json:"data" bson:"data"`
	Timestamp    int64       `json:"timestamp" bson:"timestamp"`
	CreationTime time.Time   `json:"_creationTime" bson:"_creationTime"`
}

// RateLimit represents rate limiting entries
type RateLimit struct {
	ID          string    `json:"_id" bson:"_id"`
	Identifier  string    `json:"identifier" bson:"identifier"` // Can be userId, IP, etc.
	Endpoint    *string   `json:"endpoint,omitempty" bson:"endpoint,omitempty"`
	Timestamp   int64     `json:"timestamp" bson:"timestamp"`
	CreationTime time.Time `json:"_creationTime" bson:"_creationTime"`
}

// WhatsAppMessage represents WhatsApp message logs
type WhatsAppMessage struct {
	ID          string    `json:"_id" bson:"_id"`
	Phone       string    `json:"phone" bson:"phone"`
	Text        string    `json:"text" bson:"text"`
	Timestamp   int64     `json:"timestamp" bson:"timestamp"`
	CreationTime time.Time `json:"_creationTime" bson:"_creationTime"`
}

// Report represents user reports
type Report struct {
	ID          string    `json:"_id" bson:"_id"`
	UserID      string    `json:"userId" bson:"userId"`
	BusinessID  *string   `json:"businessId,omitempty" bson:"businessId,omitempty"`
	SpotlightID *string   `json:"spotlightId,omitempty" bson:"spotlightId,omitempty"`
	Category    string    `json:"category" bson:"category"`
	Description *string   `json:"description,omitempty" bson:"description,omitempty"`
	Status      string    `json:"status" bson:"status"` // pending, reviewing, resolved, dismissed
	CreatedAt   int64     `json:"createdAt" bson:"createdAt"`
	ReviewedAt  *int64    `json:"reviewedAt,omitempty" bson:"reviewedAt,omitempty"`
	ReviewedBy  *string   `json:"reviewedBy,omitempty" bson:"reviewedBy,omitempty"`
	AdminNotes  *string   `json:"adminNotes,omitempty" bson:"adminNotes,omitempty"`
	CreationTime time.Time `json:"_creationTime" bson:"_creationTime"`
}

// PhoneOTP represents phone OTP verification
type PhoneOTP struct {
	ID          string    `json:"_id" bson:"_id"`
	PhoneNumber string    `json:"phoneNumber" bson:"phoneNumber"`
	Code        string    `json:"code" bson:"code"`
	Expires     int64     `json:"expires" bson:"expires"`
	RequestedAt int64     `json:"requestedAt" bson:"requestedAt"`
	CreationTime time.Time `json:"_creationTime" bson:"_creationTime"`
}

// Session represents user authentication sessions
type Session struct {
	ID               string    `json:"_id" bson:"_id"`
	UserID           string    `json:"userId" bson:"userId"`
	Token            string    `json:"token" bson:"token"`                         // short-lived session token (15 min)
	ExpiresAt        int64     `json:"expiresAt" bson:"expiresAt"`                 // session token expiry (15 min from last refresh)
	RefreshToken     *string   `json:"refreshToken,omitempty" bson:"refreshToken,omitempty"` // long-lived refresh token (90-day sliding)
	RefreshExpiresAt *int64    `json:"refreshExpiresAt,omitempty" bson:"refreshExpiresAt,omitempty"` // sliding 90-day expiry
	CreatedAt        int64     `json:"createdAt" bson:"createdAt"`
	DeviceInfo       *string   `json:"deviceInfo,omitempty" bson:"deviceInfo,omitempty"`
	CreationTime     time.Time `json:"_creationTime" bson:"_creationTime"`
}

// Report status constants
const (
	ReportStatusPending    = "pending"
	ReportStatusReviewing  = "reviewing"
	ReportStatusResolved   = "resolved"
	ReportStatusDismissed  = "dismissed"
)
