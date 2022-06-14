CREATE TABLE transactions (
    id            int UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    product_id    int UNSIGNED NOT NULL DEFAULT 0,
    qty           int NOT NULL DEFAULT 0,
    total_price   int UNSIGNED NOT NULL DEFAULT 0,
    email         varchar(255) NOT NULL,
    created_at    int UNSIGNED NOT NULL DEFAULT 0,
    updated_at    int UNSIGNED NOT NULL DEFAULT 0
);