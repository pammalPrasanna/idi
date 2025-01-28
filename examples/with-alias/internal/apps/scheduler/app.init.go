package scheduler

import (
	rest "with-alias/internal/apps/scheduler/internal/adapters/httprouter"
	"with-alias/internal/apps/scheduler/internal/adapters/sqlite3"
	"with-alias/internal/apps/scheduler/internal/application"
	"with-alias/internal/lib"
)

var schedulerApp *application.Scheduler

func SchedulerApp() *application.Scheduler {
	return schedulerApp
}

func InitApp(rootApp lib.IApp) {

	// create app repository
	schedulerRepo := sqlite3.NewRepository(rootApp.Sqlite3(), rootApp.Logger())
	// create app with repository
	schedulerApp = application.New(rootApp, schedulerRepo)

	// register routes if needed
	rest.RegisterRoutes(rootApp, schedulerApp)
}
