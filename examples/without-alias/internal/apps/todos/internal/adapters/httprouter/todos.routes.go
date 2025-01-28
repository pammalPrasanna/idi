package rest

import (
	"net/http"

	"without-alias/internal/apps/todos/internal/application"
	"without-alias/internal/lib"
)

func RegisterRoutes(rootApp lib.IApp, app *application.Todos) {
	ctrlr := newTodosController(rootApp, app)
	mux := ctrlr.Mux()

	mux.HandlerFunc(http.MethodGet, "/todos", ctrlr.findTodosH)
	mux.HandlerFunc(http.MethodGet, "/todos/:id", ctrlr.getTodosH)
	mux.HandlerFunc(http.MethodPost, "/todos", ctrlr.createTodosH)
	mux.HandlerFunc(http.MethodPut, "/todos", ctrlr.updateTodosH)
	mux.HandlerFunc(http.MethodDelete, "/todos", ctrlr.deleteTodosH)
}
