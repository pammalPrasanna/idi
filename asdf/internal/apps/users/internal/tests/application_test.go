package users_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"asdf/internal/apps/users/internal/application"

	"asdf/internal/lib"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

// boundry value analysis test
type bvat struct {
	tcDesc      string
	errMsg      string
	length      int
	allowLength bool
}

func createPasswordTests(t *testing.T) []bvat {
	t.Helper()

	tests := []bvat{
		{
			tcDesc:      "invalid password of length 7",
			errMsg:      "password should be minimum 8 characters",
			length:      7,
			allowLength: false,
		},
		{
			tcDesc:      "valid password of length 8",
			errMsg:      "",
			length:      8,
			allowLength: true,
		},
		{
			tcDesc:      "valid password of length 9",
			errMsg:      "",
			length:      9,
			allowLength: true,
		},
		{
			tcDesc:      "valid password of length 63",
			errMsg:      "",
			length:      63,
			allowLength: true,
		},
		{
			tcDesc:      "valid password of length 64",
			errMsg:      "",
			length:      64,
			allowLength: true,
		},
		{
			tcDesc:      "invalid password of length 65",
			errMsg:      "password should be maximum 64 characters",
			length:      65,
			allowLength: false,
		},
	}
	return tests
}

func createUsernameTests(t *testing.T) []bvat {
	t.Helper()

	tests := []bvat{
		{
			tcDesc:      "invalid username of length 1",
			errMsg:      "username should be minimum 2 characters",
			length:      1,
			allowLength: false,
		},
		{
			tcDesc:      "valid username of length 2",
			errMsg:      "",
			length:      2,
			allowLength: true,
		},
		{
			tcDesc:      "valid username of length 3",
			errMsg:      "",
			length:      3,
			allowLength: true,
		},
		{
			tcDesc:      "valid username of length 63",
			errMsg:      "",
			length:      63,
			allowLength: true,
		},
		{
			tcDesc:      "valid username of length 64",
			errMsg:      "",
			length:      64,
			allowLength: true,
		},
		{
			tcDesc:      "invalid username of length 65",
			errMsg:      "username should be maximum 64 characters",
			length:      65,
			allowLength: false,
		},
	}
	return tests
}

func TestUsersDomain(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := NewMockIUsersRepository(ctrl)
	usersApp := application.New(rootApp, mockRepo)
	mockRepo.EXPECT().CreateUser(gomock.Any(), gomock.Any()).Return(int64(1), nil).AnyTimes()

	for _, v := range createUsernameTests(t) {
		t.Run(fmt.Sprintf("-- CreateUser with - %s", v.tcDesc), func(t *testing.T) {
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
			defer cancel()

			user := createValidUser(t)
			user.Username = randString(v.length)

			id, err := usersApp.CreateUser(ctx, user)

			if !v.allowLength {

				assert.Negativef(t, id, "want -1 got %d", id)
				assert.ErrorAsf(t, err, &lib.ErrInvalidData{}, "want ErrInvalidData, got %v", err)

				e := err.(lib.ErrInvalidData).GetErrors()

				errList, ok := e["username"]
				assert.Equal(t, true, ok)
				assert.Greaterf(t, len(errList), 0, "want length > 0, got %d")

				assert.Equal(t, v.errMsg, errList[0])

			} else {
				assert.Nilf(t, err, "want nil, got '%v'", err)
				assert.Equal(t, int64(1), id)
			}
		})
	}

	t.Run("-- email", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
		defer cancel()

		user := createValidUser(t)
		user.Email = "invalidEmail"

		id, err := usersApp.CreateUser(ctx, user)

		assert.Negativef(t, id, "want -1 got %d", id)
		assert.ErrorAsf(t, err, &lib.ErrInvalidData{}, "want ErrInvalidData, got %v", err)

		e := err.(lib.ErrInvalidData).GetErrors()

		errList, ok := e["email"]
		assert.Equal(t, true, ok)
		assert.Greaterf(t, len(errList), 0, "want length > 0, got %d")

		assert.Equal(t, "invalid email", errList[0])
	})

	for _, v := range createPasswordTests(t) {
		t.Run(fmt.Sprintf("-- CreateUser with - %s", v.tcDesc), func(t *testing.T) {
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
			defer cancel()

			user := createValidUser(t)
			user.Password = randString(v.length)

			id, err := usersApp.CreateUser(ctx, user)

			if !v.allowLength {
				assert.Negativef(t, id, "want -1 got %d", id)
				assert.ErrorAsf(t, err, &lib.ErrInvalidData{}, "want ErrInvalidData, got %v", err)

				e := err.(lib.ErrInvalidData).GetErrors()

				errList, ok := e["password"]
				assert.Equal(t, true, ok)
				assert.Greaterf(t, len(errList), 0, "want length > 0, got %d")

				assert.Equal(t, v.errMsg, errList[0])
			} else {
				assert.Nilf(t, err, "want nil, got '%v'", err)
				assert.Equal(t, int64(1), id)
			}
		})
	}
}
