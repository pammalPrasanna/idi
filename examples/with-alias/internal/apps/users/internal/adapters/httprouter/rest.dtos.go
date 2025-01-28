package rest

import "with-alias/internal/dtos"

type (
	FindUsersResponse struct {
		Users []*dtos.User `json:"users"`
	}
	CreateUserResponse struct {
		UserID int64 `json:"User_id"`
	}
	GetUserResponse struct {
		User *dtos.User `json:"user"`
	}
)
