package api

import (
	"with-alias/internal/apps/users"

	"with-alias/internal/apps/scheduler"

	"with-alias/internal/lib"
)

func Main(rootApp lib.IApp) error {

	users.InitApp(rootApp)

	scheduler.InitApp(rootApp)

	return rootApp.ServeHTTP()
}
