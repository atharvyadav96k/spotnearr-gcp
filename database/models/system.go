package models

import (
	"time"
)

// SystemConfig represents system configuration
type SystemConfig struct {
	ID          string    `json:"_id"`
	Key         string    `json:"key"`
	Value       interface{} `json:"value"` // number, string, or boolean
	Description *string   `json:"description,omitempty"`
	UpdatedAt   int64     `json:"updatedAt"`
	CreationTime time.Time `json:"_creationTime"`
}

// Cache represents distributed cache entries
type Cache struct {
	Key          string      `json:"key"`
	Data         interface{} `json:"data"`
	Timestamp    int64       `json:"timestamp"`
	CreationTime time.Time   `json:"_creationTime"`
}

// RateLimit represents rate limiting entries
type RateLimit struct {
	ID          string    `json:"_id"`
	Identifier  string    `json:"identifier"` // Can be userId, IP, etc.
	Endpoint    *string   `json:"endpoint,omitempty"`
	Timestamp   int64     `json:"timestamp"`
	CreationTime time.Time `json:"_creationTime"`
}

// WhatsAppMessage represents WhatsApp message logs
type WhatsAppMessage struct {
	ID          string    `json:"_id"`
	Phone       string    `json:"phone"`
	Text        string    `json:"text"`
	Timestamp   int64     `json:"timestamp"`
	CreationTime time.Time `json:"_creationTime"`
}

// Report represents user reports
type Report struct {
	ID          string    `json:"_id"`
	UserID      string    `json:"userId"`
	BusinessID  *string   `json:"businessId,omitempty"`
	SpotlightID *string   `json:"spotlightId,omitempty"`
	Category    string    `json:"category"`
	Description *string   `json:"description,omitempty"`
	Status      string    `json:"status"` // pending, reviewing, resolved, dismissed
	CreatedAt   int64     `json:"createdAt"`
	ReviewedAt  *int64    `json:"reviewedAt,omitempty"`
	ReviewedBy  *string   `json:"reviewedBy,omitempty"`
	AdminNotes  *string   `json:"adminNotes,omitempty"`
	CreationTime time.Time `json:"_creationTime"`
}

// PhoneOTP represents phone OTP verification
type PhoneOTP struct {
	ID          string    `json:"_id"`
	PhoneNumber string    `json:"phoneNumber"`
	Code        string    `json:"code"`
	Expires     int64     `json:"expires"`
	RequestedAt int64     `json:"requestedAt"`
	CreationTime time.Time `json:"_creationTime"`
}

// Session represents user authentication sessions
type Session struct {
	ID               string    `json:"_id"`
	UserID           string    `json:"userId"`
	Token            string    `json:"token"`                         // short-lived session token (15 min)
	ExpiresAt        int64     `json:"expiresAt"`                 // session token expiry (15 min from last refresh)
	RefreshToken     *string   `json:"refreshToken,omitempty"` // long-lived refresh token (90-day sliding)
	RefreshExpiresAt *int64    `json:"refreshExpiresAt,omitempty"` // sliding 90-day expiry
	CreatedAt        int64     `json:"createdAt"`
	DeviceInfo       *string   `json:"deviceInfo,omitempty"`
	CreationTime     time.Time `json:"_creationTime"`
}

// Report status constants
const (
	ReportStatusPending    = "pending"
	ReportStatusReviewing  = "reviewing"
	ReportStatusResolved   = "resolved"
	ReportStatusDismissed  = "dismissed"
)
