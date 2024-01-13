package main

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/niicode/go-todo-app/internal/db"
)

type Todo struct {
	ID          uuid.UUID      `json:"id"`
	Title       string         `json:"title"`
	Description sql.NullString `json:"description"`
	Status      string         `json:"status"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
}

func databaseTodoToTodo(dbTodo db.Todo) Todo {
	return Todo{
		ID:          dbTodo.ID,
		Title:       dbTodo.Title,
		Description: dbTodo.Description,
		Status:      dbTodo.Status,
		CreatedAt:   dbTodo.CreatedAt,
		UpdatedAt:   dbTodo.UpdatedAt,
	}
}
