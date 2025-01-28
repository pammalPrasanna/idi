package rest

import (
	"net/http"

	"with-alias/internal/apps/scheduler/internal/application"
	"with-alias/internal/lib"
)

func RegisterRoutes(rootApp lib.IApp, app *application.Scheduler) {
	ctrlr := newSchedulerController(rootApp, app)
	mux := ctrlr.Mux()

	mux.HandlerFunc(http.MethodGet, "/scheduler", ctrlr.findSchedulerH)
	mux.HandlerFunc(http.MethodGet, "/scheduler/:id", ctrlr.getSchedulerH)
	mux.HandlerFunc(http.MethodPost, "/scheduler", ctrlr.createSchedulerH)
	mux.HandlerFunc(http.MethodPut, "/scheduler", ctrlr.updateSchedulerH)
	mux.HandlerFunc(http.MethodDelete, "/scheduler", ctrlr.deleteSchedulerH)
}
