package models

import (
	"time"

	"github.com/atharvyadav96k/spotnearr-gcp/app/utils"
	"github.com/google/uuid"
)

type ClaimStatus string

const (
	ClaimPending   ClaimStatus = "pending"
	ClaimAccepted  ClaimStatus = "accepted"
	ClaimRejected  ClaimStatus = "rejected"
	ClaimCompleted ClaimStatus = "completed"
	ClaimCancelled ClaimStatus = "cancelled"
)

type ProductClaim struct {
	ID          uuid.UUID   `db:"id"           json:"id"`
	UserID      uuid.UUID   `db:"user_id"      json:"user_id"`
	BusinessID  uuid.UUID   `db:"business_id"  json:"business_id"`
	InventoryID uuid.UUID   `db:"inventory_id" json:"inventory_id"`
	Quantity    int         `db:"quantity"     json:"quantity"`
	Status      ClaimStatus `db:"status"       json:"status"`
	Note        *string     `db:"note"         json:"note,omitempty"`
	ClaimedAt   time.Time   `db:"claimed_at"   json:"claimed_at"`
	UpdatedAt   time.Time   `db:"updated_at"   json:"updated_at"`
}

type ProductClaimResponse struct {
	ID          uuid.UUID   `json:"id"`
	UserID      uuid.UUID   `json:"user_id"`
	BusinessID  uuid.UUID   `json:"business_id"`
	InventoryID uuid.UUID   `json:"inventory_id"`
	Quantity    int         `json:"quantity"`
	Status      ClaimStatus `json:"status"`
	Note        *string     `json:"note,omitempty"`
	ClaimedAt   time.Time   `json:"claimed_at"`
	UpdatedAt   time.Time   `json:"updated_at"`
}

func (c ProductClaim) Validate() error {
	ve := &utils.ValidationErrors{}
	if c.UserID == (uuid.UUID{}) {
		ve.Add("user_id", "user_id is required")
	}
	if c.BusinessID == (uuid.UUID{}) {
		ve.Add("business_id", "business_id is required")
	}
	if c.InventoryID == (uuid.UUID{}) {
		ve.Add("inventory_id", "inventory_id is required")
	}
	if c.Quantity <= 0 {
		ve.Add("quantity", "quantity must be greater than 0")
	}
	if ve.HasErrors() {
		return ve
	}
	return nil
}

func (c ProductClaim) ToResponse() any {
	return ProductClaimResponse{
		ID:          c.ID,
		UserID:      c.UserID,
		BusinessID:  c.BusinessID,
		InventoryID: c.InventoryID,
		Quantity:    c.Quantity,
		Status:      c.Status,
		Note:        c.Note,
		ClaimedAt:   c.ClaimedAt,
		UpdatedAt:   c.UpdatedAt,
	}
}
