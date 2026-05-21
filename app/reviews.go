package app

import (
	"context"

	"github.com/atharvyadav96k/spotnearr-gcp/app/database/models"
	"github.com/google/uuid"
)

func (a *App) CreateReview(review models.BusinessReview) (*models.BusinessReview, error) {
	query := `
		INSERT INTO business_reviews (business_id, user_id, rating, comment)
		VALUES ($1, $2, $3, $4)
		RETURNING id, is_verified, created_at`
	err := a.GetDB().QueryRow(context.Background(), query,
		review.BusinessID, review.UserID, review.Rating, review.Comment,
	).Scan(&review.ID, &review.IsVerified, &review.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &review, nil
}

func (a *App) GetReviewsByBusinessID(businessID uuid.UUID) ([]models.BusinessReview, error) {
	query := `
		SELECT id, business_id, user_id, rating, comment, is_verified, created_at
		FROM business_reviews WHERE business_id = $1 ORDER BY created_at DESC`
	rows, err := a.GetDB().Query(context.Background(), query, businessID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var reviews []models.BusinessReview
	for rows.Next() {
		var r models.BusinessReview
		if err := rows.Scan(&r.ID, &r.BusinessID, &r.UserID, &r.Rating, &r.Comment, &r.IsVerified, &r.CreatedAt); err != nil {
			return nil, err
		}
		reviews = append(reviews, r)
	}
	return reviews, rows.Err()
}
