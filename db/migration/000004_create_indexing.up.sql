CREATE UNIQUE INDEX brands_index ON brands (created_at, updated_at);
CREATE UNIQUE INDEX products_index ON products (brand_id, created_at, updated_at);
CREATE UNIQUE INDEX transactions_index ON transactions (product_id, email, created_at, updated_at);