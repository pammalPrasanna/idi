package lib

import (
	"database/sql"
	"net/http"
	"time"

	infra "without-alias/internal/infrastructure"
	
	"without-alias/internal/lib/auth"
	
	"github.com/julienschmidt/httprouter"
)

type idi struct {
	*config

	// rest.go
	logger *slogLogger

	// server.go
	router *httprouter.Router
	server *http.Server


	postgres *sql.DB

	// auth
	customGetUserByID func(id int) (*IUser, error)
	auth.IAuth
}

var (
	_ IApp = (*idi)(nil)
	_idi *idi
)

func Idi() (*idi, error) {
	if _idi != nil {
		return _idi, nil
	}

	_idi = &idi{}

	cfg, err := configure()
	if err != nil {
		return nil, err
	}
	_idi.config = cfg

	_idi.logger = newLogger(nil)
	_idi.router = httprouter.New()

	// open DB connection(s)
	conn, err := infra.Postgres()
	if err != nil {
		return nil, err
	}
	_idi.postgres = conn


	
	paseto, err := auth.NewPasetoMaker(_idi.tokenExpiration, _idi.symmetricKey, _idi.baseURL)
	if err != nil {
		return nil, err
	}
	_idi.IAuth = paseto
	

	server, err := newServer(_idi.port, _idi.router, _idi.logger)
	if err != nil {
		return nil, err
	}
	_idi.server = server

	// global middlewares
	server.Handler = _idi.recoverPanicM(server.Handler)
	server.Handler = _idi.loggerM(server.Handler)
	server.Handler = _idi.corsM(server.Handler)

	// static config
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	_idi.router.Handler(http.MethodGet, "/static/*filepath", http.StripPrefix("/static", fileServer))

	return _idi, nil
}

func (i *idi) Postgres() *sql.DB {
	return i.postgres
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

func (i *idi) ContextTime() time.Duration {
	return i.contextTimeout
}
