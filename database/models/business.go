package models

import (
	"time"
)

// Business represents a business in the system
type Business struct {
	ID                     string    `json:"_id" bson:"_id"`
	OwnerID                string    `json:"ownerId" bson:"ownerId"`
	Name                   string    `json:"name" bson:"name"`
	BusinessName           *string   `json:"business_name,omitempty" bson:"business_name,omitempty"`
	Lat                    float64   `json:"lat" bson:"lat"`
	Lng                    float64   `json:"lng" bson:"lng"`
	Location               *string   `json:"location,omitempty" bson:"location,omitempty"`
	ProfileImage           *string   `json:"profileImage,omitempty" bson:"profileImage,omitempty"`
	ProfileImagePublicID   *string   `json:"profileImagePublicId,omitempty" bson:"profileImagePublicId,omitempty"`
	CoverImage             *string   `json:"coverImage,omitempty" bson:"coverImage,omitempty"`
	CoverImagePublicID     *string   `json:"coverImagePublicId,omitempty" bson:"coverImagePublicId,omitempty"`
	Address                *string   `json:"address,omitempty" bson:"address,omitempty"`
	WhatsappNo             *string   `json:"whatsappNo,omitempty" bson:"whatsappNo,omitempty"`
	Bio                    *string   `json:"bio,omitempty" bson:"bio,omitempty"`
	CreatedAt              int64     `json:"createdAt" bson:"createdAt"`
	Rating                 *float64  `json:"rating,omitempty" bson:"rating,omitempty"`
	ReviewCount            *int64    `json:"reviewCount,omitempty" bson:"reviewCount,omitempty"`
	CategoryID             *string   `json:"categoryId,omitempty" bson:"categoryId,omitempty"`
	SubcategoryIDs         []string  `json:"subcategoryIds,omitempty" bson:"subcategoryIds,omitempty"`
	IsBookable             *bool     `json:"isBookable,omitempty" bson:"isBookable,omitempty"`
	HasOffers              *bool     `json:"hasOffers,omitempty" bson:"hasOffers,omitempty"`
	Geohash6               string    `json:"geohash_6" bson:"geohash_6"`
	Geohash5               string    `json:"geohash_5" bson:"geohash_5"`
	RecommendationScore    float64   `json:"recommendationScore" bson:"recommendationScore"`
	IsCelebrity            *bool     `json:"isCelebrity,omitempty" bson:"isCelebrity,omitempty"`
	ProfessionalExpiresAt  *int64    `json:"professionalExpiresAt,omitempty" bson:"professionalExpiresAt,omitempty"`
	PremiumExpiresAt       *int64    `json:"premiumExpiresAt,omitempty" bson:"premiumExpiresAt,omitempty"`
	ActivePlanTier         *string   `json:"activePlanTier,omitempty" bson:"activePlanTier,omitempty"` // NONE, PROFESSIONAL, PREMIUM
	PlanRank               *int64    `json:"planRank,omitempty" bson:"planRank,omitempty"`
	ActiveSpotlightCount   *int64    `json:"activeSpotlightCount,omitempty" bson:"activeSpotlightCount,omitempty"`
	OfferingCount          *int64    `json:"offeringCount,omitempty" bson:"offeringCount,omitempty"`
	WeeklyVisits           *int64    `json:"weeklyVisits,omitempty" bson:"weeklyVisits,omitempty"`
	FollowerCount          *int64    `json:"followerCount,omitempty" bson:"followerCount,omitempty"`
	HasUsedFreeSpotlight   *bool     `json:"hasUsedFreeSpotlight,omitempty" bson:"hasUsedFreeSpotlight,omitempty"`
	CreationTime           time.Time `json:"_creationTime" bson:"_creationTime"`
}

// PlanTier constants
const (
	PlanTierNone        = "NONE"
	PlanTierProfessional = "PROFESSIONAL"
	PlanTierPremium     = "PREMIUM"
)
