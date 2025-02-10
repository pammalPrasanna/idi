package users

import (
	rest "with-alias/internal/apps/users/internal/adapters/httprouter"
	"with-alias/internal/apps/users/internal/adapters/sqlite3"
	"with-alias/internal/lib"
	"with-alias/internal/apps/users/internal/application"
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
