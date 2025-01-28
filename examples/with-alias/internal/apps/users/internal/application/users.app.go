package application

import (
	"context"

	"with-alias/internal/apps/users/internal/application/domain"
	"with-alias/internal/apps/users/internal/application/ports"
	"with-alias/internal/dtos"
	"with-alias/internal/lib"
)

type (
	Users struct {
		RootApp lib.IApp

		Repo ports.IUsersRepository
	}

	IUsers interface {
		GetUser(ctx context.Context, arg *dtos.GetUserParams) (todo *dtos.User, err error)
		FindUsers(ctx context.Context, arg *dtos.FindUsersParams) (users []*dtos.User, err error)
		CreateUser(ctx context.Context, arg *dtos.CreateUserParams) (id int64, err error)
		UpdateUser(ctx context.Context, arg *dtos.UpdateUserParams) error
		DeleteUser(ctx context.Context, arg *dtos.DeleteUserParams) error
	}
)

var _ IUsers = (*Users)(nil)

func New(rootApp lib.IApp, db ports.IUsersRepository) *Users {
	return &Users{
		RootApp: rootApp,
		Repo:    db,
	}
}

func (t *Users) GetUser(ctx context.Context, arg *dtos.GetUserParams) (user *dtos.User, err error) {
	return t.Repo.GetUser(ctx, arg)
}

func (t *Users) FindUsers(ctx context.Context, arg *dtos.FindUsersParams) (users []*dtos.User, err error) {
	return t.Repo.FindUsers(ctx, arg)
}

func (t *Users) CreateUser(ctx context.Context, arg *dtos.CreateUserParams) (id int64, err error) {
	_, err = domain.NewUser(arg.Task, arg.Description, arg.DueDate)
	if err != nil {
		return 0, err
	}
	return t.Repo.CreateUser(ctx, arg)
}

func (t *Users) UpdateUser(ctx context.Context, arg *dtos.UpdateUserParams) error {
	_, err := domain.NewUser(arg.Task, arg.Description, arg.DueDate)
	if err != nil {
		return err
	}
	return t.Repo.UpdateUser(ctx, arg)
}

func (t *Users) DeleteUser(ctx context.Context, arg *dtos.DeleteUserParams) error {
	return t.Repo.DeleteUser(ctx, arg)
}
