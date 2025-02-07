package rest

import (
	"net/http"

	"asdf/internal/apps/users/internal/application"
	"asdf/internal/lib"
)

func RegisterRoutes(rootApp lib.IApp, app *application.Users) {
	ctrlr := NewUsersController(rootApp, app)
	mux := ctrlr.Mux()

	mux.HandlerFunc(http.MethodGet, "/users", ctrlr.FindUsersH)
	mux.HandlerFunc(http.MethodGet, "/users/:id", ctrlr.GetUsersH)
	mux.HandlerFunc(http.MethodPost, "/users", ctrlr.CreateUsersH)
	mux.HandlerFunc(http.MethodPut, "/users", ctrlr.UpdateUsersH)
	mux.HandlerFunc(http.MethodDelete, "/users", ctrlr.DeleteUsersH)
}
