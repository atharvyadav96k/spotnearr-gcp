package app

import (
	"context"

	"github.com/atharvyadav96k/spotnearr-gcp/app/database/models"
)

func (a *App) RegisterUser(user models.User) error {
	query := `INSERT INTO users (full_name, email, phone, password_hash, role) VALUES ($1, $2, $3, $4, $5)`
	_, err := a.GetDB().Exec(context.Background(), query, user.FullName, user.Email, user.Phone, user.PasswordHash, user.Role)
	return err
}

func (a *App) IsUserVerified(userID string) (bool, error) {
	query := `SELECT is_verified FROM users WHERE id = $1`
	var isVerified bool
	err := a.GetDB().QueryRow(context.Background(), query, userID).Scan(&isVerified)
	if err != nil {
		return false, err
	}
	return isVerified, nil
}

func (a *App) GetUserByEmail(email string) (*models.User, error) {
	query := `SELECT id, full_name, email, phone, password_hash, role, avatar_url, is_verified, is_active, created_at, updated_at FROM users WHERE email = $1`
	var user models.User
	err := a.GetDB().QueryRow(context.Background(), query, email).Scan(
		&user.ID,
		&user.FullName,
		&user.Email,
		&user.Phone,
		&user.PasswordHash,
		&user.Role,
		&user.AvatarURL,
		&user.IsVerified,
		&user.IsActive,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (a *App) GetUserByPhone(phone string) (*models.User, error) {
	query := `SELECT id, full_name, email, phone, password_hash, role, avatar_url, is_verified, is_active, created_at, updated_at FROM users WHERE phone = $1`
	var user models.User
	err := a.GetDB().QueryRow(context.Background(), query, phone).Scan(
		&user.ID,
		&user.FullName,
		&user.Email,
		&user.Phone,
		&user.PasswordHash,
		&user.Role,
		&user.AvatarURL,
		&user.IsVerified,
		&user.IsActive,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
