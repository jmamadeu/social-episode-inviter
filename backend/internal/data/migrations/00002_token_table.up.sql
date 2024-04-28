CREATE TYPE "token_type_enum" AS ENUM ('email', 'api'); 

CREATE TABLE "tokens" (
    "id" VARCHAR PRIMARY KEY,
    "created_at" TIMESTAMPTZ,
    "updated_at" TIMESTAMPTZ,
    "type" token_type_enum,
    "email_token" VARCHAR,
    "valid" BOOLEAN,
    "expiration" TIMESTAMPS,
    "user_id" VARCHAR NOT NULL,

    FOREIGN KEY (user_id) REFERENCES users(id)
);