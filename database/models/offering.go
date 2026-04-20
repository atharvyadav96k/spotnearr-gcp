package models

import (
	"time"
)

// Offering represents a product or service offered by a business
type Offering struct {
	ID            string    `json:"_id"`
	BusinessID    string    `json:"businessId"`
	Type          string    `json:"type"`         // Relaxed for dynamic support
	GroupID       *string   `json:"groupId,omitempty"`
	Name          string    `json:"name"`         // "Chicken Biryani", "Haircut", "Pens"
	Description   *string   `json:"description,omitempty"`
	Price         *float64  `json:"price,omitempty"`
	ImageURL      *string   `json:"imageUrl,omitempty"`
	ImagePublicID *string   `json:"imagePublicId,omitempty"`
	IconName      *string   `json:"iconName,omitempty"`
	ProductTagIDs []string  `json:"productTagIds,omitempty"`
	Geohash5      *string   `json:"geohash_5,omitempty"`
	IsActive      bool      `json:"isActive"`
	Priority      *int64    `json:"priority,omitempty"`
	CreationTime  time.Time `json:"_creationTime"`
}
