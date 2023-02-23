-- +migrate Up
CREATE TABLE "sessions" (
                            "id" varchar NOT NULL PRIMARY KEY,
                            "user_id" varchar NOT NULL,
                            "refresh_token" varchar NOT NULL,
                            "user_agent" varchar NOT NULL,
                            "client_ip" varchar NOT NULL,
                            "is_blocked" boolean NOT NULL DEFAULT false,
                            "expires_at" bigint NOT NULL,
                            "created_at" timestamptz NOT NULL
);

ALTER TABLE "sessions" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

-- +migrate Down
DROP TABLE iF EXISTS sessions;
