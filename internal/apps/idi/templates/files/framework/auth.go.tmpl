package lib

import (
	"context"
	"errors"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

type contextKey string

// should not use built-in type string as key for value; define your own type to avoid collisions (SA1029)
const authenticatedUserContextKey = contextKey("authenticatedUser")

func (i *idi) contextSetAuthenticatedUser(r *http.Request, user *IUser) *http.Request {
	ctx := context.WithValue(r.Context(), authenticatedUserContextKey, user)
	return r.WithContext(ctx)
}

func (i *idi) ContextGetAuthenticatedUser(r *http.Request) *IUser {
	return r.Context().Value(authenticatedUserContextKey).(*IUser)
}

func (i *idi) SetUserByIDMethod(fn func(id int) (*IUser, error)) {
	i.customGetUserByID = fn
}

func (i *idi) Hash(plaintextPassword string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(plaintextPassword), 12)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

func (i *idi) CompareHashAndPassword(plaintextPassword, hashedPassword string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plaintextPassword))
	if err != nil {
		switch {
		case errors.Is(err, bcrypt.ErrMismatchedHashAndPassword):
			return false, nil
		default:
			return false, err
		}
	}

	return true, nil
}
