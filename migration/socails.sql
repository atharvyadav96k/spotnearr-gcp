-- =========================================================
-- ENUMS
-- =========================================================

CREATE TYPE notification_type AS ENUM (
    'spotlight_posted',
    'offer_live',
    'new_product',
    'review'
);

CREATE TYPE notification_ref_type AS ENUM (
    'spotlight',
    'offer',
    'product'
);

-- =========================================================
-- USER FOLLOW BUSINESS
-- =========================================================

CREATE TABLE user_follow_businesses (
    user_id UUID NOT NULL
    REFERENCES users(id)
    ON DELETE CASCADE,

    business_id UUID NOT NULL
    REFERENCES businesses(id)
    ON DELETE CASCADE,

    created_at TIMESTAMP NOT NULL DEFAULT NOW(),

    PRIMARY KEY(user_id, business_id)
);

-- =========================================================
-- USER SAVED PRODUCTS
-- =========================================================

CREATE TABLE user_saved_products (
    user_id UUID NOT NULL
    REFERENCES users(id)
    ON DELETE CASCADE,

    product_id UUID NOT NULL
    REFERENCES products(id)
    ON DELETE CASCADE,

    created_at TIMESTAMP NOT NULL DEFAULT NOW(),

    PRIMARY KEY(user_id, product_id)
);

-- =========================================================
-- BUSINESS REVIEWS
-- =========================================================

CREATE TABLE business_reviews (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

    business_id UUID NOT NULL
    REFERENCES businesses(id)
    ON DELETE CASCADE,

    user_id UUID NOT NULL
    REFERENCES users(id)
    ON DELETE CASCADE,

    rating INTEGER NOT NULL
    CHECK (rating >= 1 AND rating <= 5),

    comment TEXT,

    is_verified BOOLEAN NOT NULL DEFAULT FALSE,

    created_at TIMESTAMP NOT NULL DEFAULT NOW(),

    CONSTRAINT unique_user_business_review
    UNIQUE(user_id, business_id)
);

-- =========================================================
-- NOTIFICATIONS
-- =========================================================

CREATE TABLE notifications (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

    user_id UUID NOT NULL
    REFERENCES users(id)
    ON DELETE CASCADE,

    type notification_type NOT NULL,

    title TEXT NOT NULL,

    body TEXT,

    ref_id UUID,

    ref_type notification_ref_type,

    is_read BOOLEAN NOT NULL DEFAULT FALSE,

    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- =========================================================
-- INDEXES
-- =========================================================

CREATE INDEX idx_follow_business_user_id
ON user_follow_businesses(user_id);

CREATE INDEX idx_follow_business_business_id
ON user_follow_businesses(business_id);

CREATE INDEX idx_saved_products_user_id
ON user_saved_products(user_id);

CREATE INDEX idx_saved_products_product_id
ON user_saved_products(product_id);

CREATE INDEX idx_business_reviews_business_id
ON business_reviews(business_id);

CREATE INDEX idx_business_reviews_user_id
ON business_reviews(user_id);

CREATE INDEX idx_notifications_user_id
ON notifications(user_id);

CREATE INDEX idx_notifications_is_read
ON notifications(is_read);

CREATE INDEX idx_notifications_created_at
ON notifications(created_at DESC);