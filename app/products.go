package app

import (
	"context"
	"fmt"
	"strings"

	"github.com/atharvyadav96k/spotnearr-gcp/app/database/models"
	"github.com/google/uuid"
)

func (a *App) CreateProduct(product models.Product) (*models.Product, error) {
	query := `INSERT INTO products (category_id, name, description, unit, is_available, is_active, tags) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id, created_at, updated_at`

	err := a.GetDB().QueryRow(context.Background(), query, product.CategoryID, product.Name, product.Description, product.Unit, product.IsAvailable, product.IsActive, product.Tags).Scan(&product.ID, &product.CreatedAt, &product.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (a *App) CreateProductInventories(inventories []models.ProductInventory) error {
	query := `INSERT INTO product_inventory (product_id, location_id, price, discounted_price, stock, is_available) VALUES `

	args := []interface{}{}
	valueStrings := []string{}

	for i, inventory := range inventories {
		n := i * 6

		valueStrings = append(valueStrings, fmt.Sprintf("($%d, $%d, $%d, $%d, $%d, $%d)", n+1, n+2, n+3, n+4, n+5, n+6))

		args = append(args, inventory.ProductID, inventory.LocationID, inventory.Price, inventory.DiscountedPrice, inventory.Stock, inventory.IsAvailable)
	}

	query += strings.Join(valueStrings, ", ")

	_, err := a.GetDB().Exec(context.Background(), query, args...)

	return err
}

func (a *App) CreateMedia(media models.Media) (*models.Media, error) {
	query := `INSERT INTO media (url, type, thumbnail_url, hash) VALUES ($1, $2, $3, $4) RETURNING id, created_at`

	err := a.GetDB().QueryRow(context.Background(), query, media.URL, media.Type, media.ThumbnailURL, media.Hash).Scan(&media.ID, &media.CreatedAt)

	if err != nil {
		return nil, err
	}

	return &media, nil
}

func (a *App) AttachMediaToProduct(productMedia models.ProductMedia) error {
	query := `INSERT INTO product_media (product_id, media_id, is_primary, sort_order) VALUES ($1, $2, $3, $4)`

	_, err := a.GetDB().Exec(context.Background(), query, productMedia.ProductID, productMedia.MediaID, productMedia.IsPrimary, productMedia.SortOrder)

	return err
}

func (a *App) GetNearbyProducts(latitude float64, longitude float64, density int) ([]map[string]interface{}, error) {
	radius := float64(density * 3)

	query := `SELECT p.id, p.name, p.description, p.unit, p.tags, pi.price, pi.discounted_price, pi.stock, bl.id, bl.branch_name, bl.latitude, bl.longitude, m.url, m.thumbnail_url, (6371 * acos(cos(radians($1)) * cos(radians(bl.latitude)) * cos(radians(bl.longitude) - radians($2)) + sin(radians($1)) * sin(radians(bl.latitude)))) AS distance FROM product_inventory pi JOIN products p ON p.id = pi.product_id JOIN business_locations bl ON bl.id = pi.location_id LEFT JOIN product_media pm ON pm.product_id = p.id AND pm.is_primary = true LEFT JOIN media m ON m.id = pm.media_id WHERE pi.stock > 0 AND pi.is_available = true AND p.is_active = true AND CURRENT_TIME BETWEEN bl.opening_time AND bl.closing_time AND LOWER(TRIM(TO_CHAR(CURRENT_DATE, 'Day'))) = ANY(bl.working_days) AND (6371 * acos(cos(radians($1)) * cos(radians(bl.latitude)) * cos(radians(bl.longitude) - radians($2)) + sin(radians($1)) * sin(radians(bl.latitude)))) <= $3 ORDER BY distance ASC`

	rows, err := a.GetDB().Query(context.Background(), query, latitude, longitude, radius)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []map[string]interface{}

	for rows.Next() {
		var (
			productID       uuid.UUID
			name            string
			description     *string
			unit            string
			tags            []string
			price           int64
			discountedPrice int64
			stock           int
			locationID      uuid.UUID
			branchName      string
			lat             float64
			lng             float64
			imageURL        *string
			thumbnailURL    *string
			distance        float64
		)

		err := rows.Scan(&productID, &name, &description, &unit, &tags, &price, &discountedPrice, &stock, &locationID, &branchName, &lat, &lng, &imageURL, &thumbnailURL, &distance)
		if err != nil {
			return nil, err
		}

		products = append(products, map[string]interface{}{
			"product_id":       productID,
			"name":             name,
			"description":      description,
			"unit":             unit,
			"tags":             tags,
			"price":            price,
			"discounted_price": discountedPrice,
			"stock":            stock,
			"location_id":      locationID,
			"branch_name":      branchName,
			"latitude":         lat,
			"longitude":        lng,
			"image_url":        imageURL,
			"thumbnail_url":    thumbnailURL,
			"distance_km":      distance,
		})
	}

	return products, nil
}
