package app

import (
	"context"
	"errors"

	"github.com/atharvyadav96k/spotnearr-gcp/app/database/models"
	"github.com/google/uuid"
)

func (a *App) CreateClaim(userID, businessID, inventoryID uuid.UUID, qty int, note *string) (*models.ProductClaim, error) {
	ctx := context.Background()

	tx, err := a.GetDB().Begin(ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback(ctx)

	// Atomically reserve stock — fails if insufficient
	tag, err := tx.Exec(ctx,
		`UPDATE product_inventory SET stock = stock - $1 WHERE id = $2 AND stock >= $1`,
		qty, inventoryID,
	)
	if err != nil {
		return nil, err
	}
	if tag.RowsAffected() == 0 {
		return nil, errors.New("out of stock")
	}

	var c models.ProductClaim
	err = tx.QueryRow(ctx, `
		INSERT INTO product_claims (user_id, business_id, inventory_id, quantity, note)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, user_id, business_id, inventory_id, quantity, status, note, claimed_at, updated_at`,
		userID, businessID, inventoryID, qty, note,
	).Scan(&c.ID, &c.UserID, &c.BusinessID, &c.InventoryID, &c.Quantity, &c.Status, &c.Note, &c.ClaimedAt, &c.UpdatedAt)
	if err != nil {
		return nil, err
	}

	if err := tx.Commit(ctx); err != nil {
		return nil, err
	}
	return &c, nil
}

func (a *App) GetClaimByID(claimID uuid.UUID) (*models.ProductClaim, error) {
	var c models.ProductClaim
	err := a.GetDB().QueryRow(context.Background(), `
		SELECT id, user_id, business_id, inventory_id, quantity, status, note, claimed_at, updated_at
		FROM product_claims WHERE id = $1`,
		claimID,
	).Scan(&c.ID, &c.UserID, &c.BusinessID, &c.InventoryID, &c.Quantity, &c.Status, &c.Note, &c.ClaimedAt, &c.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &c, nil
}

func (a *App) GetClaimsByUserID(userID uuid.UUID) ([]models.ProductClaim, error) {
	rows, err := a.GetDB().Query(context.Background(), `
		SELECT id, user_id, business_id, inventory_id, quantity, status, note, claimed_at, updated_at
		FROM product_claims WHERE user_id = $1 ORDER BY claimed_at DESC`,
		userID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var claims []models.ProductClaim
	for rows.Next() {
		var c models.ProductClaim
		if err := rows.Scan(&c.ID, &c.UserID, &c.BusinessID, &c.InventoryID, &c.Quantity, &c.Status, &c.Note, &c.ClaimedAt, &c.UpdatedAt); err != nil {
			return nil, err
		}
		claims = append(claims, c)
	}
	return claims, rows.Err()
}

func (a *App) GetClaimsByBusinessID(businessID uuid.UUID) ([]models.ProductClaim, error) {
	rows, err := a.GetDB().Query(context.Background(), `
		SELECT id, user_id, business_id, inventory_id, quantity, status, note, claimed_at, updated_at
		FROM product_claims WHERE business_id = $1 ORDER BY claimed_at DESC`,
		businessID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var claims []models.ProductClaim
	for rows.Next() {
		var c models.ProductClaim
		if err := rows.Scan(&c.ID, &c.UserID, &c.BusinessID, &c.InventoryID, &c.Quantity, &c.Status, &c.Note, &c.ClaimedAt, &c.UpdatedAt); err != nil {
			return nil, err
		}
		claims = append(claims, c)
	}
	return claims, rows.Err()
}

// SetClaimStatus updates status for accept and complete transitions (no stock change).
func (a *App) SetClaimStatus(claimID uuid.UUID, status models.ClaimStatus) error {
	_, err := a.GetDB().Exec(context.Background(),
		`UPDATE product_claims SET status = $1, updated_at = NOW() WHERE id = $2`,
		status, claimID,
	)
	return err
}

// RestoreAndSetClaimStatus restores reserved stock and updates status for reject/cancel transitions.
func (a *App) RestoreAndSetClaimStatus(claimID uuid.UUID, status models.ClaimStatus) error {
	ctx := context.Background()

	tx, err := a.GetDB().Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	var inventoryID uuid.UUID
	var qty int
	err = tx.QueryRow(ctx,
		`SELECT inventory_id, quantity FROM product_claims WHERE id = $1`,
		claimID,
	).Scan(&inventoryID, &qty)
	if err != nil {
		return err
	}

	_, err = tx.Exec(ctx,
		`UPDATE product_inventory SET stock = stock + $1 WHERE id = $2`,
		qty, inventoryID,
	)
	if err != nil {
		return err
	}

	_, err = tx.Exec(ctx,
		`UPDATE product_claims SET status = $1, updated_at = NOW() WHERE id = $2`,
		status, claimID,
	)
	if err != nil {
		return err
	}

	return tx.Commit(ctx)
}
