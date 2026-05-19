package app

import (
	"context"

	"github.com/atharvyadav96k/spotnearr-gcp/app/database/models"
	"github.com/google/uuid"
)

func (a *App) CreateSpotlight(spotlight models.Spotlight) (*models.Spotlight, error) {
	query := `INSERT INTO spotlights (business_id, type, status, title, description, media_url, media_type, thumbnail_url, product_id, offer_id, expires_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) RETURNING id, view_count, like_count, share_count, comment_count, created_at, updated_at`

	err := a.GetDB().QueryRow(context.Background(), query, spotlight.BusinessID, spotlight.Type, spotlight.Status, spotlight.Title, spotlight.Description, spotlight.MediaURL, spotlight.MediaType, spotlight.ThumbnailURL, spotlight.ProductID, spotlight.OfferID, spotlight.ExpiresAt).Scan(&spotlight.ID, &spotlight.ViewCount, &spotlight.LikeCount, &spotlight.ShareCount, &spotlight.CommentCount, &spotlight.CreatedAt, &spotlight.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return &spotlight, nil
}

func (a *App) UpdateSpotlight(spotlight models.Spotlight) (*models.Spotlight, error) {
	query := `UPDATE spotlights SET type = $1, status = $2, title = $3, description = $4, media_url = $5, media_type = $6, thumbnail_url = $7, product_id = $8, offer_id = $9, expires_at = $10, updated_at = NOW() WHERE id = $11 RETURNING updated_at`

	err := a.GetDB().QueryRow(context.Background(), query, spotlight.Type, spotlight.Status, spotlight.Title, spotlight.Description, spotlight.MediaURL, spotlight.MediaType, spotlight.ThumbnailURL, spotlight.ProductID, spotlight.OfferID, spotlight.ExpiresAt, spotlight.ID).Scan(&spotlight.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return &spotlight, nil
}

func (a *App) DeleteSpotlight(id uuid.UUID) error {
	query := `DELETE FROM spotlights WHERE id = $1`

	_, err := a.GetDB().Exec(context.Background(), query, id)

	return err
}

func (a *App) GetSpotlightByID(id uuid.UUID) (*models.Spotlight, error) {
	query := `SELECT id, business_id, type, status, title, description, media_url, media_type, thumbnail_url, product_id, offer_id, view_count, like_count, share_count, comment_count, expires_at, created_at, updated_at FROM spotlights WHERE id = $1`

	var spotlight models.Spotlight

	err := a.GetDB().QueryRow(context.Background(), query, id).Scan(&spotlight.ID, &spotlight.BusinessID, &spotlight.Type, &spotlight.Status, &spotlight.Title, &spotlight.Description, &spotlight.MediaURL, &spotlight.MediaType, &spotlight.ThumbnailURL, &spotlight.ProductID, &spotlight.OfferID, &spotlight.ViewCount, &spotlight.LikeCount, &spotlight.ShareCount, &spotlight.CommentCount, &spotlight.ExpiresAt, &spotlight.CreatedAt, &spotlight.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return &spotlight, nil
}

func (a *App) GetSpotlightsByBusinessID(businessID uuid.UUID) ([]models.Spotlight, error) {
	query := `SELECT id, business_id, type, status, title, description, media_url, media_type, thumbnail_url, product_id, offer_id, view_count, like_count, share_count, comment_count, expires_at, created_at, updated_at FROM spotlights WHERE business_id = $1 ORDER BY created_at DESC`

	rows, err := a.GetDB().Query(context.Background(), query, businessID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var spotlights []models.Spotlight

	for rows.Next() {
		var spotlight models.Spotlight

		err := rows.Scan(&spotlight.ID, &spotlight.BusinessID, &spotlight.Type, &spotlight.Status, &spotlight.Title, &spotlight.Description, &spotlight.MediaURL, &spotlight.MediaType, &spotlight.ThumbnailURL, &spotlight.ProductID, &spotlight.OfferID, &spotlight.ViewCount, &spotlight.LikeCount, &spotlight.ShareCount, &spotlight.CommentCount, &spotlight.ExpiresAt, &spotlight.CreatedAt, &spotlight.UpdatedAt)

		if err != nil {
			return nil, err
		}

		spotlights = append(spotlights, spotlight)
	}

	return spotlights, nil
}
