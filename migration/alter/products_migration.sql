-- =========================================================
-- REMOVE OLD PRODUCT DEPENDENCIES
-- =========================================================

DROP INDEX IF EXISTS idx_products_business_id;

ALTER TABLE products
DROP CONSTRAINT IF EXISTS products_business_id_fkey;

ALTER TABLE products
DROP COLUMN IF EXISTS business_id;

ALTER TABLE products
DROP COLUMN IF EXISTS sku;

ALTER TABLE products
DROP COLUMN IF EXISTS barcode;

-- =========================================================
-- REMOVE OLD PRODUCT IMAGES TABLE
-- =========================================================

DROP INDEX IF EXISTS idx_product_images_product_id;

DROP TABLE IF EXISTS product_images;

-- =========================================================
-- ADD GEOHASH TO BUSINESS LOCATIONS
-- =========================================================

ALTER TABLE business_locations
ADD COLUMN IF NOT EXISTS geohash TEXT;

CREATE INDEX IF NOT EXISTS idx_business_locations_geohash
ON business_locations(geohash);

-- =========================================================
-- CREATE MEDIA TABLE
-- =========================================================

CREATE TABLE IF NOT EXISTS media (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

    url TEXT NOT NULL,

    type TEXT NOT NULL,

    thumbnail_url TEXT,

    hash TEXT,

    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- =========================================================
-- CREATE PRODUCT MEDIA TABLE
-- =========================================================

CREATE TABLE IF NOT EXISTS product_media (
    product_id UUID NOT NULL REFERENCES products(id) ON DELETE CASCADE,

    media_id UUID NOT NULL REFERENCES media(id) ON DELETE CASCADE,

    is_primary BOOLEAN NOT NULL DEFAULT FALSE,

    sort_order INTEGER NOT NULL DEFAULT 0,

    PRIMARY KEY(product_id, media_id)
);

-- =========================================================
-- INDEXES
-- =========================================================

CREATE INDEX IF NOT EXISTS idx_media_hash
ON media(hash);

CREATE INDEX IF NOT EXISTS idx_product_media_product_id
ON product_media(product_id);

CREATE INDEX IF NOT EXISTS idx_product_media_media_id
ON product_media(media_id);