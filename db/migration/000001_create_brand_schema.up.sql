CREATE TABLE brands (
    id            int UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    name          varchar(255) NOT NULL,
    description   text,
    created_at    int UNSIGNED NOT NULL DEFAULT 0,
    updated_at    int UNSIGNED NOT NULL DEFAULT 0
);