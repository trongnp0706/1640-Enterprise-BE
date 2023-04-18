-- +migrate Up
CREATE TABLE "votes" (
                            "id" varchar NOT NULL PRIMARY KEY,
                            "user_id" varchar NOT NULL,
                            "idea_id" varchar NOT NULL,
                            "vote" varchar NOT NULL
                            
);

ALTER TABLE "votes" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");
ALTER TABLE "votes" ADD FOREIGN KEY ("idea_id") REFERENCES "ideas" ("id");

-- +migrate Down
DROP TABLE iF EXISTS votes;
