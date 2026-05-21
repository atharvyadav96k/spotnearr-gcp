CREATE TABLE product_claims (
    id           UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id      UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    business_id  UUID NOT NULL REFERENCES businesses(id) ON DELETE CASCADE,
    inventory_id UUID NOT NULL REFERENCES product_inventory(id) ON DELETE CASCADE,
    quantity     INT NOT NULL DEFAULT 1,
    status       TEXT NOT NULL DEFAULT 'pending',
    note         TEXT,
    claimed_at   TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at   TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
CREATE INDEX ON product_claims(user_id);
CREATE INDEX ON product_claims(business_id);
CREATE INDEX ON product_claims(status);