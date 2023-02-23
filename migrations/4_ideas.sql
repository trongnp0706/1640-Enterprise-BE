-- +migrate Up
CREATE TABLE "ideas" (
                         "id" varchar NOT NULL,
                         "title" varchar NOT NULL,
                         "content" varchar NOT NULL,
                         "view_count" integer NOT NULL,
                         "image_array" varchar NOT NULL,
                         "upvote_count" integer NOT NULL,
                         "downvote_count" integer NOT NULL,
                         "is_anonymous" BOOLEAN NOT NULL,
                         "user_id" varchar NOT NULL,
                         "category_id" varchar NOT NULL,
                         "academic_year" varchar NOT NULL,
                         CONSTRAINT "ideas_pk" PRIMARY KEY ("id")
);

ALTER TABLE "ideas" ADD FOREIGN KEY ("user_id") REFERENCES "users"("id");

-- +migrate Down
DROP TABLE IF EXISTS ideas CASCADE;