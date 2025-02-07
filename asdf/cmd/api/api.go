package api

import (
	"asdf/internal/apps/users"

	"asdf/internal/lib"
)

func Main(rootApp lib.IApp) error {

	users.InitApp(rootApp)

	return rootApp.ServeHTTP()
}
