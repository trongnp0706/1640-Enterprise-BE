-- +migrate Up
CREATE TABLE "users" (
                         "id" varchar NOT NULL,
                         "username" varchar NOT NULL,
                         "email" varchar NOT NULL UNIQUE,
                         "password" varchar NOT NULL,
                         "role_ticker" varchar NOT NULL,
                         "department_id" varchar NOT NULL,
                         CONSTRAINT "users_pk" PRIMARY KEY ("id")
);

-- +migrate Down
DROP TABLE IF EXISTS users CASCADE;