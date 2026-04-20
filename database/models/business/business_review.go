package business_models

import (
	"time"

	"github.com/atharvyadav96k/spotnearr-app/database/enums"
)

type Review struct {
	ID          string      `json:"_id"`
	UserID      string      `json:"userId"`
	BusinessID  string      `json:"businessId"`
	Rating      enums.Stars `json:"rating"`
	Description string      `json:"content"`
	ImageURL    string      `json:"imageUrl,omitempty"`
	CreatedAt   time.Time   `json:"createdAt"`
	UpdatedAt   time.Time   `json:"updateAt"`
}
