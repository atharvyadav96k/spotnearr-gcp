package app

import (
	"context"

	"github.com/atharvyadav96k/spotnearr-gcp/app/database/models"
	"github.com/google/uuid"
)

func (a *App) AddPlan(plan models.Plan) (*models.Plan, error) {
	query := `
		INSERT INTO plans (name, tier, price, description)
		VALUES ($1, $2, $3, $4)
		RETURNING id, is_active, created_at`

	err := a.GetDB().QueryRow(context.Background(), query,
		plan.Name, plan.Tier, plan.Price, plan.Description,
	).Scan(&plan.ID, &plan.IsActive, &plan.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &plan, nil
}

func (a *App) GetPlans() ([]models.Plan, error) {
	query := `SELECT id, name, tier, price, description, is_active, created_at FROM plans WHERE is_active = true ORDER BY price ASC`

	rows, err := a.GetDB().Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var plans []models.Plan
	for rows.Next() {
		var p models.Plan
		if err := rows.Scan(&p.ID, &p.Name, &p.Tier, &p.Price, &p.Description, &p.IsActive, &p.CreatedAt); err != nil {
			return nil, err
		}
		plans = append(plans, p)
	}
	return plans, rows.Err()
}

func (a *App) GetBusinessActivePlan(businessID uuid.UUID) (*models.BusinessSubscription, error) {
	query := `
		SELECT
			bs.id, bs.business_id, bs.plan_id, bs.tier,
			p.name, p.price,
			bs.started_at, bs.expires_at, bs.created_at
		FROM business_subscriptions bs
		LEFT JOIN plans p ON p.id = bs.plan_id
		WHERE bs.business_id = $1`

	var s models.BusinessSubscription
	err := a.GetDB().QueryRow(context.Background(), query, businessID).Scan(
		&s.ID, &s.BusinessID, &s.PlanID, &s.Tier,
		&s.PlanName, &s.Price,
		&s.StartedAt, &s.ExpiresAt, &s.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &s, nil
}

func (a *App) SetBusinessActivePlan(businessID, planID uuid.UUID) error {
	query := `
		UPDATE business_subscriptions bs
		SET
			plan_id   = p.id,
			tier      = p.tier,
			started_at = NOW(),
			expires_at = NOW() + INTERVAL '30 days'
		FROM plans p
		WHERE p.id = $2
		  AND bs.business_id = $1`

	_, err := a.GetDB().Exec(context.Background(), query, businessID, planID)
	return err
}
