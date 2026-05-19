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

func (a *App) CreateSpotlightComment(comment models.SpotlightComment) (*models.SpotlightComment, error) {
	query := `INSERT INTO spotlight_comments (spotlight_id, user_id, parent_id, content, is_deleted) VALUES ($1, $2, $3, $4, $5) RETURNING id, created_at`

	err := a.GetDB().QueryRow(context.Background(), query, comment.SpotlightID, comment.UserID, comment.ParentID, comment.Content, comment.IsDeleted).Scan(&comment.ID, &comment.CreatedAt)

	if err != nil {
		return nil, err
	}

	return &comment, nil
}

func (a *App) UpdateSpotlightComment(comment models.SpotlightComment) (*models.SpotlightComment, error) {
	query := `UPDATE spotlight_comments SET content = $1, is_deleted = $2 WHERE id = $3`

	_, err := a.GetDB().Exec(context.Background(), query, comment.Content, comment.IsDeleted, comment.ID)

	if err != nil {
		return nil, err
	}

	return &comment, nil
}

func (a *App) DeleteSpotlightComment(id uuid.UUID) error {
	query := `DELETE FROM spotlight_comments WHERE id = $1`

	_, err := a.GetDB().Exec(context.Background(), query, id)

	return err
}

func (a *App) GetSpotlightCommentsBySpotlightID(spotlightID uuid.UUID) ([]models.SpotlightComment, error) {
	query := `SELECT id, spotlight_id, user_id, parent_id, content, is_deleted, created_at FROM spotlight_comments WHERE spotlight_id = $1 ORDER BY created_at DESC`

	rows, err := a.GetDB().Query(context.Background(), query, spotlightID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var comments []models.SpotlightComment

	for rows.Next() {
		var comment models.SpotlightComment

		err := rows.Scan(&comment.ID, &comment.SpotlightID, &comment.UserID, &comment.ParentID, &comment.Content, &comment.IsDeleted, &comment.CreatedAt)

		if err != nil {
			return nil, err
		}

		comments = append(comments, comment)
	}

	return comments, nil
}

func (a *App) CreateSpotlightFeedItem(feed models.SpotlightFeedItem) (*models.SpotlightFeedItem, error) {
	query := `INSERT INTO spotlight_feed_items (user_id, spotlight_id, is_seen, seen_at) VALUES ($1, $2, $3, $4) RETURNING id, created_at`

	err := a.GetDB().QueryRow(context.Background(), query, feed.UserID, feed.SpotlightID, feed.IsSeen, feed.SeenAt).Scan(&feed.ID, &feed.CreatedAt)

	if err != nil {
		return nil, err
	}

	return &feed, nil
}

func (a *App) UpdateSpotlightFeedItem(feed models.SpotlightFeedItem) (*models.SpotlightFeedItem, error) {
	query := `UPDATE spotlight_feed_items SET is_seen = $1, seen_at = $2 WHERE id = $3`

	_, err := a.GetDB().Exec(context.Background(), query, feed.IsSeen, feed.SeenAt, feed.ID)

	if err != nil {
		return nil, err
	}

	return &feed, nil
}

func (a *App) GetSpotlightFeedByUserID(userID uuid.UUID) ([]models.SpotlightFeedItem, error) {
	query := `SELECT id, user_id, spotlight_id, is_seen, seen_at, created_at FROM spotlight_feed_items WHERE user_id = $1 ORDER BY created_at DESC`

	rows, err := a.GetDB().Query(context.Background(), query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var feeds []models.SpotlightFeedItem

	for rows.Next() {
		var feed models.SpotlightFeedItem

		err := rows.Scan(&feed.ID, &feed.UserID, &feed.SpotlightID, &feed.IsSeen, &feed.SeenAt, &feed.CreatedAt)

		if err != nil {
			return nil, err
		}

		feeds = append(feeds, feed)
	}

	return feeds, nil
}
