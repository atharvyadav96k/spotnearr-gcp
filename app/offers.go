package app

import (
	"context"

	"github.com/atharvyadav96k/spotnearr-gcp/app/database/models"
	"github.com/google/uuid"
)

func (a *App) CreateOffer(offer models.Offer) (*models.Offer, error) {
	query := `
		INSERT INTO offers (business_id, title, description, discount_type, discount_value, min_order_value, coupon_code, banner_url, is_active, starts_at, expires_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
		RETURNING id, created_at, updated_at`
	err := a.GetDB().QueryRow(context.Background(), query,
		offer.BusinessID, offer.Title, offer.Description,
		offer.DiscountType, offer.DiscountValue, offer.MinOrderValue,
		offer.CouponCode, offer.BannerURL, offer.IsActive,
		offer.StartsAt, offer.ExpiresAt,
	).Scan(&offer.ID, &offer.CreatedAt, &offer.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &offer, nil
}

func (a *App) GetOffersByBusinessID(businessID uuid.UUID) ([]models.Offer, error) {
	query := `
		SELECT id, business_id, title, description, discount_type, discount_value, min_order_value,
		       coupon_code, banner_url, is_active, starts_at, expires_at, created_at, updated_at
		FROM offers WHERE business_id = $1 ORDER BY created_at DESC`
	rows, err := a.GetDB().Query(context.Background(), query, businessID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var offers []models.Offer
	for rows.Next() {
		var o models.Offer
		if err := rows.Scan(
			&o.ID, &o.BusinessID, &o.Title, &o.Description,
			&o.DiscountType, &o.DiscountValue, &o.MinOrderValue,
			&o.CouponCode, &o.BannerURL, &o.IsActive,
			&o.StartsAt, &o.ExpiresAt, &o.CreatedAt, &o.UpdatedAt,
		); err != nil {
			return nil, err
		}
		offers = append(offers, o)
	}
	return offers, rows.Err()
}

func (a *App) UpdateOffer(offer models.Offer) (*models.Offer, error) {
	query := `
		UPDATE offers
		SET title = $1, description = $2, discount_type = $3, discount_value = $4,
		    min_order_value = $5, coupon_code = $6, banner_url = $7, is_active = $8,
		    starts_at = $9, expires_at = $10, updated_at = NOW()
		WHERE id = $11 AND business_id = $12
		RETURNING updated_at`
	err := a.GetDB().QueryRow(context.Background(), query,
		offer.Title, offer.Description, offer.DiscountType, offer.DiscountValue,
		offer.MinOrderValue, offer.CouponCode, offer.BannerURL, offer.IsActive,
		offer.StartsAt, offer.ExpiresAt, offer.ID, offer.BusinessID,
	).Scan(&offer.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &offer, nil
}

func (a *App) DeleteOffer(id, businessID uuid.UUID) error {
	_, err := a.GetDB().Exec(context.Background(),
		`DELETE FROM offers WHERE id = $1 AND business_id = $2`, id, businessID)
	return err
}

func (a *App) GetActiveOffersByBusinessID(businessID uuid.UUID) ([]models.Offer, error) {
	query := `
		SELECT id, business_id, title, description, discount_type, discount_value, min_order_value,
		       coupon_code, banner_url, is_active, starts_at, expires_at, created_at, updated_at
		FROM offers
		WHERE business_id = $1 AND is_active = true AND expires_at > NOW()
		ORDER BY created_at DESC`
	rows, err := a.GetDB().Query(context.Background(), query, businessID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var offers []models.Offer
	for rows.Next() {
		var o models.Offer
		if err := rows.Scan(
			&o.ID, &o.BusinessID, &o.Title, &o.Description,
			&o.DiscountType, &o.DiscountValue, &o.MinOrderValue,
			&o.CouponCode, &o.BannerURL, &o.IsActive,
			&o.StartsAt, &o.ExpiresAt, &o.CreatedAt, &o.UpdatedAt,
		); err != nil {
			return nil, err
		}
		offers = append(offers, o)
	}
	return offers, rows.Err()
}
