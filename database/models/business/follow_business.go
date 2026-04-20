package business_models

import "time"

type FollowBusiness struct {
	ID           string    `json:"_id"`
	UserID       string    `json:"userId"`
	BusinessID   string    `json:"businessId"`
	CreationTime time.Time `json:"_creationTime"`
}
