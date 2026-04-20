package models

import (
	"time"
)

// Spotlight represents a spotlight/promotional content
type Spotlight struct {
	ID               string     `json:"_id"`
	BusinessID       string     `json:"businessId"`
	Type             string     `json:"type"` // offer, new, special, update
	TypeID           *string    `json:"typeId,omitempty"`
	CategoryGroup    *string    `json:"categoryGroup,omitempty"`
	GroupID          *string    `json:"groupId,omitempty"`
	Priority         *int64     `json:"priority,omitempty"`
	Title            string     `json:"title"`
	ImageURL         *string    `json:"imageUrl,omitempty"`
	ImagePublicID    *string    `json:"imagePublicId,omitempty"`
	CreatedAt        int64      `json:"createdAt"`
	ExpiryDate       int64      `json:"expiryDate"`
	Geohash6         *string    `json:"geohash_6,omitempty"`
	Geohash5         *string    `json:"geohash_5,omitempty"`
	BusinessName     string     `json:"businessName"`
	BusinessImage    *string    `json:"businessImage,omitempty"`
	LikeCount        *int64     `json:"likeCount,omitempty"`
	TrendingScore    *float64   `json:"trendingScore,omitempty"`
	LastDecayFactor  *float64   `json:"lastDecayFactor,omitempty"`
	LastTrendingAlert *int64    `json:"lastTrendingAlert,omitempty"`
	PlanRank         *int64     `json:"planRank,omitempty"`
	CreationTime     time.Time  `json:"_creationTime"`
}

// SpotlightLike represents a like on a spotlight
type SpotlightLike struct {
	ID         string    `json:"_id"`
	UserID     string    `json:"userId"`
	SpotlightID string   `json:"spotlightId"`
	CreationTime time.Time `json:"_creationTime"`
}

// SeenSpotlight represents when a user has seen a spotlight
type SeenSpotlight struct {
	ID          string    `json:"_id"`
	UserID      string    `json:"userId"`
	SpotlightID string    `json:"spotlightId"`
	CreatedAt   int64     `json:"createdAt"`
	CreationTime time.Time `json:"_creationTime"`
}

// SpotlightType represents the type of spotlight
type SpotlightType struct {
	ID          string    `json:"_id"`
	Name        string    `json:"name"`        // "Offer"
	Slug        string    `json:"slug"`        // "offer"
	Icon        string    `json:"icon"`        // Ionicons name
	Color       string    `json:"color"`      // tailwind color prefix
	IsActive    bool      `json:"isActive"`
	Priority    int64     `json:"priority"`
	CreationTime time.Time `json:"_creationTime"`
}
