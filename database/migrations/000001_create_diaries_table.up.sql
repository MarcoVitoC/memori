BEGIN;
CREATE TABLE IF NOT EXISTS "diaries"(
    "id" UUID PRIMARY KEY,
    "content" TEXT NOT NULL,
    "created_at" TIMESTAMP,
    "updated_at" TIMESTAMP
);
COMMIT;