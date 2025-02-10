package users_test

import (
	"context"
	"fmt"
	
	"regexp"
	
	"testing"
	"time"

	"asdf/internal/apps/users/internal/adapters/sqlite3"
	"asdf/internal/apps/users/internal/application"

	
	"asdf/internal/apps/users/internal/application/domain"
	
	"asdf/internal/dtos"
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
		t.Run(fmt.Sprintf("-- failed to create user with - %s", v.tcDesc), func(t *testing.T) {
			t.Parallel()
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

	t.Run("-- failed to create user with invalid email", func(t *testing.T) {
		t.Parallel()
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
		t.Run(fmt.Sprintf("-- failed to create user with - %s", v.tcDesc), func(t *testing.T) {
			t.Parallel()
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


func TestEmail(t *testing.T) {

	invalidEmail:= "gAYhYemail.com"
	validEmail:="gAYhY@email.com"

	v := lib.NewValidator()
	domain.IsValidEmail(v, invalidEmail)
	err := v.Errors().(lib.ErrInvalidData).GetErrors()
	assert.Equal(t, 1, len(err))

	var emailRx = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

	assert.False(t, emailRx.MatchString(invalidEmail))
	assert.True(t, emailRx.MatchString(validEmail))

}


func TestUsersApplication(t *testing.T) {
	t.Parallel()
	if !INTEGRATION_TESTS {
		t.Errorf("skipping tests | INTEGRATION_TESTS = '%v'", INTEGRATION_TESTS)
		t.FailNow()
	}
	assert.NotNil(t, dbConn, "dbConn is nil")

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockLogger := NewMockILogger(ctrl)
	repo := sqlite3.NewRepository(dbConn, mockLogger)
	usersApp := application.New(rootApp, repo)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	// create user
	tu := createValidUser(t)
	id, err := usersApp.CreateUser(ctx, tu)
	assert.Nilf(t, err, "email used %s", tu.Email)
	assert.GreaterOrEqual(t, id, int64(1))

	// get all users
	users, err := usersApp.FindUsers(ctx, &dtos.FindUsersParams{})
	assert.Nil(t, err)
	assert.NotNil(t, users)
	assert.GreaterOrEqual(t, len(users), 0)

	// get created  user
	gotUser1, err := usersApp.GetUser(ctx, &dtos.GetUserParams{ID: id})
	assert.Nil(t, err)
	assert.NotNil(t, gotUser1)

	// update user's username
	updatedUser1 := &dtos.UpdateUserParams{
		ID:       id,
		Username: addrOfStr(randomUsername()),
	}
	err = usersApp.UpdateUser(ctx, updatedUser1)
	assert.Nil(t, err)

	// verify update username
	gotUser2, err := usersApp.GetUser(ctx, &dtos.GetUserParams{ID: id})
	assert.Nil(t, err)
	assert.NotNil(t, gotUser2)
	assert.Equal(t, *updatedUser1.Username, gotUser2.Username)

	// update user's email
	updatedUser2 := &dtos.UpdateUserParams{
		ID:    id,
		Email: addrOfStr(randomEmail()),
	}
	err = usersApp.UpdateUser(ctx, updatedUser2)
	assert.Nil(t, err)

	// verify update email
	gotUser3, err := usersApp.GetUser(ctx, &dtos.GetUserParams{ID: id})
	assert.Nil(t, err)
	assert.NotNil(t, gotUser3)
	assert.Equal(t, *updatedUser2.Email, gotUser3.Email)

	// delete the user
	err = usersApp.DeleteUser(ctx, &dtos.DeleteUserParams{
		ID: gotUser3.ID,
	})
	assert.Nil(t, err)
}
