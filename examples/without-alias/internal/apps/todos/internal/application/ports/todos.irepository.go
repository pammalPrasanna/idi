package ports

import (
	"context"

	"without-alias/internal/dtos"
)

type ITodosRepository interface {
	GetTodo(ctx context.Context, arg *dtos.GetTodoParams) (todo *dtos.Todo, err error)
	FindTodos(ctx context.Context, arg *dtos.FindTodosParams) (todos []*dtos.Todo, err error)
	CreateTodo(ctx context.Context, arg *dtos.CreateTodoParams) (id int64, err error)
	UpdateTodo(ctx context.Context, arg *dtos.UpdateTodoParams) error
	DeleteTodo(ctx context.Context, arg *dtos.DeleteTodoParams) error
}
