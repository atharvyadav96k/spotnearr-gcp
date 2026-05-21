package app

import (
	"context"

	"github.com/atharvyadav96k/spotnearr-gcp/app/database/models"
	"github.com/google/uuid"
)

func (a *App) LikeBusiness(userID, businessID uuid.UUID) error {
	query := `INSERT INTO user_liked_businesses (user_id, business_id) VALUES ($1, $2) ON CONFLICT DO NOTHING`
	_, err := a.GetDB().Exec(context.Background(), query, userID, businessID)
	return err
}

func (a *App) UnlikeBusiness(userID, businessID uuid.UUID) error {
	query := `DELETE FROM user_liked_businesses WHERE user_id = $1 AND business_id = $2`
	_, err := a.GetDB().Exec(context.Background(), query, userID, businessID)
	return err
}

func (a *App) FollowBusiness(userID, businessID uuid.UUID) error {
	query := `INSERT INTO user_follow_businesses (user_id, business_id) VALUES ($1, $2) ON CONFLICT DO NOTHING`
	_, err := a.GetDB().Exec(context.Background(), query, userID, businessID)
	return err
}

func (a *App) UnfollowBusiness(userID, businessID uuid.UUID) error {
	query := `DELETE FROM user_follow_businesses WHERE user_id = $1 AND business_id = $2`
	_, err := a.GetDB().Exec(context.Background(), query, userID, businessID)
	return err
}

func (a *App) GetFollowingBusinesses(userID uuid.UUID) ([]models.UserFollowBusiness, error) {
	query := `SELECT user_id, business_id, created_at FROM user_follow_businesses WHERE user_id = $1 ORDER BY created_at DESC`
	rows, err := a.GetDB().Query(context.Background(), query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var follows []models.UserFollowBusiness
	for rows.Next() {
		var f models.UserFollowBusiness
		if err := rows.Scan(&f.UserID, &f.BusinessID, &f.CreatedAt); err != nil {
			return nil, err
		}
		follows = append(follows, f)
	}
	return follows, rows.Err()
}

func (a *App) LikeProduct(userID, productID uuid.UUID) error {
	query := `INSERT INTO user_liked_products (user_id, product_id) VALUES ($1, $2) ON CONFLICT DO NOTHING`
	_, err := a.GetDB().Exec(context.Background(), query, userID, productID)
	return err
}

func (a *App) UnlikeProduct(userID, productID uuid.UUID) error {
	query := `DELETE FROM user_liked_products WHERE user_id = $1 AND product_id = $2`
	_, err := a.GetDB().Exec(context.Background(), query, userID, productID)
	return err
}

func (a *App) GetLikedProducts(userID uuid.UUID) ([]models.UserLikedProduct, error) {
	query := `SELECT user_id, product_id, created_at FROM user_liked_products WHERE user_id = $1 ORDER BY created_at DESC`
	rows, err := a.GetDB().Query(context.Background(), query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var likes []models.UserLikedProduct
	for rows.Next() {
		var l models.UserLikedProduct
		if err := rows.Scan(&l.UserID, &l.ProductID, &l.CreatedAt); err != nil {
			return nil, err
		}
		likes = append(likes, l)
	}
	return likes, rows.Err()
}

func (a *App) FollowUser(followerID, followingID uuid.UUID) error {
	query := `INSERT INTO user_follow_users (follower_id, following_id) VALUES ($1, $2) ON CONFLICT DO NOTHING`
	_, err := a.GetDB().Exec(context.Background(), query, followerID, followingID)
	return err
}

func (a *App) UnfollowUser(followerID, followingID uuid.UUID) error {
	query := `DELETE FROM user_follow_users WHERE follower_id = $1 AND following_id = $2`
	_, err := a.GetDB().Exec(context.Background(), query, followerID, followingID)
	return err
}

func (a *App) GetUserFollowers(userID uuid.UUID) ([]models.UserFollowUser, error) {
	query := `SELECT follower_id, following_id, created_at FROM user_follow_users WHERE following_id = $1 ORDER BY created_at DESC`

	rows, err := a.GetDB().Query(context.Background(), query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var followers []models.UserFollowUser
	for rows.Next() {
		var f models.UserFollowUser
		if err := rows.Scan(&f.FollowerID, &f.FollowingID, &f.CreatedAt); err != nil {
			return nil, err
		}
		followers = append(followers, f)
	}
	return followers, nil
}

func (a *App) GetUserFollowing(userID uuid.UUID) ([]models.UserFollowUser, error) {
	query := `SELECT follower_id, following_id, created_at FROM user_follow_users WHERE follower_id = $1 ORDER BY created_at DESC`

	rows, err := a.GetDB().Query(context.Background(), query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var following []models.UserFollowUser
	for rows.Next() {
		var f models.UserFollowUser
		if err := rows.Scan(&f.FollowerID, &f.FollowingID, &f.CreatedAt); err != nil {
			return nil, err
		}
		following = append(following, f)
	}
	return following, nil
}

func (a *App) SaveProduct(save models.UserSavedProduct) error {
	query := `INSERT INTO user_saved_products (user_id, product_id) VALUES ($1, $2) ON CONFLICT DO NOTHING`
	_, err := a.GetDB().Exec(context.Background(), query, save.UserID, save.ProductID)
	return err
}

func (a *App) UnsaveProduct(userID, productID uuid.UUID) error {
	query := `DELETE FROM user_saved_products WHERE user_id = $1 AND product_id = $2`
	_, err := a.GetDB().Exec(context.Background(), query, userID, productID)
	return err
}

func (a *App) GetSavedProducts(userID uuid.UUID) ([]models.UserSavedProduct, error) {
	query := `SELECT user_id, product_id, created_at FROM user_saved_products WHERE user_id = $1 ORDER BY created_at DESC`

	rows, err := a.GetDB().Query(context.Background(), query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var saved []models.UserSavedProduct
	for rows.Next() {
		var s models.UserSavedProduct
		if err := rows.Scan(&s.UserID, &s.ProductID, &s.CreatedAt); err != nil {
			return nil, err
		}
		saved = append(saved, s)
	}
	return saved, nil
}

func (a *App) GetLikedBusinesses(userID uuid.UUID) ([]models.UserLikedBusiness, error) {
	query := `SELECT user_id, business_id, created_at FROM user_liked_businesses WHERE user_id = $1 ORDER BY created_at DESC`
	rows, err := a.GetDB().Query(context.Background(), query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var likes []models.UserLikedBusiness
	for rows.Next() {
		var l models.UserLikedBusiness
		if err := rows.Scan(&l.UserID, &l.BusinessID, &l.CreatedAt); err != nil {
			return nil, err
		}
		likes = append(likes, l)
	}
	return likes, rows.Err()
}

func (a *App) GetSavedSpotlightsByUserID(userID uuid.UUID) ([]models.Spotlight, error) {
	query := `
		SELECT s.id, s.business_id, s.type, s.status, s.title, s.description,
		       s.media_url, s.media_type, s.thumbnail_url, s.product_id, s.offer_id,
		       s.view_count, s.like_count, s.share_count, s.comment_count,
		       s.expires_at, s.created_at, s.updated_at
		FROM spotlights s
		INNER JOIN spotlight_interactions si ON s.id = si.spotlight_id
		WHERE si.user_id = $1 AND si.type = 'save'
		ORDER BY si.created_at DESC`

	rows, err := a.GetDB().Query(context.Background(), query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var spotlights []models.Spotlight
	for rows.Next() {
		var s models.Spotlight
		if err := rows.Scan(
			&s.ID, &s.BusinessID, &s.Type, &s.Status, &s.Title, &s.Description,
			&s.MediaURL, &s.MediaType, &s.ThumbnailURL, &s.ProductID, &s.OfferID,
			&s.ViewCount, &s.LikeCount, &s.ShareCount, &s.CommentCount,
			&s.ExpiresAt, &s.CreatedAt, &s.UpdatedAt,
		); err != nil {
			return nil, err
		}
		spotlights = append(spotlights, s)
	}
	return spotlights, nil
}
