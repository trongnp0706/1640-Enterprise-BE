-- +migrate Up
CREATE TABLE "comments" (
                            "id" varchar NOT NULL,
                            "content" varchar NOT NULL,
                            "is_anonymous" BOOLEAN NOT NULL,
                            "user_id" varchar NOT NULL,
                            "idea_id" varchar NOT NULL,
                            CONSTRAINT "comments_pk" PRIMARY KEY ("id")
);

ALTER TABLE "comments" ADD CONSTRAINT "comments_fk0" FOREIGN KEY ("user_id") REFERENCES "users"("id");
ALTER TABLE "comments" ADD CONSTRAINT "comments_fk1" FOREIGN KEY ("idea_id") REFERENCES "ideas"("id");

-- +migrate Down
DROP TABLE IF EXISTS comments CASCADE;
