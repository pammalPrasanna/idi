package rest

import (
	"context"
	"errors"
	"net/http"

	"without-alias/internal/apps/todos/internal/application"
	"without-alias/internal/dtos"
	"without-alias/internal/lib"
)

type TodosController struct {
	lib.IApp
	app *application.Todos
}

func newTodosController(rootApp lib.IApp, app *application.Todos) *TodosController {
	return &TodosController{
		rootApp,
		app,
	}
}

func (tc *TodosController) findTodosH(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), tc.ContextTime())
	defer cancel()

	// parse -->  page, sort, filter etc in FindTodosParams

	todos, err := tc.app.FindTodos(ctx, &dtos.FindTodosParams{})
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

	tc.JSON(w, http.StatusOK, &FindTodosResponse{
		Todos: todos,
	})
}

func (tc *TodosController) createTodosH(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), tc.ContextTime())
	defer cancel()

	args := &dtos.CreateTodoParams{}
	if err := tc.DecodeJSON(w, r, args); err != nil {
		tc.BadRequest(w, r, err)
		return
	}

	id, err := tc.app.CreateTodo(ctx, args)
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

	tc.JSON(w, http.StatusOK, &CreateTodoResponse{
		TodoID: id,
	})
}

func (tc *TodosController) getTodosH(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), tc.ContextTime())
	defer cancel()

	id, err := tc.ParseIntFromRequest("id", r)
	if err != nil {
		tc.BadRequest(w, r, err)
		return
	}

	todo, err := tc.app.GetTodo(ctx, &dtos.GetTodoParams{
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

	tc.JSON(w, http.StatusOK, &GetTodoResponse{
		Todo: todo,
	})
}

func (tc *TodosController) updateTodosH(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), tc.ContextTime())
	defer cancel()

	id, err := tc.ParseIntFromRequest("id", r)
	if err != nil {
		tc.BadRequest(w, r, err)
		return
	}

	args := &dtos.UpdateTodoParams{}
	if err := tc.DecodeJSON(w, r, args); err != nil {
		tc.BadRequest(w, r, err)
		return
	}
	args.ID = id

	err = tc.app.UpdateTodo(ctx, args)
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

func (tc *TodosController) deleteTodosH(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), tc.ContextTime())
	defer cancel()

	id, err := tc.ParseIntFromRequest("id", r)
	if err != nil {
		tc.BadRequest(w, r, err)
		return
	}

	err = tc.app.DeleteTodo(ctx, &dtos.DeleteTodoParams{
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
