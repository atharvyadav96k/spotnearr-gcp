-- =========================================================
-- ENUMS
-- =========================================================

CREATE TYPE spotlight_type AS ENUM (
    'product',
    'offer',
    'general'
);

CREATE TYPE spotlight_status AS ENUM (
    'draft',
    'published',
    'expired'
);

CREATE TYPE spotlight_media_type AS ENUM (
    'image',
    'video'
);

CREATE TYPE spotlight_interaction_type AS ENUM (
    'like',
    'share',
    'save'
);

-- =========================================================
-- SPOTLIGHTS
-- =========================================================

CREATE TABLE spotlights (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

    business_id UUID NOT NULL
    REFERENCES businesses(id)
    ON DELETE CASCADE,

    type spotlight_type NOT NULL,

    status spotlight_status NOT NULL DEFAULT 'draft',

    title TEXT NOT NULL,

    description TEXT,

    media_url TEXT NOT NULL,

    media_type spotlight_media_type NOT NULL,

    thumbnail_url TEXT,

    product_id UUID
    REFERENCES products(id)
    ON DELETE SET NULL,

    offer_id UUID,

    view_count INTEGER NOT NULL DEFAULT 0,
    like_count INTEGER NOT NULL DEFAULT 0,
    share_count INTEGER NOT NULL DEFAULT 0,
    comment_count INTEGER NOT NULL DEFAULT 0,

    expires_at TIMESTAMP,

    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- =========================================================
-- SPOTLIGHT FEED ITEMS
-- =========================================================

CREATE TABLE spotlight_feed_items (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

    user_id UUID NOT NULL
    REFERENCES users(id)
    ON DELETE CASCADE,

    spotlight_id UUID NOT NULL
    REFERENCES spotlights(id)
    ON DELETE CASCADE,

    is_seen BOOLEAN NOT NULL DEFAULT FALSE,

    seen_at TIMESTAMP,

    created_at TIMESTAMP NOT NULL DEFAULT NOW(),

    CONSTRAINT unique_feed_item
    UNIQUE(user_id, spotlight_id)
);

-- =========================================================
-- SPOTLIGHT INTERACTIONS
-- =========================================================

CREATE TABLE spotlight_interactions (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

    user_id UUID NOT NULL
    REFERENCES users(id)
    ON DELETE CASCADE,

    spotlight_id UUID NOT NULL
    REFERENCES spotlights(id)
    ON DELETE CASCADE,

    type spotlight_interaction_type NOT NULL,

    created_at TIMESTAMP NOT NULL DEFAULT NOW(),

    CONSTRAINT unique_interaction
    UNIQUE(user_id, spotlight_id, type)
);

-- =========================================================
-- SPOTLIGHT COMMENTS
-- =========================================================

CREATE TABLE spotlight_comments (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

    spotlight_id UUID NOT NULL
    REFERENCES spotlights(id)
    ON DELETE CASCADE,

    user_id UUID NOT NULL
    REFERENCES users(id)
    ON DELETE CASCADE,

    parent_id UUID
    REFERENCES spotlight_comments(id)
    ON DELETE CASCADE,

    content TEXT NOT NULL,

    is_deleted BOOLEAN NOT NULL DEFAULT FALSE,

    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- =========================================================
-- INDEXES
-- =========================================================

CREATE INDEX idx_spotlights_business_id
ON spotlights(business_id);

CREATE INDEX idx_spotlights_product_id
ON spotlights(product_id);

CREATE INDEX idx_spotlights_status
ON spotlights(status);

CREATE INDEX idx_spotlights_created_at
ON spotlights(created_at DESC);

CREATE INDEX idx_feed_user_id
ON spotlight_feed_items(user_id);

CREATE INDEX idx_interactions_spotlight_id
ON spotlight_interactions(spotlight_id);

CREATE INDEX idx_comments_spotlight_id
ON spotlight_comments(spotlight_id);

CREATE INDEX idx_comments_parent_id
ON spotlight_comments(parent_id);