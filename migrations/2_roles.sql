-- +migrate Up
CREATE TABLE "roles" (
                         "ticker" varchar NOT NULL,
                         "role_name" varchar NOT NULL,
                         CONSTRAINT "roles_pk" PRIMARY KEY ("ticker")
);

ALTER TABLE "users" ADD FOREIGN KEY ("role_ticker") REFERENCES "roles"("ticker");

-- +migrate Down
DROP TABLE IF EXISTS roles CASCADE;