CREATE TABLE user_follow_businesses (
    user_id     UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    business_id UUID NOT NULL REFERENCES businesses(id) ON DELETE CASCADE,
    created_at  TIMESTAMP NOT NULL DEFAULT NOW(),
    PRIMARY KEY (user_id, business_id)
);

CREATE TABLE user_liked_products (
    user_id    UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    product_id UUID NOT NULL REFERENCES products(id) ON DELETE CASCADE,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    PRIMARY KEY (user_id, product_id)
);

CREATE INDEX idx_user_follow_businesses_user ON user_follow_businesses(user_id);
CREATE INDEX idx_user_liked_products_user ON user_liked_products(user_id);
