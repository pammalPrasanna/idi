package application

import (
	"context"

	"asdf/internal/apps/users/internal/application/ports"
	"asdf/internal/dtos"
	"asdf/internal/lib"
)

type Users struct {
	lib.IApp

	Repo ports.IUsersRepository
}

func New(rootApp lib.IApp, repo ports.IUsersRepository) *Users {
	u := &Users{}
	u.IApp = rootApp
	u.Repo = repo
	return u
}

func (u *Users) ListUsers(ctx context.Context, args *dtos.FindUsersParams) ([]*dtos.User, error) {
	// validate args
	users, err := u.Repo.FindUsers(ctx, args)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (u *Users) CreateUser(ctx context.Context, args *dtos.CreateUsersParams) (int64, error) {
	// validate args
	userID, err := u.Repo.CreateUsers(ctx, args)
	if err != nil {
		return 0, err
	}
	return userID, nil
}

func (u *Users) GetUser(ctx context.Context, args *dtos.GetUsersParams) (*dtos.User, error) {
	// validate args
	user, err := u.Repo.GetUsers(ctx, args)
	if err != nil {
		return nil, err
	}
	return user, nil
}
