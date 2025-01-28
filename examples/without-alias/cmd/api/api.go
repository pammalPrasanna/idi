package api

import (
	"without-alias/internal/apps/users"

	"without-alias/internal/apps/todos"

	"without-alias/internal/lib"
)

func Main(rootApp lib.IApp) error {

	users.InitApp(rootApp)

	todos.InitApp(rootApp)

	return rootApp.ServeHTTP()
}
