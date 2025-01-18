package lib

import (
	"database/sql"
	"net/http"
	"time"

	infra "asdf/internal/infrastructure"

	"github.com/julienschmidt/httprouter"
)

type idi struct {
	*configs

	// rest.go
	logger ILogger

	// server.go
	router *httprouter.Router
	server *http.Server

	// sqlite3
	sqlite3 *sql.DB

	// auth
	customGetUserByID func(id int) (*IUser, error)
}

var _ IApp = (*idi)(nil)

func Idi() (*idi, error) {
	idi := &idi{}

	cfg, err := configure()
	if err != nil {
		return nil, err
	}
	idi.configs = cfg
	idi.logger = newLogger(nil)

	// open DB connection(s)
	conn, err := infra.Sqlite3()
	if err != nil {
		return nil, err
	}
	idi.sqlite3 = conn

	idi.router = httprouter.New()
	server, err := newServer(idi.port, idi.router, idi.logger)
	if err != nil {
		return nil, err
	}
	idi.server = server

	// global middlewares

	server.Handler = idi.recoverPanicM(server.Handler)
	server.Handler = idi.loggerM(server.Handler)
	server.Handler = idi.corsM(server.Handler)

	// static config
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	idi.router.Handler(http.MethodGet, "/static/*filepath", http.StripPrefix("/static", fileServer))

	return idi, nil
}

func (i *idi) Sqlite3() *sql.DB {
	return i.sqlite3
}

func (i *idi) Mux() *httprouter.Router {
	return i.router
}

func (i *idi) Logger() ILogger {
	return i.logger
}

func (i *idi) GetUserById(id int) (*IUser, error) {
	if i.customGetUserByID == nil {
		i.logger.Warn("unimplemented method: GetUserById")
		return &IUser{ID: 2}, nil
	} else {
		return i.customGetUserByID(id)
	}
}

func (i *idi) ContextTimeout() time.Duration {
	return i.contextTimeout
}
