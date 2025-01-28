package rest

import (
	"context"
	"errors"
	"net/http"

	"with-alias/internal/apps/scheduler/internal/application"
	"with-alias/internal/dtos"
	"with-alias/internal/lib"
)

type SchedulerController struct {
	lib.IApp
	app *application.Scheduler
}

func newSchedulerController(rootApp lib.IApp, app *application.Scheduler) *SchedulerController {
	return &SchedulerController{
		rootApp,
		app,
	}
}

func (tc *SchedulerController) findSchedulerH(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), tc.ContextTime())
	defer cancel()

	// parse -->  page, sort, filter etc in FindSchedulerParams

	scheduler, err := tc.app.FindScheduler(ctx, &dtos.FindSchedulerParams{})
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

	tc.JSON(w, http.StatusOK, &FindSchedulerResponse{
		Scheduler: scheduler,
	})
}

func (tc *SchedulerController) createSchedulerH(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), tc.ContextTime())
	defer cancel()

	args := &dtos.CreateSchedulerParams{}
	if err := tc.DecodeJSON(w, r, args); err != nil {
		tc.BadRequest(w, r, err)
		return
	}

	id, err := tc.app.CreateScheduler(ctx, args)
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

	tc.JSON(w, http.StatusOK, &CreateSchedulerResponse{
		SchedulerID: id,
	})
}

func (tc *SchedulerController) getSchedulerH(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), tc.ContextTime())
	defer cancel()

	id, err := tc.ParseIntFromRequest("id", r)
	if err != nil {
		tc.BadRequest(w, r, err)
		return
	}

	todo, err := tc.app.GetScheduler(ctx, &dtos.GetSchedulerParams{
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

	tc.JSON(w, http.StatusOK, &GetSchedulerResponse{
		Scheduler: todo,
	})
}

func (tc *SchedulerController) updateSchedulerH(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), tc.ContextTime())
	defer cancel()

	id, err := tc.ParseIntFromRequest("id", r)
	if err != nil {
		tc.BadRequest(w, r, err)
		return
	}

	args := &dtos.UpdateSchedulerParams{}
	if err := tc.DecodeJSON(w, r, args); err != nil {
		tc.BadRequest(w, r, err)
		return
	}
	args.ID = id

	err = tc.app.UpdateScheduler(ctx, args)
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

func (tc *SchedulerController) deleteSchedulerH(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), tc.ContextTime())
	defer cancel()

	id, err := tc.ParseIntFromRequest("id", r)
	if err != nil {
		tc.BadRequest(w, r, err)
		return
	}

	err = tc.app.DeleteScheduler(ctx, &dtos.DeleteSchedulerParams{
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
