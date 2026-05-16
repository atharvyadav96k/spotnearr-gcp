package app

import (
	"context"

	"github.com/atharvyadav96k/spotnearr-gcp/app/database/models"
)

func (a *App) CreateBusinessCategory(category models.BusinessCategory) (*models.BusinessCategory, error) {
	query := `INSERT INTO business_categories (name, icon_url) VALUES ($1, $2) RETURNING id, created_at`
	err := a.GetDB().QueryRow(context.Background(), query, category.Name, category.IconURL).Scan(&category.ID, &category.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &category, nil
}

func (a *App) GetBusinessCategories() ([]models.BusinessCategory, error) {
	query := `SELECT id, name, icon_url, created_at FROM business_categories ORDER BY created_at DESC`

	rows, err := a.GetDB().Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []models.BusinessCategory

	for rows.Next() {
		var category models.BusinessCategory

		err := rows.Scan(&category.ID, &category.Name, &category.IconURL, &category.CreatedAt)
		if err != nil {
			return nil, err
		}

		categories = append(categories, category)
	}

	return categories, nil
}
