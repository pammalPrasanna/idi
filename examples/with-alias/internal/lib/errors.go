package lib

import (
	"errors"
	"strings"
)

var (
	ErrNoRecord           = errors.New("no record found")
	ErrInvalidParameterID = errors.New("invalid parameter: id")
)

type (
	ErrInvalidData map[string][]string
	ErrNotUnique   struct {
		Msg   string
		field string
	}
	ErrInvalidJSON struct {
		Msg string
	}
)

var _ error = (*ErrInvalidData)(nil)

func (e ErrInvalidData) Error() string {
	return "invalid data"
}

func (e ErrInvalidData) GetErrors() map[string][]string {
	return e
}

var _ error = (*ErrNotUnique)(nil)

func (e ErrNotUnique) Error() string {
	if e.field == "" {
		sp := strings.Split(e.Msg, ".")
		e.field = sp[len(sp)-1]
	}
	return e.field + " already exists"
}

var _ error = (*ErrInvalidJSON)(nil)

func (e ErrInvalidJSON) Error() string {
	return e.Msg
}
