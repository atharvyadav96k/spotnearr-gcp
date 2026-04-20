package models

import (
	"time"
)

// SubscriptionPlan represents available subscription plans
type SubscriptionPlan struct {
	ID            string    `json:"_id" bson:"_id"`
	Tier          *string   `json:"tier,omitempty" bson:"tier,omitempty"` // PROFESSIONAL, PREMIUM
	Duration      int64     `json:"duration" bson:"duration"`             // 3, 6, 12 (months)
	Price         int64     `json:"price" bson:"price"`                   // Price in paise
	MaxSpotlights int64     `json:"maxSpotlights" bson:"maxSpotlights"`
	MaxOfferings  int64     `json:"maxOfferings" bson:"maxOfferings"`
	Rank          int64     `json:"rank" bson:"rank"`                     // 2=Premium, 1=Professional
	IsActive      bool      `json:"isActive" bson:"isActive"`
	CreationTime  time.Time `json:"_creationTime" bson:"_creationTime"`
}

// Subscription represents a business subscription
type Subscription struct {
	ID         string    `json:"_id" bson:"_id"`
	BusinessID string    `json:"businessId" bson:"businessId"`
	PlanTier   string    `json:"planTier" bson:"planTier"`   // PROFESSIONAL, PREMIUM
	Duration   int64     `json:"duration" bson:"duration"`   // 3, 6, 12 (months)
	Price      int64     `json:"price" bson:"price"`       // Price paid in paise
	PaymentID  string    `json:"paymentId" bson:"paymentId"` // Razorpay payment ID
	OrderID    string    `json:"orderId" bson:"orderId"`     // Razorpay order ID
	Status     string    `json:"status" bson:"status"`       // pending, active, expired, cancelled
	StartedAt  int64     `json:"startedAt" bson:"startedAt"`
	ExpiresAt  int64     `json:"expiresAt" bson:"expiresAt"`
	CreatedAt  int64     `json:"createdAt" bson:"createdAt"`
	CreationTime time.Time `json:"_creationTime" bson:"_creationTime"`
}

// Subscription status constants
const (
	SubscriptionStatusPending   = "pending"
	SubscriptionStatusActive    = "active"
	SubscriptionStatusExpired   = "expired"
	SubscriptionStatusCancelled = "cancelled"
)
