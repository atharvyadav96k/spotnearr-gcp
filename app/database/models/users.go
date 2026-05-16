package models

import (
	"time"

	"github.com/google/uuid"
)

type UserRole string

const (
	RoleCustomer UserRole = "customer"
	RoleBusiness UserRole = "business"
	RoleAdmin    UserRole = "admin"
)

type User struct {
	ID uuid.UUID `db:"id" json:"id"`

	FullName string `db:"full_name" json:"full_name"`

	Email *string `db:"email" json:"email,omitempty"`
	Phone *string `db:"phone" json:"phone,omitempty"`

	PasswordHash *string `db:"password_hash" json:"password"`

	Role UserRole `db:"role" json:"role"`

	AvatarURL *string `db:"avatar_url" json:"avatar_url,omitempty"`

	IsVerified bool `db:"is_verified" json:"is_verified"`
	IsActive   bool `db:"is_active" json:"is_active"`

	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`

	DeletedAt *time.Time `db:"deleted_at" json:"-"`
}
