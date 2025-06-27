-- Users
CREATE TABLE core.user
(
    id            UUID PRIMARY KEY,
    first_name    TEXT        NOT NULL,
    last_name     TEXT        NOT NULL,
    username      TEXT UNIQUE NOT NULL,
    email         TEXT UNIQUE NOT NULL,
    password_hash TEXT        NOT NULL,
    role          TEXT        NOT NULL DEFAULT 'user',   -- or 'admin'
    status        TEXT        NOT NULL DEFAULT 'active', -- 'active', 'banned', 'suspended'
    created_at    TIMESTAMPTZ          DEFAULT now()
);

-- Wallet
CREATE TABLE core.wallet
(
    id              UUID PRIMARY KEY,
    user_id         UUID REFERENCES core.user (id),
    points_balance  INT NOT NULL DEFAULT 0,
    credits_balance INT NOT NULL DEFAULT 0,
    updated_at      TIMESTAMPTZ  DEFAULT now()
);

-- Credit Packages
CREATE TABLE core.credit_package
(
    id            UUID PRIMARY KEY,
    name          TEXT           NOT NULL,
    price_egp     NUMERIC(10, 2) NOT NULL,
    credits       INT            NOT NULL,
    reward_points INT            NOT NULL,
    is_active     BOOLEAN     DEFAULT TRUE,
    created_at    TIMESTAMPTZ DEFAULT now()
);

-- Purchases
CREATE TABLE core.purchase
(
    id                UUID PRIMARY KEY,
    user_id           UUID REFERENCES core.user (id),
    credit_package_id UUID REFERENCES core.credit_package (id) ON DELETE SET NULL,
    created_at        TIMESTAMPTZ DEFAULT now()
);

-- Categories
CREATE TABLE core.category
(

    id                 UUID PRIMARY KEY,
    parent_category_id UUID REFERENCES core.category (id) ON DELETE SET NULL,
    name               TEXT NOT NULL,
    description        TEXT,
    created_at         TIMESTAMPTZ DEFAULT now()
);

-- product
CREATE TABLE core.product
(
    id                UUID PRIMARY KEY,
    category_id       UUID REFERENCES core.category (id) ON DELETE SET NULL,
    name              TEXT NOT NULL,
    description       TEXT,
    redemption_points INT  NOT NULL,
    stock_quantity    INT  NOT NULL,
    is_offer          BOOLEAN     DEFAULT FALSE,
    image_url         TEXT,
    tags              TEXT[], -- PostgreSQL array
    is_active         BOOLEAN     DEFAULT TRUE,
    created_at        TIMESTAMPTZ DEFAULT now()
);

-- Redemptions
CREATE TABLE core.redeem
(
    id         UUID PRIMARY KEY,
    user_id    UUID REFERENCES core.user (id),
    product_id UUID REFERENCES core.product (id) ON DELETE SET NULL,
    quantity   INT  NOT NULL,
    created_at TIMESTAMPTZ DEFAULT now()
);
