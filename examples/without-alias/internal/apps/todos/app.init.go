package todos

import (
	rest "without-alias/internal/apps/todos/internal/adapters/httprouter"
	"without-alias/internal/apps/todos/internal/adapters/postgres"
	"without-alias/internal/apps/todos/internal/application"
	"without-alias/internal/lib"
)

var todosApp *application.Todos

func TodosApp() *application.Todos {
	return todosApp
}

func InitApp(rootApp lib.IApp) {

	// create app repository
	todosRepo := postgres.NewRepository(rootApp.Postgres(), rootApp.Logger())
	// create app with repository
	todosApp = application.New(rootApp, todosRepo)

	// register routes if needed
	rest.RegisterRoutes(rootApp, todosApp)
}
