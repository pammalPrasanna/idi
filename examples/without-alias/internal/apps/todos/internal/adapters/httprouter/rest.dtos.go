package rest

import "without-alias/internal/dtos"

type (
	FindTodosResponse struct {
		Todos []*dtos.Todo `json:"todos"`
	}
	CreateTodoResponse struct {
		TodoID int64 `json:"Todo_id"`
	}
	GetTodoResponse struct {
		Todo *dtos.Todo `json:"todo"`
	}
)
