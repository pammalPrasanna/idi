package rest

import (
	"context"
	"errors"
	"net/http"

	"without-alias/internal/apps/users/internal/application"
	"without-alias/internal/dtos"
	"without-alias/internal/lib"
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

func (tc *UsersController) findUsersH(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), tc.ContextTime())
	defer cancel()

	// parse -->  page, sort, filter etc in FindUsersParams

	users, err := tc.app.FindUsers(ctx, &dtos.FindUsersParams{})
	if err != nil {
		// data error
		if errors.As(err, &lib.ErrInvalidData{}) {
			tc.UnprocessableEntity(w, r, err)
			return
		}
		// db error
		tc.ServerError(w, r, err)
		return
	}

	tc.JSON(w, http.StatusOK, &FindUsersResponse{
		Users: users,
	})
}

func (tc *UsersController) createUsersH(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), tc.ContextTime())
	defer cancel()

	args := &dtos.CreateUserParams{}
	if err := tc.DecodeJSON(w, r, args); err != nil {
		tc.BadRequest(w, r, err)
		return
	}

	id, err := tc.app.CreateUser(ctx, args)
	if err != nil {
		// data error
		if errors.As(err, &lib.ErrInvalidData{}) {
			e := err.(lib.ErrInvalidData)
			tc.UnprocessableEntity(w, r, &dtos.HTTPErrs{
				Errors: e.GetErrors(),
			})
			return
		}
	}

	tc.JSON(w, http.StatusOK, &CreateUserResponse{
		UserID: id,
	})
}

func (tc *UsersController) getUsersH(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), tc.ContextTime())
	defer cancel()

	id, err := tc.ParseIntFromRequest("id", r)
	if err != nil {
		tc.BadRequest(w, r, err)
		return
	}

	todo, err := tc.app.GetUser(ctx, &dtos.GetUserParams{
		ID: id,
	})
	if err != nil {
		if errors.Is(err, lib.ErrNoRecord) {
			tc.JSON(w, http.StatusNotFound, &dtos.HTTPErrMsg{
				Error: "todo not found",
			})
			return
		}
		tc.ServerError(w, r, err)
		return
	}

	tc.JSON(w, http.StatusOK, &GetUserResponse{
		User: todo,
	})
}

func (tc *UsersController) updateUsersH(w http.ResponseWriter, r *http.Request) {
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
		if errors.As(err, &lib.ErrInvalidData{}) {
			e := err.(lib.ErrInvalidData)
			tc.UnprocessableEntity(w, r, &dtos.HTTPErrs{
				Errors: e.GetErrors(),
			})
			return
		} else if errors.Is(err, lib.ErrNoRecord) {
			tc.NotFound(w, r)
			return
		}
		tc.ServerError(w, r, err)
		return
	}

	tc.JSON(w, http.StatusOK, &dtos.HTTPMsg{
		Message: "updated successfully",
	})
}

func (tc *UsersController) deleteUsersH(w http.ResponseWriter, r *http.Request) {
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
		if errors.Is(err, lib.ErrNoRecord) {
			tc.NotFound(w, r)
			return
		}
		tc.ServerError(w, r, err)
		return
	}

	tc.JSON(w, http.StatusOK, &dtos.HTTPMsg{
		Message: "deleted successfully",
	})
}
