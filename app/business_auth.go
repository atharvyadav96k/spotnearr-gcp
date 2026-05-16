package app

import (
	"context"

	"github.com/atharvyadav96k/spotnearr-gcp/app/database/models"
)

func (a *App) RegisterBusiness(business models.Business) (*models.Business, error) {
	query := `INSERT INTO businesses (owner_id, name, description, category_id, email, phone, website, logo_url, cover_url) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id, is_verified, rating, total_reviews, created_at, updated_at`

	err := a.GetDB().QueryRow(context.Background(), query, business.OwnerID, business.Name, business.Description, business.CategoryID, business.Email, business.Phone, business.Website, business.LogoURL, business.CoverURL).Scan(&business.ID, &business.IsVerified, &business.Rating, &business.TotalReviews, &business.CreatedAt, &business.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return &business, nil
}
