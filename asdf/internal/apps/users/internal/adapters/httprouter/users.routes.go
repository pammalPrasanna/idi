package rest

import (
	"net/http"

	"asdf/internal/apps/users/internal/application"
	"asdf/internal/lib"
)

func RegisterRoutes(rootApp lib.IApp, app *application.Users) {
	ctrlr := newUsersController(rootApp, app)
	mux := ctrlr.Mux()

	mux.HandlerFunc(http.MethodGet, "/users", ctrlr.findUsersC)
	mux.HandlerFunc(http.MethodGet, "/users/:id", ctrlr.getUsersC)
	mux.HandlerFunc(http.MethodPost, "/users", ctrlr.createUsersC)
	mux.HandlerFunc(http.MethodPut, "/users", ctrlr.updateUsersC)
	mux.HandlerFunc(http.MethodDelete, "/users", ctrlr.deleteUsersC)
}
