-- name: CreateDepartment :one
INSERT INTO departments(
    id, department_name
)
VALUES (
           $1,  $2
       )
    RETURNING *;

-- name: GetAllDepartments :many
SELECT id, department_name FROM departments;

-- name: GetDepartmentById :one
SELECT * FROM departments WHERE id = $1;

-- name: UpdateDepartment :one
UPDATE departments
SET department_name = $1
WHERE id = $2
    RETURNING *;

-- name: DeleteDepartment :one
DELETE FROM departments
WHERE id = $1
    RETURNING *;
