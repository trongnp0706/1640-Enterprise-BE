-- +migrate Up
CREATE TABLE "categories" (
                              "id" varchar NOT NULL,
                              "category_name" varchar NOT NULL,
                              CONSTRAINT "categories_pk" PRIMARY KEY ("id")
);

ALTER TABLE "ideas" ADD FOREIGN KEY ("category_id") REFERENCES "categories"("id");

-- +migrate Down
DROP TABLE IF EXISTS categories CASCADE;