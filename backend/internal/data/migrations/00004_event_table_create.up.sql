CREATE TABLE IF NOT EXISTS "events"(
    "id" VARCHAR PRIMARY KEY,
    "name" VARCHAR NOT NULL,
    "location" VARCHAR,
    "description" VARCHAR,
    "media_channel_id" VARCHAR NOT NULL,
    "date" TIMESTAMPTZ,
    FOREIGN KEY ("media_channel_id") REFERENCES "media_channels"("id") ON DELETE NO ACTION
);
