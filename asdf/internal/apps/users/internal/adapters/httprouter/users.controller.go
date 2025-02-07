package rest

import (
	"context"
	"errors"
	"net/http"

	"asdf/internal/apps/users/internal/application"
	"asdf/internal/dtos"
	"asdf/internal/lib"
)

type UsersController struct {
	lib.IApp
	app *application.Users
}

func NewUsersController(rootApp lib.IApp, app *application.Users) *UsersController {
	return &UsersController{
		rootApp,
		app,
	}
}

func (tc *UsersController) FindUsersH(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), tc.ContextTime())
	defer cancel()

	// parse -->  page, sort, filter etc in FindUsersParams

	users, err := tc.app.FindUsers(ctx, &dtos.FindUsersParams{})
	if err != nil {
		tc.handleError(w, r, err)
		return
	}

	tc.JSON(w, http.StatusOK, &FindUsersResponse{
		Users: users,
	})
}

func (tc *UsersController) CreateUsersH(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), tc.ContextTime())
	defer cancel()

	args := &dtos.CreateUserParams{}
	if err := tc.DecodeJSON(w, r, args); err != nil {
		tc.BadRequest(w, r, err)
		return
	}

	id, err := tc.app.CreateUser(ctx, args)
	if err != nil {
		tc.handleError(w, r, err)
		return
	}

	tc.JSON(w, http.StatusOK, &CreateUserResponse{
		UserID: id,
	})
}

func (tc *UsersController) GetUsersH(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), tc.ContextTime())
	defer cancel()

	id, err := tc.ParseIntFromRequest("id", r)
	if err != nil {
		tc.BadRequest(w, r, err)
		return
	}

	user, err := tc.app.GetUser(ctx, &dtos.GetUserParams{
		ID: id,
	})
	if err != nil {
		tc.handleError(w, r, err)
		return
	}

	tc.JSON(w, http.StatusOK, &GetUserResponse{
		User: user,
	})
}

func (tc *UsersController) UpdateUsersH(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), tc.ContextTime())
	defer cancel()

	id, err := tc.ParseIntFromRequest("id", r)
	if err != nil {
		tc.BadRequest(w, r, err)
		return
	}

	args := &dtos.UpdateUserParams{}
	if err := tc.DecodeJSON(w, r, args); err != nil {
		tc.BadRequest(w, r, err)
		return
	}
	args.ID = id

	err = tc.app.UpdateUser(ctx, args)
	if err != nil {
		tc.handleError(w, r, err)
		return
	}

	tc.JSON(w, http.StatusOK, &dtos.HTTPMsg{
		Message: "updated successfully",
	})
}

func (tc *UsersController) DeleteUsersH(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), tc.ContextTime())
	defer cancel()

	id, err := tc.ParseIntFromRequest("id", r)
	if err != nil {
		tc.BadRequest(w, r, err)
		return
	}

	err = tc.app.DeleteUser(ctx, &dtos.DeleteUserParams{
		ID: id,
	})
	if err != nil {
		tc.handleError(w, r, err)
		return
	}

	tc.JSON(w, http.StatusOK, &dtos.HTTPMsg{
		Message: "deleted successfully",
	})
}

func (tc *UsersController) handleError(w http.ResponseWriter, r *http.Request, err error) {
	switch {
	case errors.As(err, &lib.ErrInvalidData{}):
		e := err.(lib.ErrInvalidData)
		tc.UnprocessableEntity(w, r, &dtos.HTTPErrs{
			Errors: e.GetErrors(),
		})
		return
	case errors.Is(err, lib.ErrNoRecord):
		tc.NotFound(w, r)
		return
	default:
		tc.ServerError(w, r, err)
		return
	}
}
