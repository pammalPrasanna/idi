package users_test

import (
	"context"
	"testing"
	"time"

	"asdf/internal/apps/users/internal/adapters/sqlite3"
	"asdf/internal/dtos"
	"asdf/internal/lib"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func createValidUser(t *testing.T) *dtos.CreateUserParams {
	t.Helper()

	return &dtos.CreateUserParams{
		Username: randomUsername(),
		Email:    randomEmail(),
		Password: "adsfi9he",
	}
}

func TestUsersRepository_CRUD(t *testing.T) {
	t.Parallel()
	if !INTEGRATION_TESTS {
		t.Errorf("skipping tests | INTEGRATION_TESTS = '%v'", INTEGRATION_TESTS)
		t.FailNow()
	}
	assert.NotNil(t, dbConn, "db connection is nil")

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockLogger := NewMockILogger(ctrl)
	t.Run("-- create user with valid data", func(t *testing.T) {
		repo := sqlite3.NewRepository(dbConn, mockLogger)

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
		defer cancel()

		// create user
		id, err := repo.CreateUser(ctx, createValidUser(t))

		assert.GreaterOrEqualf(t, id, int64(1), "want >= 1, got %d", id)
		assert.Nilf(t, err, "want nil, got %v", err)
	})

	t.Run("-- get user with valid id", func(t *testing.T) {
		repo := sqlite3.NewRepository(dbConn, mockLogger)

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
		defer cancel()

		validUser := createValidUser(t)
		// create user
		id, err := repo.CreateUser(ctx, validUser)

		assert.GreaterOrEqualf(t, id, int64(1), "want >= 1, got %d", id)
		assert.Nilf(t, err, "want nil, got %v", err)

		// verify the created user
		gotUser, err := repo.GetUser(ctx, &dtos.GetUserParams{
			ID: id,
		})
		assert.Nilf(t, err, "want nil, got %v", err)

		assert.Equalf(t, validUser.Email, gotUser.Email, "want '%s', got '%s'", validUser.Email, gotUser.Email)
		assert.Equalf(t, validUser.Username, gotUser.Username, "want '%s', got '%s'", validUser.Email, gotUser.Email)
	})

	t.Run("-- update user with valid id", func(t *testing.T) {
		repo := sqlite3.NewRepository(dbConn, mockLogger)

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
		defer cancel()

		validUser := createValidUser(t)

		// create user
		id, err := repo.CreateUser(ctx, validUser)

		assert.GreaterOrEqualf(t, id, int64(1), "want >= 1, got %d", id)
		assert.Nilf(t, err, "want nil, got %v", err)

		updatedUser := &dtos.UpdateUserParams{
			ID:       id,
			Username: "updated_username",
			Email:    "updated@email.com",
		}

		// update user
		err = repo.UpdateUser(ctx, updatedUser)
		assert.Nilf(t, err, "want nil, got %v", err)

		// verify the update
		gotUser, err := repo.GetUser(ctx, &dtos.GetUserParams{
			ID: id,
		})
		assert.Nilf(t, err, "want nil, got %v", err)

		assert.Equalf(t, updatedUser.Email, gotUser.Email, "want '%s', got '%s'", updatedUser.Email, gotUser.Email)
		assert.Equalf(t, updatedUser.Username, gotUser.Username, "want '%s', got '%s'", updatedUser.Email, gotUser.Email)
	})

	t.Run("-- delete user with valid id", func(t *testing.T) {
		repo := sqlite3.NewRepository(dbConn, mockLogger)

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
		defer cancel()

		validUser := createValidUser(t)

		// create user
		id, err := repo.CreateUser(ctx, validUser)

		assert.GreaterOrEqualf(t, id, int64(1), "want >= 1, got %d", id)
		assert.Nilf(t, err, "want nil, got %v", err)

		// delete user
		err = repo.DeleteUser(ctx, &dtos.DeleteUserParams{
			ID: id,
		})
		assert.Nilf(t, err, "want nil, got %v", err)

		// verify the delete
		gotUser, err := repo.GetUser(ctx, &dtos.GetUserParams{
			ID: id,
		})
		assert.Nilf(t, gotUser, "want nil, got %v", gotUser)
		assert.ErrorIs(t, err, lib.ErrNoRecord, "want lib.ErrNoRecord, got %s", err)
	})
}

func TestUsersRepository_NonExistentID(t *testing.T) {
	t.Parallel()
	if !INTEGRATION_TESTS {
		t.Errorf("skipping tests | INTEGRATION_TESTS = '%v'", INTEGRATION_TESTS)
		t.FailNow()
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockLogger := NewMockILogger(ctrl)

	t.Run("-- read user with invalid id", func(t *testing.T) {
		repo := sqlite3.NewRepository(dbConn, mockLogger)

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
		defer cancel()

		gotUser, err := repo.GetUser(ctx, &dtos.GetUserParams{
			ID: -1,
		})
		assert.Nilf(t, gotUser, "want nil, got %v", gotUser)
		assert.NotNil(t, err)
		assert.ErrorIs(t, lib.ErrNoRecord, err)
	})
	t.Run("-- update user with invalid id", func(t *testing.T) {
		repo := sqlite3.NewRepository(dbConn, mockLogger)

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
		defer cancel()

		err := repo.UpdateUser(ctx, &dtos.UpdateUserParams{
			ID:       -1,
			Username: "user1",
			Email:    "user1@email.com",
		})

		assert.NotNil(t, err)
		assert.ErrorIs(t, lib.ErrNoRecord, err)
	})
	t.Run("-- delete user with invalid id", func(t *testing.T) {
		repo := sqlite3.NewRepository(dbConn, mockLogger)

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
		defer cancel()

		err := repo.DeleteUser(ctx, &dtos.DeleteUserParams{
			ID: -1,
		})

		assert.NotNil(t, err)
		assert.ErrorIs(t, lib.ErrNoRecord, err)
	})
}

func TestUsersRepository_ConstrainsValidation(t *testing.T) {
	t.Parallel()
	if !INTEGRATION_TESTS {
		t.Errorf("skipping tests | INTEGRATION_TESTS = '%v'", INTEGRATION_TESTS)
		t.FailNow()
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockLogger := NewMockILogger(ctrl)

	t.Run("-- unique email", func(t *testing.T) {
		repo := sqlite3.NewRepository(dbConn, mockLogger)

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
		defer cancel()

		user := createValidUser(t)
		repo.CreateUser(ctx, user)

		_, err := repo.CreateUser(ctx, user)

		assert.NotNil(t, err, "want unique constraint error, got nil", err)
		assert.Equal(t, "email already exists", err.Error())
	})

	t.Run("-- unique username", func(t *testing.T) {
		repo := sqlite3.NewRepository(dbConn, mockLogger)

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
		defer cancel()

		user := createValidUser(t)
		repo.CreateUser(ctx, user)

		user.Email = randomEmail()

		_, err := repo.CreateUser(ctx, user)

		assert.NotNil(t, err, "want unique constraint error, got nil", err)
		assert.Equal(t, "username already exists", err.Error())
	})

	t.Run("-- empty username", func(t *testing.T) {
		repo := sqlite3.NewRepository(dbConn, mockLogger)

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
		defer cancel()

		user := createValidUser(t)
		user.Username = ""

		id, err := repo.CreateUser(ctx, user)
		assert.Negative(t, id)
		assert.NotNilf(t, err, "want CHECK constraint failed, got: '%v'", err)
	})
	t.Run("-- empty email", func(t *testing.T) {
		repo := sqlite3.NewRepository(dbConn, mockLogger)

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
		defer cancel()

		user := createValidUser(t)
		user.Email = ""

		id, err := repo.CreateUser(ctx, user)
		assert.Negative(t, id)
		assert.NotNilf(t, err, "want CHECK constraint failed, got: '%v'", err)
	})
	t.Run("-- empty password", func(t *testing.T) {
		repo := sqlite3.NewRepository(dbConn, mockLogger)

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
		defer cancel()

		user := createValidUser(t)
		user.Password = ""

		id, err := repo.CreateUser(ctx, user)
		assert.Negative(t, id)
		assert.NotNilf(t, err, "want CHECK constraint failed, got: '%v'", err)
	})
}

func TestUsersRepository_HashingPassword(t *testing.T) {
}

func TestUserRepository_ErrorHandling(t *testing.T) {
}

func TestUsersRepository_SQLInjection(t *testing.T) {
}
