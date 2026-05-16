package app

import (
	"context"

	"github.com/atharvyadav96k/spotnearr-gcp/app/database/models"
	"github.com/atharvyadav96k/spotnearr-gcp/app/utils"
	"github.com/google/uuid"
)

func (a *App) CreateBusinessLocation(location models.BusinessLocation) (*models.BusinessLocation, error) {
	location.Geohash = utils.GenerateGeohash(location.Latitude, location.Longitude, 7)

	query := `INSERT INTO business_locations (business_id, branch_name, address_line1, address_line2, city, state, pin_code, latitude, longitude, geohash, is_main, opening_time, closing_time, working_days) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14) RETURNING id, created_at, updated_at`

	err := a.GetDB().QueryRow(context.Background(), query, location.BusinessID, location.BranchName, location.AddressLine1, location.AddressLine2, location.City, location.State, location.PinCode, location.Latitude, location.Longitude, location.Geohash, location.IsMain, location.OpeningTime, location.ClosingTime, location.WorkingDays).Scan(&location.ID, &location.CreatedAt, &location.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return &location, nil
}

func (a *App) GetBusinessesNearby(latitude float64, longitude float64, density int) ([]models.BusinessLocation, error) {
	radius := float64(density * 3)

	query := `SELECT * FROM (SELECT id, business_id, branch_name, address_line1, address_line2, city, state, pin_code, latitude, longitude, geohash, is_main, opening_time, closing_time, working_days, created_at, updated_at, (6371 * acos(cos(radians($1)) * cos(radians(latitude)) * cos(radians(longitude) - radians($2)) + sin(radians($1)) * sin(radians(latitude)))) AS distance FROM business_locations) nearby WHERE distance <= $3 AND CURRENT_TIME BETWEEN opening_time AND closing_time AND LOWER(TRIM(TO_CHAR(CURRENT_DATE, 'Day'))) = ANY(working_days) ORDER BY distance ASC`

	rows, err := a.GetDB().Query(context.Background(), query, latitude, longitude, radius)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var locations []models.BusinessLocation

	for rows.Next() {
		var location models.BusinessLocation
		var distance float64

		err := rows.Scan(&location.ID, &location.BusinessID, &location.BranchName, &location.AddressLine1, &location.AddressLine2, &location.City, &location.State, &location.PinCode, &location.Latitude, &location.Longitude, &location.Geohash, &location.IsMain, &location.OpeningTime, &location.ClosingTime, &location.WorkingDays, &location.CreatedAt, &location.UpdatedAt, &distance)

		if err != nil {
			return nil, err
		}

		locations = append(locations, location)
	}

	return locations, nil
}

func (a *App) GetBusinessLocationByID(id uuid.UUID) (*models.BusinessLocation, error) {
	query := `SELECT id, business_id, branch_name, address_line1, address_line2, city, state, pin_code, latitude, longitude, geohash, is_main, opening_time, closing_time, working_days, created_at, updated_at FROM business_locations WHERE id = $1`

	var location models.BusinessLocation

	err := a.GetDB().QueryRow(context.Background(), query, id).Scan(&location.ID, &location.BusinessID, &location.BranchName, &location.AddressLine1, &location.AddressLine2, &location.City, &location.State, &location.PinCode, &location.Latitude, &location.Longitude, &location.Geohash, &location.IsMain, &location.OpeningTime, &location.ClosingTime, &location.WorkingDays, &location.CreatedAt, &location.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return &location, nil
}

func (a *App) UpdateBusinessLocation(location models.BusinessLocation) (*models.BusinessLocation, error) {
	location.Geohash = utils.GenerateGeohash(location.Latitude, location.Longitude, 7)

	query := `UPDATE business_locations SET branch_name = $1, address_line1 = $2, address_line2 = $3, city = $4, state = $5, pin_code = $6, latitude = $7, longitude = $8, geohash = $9, is_main = $10, opening_time = $11, closing_time = $12, working_days = $13, updated_at = NOW() WHERE id = $14 RETURNING updated_at`

	err := a.GetDB().QueryRow(context.Background(), query, location.BranchName, location.AddressLine1, location.AddressLine2, location.City, location.State, location.PinCode, location.Latitude, location.Longitude, location.Geohash, location.IsMain, location.OpeningTime, location.ClosingTime, location.WorkingDays, location.ID).Scan(&location.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return &location, nil
}

func (a *App) DeleteBusinessLocation(id uuid.UUID) error {
	query := `DELETE FROM business_locations WHERE id = $1`

	_, err := a.GetDB().Exec(context.Background(), query, id)

	return err
}
