package models

import (
	"time"
)

// Offering represents a product or service offered by a business
type Offering struct {
	ID            string    `json:"_id" bson:"_id"`
	BusinessID    string    `json:"businessId" bson:"businessId"`
	Type          string    `json:"type" bson:"type"`         // Relaxed for dynamic support
	GroupID       *string   `json:"groupId,omitempty" bson:"groupId,omitempty"`
	Name          string    `json:"name" bson:"name"`         // "Chicken Biryani", "Haircut", "Pens"
	Description   *string   `json:"description,omitempty" bson:"description,omitempty"`
	Price         *float64  `json:"price,omitempty" bson:"price,omitempty"`
	ImageURL      *string   `json:"imageUrl,omitempty" bson:"imageUrl,omitempty"`
	ImagePublicID *string   `json:"imagePublicId,omitempty" bson:"imagePublicId,omitempty"`
	IconName      *string   `json:"iconName,omitempty" bson:"iconName,omitempty"`
	ProductTagIDs []string  `json:"productTagIds,omitempty" bson:"productTagIds,omitempty"`
	Geohash5      *string   `json:"geohash_5,omitempty" bson:"geohash_5,omitempty"`
	IsActive      bool      `json:"isActive" bson:"isActive"`
	Priority      *int64    `json:"priority,omitempty" bson:"priority,omitempty"`
	CreationTime  time.Time `json:"_creationTime" bson:"_creationTime"`
}
