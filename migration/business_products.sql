-- Enable UUID generation
CREATE EXTENSION IF NOT EXISTS "pgcrypto";

-- =========================================================
-- BUSINESS STATUS ENUM
-- =========================================================

CREATE TYPE business_status AS ENUM (
    'pending',
    'active',
    'suspended'
);

-- =========================================================
-- BUSINESS CATEGORIES
-- =========================================================

CREATE TABLE business_categories (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

    name TEXT NOT NULL UNIQUE,
    icon_url TEXT,

    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- =========================================================
-- BUSINESSES
-- =========================================================

CREATE TABLE businesses (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

    owner_id UUID NOT NULL,
    category_id UUID NOT NULL REFERENCES business_categories(id),

    name TEXT NOT NULL,
    description TEXT,

    email TEXT,
    phone TEXT,
    website TEXT,

    logo_url TEXT,
    cover_url TEXT,

    status business_status NOT NULL DEFAULT 'pending',

    is_verified BOOLEAN NOT NULL DEFAULT FALSE,

    rating DOUBLE PRECISION NOT NULL DEFAULT 0,
    total_reviews INTEGER NOT NULL DEFAULT 0,

    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- =========================================================
-- BUSINESS LOCATIONS
-- =========================================================

CREATE TABLE business_locations (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

    business_id UUID NOT NULL REFERENCES businesses(id) ON DELETE CASCADE,

    branch_name TEXT NOT NULL,

    address_line1 TEXT NOT NULL,
    address_line2 TEXT,

    city TEXT NOT NULL,
    state TEXT NOT NULL,
    pin_code TEXT NOT NULL,

    latitude DOUBLE PRECISION NOT NULL,
    longitude DOUBLE PRECISION NOT NULL,

    is_main BOOLEAN NOT NULL DEFAULT FALSE,

    opening_time TIME NOT NULL,
    closing_time TIME NOT NULL,

    working_days TEXT[] NOT NULL DEFAULT ARRAY[]::TEXT[],

    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- =========================================================
-- PRODUCT CATEGORIES
-- =========================================================

CREATE TABLE product_categories (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

    name TEXT NOT NULL,

    parent_id UUID REFERENCES product_categories(id) ON DELETE SET NULL,

    icon_url TEXT,

    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- =========================================================
-- PRODUCTS
-- =========================================================

CREATE TABLE products (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

    business_id UUID NOT NULL REFERENCES businesses(id) ON DELETE CASCADE,

    category_id UUID NOT NULL REFERENCES product_categories(id),

    name TEXT NOT NULL,
    description TEXT,

    sku TEXT UNIQUE,
    barcode TEXT UNIQUE,

    unit TEXT NOT NULL,

    is_available BOOLEAN NOT NULL DEFAULT TRUE,
    is_active BOOLEAN NOT NULL DEFAULT TRUE,

    tags TEXT[] NOT NULL DEFAULT ARRAY[]::TEXT[],

    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- =========================================================
-- PRODUCT INVENTORY
-- =========================================================

CREATE TABLE product_inventory (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

    product_id UUID NOT NULL REFERENCES products(id) ON DELETE CASCADE,

    location_id UUID NOT NULL REFERENCES business_locations(id) ON DELETE CASCADE,

    price BIGINT NOT NULL,
    discounted_price BIGINT NOT NULL DEFAULT 0,

    stock INTEGER NOT NULL DEFAULT 0,

    is_available BOOLEAN NOT NULL DEFAULT TRUE,

    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),

    CONSTRAINT unique_product_location UNIQUE(product_id, location_id)
);

-- =========================================================
-- PRODUCT IMAGES
-- =========================================================

CREATE TABLE product_images (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

    product_id UUID NOT NULL REFERENCES products(id) ON DELETE CASCADE,

    url TEXT NOT NULL,

    is_primary BOOLEAN NOT NULL DEFAULT FALSE,
    sort_order INTEGER NOT NULL DEFAULT 0,

    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- =========================================================
-- INDEXES
-- =========================================================

CREATE INDEX idx_businesses_owner_id
ON businesses(owner_id);

CREATE INDEX idx_businesses_category_id
ON businesses(category_id);

CREATE INDEX idx_business_locations_business_id
ON business_locations(business_id);

CREATE INDEX idx_business_locations_coordinates
ON business_locations(latitude, longitude);

CREATE INDEX idx_products_business_id
ON products(business_id);

CREATE INDEX idx_products_category_id
ON products(category_id);

CREATE INDEX idx_product_inventory_product_id
ON product_inventory(product_id);

CREATE INDEX idx_product_inventory_location_id
ON product_inventory(location_id);

CREATE INDEX idx_product_images_product_id
ON product_images(product_id);