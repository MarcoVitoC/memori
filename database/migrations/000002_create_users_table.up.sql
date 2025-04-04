BEGIN;
CREATE TABLE IF NOT EXISTS "users"(
    "id" UUID PRIMARY KEY,
    "username" VARCHAR(255) NOT NULL,
    "email" VARCHAR(100) NOT NULL,
    "password" VARCHAR(100) NOT NULL,
    "created_at" TIMESTAMP,
    "updated_at" TIMESTAMP
);
COMMIT;