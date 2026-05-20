CREATE TABLE user_follow_users (
    follower_id UUID NOT NULL
    REFERENCES users(id)
    ON DELETE CASCADE,

    following_id UUID NOT NULL
    REFERENCES users(id)
    ON DELETE CASCADE,

    created_at TIMESTAMP NOT NULL DEFAULT NOW(),

    PRIMARY KEY(follower_id, following_id),

    CONSTRAINT no_self_follow CHECK (follower_id != following_id)
);

CREATE INDEX idx_follow_user_follower_id
ON user_follow_users(follower_id);

CREATE INDEX idx_follow_user_following_id
ON user_follow_users(following_id);
