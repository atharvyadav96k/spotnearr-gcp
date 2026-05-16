-- =========================================================
-- ENUMS
-- =========================================================

CREATE TYPE offer_discount_type AS ENUM (
    'percentage',
    'flat'
);

-- =========================================================
-- OFFERS
-- =========================================================

CREATE TABLE offers (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

    business_id UUID NOT NULL
    REFERENCES businesses(id)
    ON DELETE CASCADE,

    title TEXT NOT NULL,

    description TEXT,

    discount_type offer_discount_type NOT NULL,

    discount_value BIGINT NOT NULL,

    min_order_value BIGINT NOT NULL DEFAULT 0,

    coupon_code TEXT UNIQUE,

    banner_url TEXT,

    is_active BOOLEAN NOT NULL DEFAULT TRUE,

    starts_at TIMESTAMP NOT NULL,

    expires_at TIMESTAMP NOT NULL,

    created_at TIMESTAMP NOT NULL DEFAULT NOW(),

    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),

    CONSTRAINT valid_offer_dates
    CHECK (expires_at > starts_at),

    CONSTRAINT valid_discount_value
    CHECK (discount_value >= 0),

    CONSTRAINT valid_min_order_value
    CHECK (min_order_value >= 0)
);

-- =========================================================
-- OFFER PRODUCTS (MANY TO MANY)
-- =========================================================

CREATE TABLE offer_products (
    offer_id UUID NOT NULL
    REFERENCES offers(id)
    ON DELETE CASCADE,

    product_id UUID NOT NULL
    REFERENCES products(id)
    ON DELETE CASCADE,

    PRIMARY KEY(offer_id, product_id)
);

-- =========================================================
-- INDEXES
-- =========================================================

CREATE INDEX idx_offers_business_id
ON offers(business_id);

CREATE INDEX idx_offers_is_active
ON offers(is_active);

CREATE INDEX idx_offers_expires_at
ON offers(expires_at);

CREATE INDEX idx_offer_products_offer_id
ON offer_products(offer_id);

CREATE INDEX idx_offer_products_product_id
ON offer_products(product_id);