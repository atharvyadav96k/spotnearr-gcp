package app

import (
	"context"

	"github.com/atharvyadav96k/spotnearr-gcp/app/database/models"
	"github.com/google/uuid"
)

func (a *App) CreateNotification(n models.Notification) (*models.Notification, error) {
	query := `
		INSERT INTO notifications (user_id, type, title, body, ref_id, ref_type)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id, is_read, created_at`
	err := a.GetDB().QueryRow(context.Background(), query,
		n.UserID, n.Type, n.Title, n.Body, n.RefID, n.RefType,
	).Scan(&n.ID, &n.IsRead, &n.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &n, nil
}

func (a *App) GetNotificationsByUserID(userID uuid.UUID) ([]models.Notification, error) {
	query := `
		SELECT id, user_id, type, title, body, ref_id, ref_type, is_read, created_at
		FROM notifications WHERE user_id = $1 ORDER BY created_at DESC`
	rows, err := a.GetDB().Query(context.Background(), query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var notifications []models.Notification
	for rows.Next() {
		var n models.Notification
		if err := rows.Scan(
			&n.ID, &n.UserID, &n.Type, &n.Title, &n.Body,
			&n.RefID, &n.RefType, &n.IsRead, &n.CreatedAt,
		); err != nil {
			return nil, err
		}
		notifications = append(notifications, n)
	}
	return notifications, rows.Err()
}

func (a *App) MarkNotificationRead(id, userID uuid.UUID) error {
	_, err := a.GetDB().Exec(context.Background(),
		`UPDATE notifications SET is_read = true WHERE id = $1 AND user_id = $2`, id, userID)
	return err
}

func (a *App) MarkAllNotificationsRead(userID uuid.UUID) error {
	_, err := a.GetDB().Exec(context.Background(),
		`UPDATE notifications SET is_read = true WHERE user_id = $1 AND is_read = false`, userID)
	return err
}
