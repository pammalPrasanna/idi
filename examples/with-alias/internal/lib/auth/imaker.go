package auth

import "errors"

type IAuth interface {
	CreateToken(userID int) (*Token, error)
	VerifyToken(token string) (string, error)
}

var ErrInvalidToken = errors.New("invalid authentication token")

type Token struct {
	Token  string
	Expiry string
}
