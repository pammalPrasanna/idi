package api

import (
	"asdf/internal/apps/users"

	"asdf/internal/lib"
)

func Main(rootApp lib.IApp) error {

	// handle DB close on application shutdown
	defer func() {
		if err := rootApp.Sqlite3().Close(); err != nil {
			rootApp.Logger().Error("unable to close Sqlite3", "Sqlite3.Close()", err)
		}
	}()

	// inject root app to other apps
	users.InitApp(rootApp)

	return rootApp.ServeHTTP()
}
