package lib

import "errors"

type ErrInvalidData map[string][]string

var _ error = (*ErrInvalidData)(nil)

func (e ErrInvalidData) Error() string {
	return "Invalid data"
}

func (e ErrInvalidData) GetErrors() map[string][]string {
	return e
}


var ErrNoRecord = errors.New("no matching records found")