package dtos

import "without-alias/internal/lib"

type (
	Todo struct {
		ID          int64  `json:"id"`
		Task        string `json:"task"`
		Description string `json:"description"`
		Archived    bool   `json:"archived"`
		DueDate     string `json:"due_date"`
		CreatedAt   string `json:"created_at"`
		UpdatedAt   string `json:"updated_at"`
	}
	GetTodoParams struct {
		ID int64
	}
	FindTodosParams  struct{}
	CreateTodoParams struct {
		Task        string    `json:"task"`
		Description string    `json:"description"`
		DueDate     lib.ITime `json:"due_date"`
	}
	UpdateTodoParams struct {
		ID          int64     `json:"id"`
		Description string    `json:"description"`
		Task        string    `json:"task"`
		DueDate     lib.ITime `json:"due_date"`
	}
	DeleteTodoParams struct {
		ID int64
	}
)
