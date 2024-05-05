CREATE TABLE "media_channels" (
    "id" VARCHAR PRIMARY KEY,
    "name" VARCHAR,
    "description" VARCHAR,
    "banner_url" VARCHAR,
    "owner_id" VARCHAR NOT NULL,
    "username" VARCHAR NOT NULL UNIQUE,

    FOREIGN KEY ("owner_id") REFERENCES "users"("id")
);