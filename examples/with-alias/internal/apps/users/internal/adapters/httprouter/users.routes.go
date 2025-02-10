package rest

import (
	"net/http"

	"with-alias/internal/apps/users/internal/application"
	"with-alias/internal/lib"

)

func RegisterRoutes(rootApp lib.IApp, app *application.Users) {
	ctrlr := NewUsersController(rootApp, app)
	mux := ctrlr.Mux()

	mux.HandlerFunc(http.MethodGet, "/users", ctrlr.FindUsersH)
	mux.HandlerFunc(http.MethodGet, "/users/:id", ctrlr.GetUserH)
	mux.HandlerFunc(http.MethodPost, "/users", ctrlr.CreateUserH)
	mux.HandlerFunc(http.MethodPatch, "/users", ctrlr.PatchUserH)
	mux.HandlerFunc(http.MethodDelete, "/users", ctrlr.DeleteUserH)
}
