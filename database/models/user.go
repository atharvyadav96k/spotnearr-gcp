package models

import (
	"time"
)

// User represents a user in the system
type User struct {
	ID                   string    `json:"_id"`
	Email                string    `json:"email"`
	AuthUserID           string    `json:"authUserId"`
	ImageURL             *string   `json:"imageUrl,omitempty"`
	ImageURLPublicID     *string   `json:"imageUrlPublicId,omitempty"`
	Username             *string   `json:"username,omitempty"`
	TokenIdentifier      *string   `json:"tokenIdentifier,omitempty"`
	Name                 *string   `json:"name,omitempty"`
	Bio                  *string   `json:"bio,omitempty"`
	WebsiteURL           *string   `json:"websiteUrl,omitempty"`
	PushToken            *string   `json:"pushToken,omitempty"`
	LastKnownGeohash5    *string   `json:"lastKnownGeohash5,omitempty"`
	Role                 *string   `json:"role,omitempty"` // USER, SERVICE_PROVIDER, ADMIN
	ProfileCompleted     *bool     `json:"profileCompleted,omitempty"`
	PhoneNumber          *string   `json:"phoneNumber,omitempty"`
	PhoneVerified        *bool     `json:"phoneVerified,omitempty"`
	CreationTime         time.Time `json:"_creationTime"`
}

// UserRole constants
const (
	UserRoleUser            = "USER"
	UserRoleServiceProvider = "SERVICE_PROVIDER"
	UserRoleAdmin           = "ADMIN"
)
