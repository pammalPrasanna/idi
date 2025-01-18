package ports

import (
	"context"

	"asdf/internal/dtos"
)

type IUsersRepository interface {
	GetUsers(ctx context.Context, arg *dtos.GetUsersParams) (user *dtos.User, err error)
	FindUsers(ctx context.Context, arg *dtos.FindUsersParams) ([]*dtos.User, error)
	CreateUsers(ctx context.Context, arg *dtos.CreateUsersParams) (lastInsertId int64, err error)
	UpdateUsers(ctx context.Context, arg *dtos.UpdateUsersParams) error
	DeleteUsers(ctx context.Context, arg *dtos.DeleteUsersParams) error
}
