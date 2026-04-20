package models

import (
	"time"
)

// User represents a user in the system
type User struct {
	ID                   string    `json:"_id" bson:"_id"`
	Email                string    `json:"email" bson:"email"`
	AuthUserID           string    `json:"authUserId" bson:"authUserId"`
	ImageURL             *string   `json:"imageUrl,omitempty" bson:"imageUrl,omitempty"`
	ImageURLPublicID     *string   `json:"imageUrlPublicId,omitempty" bson:"imageUrlPublicId,omitempty"`
	Username             *string   `json:"username,omitempty" bson:"username,omitempty"`
	TokenIdentifier      *string   `json:"tokenIdentifier,omitempty" bson:"tokenIdentifier,omitempty"`
	Name                 *string   `json:"name,omitempty" bson:"name,omitempty"`
	Bio                  *string   `json:"bio,omitempty" bson:"bio,omitempty"`
	WebsiteURL           *string   `json:"websiteUrl,omitempty" bson:"websiteUrl,omitempty"`
	PushToken            *string   `json:"pushToken,omitempty" bson:"pushToken,omitempty"`
	LastKnownGeohash5    *string   `json:"lastKnownGeohash5,omitempty" bson:"lastKnownGeohash5,omitempty"`
	Role                 *string   `json:"role,omitempty" bson:"role,omitempty"` // USER, SERVICE_PROVIDER, ADMIN
	ProfileCompleted     *bool     `json:"profileCompleted,omitempty" bson:"profileCompleted,omitempty"`
	PhoneNumber          *string   `json:"phoneNumber,omitempty" bson:"phoneNumber,omitempty"`
	PhoneVerified        *bool     `json:"phoneVerified,omitempty" bson:"phoneVerified,omitempty"`
	CreationTime         time.Time `json:"_creationTime" bson:"_creationTime"`
}

// UserRole constants
const (
	UserRoleUser            = "USER"
	UserRoleServiceProvider = "SERVICE_PROVIDER"
	UserRoleAdmin           = "ADMIN"
)
