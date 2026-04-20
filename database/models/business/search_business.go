package business_models

import (
	"time"

	"github.com/atharvyadav96k/spotnearr-app/database/enums"
)

type SearchBusiness struct {
	ID                   string     `json:"id" firestore:"id"`
	BusinessID           string     `json:"business_id" firestore:"id"`
	Lat                  float64    `json:"lat"`
	Lng                  float64    `json:"lng"`
	Geohash6             string     `json:"geohash_6"`
	Geohash5             string     `json:"geohash_5"`
	RecommendationScore  float64    `json:"recommendationScore"`
	HasUsedFreeSpotlight bool       `json:"hasUsedFreeSpotlight,omitempty"`
	HasOffers            bool       `json:"hasOffers"`
	CategoryID           string     `json:"categoryId,omitempty"`
	SubcategoryIDs       []string   `json:"subcategoryIds,omitempty"`
	PlanExpireDate       time.Time  `json:"plan_expiry_date"`
	ActivePlanTier       enums.Plan `json:"activePlanTier,omitempty"`
}
