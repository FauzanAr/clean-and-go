CREATE TABLE products (
    id            int UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    brand_id      int UNSIGNED NOT NULL DEFAULT 0,
    name          varchar(255) NOT NULL,
    description   text,
    price         int UNSIGNED NOT NULL DEFAULT 0,
    stock         int NOT NULL DEFAULT 0,
    created_at    int UNSIGNED NOT NULL DEFAULT 0,
    updated_at    int UNSIGNED NOT NULL DEFAULT 0
);