-- =========================================================
-- EXTENSION
-- =========================================================

CREATE EXTENSION IF NOT EXISTS "pgcrypto";

-- =========================================================
-- USER ROLE ENUM
-- =========================================================

CREATE TYPE user_role AS ENUM (
    'customer',
    'business',
    'admin'
);

-- =========================================================
-- USERS TABLE
-- =========================================================

CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

    full_name TEXT NOT NULL,

    email TEXT UNIQUE,
    phone TEXT UNIQUE,

    password_hash TEXT,

    role user_role NOT NULL DEFAULT 'customer',

    avatar_url TEXT,

    is_verified BOOLEAN NOT NULL DEFAULT FALSE,
    is_active BOOLEAN NOT NULL DEFAULT TRUE,

    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),

    deleted_at TIMESTAMP
);

-- =========================================================
-- INDEXES
-- =========================================================

CREATE INDEX idx_users_role
ON users(role);

CREATE INDEX idx_users_created_at
ON users(created_at);

CREATE INDEX idx_users_deleted_at
ON users(deleted_at);