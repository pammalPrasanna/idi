package domain

import (
	"regexp"
	"strings"

	"asdf/internal/lib"
)

type User struct {
	Username string
	Email    string
	Password string
}

var emailRx = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zAZ0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

func IsValidUsername(v *lib.Validator, username string) {
	username = strings.TrimSpace(username)
	v.Check("username", len(username) < 2, "username should be minimum 2 characters")
}

func IsValidEmail(v *lib.Validator, email string) {
	v.Check("email", !emailRx.MatchString(email), "invalid email")
}

func IsValidPassword(v *lib.Validator, password string) {
	v.Check("password", len(password) < 8, "password should be minimum 8 characters")
	v.Check("password", len(password) > 64, "password should be maximum 64 characters")
}
