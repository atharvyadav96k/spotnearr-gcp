package models

import (
	"time"
)

// Spotlight represents a spotlight/promotional content
type Spotlight struct {
	ID               string     `json:"_id" bson:"_id"`
	BusinessID       string     `json:"businessId" bson:"businessId"`
	Type             string     `json:"type" bson:"type"` // offer, new, special, update
	TypeID           *string    `json:"typeId,omitempty" bson:"typeId,omitempty"`
	CategoryGroup    *string    `json:"categoryGroup,omitempty" bson:"categoryGroup,omitempty"`
	GroupID          *string    `json:"groupId,omitempty" bson:"groupId,omitempty"`
	Priority         *int64     `json:"priority,omitempty" bson:"priority,omitempty"`
	Title            string     `json:"title" bson:"title"`
	ImageURL         *string    `json:"imageUrl,omitempty" bson:"imageUrl,omitempty"`
	ImagePublicID    *string    `json:"imagePublicId,omitempty" bson:"imagePublicId,omitempty"`
	CreatedAt        int64      `json:"createdAt" bson:"createdAt"`
	ExpiryDate       int64      `json:"expiryDate" bson:"expiryDate"`
	Geohash6         *string    `json:"geohash_6,omitempty" bson:"geohash_6,omitempty"`
	Geohash5         *string    `json:"geohash_5,omitempty" bson:"geohash_5,omitempty"`
	BusinessName     string     `json:"businessName" bson:"businessName"`
	BusinessImage    *string    `json:"businessImage,omitempty" bson:"businessImage,omitempty"`
	LikeCount        *int64     `json:"likeCount,omitempty" bson:"likeCount,omitempty"`
	TrendingScore    *float64   `json:"trendingScore,omitempty" bson:"trendingScore,omitempty"`
	LastDecayFactor  *float64   `json:"lastDecayFactor,omitempty" bson:"lastDecayFactor,omitempty"`
	LastTrendingAlert *int64    `json:"lastTrendingAlert,omitempty" bson:"lastTrendingAlert,omitempty"`
	PlanRank         *int64     `json:"planRank,omitempty" bson:"planRank,omitempty"`
	CreationTime     time.Time  `json:"_creationTime" bson:"_creationTime"`
}

// SpotlightLike represents a like on a spotlight
type SpotlightLike struct {
	ID         string    `json:"_id" bson:"_id"`
	UserID     string    `json:"userId" bson:"userId"`
	SpotlightID string   `json:"spotlightId" bson:"spotlightId"`
	CreationTime time.Time `json:"_creationTime" bson:"_creationTime"`
}

// SeenSpotlight represents when a user has seen a spotlight
type SeenSpotlight struct {
	ID          string    `json:"_id" bson:"_id"`
	UserID      string    `json:"userId" bson:"userId"`
	SpotlightID string    `json:"spotlightId" bson:"spotlightId"`
	CreatedAt   int64     `json:"createdAt" bson:"createdAt"`
	CreationTime time.Time `json:"_creationTime" bson:"_creationTime"`
}

// SpotlightType represents the type of spotlight
type SpotlightType struct {
	ID          string    `json:"_id" bson:"_id"`
	Name        string    `json:"name" bson:"name"`        // "Offer"
	Slug        string    `json:"slug" bson:"slug"`        // "offer"
	Icon        string    `json:"icon" bson:"icon"`        // Ionicons name
	Color       string    `json:"color" bson:"color"`      // tailwind color prefix
	IsActive    bool      `json:"isActive" bson:"isActive"`
	Priority    int64     `json:"priority" bson:"priority"`
	CreationTime time.Time `json:"_creationTime" bson:"_creationTime"`
}
