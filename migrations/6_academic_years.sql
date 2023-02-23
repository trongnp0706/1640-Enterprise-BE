-- +migrate Up
CREATE TABLE "academic_years" (
                                  "academic_year" varchar NOT NULL,
                                  "closure_date" TIMESTAMP NOT NULL,
                                  CONSTRAINT "academic_years_pk" PRIMARY KEY ("academic_year")
);

ALTER TABLE "ideas" ADD CONSTRAINT "ideas_fk2" FOREIGN KEY ("academic_year") REFERENCES "academic_years"("academic_year");

-- +migrate Down
DROP TABLE IF EXISTS academic_years CASCADE;