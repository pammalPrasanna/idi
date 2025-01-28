package rest

import (
	"net/http"

	"with-alias/internal/apps/users/internal/application"
	"with-alias/internal/lib"
)

func RegisterRoutes(rootApp lib.IApp, app *application.Users) {
	ctrlr := newUsersController(rootApp, app)
	mux := ctrlr.Mux()

	mux.HandlerFunc(http.MethodGet, "/users", ctrlr.findUsersH)
	mux.HandlerFunc(http.MethodGet, "/users/:id", ctrlr.getUsersH)
	mux.HandlerFunc(http.MethodPost, "/users", ctrlr.createUsersH)
	mux.HandlerFunc(http.MethodPut, "/users", ctrlr.updateUsersH)
	mux.HandlerFunc(http.MethodDelete, "/users", ctrlr.deleteUsersH)
}
