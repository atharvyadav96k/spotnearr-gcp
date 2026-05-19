package app

import (
	"context"

	"github.com/atharvyadav96k/spotnearr-gcp/app/database/models"
	"github.com/google/uuid"
)

func (a *App) CreateSpotlightFeedItem(feed models.SpotlightFeedItem) (*models.SpotlightFeedItem, error) {
	query := `INSERT INTO spotlight_feed_items (user_id, spotlight_id, is_seen, seen_at) VALUES ($1, $2, $3, $4) RETURNING id, created_at`

	err := a.GetDB().QueryRow(context.Background(), query, feed.UserID, feed.SpotlightID, feed.IsSeen, feed.SeenAt).Scan(&feed.ID, &feed.CreatedAt)

	if err != nil {
		return nil, err
	}

	return &feed, nil
}

func (a *App) UpdateSpotlightFeedItem(feed models.SpotlightFeedItem) (*models.SpotlightFeedItem, error) {
	query := `UPDATE spotlight_feed_items SET is_seen = $1, seen_at = $2 WHERE id = $3 RETURNING user_id, spotlight_id, is_seen, seen_at, created_at`

	err := a.GetDB().QueryRow(context.Background(), query, feed.IsSeen, feed.SeenAt, feed.ID).Scan(
		&feed.UserID,
		&feed.SpotlightID,
		&feed.IsSeen,
		&feed.SeenAt,
		&feed.CreatedAt,
	)

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
