CREATE TABLE IF NOT EXISTS products (
    product_id      uuid PRIMARY KEY,
    name            VARCHAR NOT NULL,
    description     TEXT NOT NULL,
    available_stock BIGINT DEFAULT 0,
    lock_stock      BIGINT DEFAULT 0,
    pending_stock   BIGINT DEFAULT 0,
    shop_id         uuid NOT NULL,
    sold_qty        BIGINT DEFAULT 0,
    created_at      TIMESTAMP DEFAULT NOW(),
    updated_at      TIMESTAMP
);

CREATE TABLE IF NOT EXISTS product_inventory_ledgers (
  id uuid PRIMARY KEY,
  ref_id uuid NOT NULL,
  product_id uuid NOT NULL,
  stock_movement BIGINT NOT NULL,
  action_type VARCHAR NOT NULL,
  user_id uuid NOT NULL,
  created_at TIMESTAMP DEFAULT NOW()
);