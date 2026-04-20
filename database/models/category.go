package models

import (
	"time"
)

// Group represents a business category group
type Group struct {
	ID               string    `json:"_id"`
	Name             string    `json:"name"`             // "Food"
	Slug             string    `json:"slug"`             // "food"
	OfferingTitle    string    `json:"offeringTitle"`    // "Menu"
	OfferingAddLabel string    `json:"offeringAddLabel"` // "Add Dish"
	OfferingPlaceholder string  `json:"offeringPlaceholder"` // "e.g. Chicken Biryani"
	HomeTitle        string    `json:"homeTitle"`   // "Daily Needs" or "Food & Beverages"
	Icon             string    `json:"icon"`             // Ionicons name
	IsActive         bool      `json:"isActive"`
	Priority         int64     `json:"priority"`
	CreationTime     time.Time `json:"_creationTime"`
}

// BusinessCategory represents a business category
type BusinessCategory struct {
	ID            string    `json:"_id"`
	Name          string    `json:"name"`
	Slug          string    `json:"slug"`
	IsActive      bool      `json:"isActive"`
	CreatedAt     int64     `json:"createdAt"`
	Icon          *string   `json:"icon,omitempty"`
	Image         *string   `json:"image,omitempty"`
	ImagePublicID *string   `json:"imagePublicId,omitempty"`
	Group         *string   `json:"group,omitempty"` // general, service, food, fashion, lifestyle, shopping
	GroupID       *string   `json:"groupId,omitempty"`
	Priority      *int64    `json:"priority,omitempty"`
	CreationTime  time.Time `json:"_creationTime"`
}

// BusinessSubcategory represents a business subcategory
type BusinessSubcategory struct {
	ID          string    `json:"_id"`
	Name        string    `json:"name"`
	Slug        string    `json:"slug"`
	CategoryID  string    `json:"categoryId"`
	IsActive    bool      `json:"isActive"`
	CreatedAt   int64     `json:"createdAt"`
	Icon        *string   `json:"icon,omitempty"`
	Priority    *int64    `json:"priority,omitempty"`
	CreationTime time.Time `json:"_creationTime"`
}

// ProductTag represents a product tag for search
type ProductTag struct {
	ID            string    `json:"_id"`
	Name          string    `json:"name"`          // "pizza", "burger", "haircut", "manicure"
	Slug          string    `json:"slug"`          // "pizza", "burger", "haircut", "manicure" 
	SubcategoryID string    `json:"subcategoryId"`
	GroupID       string    `json:"groupId"`
	IsActive      bool      `json:"isActive"`
	Priority      int64     `json:"priority"`
	SearchKeywords []string `json:"searchKeywords"` // ["pizza", "piza", "pzza"]
	CreationTime  time.Time `json:"_creationTime"`
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
