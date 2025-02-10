package users

import (
	rest "asdf/internal/apps/users/internal/adapters/httprouter"
	"asdf/internal/apps/users/internal/adapters/sqlite3"
	"asdf/internal/lib"
	"asdf/internal/apps/users/internal/application"
)

var usersApp *application.Users

func UsersApp() *application.Users {
 	return usersApp
}

func InitApp(rootApp lib.IApp) {
	
		// create app repository
		usersRepo := sqlite3.NewRepository(rootApp.Sqlite3(), rootApp.Logger())
		// create app with repository
		usersApp = application.New(rootApp, usersRepo)
	


	// register routes if needed
	rest.RegisterRoutes(rootApp, usersApp)
}
