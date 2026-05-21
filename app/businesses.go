package app

import (
	"context"

	"github.com/atharvyadav96k/spotnearr-gcp/app/database/models"
	"github.com/google/uuid"
)

func (a *App) GetBusinessByOwnerID(ownerID uuid.UUID) (*models.Business, error) {
	query := `SELECT id, owner_id, name, description, category_id, email, phone, website, logo_url, cover_url, status, is_verified, rating, total_reviews, created_at, updated_at FROM businesses WHERE owner_id = $1`
	var b models.Business
	err := a.GetDB().QueryRow(context.Background(), query, ownerID).Scan(
		&b.ID, &b.OwnerID, &b.Name, &b.Description, &b.CategoryID,
		&b.Email, &b.Phone, &b.Website, &b.LogoURL, &b.CoverURL,
		&b.Status, &b.IsVerified, &b.Rating, &b.TotalReviews, &b.CreatedAt, &b.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &b, nil
}

func (a *App) UpdateBusinessLogo(businessID uuid.UUID, logoURL string) error {
	query := `UPDATE businesses SET logo_url = $1, updated_at = NOW() WHERE id = $2`
	_, err := a.GetDB().Exec(context.Background(), query, logoURL, businessID)
	return err
}

func (a *App) DeleteBusinessLogo(businessID uuid.UUID) error {
	query := `UPDATE businesses SET logo_url = NULL, updated_at = NOW() WHERE id = $1`
	_, err := a.GetDB().Exec(context.Background(), query, businessID)
	return err
}
