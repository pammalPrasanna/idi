package application

import (
	"context"

	"without-alias/internal/apps/todos/internal/application/domain"
	"without-alias/internal/apps/todos/internal/application/ports"
	"without-alias/internal/dtos"
	"without-alias/internal/lib"
)

type (
	Todos struct {
		RootApp lib.IApp

		Repo ports.ITodosRepository
	}

	ITodos interface {
		GetTodo(ctx context.Context, arg *dtos.GetTodoParams) (todo *dtos.Todo, err error)
		FindTodos(ctx context.Context, arg *dtos.FindTodosParams) (todos []*dtos.Todo, err error)
		CreateTodo(ctx context.Context, arg *dtos.CreateTodoParams) (id int64, err error)
		UpdateTodo(ctx context.Context, arg *dtos.UpdateTodoParams) error
		DeleteTodo(ctx context.Context, arg *dtos.DeleteTodoParams) error
	}
)

var _ ITodos = (*Todos)(nil)

func New(rootApp lib.IApp, db ports.ITodosRepository) *Todos {
	return &Todos{
		RootApp: rootApp,
		Repo:    db,
	}
}

func (t *Todos) GetTodo(ctx context.Context, arg *dtos.GetTodoParams) (todo *dtos.Todo, err error) {
	return t.Repo.GetTodo(ctx, arg)
}

func (t *Todos) FindTodos(ctx context.Context, arg *dtos.FindTodosParams) (todos []*dtos.Todo, err error) {
	return t.Repo.FindTodos(ctx, arg)
}

func (t *Todos) CreateTodo(ctx context.Context, arg *dtos.CreateTodoParams) (id int64, err error) {
	_, err = domain.NewTodo(arg.Task, arg.Description, arg.DueDate)
	if err != nil {
		return 0, err
	}
	return t.Repo.CreateTodo(ctx, arg)
}

func (t *Todos) UpdateTodo(ctx context.Context, arg *dtos.UpdateTodoParams) error {
	_, err := domain.NewTodo(arg.Task, arg.Description, arg.DueDate)
	if err != nil {
		return err
	}
	return t.Repo.UpdateTodo(ctx, arg)
}

func (t *Todos) DeleteTodo(ctx context.Context, arg *dtos.DeleteTodoParams) error {
	return t.Repo.DeleteTodo(ctx, arg)
}
