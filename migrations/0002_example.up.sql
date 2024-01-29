CREATE TABLE IF NOT EXISTS "example" (
    "id" VARCHAR(20) PRIMARY KEY,
    "name" VARCHAR(100) NOT NULL,
    "age" INTEGER NOT NULL,
    "country" VARCHAR(2) NOT NULL,
    "role" VARCHAR(50) NOT NULL,
    "settings" JSONB NOT NULL,
    "deleted_at" TIMESTAMP WITH TIME ZONE NULL
);

CREATE INDEX CONCURRENTLY IF NOT EXISTS "example_name_idx" ON "example" USING gin ("name" gin_trgm_ops);
