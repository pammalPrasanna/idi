package users_test

import (
	"math/rand"
)

const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func randString(length int) string {
	b := make([]byte, length)
	for i := range b {
		r := rand.Intn(61)
		b[i] = charset[r]
	}
	return string(b)
}

func randomEmail() string {
	return randString(5) + "@email.com"
}

func randomUsername() string {
	return randString(6)
}
