package rest

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"asdf/internal/apps/users/internal/application"
	"asdf/internal/dtos"
	"asdf/internal/lib"
)

type UsersController struct {
	lib.IApp
	app *application.Users
}

func newUsersController(rootApp lib.IApp, app *application.Users) *UsersController {
	return &UsersController{
		rootApp,
		app,
	}
}

func (tc *UsersController) findUsersC(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), tc.ContextTimeout())
	defer cancel()

	users, err := tc.app.ListUsers(ctx, &dtos.FindUsersParams{})
	if err != nil {
		tc.ServerError(w, r, err)
	}
	tc.JSON(w, http.StatusOK, dtos.FindUsersResponse{
		Users: users,
	})
}

func (tc *UsersController) createUsersC(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), tc.ContextTimeout())
	defer cancel()

	cup := &dtos.CreateUsersParams{}

	if err := tc.DecodeJSON(w, r, cup); err != nil {
		tc.BadRequest(w, r, err)
		return
	}

	userID, err := tc.app.CreateUser(ctx, &dtos.CreateUsersParams{})
	if err != nil {
		tc.ServerError(w, r, err)
	}
	tc.JSON(w, http.StatusOK, dtos.CreateUserResponse{
		UserID: userID,
	})
}

func (tc *UsersController) getUsersC(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), tc.ContextTimeout())
	defer cancel()

	id, err := tc.ParseIntFromRequest("id", r)
	if err != nil {
		tc.BadRequest(w, r, err)
		return
	}

	gup := &dtos.GetUsersParams{
		ID: id,
	}
	user, err := tc.app.GetUser(ctx, gup)
	if err != nil {
		if errors.Is(err, lib.ErrNoRecord) {
			tc.NotFound(w, r, fmt.Sprintf("user with id %d not found", gup.ID))
		} else {
			tc.ServerError(w, r, err)
		}
		return

	}
	tc.JSON(w, http.StatusOK, dtos.GetUserResponse{
		User: user,
	})
}

func (tc *UsersController) updateUsersC(w http.ResponseWriter, r *http.Request) {}

func (tc *UsersController) deleteUsersC(w http.ResponseWriter, r *http.Request) {}
