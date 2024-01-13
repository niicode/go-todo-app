-- name: CreateTodo :one

INSERT INTO todos (title, description)
VALUES ($1, $2)
RETURNING *;

-- name: UpdateTodo :one

UPDATE todos SET
    status = $2,
    updated_at = CURRENT_TIMESTAMP
WHERE id = $1
RETURNING *;

-- name: DeleteTodo :exec

DELETE FROM todos
WHERE id = $1;


-- name: GetTodo :one

SELECT * FROM todos
WHERE id = $1;

-- name: GetTodos :many

SELECT * FROM todos
ORDER BY created_at DESC;

-- name: GetTodosByStatus :many

SELECT * FROM todos
WHERE status = $1
ORDER BY created_at DESC;

