package dtos

import "with-alias/internal/lib"

type (
	User struct {
		ID          int64  `json:"id"`
		Task        string `json:"task"`
		Description string `json:"description"`
		Archived    bool   `json:"archived"`
		DueDate     string `json:"due_date"`
		CreatedAt   string `json:"created_at"`
		UpdatedAt   string `json:"updated_at"`
	}
	GetUserParams struct {
		ID int64
	}
	FindUsersParams  struct{}
	CreateUserParams struct {
		Task        string    `json:"task"`
		Description string    `json:"description"`
		DueDate     lib.ITime `json:"due_date"`
	}
	UpdateUserParams struct {
		ID          int64     `json:"id"`
		Description string    `json:"description"`
		Task        string    `json:"task"`
		DueDate     lib.ITime `json:"due_date"`
	}
	DeleteUserParams struct {
		ID int64
	}
)
