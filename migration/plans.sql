-- =========================================================
-- PLANS CATALOG
-- =========================================================

CREATE TABLE plans (
    id          UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name        VARCHAR(100) NOT NULL,
    tier        VARCHAR(10)  NOT NULL,
    price       NUMERIC(10,2) NOT NULL DEFAULT 0,
    description TEXT,
    is_active   BOOLEAN NOT NULL DEFAULT true,
    created_at  TIMESTAMP NOT NULL DEFAULT NOW()
);

-- =========================================================
-- BUSINESS SUBSCRIPTIONS
-- One row per business — default tier is 'none'
-- =========================================================

CREATE TABLE business_subscriptions (
    id          UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    business_id UUID NOT NULL REFERENCES businesses(id) ON DELETE CASCADE,
    plan_id     UUID REFERENCES plans(id),
    tier        VARCHAR(10) NOT NULL DEFAULT 'none',
    started_at  TIMESTAMP NOT NULL DEFAULT NOW(),
    expires_at  TIMESTAMP,
    created_at  TIMESTAMP NOT NULL DEFAULT NOW(),

    UNIQUE(business_id)
);

CREATE INDEX idx_business_subscriptions_business_id
ON business_subscriptions(business_id);
