CREATE TABLE "media_channels" (
    "id" VARCHAR PRIMARY KEY,
    "name" VARCHAR,
    "platform" VARCHAR,
    "url" VARCHAR,
    "banner_url" VARCHAR,
    "user_id" VARCHAR NOT NULL,

    FOREIGN KEY ("user_id") REFERENCES "users"("id")
);