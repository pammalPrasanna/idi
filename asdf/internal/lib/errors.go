package lib

import (
	"errors"
	"strings"
)

var ErrNoRecord = errors.New("no record found")

type ErrInvalidData map[string][]string

var _ error = (*ErrInvalidData)(nil)

func (e ErrInvalidData) Error() string {
	return "Invalid data"
}

func (e ErrInvalidData) GetErrors() map[string][]string {
	return e
}

type ErrNotUnique struct {
	Msg   string
	field string
}

func (e ErrNotUnique) Error() string {
	if e.field == "" {
		sp := strings.Split(e.Msg, ".")
		e.field = sp[len(sp)-1]
	}
	return e.field + " already exists"
}

var _ error = (*ErrNotUnique)(nil)
