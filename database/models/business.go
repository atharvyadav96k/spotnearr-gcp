package models

import (
	"time"
)

// Business represents a business in the system
type Business struct {
	ID                     string    `json:"_id"`
	OwnerID                string    `json:"ownerId"`
	Name                   string    `json:"name"`
	BusinessName           *string   `json:"business_name,omitempty"`
	Lat                    float64   `json:"lat"`
	Lng                    float64   `json:"lng"`
	Location               *string   `json:"location,omitempty"`
	ProfileImage           *string   `json:"profileImage,omitempty"`
	ProfileImagePublicID   *string   `json:"profileImagePublicId,omitempty"`
	CoverImage             *string   `json:"coverImage,omitempty"`
	CoverImagePublicID     *string   `json:"coverImagePublicId,omitempty"`
	Address                *string   `json:"address,omitempty"`
	WhatsappNo             *string   `json:"whatsappNo,omitempty"`
	Bio                    *string   `json:"bio,omitempty"`
	CreatedAt              int64     `json:"createdAt"`
	Rating                 *float64  `json:"rating,omitempty"`
	ReviewCount            *int64    `json:"reviewCount,omitempty"`
	CategoryID             *string   `json:"categoryId,omitempty"`
	SubcategoryIDs         []string  `json:"subcategoryIds,omitempty"`
	IsBookable             *bool     `json:"isBookable,omitempty"`
	HasOffers              *bool     `json:"hasOffers,omitempty"`
	Geohash6               string    `json:"geohash_6"`
	Geohash5               string    `json:"geohash_5"`
	RecommendationScore    float64   `json:"recommendationScore"`
	IsCelebrity            *bool     `json:"isCelebrity,omitempty"`
	ProfessionalExpiresAt  *int64    `json:"professionalExpiresAt,omitempty"`
	PremiumExpiresAt       *int64    `json:"premiumExpiresAt,omitempty"`
	ActivePlanTier         *string   `json:"activePlanTier,omitempty"` // NONE, PROFESSIONAL, PREMIUM
	PlanRank               *int64    `json:"planRank,omitempty"`
	ActiveSpotlightCount   *int64    `json:"activeSpotlightCount,omitempty"`
	OfferingCount          *int64    `json:"offeringCount,omitempty"`
	WeeklyVisits           *int64    `json:"weeklyVisits,omitempty"`
	FollowerCount          *int64    `json:"followerCount,omitempty"`
	HasUsedFreeSpotlight   *bool     `json:"hasUsedFreeSpotlight,omitempty"`
	CreationTime           time.Time `json:"_creationTime"`
}

// PlanTier constants
const (
	PlanTierNone        = "NONE"
	PlanTierProfessional = "PROFESSIONAL"
	PlanTierPremium     = "PREMIUM"
)
