package models

import (
	"time"
)

// Group represents a business category group
type Group struct {
	ID               string    `json:"_id" bson:"_id"`
	Name             string    `json:"name" bson:"name"`             // "Food"
	Slug             string    `json:"slug" bson:"slug"`             // "food"
	OfferingTitle    string    `json:"offeringTitle" bson:"offeringTitle"`    // "Menu"
	OfferingAddLabel string    `json:"offeringAddLabel" bson:"offeringAddLabel"` // "Add Dish"
	OfferingPlaceholder string  `json:"offeringPlaceholder" bson:"offeringPlaceholder"` // "e.g. Chicken Biryani"
	HomeTitle        string    `json:"homeTitle" bson:"homeTitle"`   // "Daily Needs" or "Food & Beverages"
	Icon             string    `json:"icon" bson:"icon"`             // Ionicons name
	IsActive         bool      `json:"isActive" bson:"isActive"`
	Priority         int64     `json:"priority" bson:"priority"`
	CreationTime     time.Time `json:"_creationTime" bson:"_creationTime"`
}

// BusinessCategory represents a business category
type BusinessCategory struct {
	ID            string    `json:"_id" bson:"_id"`
	Name          string    `json:"name" bson:"name"`
	Slug          string    `json:"slug" bson:"slug"`
	IsActive      bool      `json:"isActive" bson:"isActive"`
	CreatedAt     int64     `json:"createdAt" bson:"createdAt"`
	Icon          *string   `json:"icon,omitempty" bson:"icon,omitempty"`
	Image         *string   `json:"image,omitempty" bson:"image,omitempty"`
	ImagePublicID *string   `json:"imagePublicId,omitempty" bson:"imagePublicId,omitempty"`
	Group         *string   `json:"group,omitempty" bson:"group,omitempty"` // general, service, food, fashion, lifestyle, shopping
	GroupID       *string   `json:"groupId,omitempty" bson:"groupId,omitempty"`
	Priority      *int64    `json:"priority,omitempty" bson:"priority,omitempty"`
	CreationTime  time.Time `json:"_creationTime" bson:"_creationTime"`
}

// BusinessSubcategory represents a business subcategory
type BusinessSubcategory struct {
	ID          string    `json:"_id" bson:"_id"`
	Name        string    `json:"name" bson:"name"`
	Slug        string    `json:"slug" bson:"slug"`
	CategoryID  string    `json:"categoryId" bson:"categoryId"`
	IsActive    bool      `json:"isActive" bson:"isActive"`
	CreatedAt   int64     `json:"createdAt" bson:"createdAt"`
	Icon        *string   `json:"icon,omitempty" bson:"icon,omitempty"`
	Priority    *int64    `json:"priority,omitempty" bson:"priority,omitempty"`
	CreationTime time.Time `json:"_creationTime" bson:"_creationTime"`
}

// ProductTag represents a product tag for search
type ProductTag struct {
	ID            string    `json:"_id" bson:"_id"`
	Name          string    `json:"name" bson:"name"`          // "pizza", "burger", "haircut", "manicure"
	Slug          string    `json:"slug" bson:"slug"`          // "pizza", "burger", "haircut", "manicure"
	SubcategoryID string    `json:"subcategoryId" bson:"subcategoryId"`
	GroupID       string    `json:"groupId" bson:"groupId"`
	IsActive      bool      `json:"isActive" bson:"isActive"`
	Priority      int64     `json:"priority" bson:"priority"`
	SearchKeywords []string `json:"searchKeywords" bson:"searchKeywords"` // ["pizza", "piza", "pzza"]
	CreationTime  time.Time `json:"_creationTime" bson:"_creationTime"`
}

// Group constants
const (
	GroupGeneral    = "general"
	GroupService    = "service"
	GroupFood       = "food"
	GroupFashion    = "fashion"
	GroupLifestyle  = "lifestyle"
	GroupShopping   = "shopping"
)
