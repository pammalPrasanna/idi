package lib

import (
	"database/sql"
	"net/http"
	"time"

	infra "with-alias/internal/infrastructure"

	"with-alias/internal/lib/auth"

	"github.com/julienschmidt/httprouter"
)

type rootApp struct {
	*config

	// rest.go
	logger *slogLogger

	// server.go
	router *httprouter.Router
	server *http.Server

	sqlite3 *sql.DB

	// auth
	customGetUserByID func(id int) (*IUser, error)
	auth.IAuth
}

var (
	_        IApp = (*rootApp)(nil)
	_rootApp *rootApp
)

func RootApp() (*rootApp, error) {
	if _rootApp != nil {
		return _rootApp, nil
	}

	_rootApp = &rootApp{}

	cfg, err := configure()
	if err != nil {
		return nil, err
	}
	_rootApp.config = cfg

	_rootApp.logger = newLogger(nil)
	_rootApp.router = httprouter.New()

	// open DB connection(s)
	conn, err := infra.Sqlite3()
	if err != nil {
		return nil, err
	}
	_rootApp.sqlite3 = conn

	paseto, err := auth.NewPasetoMaker(_rootApp.tokenExpiration, _rootApp.symmetricKey, _rootApp.baseURL)
	if err != nil {
		return nil, err
	}
	_rootApp.IAuth = paseto

	server, err := newServer(_rootApp.port, _rootApp.router, _rootApp.logger)
	if err != nil {
		return nil, err
	}
	_rootApp.server = server

	// global middlewares
	server.Handler = _rootApp.recoverPanicM(server.Handler)
	server.Handler = _rootApp.loggerM(server.Handler)
	server.Handler = _rootApp.corsM(server.Handler)

	// static config
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	_rootApp.router.Handler(http.MethodGet, "/static/*filepath", http.StripPrefix("/static", fileServer))

	return _rootApp, nil
}

func (i *rootApp) Sqlite3() *sql.DB {
	return i.sqlite3
}

func (i *rootApp) Mux() *httprouter.Router {
	return i.router
}

func (i *rootApp) Logger() ILogger {
	return i.logger
}

func (i *rootApp) GetUserById(id int) (*IUser, error) {
	if i.customGetUserByID == nil {
		i.logger.Warn("unimplemented method: GetUserById")
		return &IUser{ID: 2}, nil
	} else {
		return i.customGetUserByID(id)
	}
}

func (i *rootApp) ContextTime() time.Duration {
	return i.contextTimeout
}
