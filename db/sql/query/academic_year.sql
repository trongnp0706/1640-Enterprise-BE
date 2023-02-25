-- name: CreateAcademicYear :one
INSERT INTO academic_years(
    academic_year, closure_date
)
VALUES (
           $1,  $2
       )
    RETURNING *;

-- name: GetAcademicYears :many
SELECT academic_year, closure_date FROM academic_years;

-- name: GetAcademicYear :one
SELECT * FROM academic_years WHERE academic_year = $1;

-- name: UpdateAcademicYear :one
UPDATE academic_years
SET academic_year = $1,
    closure_date = $2
WHERE academic_year = $3
    RETURNING *;

-- name: DeleteAcademicYear :one
DELETE FROM academic_years
WHERE academic_year = $1
    RETURNING *;
