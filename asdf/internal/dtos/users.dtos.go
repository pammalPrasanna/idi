package dtos

import "time"

type (
	User struct {
		ID             int64     `json:"id"`
		Username       string    `json:"username"`
		Email          string    `json:"email"`
		HashedPassword string    `json:"-"`
		CreatedAt      time.Time `json:"created_at"`
		UpdatedAt      time.Time `json:"updated_at"`
	}
	GetUserParams struct {
		ID    int64
		Email string
	}
	FindUsersParams  struct{}
	CreateUserParams struct {
		Username string `json:"task"`
		Email    string `json:"description"`
		Password string `json:"password"`
	}
	UpdateUserParams struct {
		Username *string `json:"task"`
		Email    *string `json:"description"`
		ID       int64  `json:"id"`
	}
	DeleteUserParams struct {
		ID int64
	}
)
