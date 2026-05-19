package app

import (
	"context"

	"github.com/atharvyadav96k/spotnearr-gcp/app/database/models"
	"github.com/google/uuid"
)

func (a *App) CreateSpotlightInteraction(interaction models.SpotlightInteraction) (*models.SpotlightInteraction, error) {
	query := `INSERT INTO spotlight_interactions (user_id, spotlight_id, type) VALUES ($1, $2, $3) RETURNING id, created_at`

	err := a.GetDB().QueryRow(context.Background(), query, interaction.UserID, interaction.SpotlightID, interaction.Type).Scan(&interaction.ID, &interaction.CreatedAt)

	if err != nil {
		return nil, err
	}

	return &interaction, nil
}

func (a *App) DeleteSpotlightInteraction(id uuid.UUID) error {
	query := `DELETE FROM spotlight_interactions WHERE id = $1`

	_, err := a.GetDB().Exec(context.Background(), query, id)

	return err
}

func (a *App) GetSpotlightInteractionsBySpotlightID(spotlightID uuid.UUID) ([]models.SpotlightInteraction, error) {
	query := `SELECT id, user_id, spotlight_id, type, created_at FROM spotlight_interactions WHERE spotlight_id = $1 ORDER BY created_at DESC`

	rows, err := a.GetDB().Query(context.Background(), query, spotlightID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var interactions []models.SpotlightInteraction

	for rows.Next() {
		var interaction models.SpotlightInteraction

		err := rows.Scan(&interaction.ID, &interaction.UserID, &interaction.SpotlightID, &interaction.Type, &interaction.CreatedAt)

		if err != nil {
			return nil, err
		}

		interactions = append(interactions, interaction)
	}

	return interactions, nil
}
