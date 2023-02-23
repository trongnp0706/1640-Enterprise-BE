-- +migrate Up
CREATE TABLE "departments"
(
    "id" varchar NOT NULL,
    "department_name" varchar NOT NULL,
    CONSTRAINT "departments_pk" PRIMARY KEY ("id")
);

ALTER TABLE "users" ADD FOREIGN KEY ("department_id") REFERENCES "departments"("id");

-- +migrate Down
DROP TABLE IF EXISTS departments CASCADE;