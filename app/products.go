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

func (a *App) UpdateProduct(product models.Product) (*models.Product, error) {
	query := `UPDATE products SET category_id = $1, name = $2, description = $3, unit = $4, is_available = $5, is_active = $6, tags = $7, updated_at = NOW() WHERE id = $8 RETURNING updated_at`

	err := a.GetDB().QueryRow(context.Background(), query, product.CategoryID, product.Name, product.Description, product.Unit, product.IsAvailable, product.IsActive, product.Tags, product.ID).Scan(&product.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (a *App) DeleteProduct(id uuid.UUID) error {
	query := `DELETE FROM products WHERE id = $1`

	_, err := a.GetDB().Exec(context.Background(), query, id)

	return err
}

func (a *App) GetProductByID(id uuid.UUID) (*models.Product, error) {
	query := `SELECT id, category_id, name, description, unit, is_available, is_active, tags, created_at, updated_at FROM products WHERE id = $1`

	var product models.Product

	err := a.GetDB().QueryRow(context.Background(), query, id).Scan(&product.ID, &product.CategoryID, &product.Name, &product.Description, &product.Unit, &product.IsAvailable, &product.IsActive, &product.Tags, &product.CreatedAt, &product.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (a *App) GetProductsByBusinessID(businessID uuid.UUID) ([]models.Product, error) {
	query := `SELECT DISTINCT p.id, p.category_id, p.name, p.description, p.unit, p.is_available, p.is_active, p.tags, p.created_at, p.updated_at FROM products p JOIN product_inventory pi ON pi.product_id = p.id JOIN business_locations bl ON bl.id = pi.location_id WHERE bl.business_id = $1 ORDER BY p.created_at DESC`

	rows, err := a.GetDB().Query(context.Background(), query, businessID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []models.Product

	for rows.Next() {
		var product models.Product

		err := rows.Scan(&product.ID, &product.CategoryID, &product.Name, &product.Description, &product.Unit, &product.IsAvailable, &product.IsActive, &product.Tags, &product.CreatedAt, &product.UpdatedAt)

		if err != nil {
			return nil, err
		}

		products = append(products, product)
	}

	return products, nil
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

func (a *App) UpdateProductInventory(inventory models.ProductInventory) (*models.ProductInventory, error) {
	query := `UPDATE product_inventory SET product_id = $1, location_id = $2, price = $3, discounted_price = $4, stock = $5, is_available = $6, updated_at = NOW() WHERE id = $7 RETURNING updated_at`

	err := a.GetDB().QueryRow(context.Background(), query, inventory.ProductID, inventory.LocationID, inventory.Price, inventory.DiscountedPrice, inventory.Stock, inventory.IsAvailable, inventory.ID).Scan(&inventory.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return &inventory, nil
}

func (a *App) DeleteProductInventory(id uuid.UUID) error {
	query := `DELETE FROM product_inventory WHERE id = $1`

	_, err := a.GetDB().Exec(context.Background(), query, id)

	return err
}

func (a *App) GetProductInventoryByID(id uuid.UUID) (*models.ProductInventory, error) {
	query := `SELECT id, product_id, location_id, price, discounted_price, stock, is_available, created_at, updated_at FROM product_inventory WHERE id = $1`

	var inventory models.ProductInventory

	err := a.GetDB().QueryRow(context.Background(), query, id).Scan(&inventory.ID, &inventory.ProductID, &inventory.LocationID, &inventory.Price, &inventory.DiscountedPrice, &inventory.Stock, &inventory.IsAvailable, &inventory.CreatedAt, &inventory.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return &inventory, nil
}
func (a *App) GetProductInventoriesByBusinessID(businessID uuid.UUID) ([]models.ProductInventory, error) {
	query := `SELECT pi.id, pi.product_id, pi.location_id, pi.price, pi.discounted_price, pi.stock, pi.is_available, pi.created_at, pi.updated_at FROM product_inventory pi JOIN business_locations bl ON bl.id = pi.location_id WHERE bl.business_id = $1 ORDER BY pi.created_at DESC`

	rows, err := a.GetDB().Query(context.Background(), query, businessID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var inventories []models.ProductInventory

	for rows.Next() {
		var inventory models.ProductInventory

		err := rows.Scan(
			&inventory.ID,
			&inventory.ProductID,
			&inventory.LocationID,
			&inventory.Price,
			&inventory.DiscountedPrice,
			&inventory.Stock,
			&inventory.IsAvailable,
			&inventory.CreatedAt,
			&inventory.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}

		inventories = append(inventories, inventory)
	}

	return inventories, nil
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

func (a *App) GetNearbyProducts(latitude float64, longitude float64, density int) ([]models.Product, error) {
	radius := float64(density * 3)

	query := `SELECT DISTINCT p.id, p.category_id, p.name, p.description, p.unit, p.is_available, p.is_active, p.tags, p.created_at, p.updated_at FROM product_inventory pi JOIN products p ON p.id = pi.product_id JOIN business_locations bl ON bl.id = pi.location_id WHERE pi.stock > 0 AND pi.is_available = true AND p.is_active = true AND CURRENT_TIME BETWEEN bl.opening_time AND bl.closing_time AND LOWER(TRIM(TO_CHAR(CURRENT_DATE, 'Day'))) = ANY(bl.working_days) AND (6371 * acos(cos(radians($1)) * cos(radians(bl.latitude)) * cos(radians(bl.longitude) - radians($2)) + sin(radians($1)) * sin(radians(bl.latitude)))) <= $3 ORDER BY p.created_at DESC`

	rows, err := a.GetDB().Query(context.Background(), query, latitude, longitude, radius)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []models.Product

	for rows.Next() {
		var product models.Product

		err := rows.Scan(
			&product.ID,
			&product.CategoryID,
			&product.Name,
			&product.Description,
			&product.Unit,
			&product.IsAvailable,
			&product.IsActive,
			&product.Tags,
			&product.CreatedAt,
			&product.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}

		products = append(products, product)
	}

	return products, nil
}
