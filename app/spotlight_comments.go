package app

import (
	"context"

	"github.com/atharvyadav96k/spotnearr-gcp/app/database/models"
	"github.com/google/uuid"
)

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
