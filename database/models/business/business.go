package business_models

import (
	"time"
)

// Business represents a business in the system
type Business struct {
	ID                   string    `json:"_id"`
	OwnerID              string    `json:"ownerId"`
	Name                 string    `json:"name"`
	BusinessName         string    `json:"business_name,omitempty"`
	LocationID           string    `json: "location_id"`
	WhatsappNo           string    `json:"whatsappNo,omitempty"`
	IsBookable           bool      `json:"isBookable"`
	IsCelebrity          bool      `json:"isCelebrity,omitempty"`
	PlanRank             int64     `json:"planRank,omitempty"`
	ActiveSpotlightCount int64     `json:"activeSpotlightCount,omitempty"`
	OfferingCount        int64     `json:"offeringCount,omitempty"`
	CreatedAt            time.Time `json:"createdAt"`
	UpdatedAt            time.Time `json:"updatedAt"`
}
