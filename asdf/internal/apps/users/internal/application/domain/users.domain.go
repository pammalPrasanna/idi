package domain

import (
	"strings"
	"unicode/utf8"
	
	
	"regexp"
	

	"asdf/internal/lib"
)

func IsValidUsername(v *lib.Validator, username string) {
	username = strings.TrimSpace(username)
	v.Check("username", utf8.RuneCountInString(username) < 2, "username should be minimum 2 characters")
	v.Check("username", utf8.RuneCountInString(username) > 64, "username should be maximum 64 characters")
}


var emailRx = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
func IsValidEmail(v *lib.Validator, email string) {
	if emailRx == nil {
		emailRx = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	}
	v.Check("email", !emailRx.MatchString(email), "invalid email")
}

func IsValidPassword(v *lib.Validator, password string) {
	v.Check("password", utf8.RuneCountInString(password) < 8, "password should be minimum 8 characters")
	v.Check("password", utf8.RuneCountInString(password) > 64, "password should be maximum 64 characters")
}

