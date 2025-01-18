package domain

import (
	"asdf/internal/dtos"
)

type Users struct{}

func NewUsers(ctp dtos.CreateUsersParams) (*Users, error) {

	return &Users{}, nil
}
