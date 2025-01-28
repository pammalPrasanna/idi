package users

import (
	rest "without-alias/internal/apps/users/internal/adapters/httprouter"
	"without-alias/internal/apps/users/internal/adapters/postgres"
	"without-alias/internal/apps/users/internal/application"
	"without-alias/internal/lib"
)

var usersApp *application.Users

func UsersApp() *application.Users {
	return usersApp
}

func InitApp(rootApp lib.IApp) {

	// create app repository
	usersRepo := postgres.NewRepository(rootApp.Postgres(), rootApp.Logger())
	// create app with repository
	usersApp = application.New(rootApp, usersRepo)

	// register routes if needed
	rest.RegisterRoutes(rootApp, usersApp)
}
