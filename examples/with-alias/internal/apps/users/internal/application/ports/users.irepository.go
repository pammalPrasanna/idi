package ports

import (
	"context"

	"with-alias/internal/dtos"
)

type IUsersRepository interface {
	GetUser(ctx context.Context, arg *dtos.GetUserParams) (user *dtos.User, err error)
	FindUsers(ctx context.Context, arg *dtos.FindUsersParams) (users []*dtos.User, err error)
	CreateUser(ctx context.Context, arg *dtos.CreateUserParams) (id int64, err error)
	UpdateUser(ctx context.Context, arg *dtos.UpdateUserParams) error
	DeleteUser(ctx context.Context, arg *dtos.DeleteUserParams) error
}
