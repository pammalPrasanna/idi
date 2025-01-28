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
	logger ILogger

	// server.go
	router *httprouter.Router
	server *http.Server

	postgres *sql.DB

	// auth
	customGetUserByID func(id int) (*IUser, error)
	auth.IAuth
}

var _ IApp = (*idi)(nil)

func Idi() (*idi, error) {
	idi := &idi{}

	cfg, err := configure()
	if err != nil {
		return nil, err
	}
	idi.config = cfg

	idi.logger = newLogger(nil)
	idi.router = httprouter.New()

	// open DB connection(s)
	conn, err := infra.Postgres()
	if err != nil {
		return nil, err
	}
	idi.postgres = conn

	jwt, err := auth.NewJWTMaker(idi.tokenExpiration, idi.jwtSecret, idi.baseURL)
	if err != nil {
		return nil, err
	}
	idi.IAuth = jwt

	server, err := newServer(idi.port, idi.router, idi.logger)
	if err != nil {
		return nil, err
	}
	idi.server = server

	// global middlewares
	server.Handler = idi.authenticateM(server.Handler)
	server.Handler = idi.recoverPanicM(server.Handler)
	server.Handler = idi.loggerM(server.Handler)
	server.Handler = idi.corsM(server.Handler)

	// static config
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	idi.router.Handler(http.MethodGet, "/static/*filepath", http.StripPrefix("/static", fileServer))

	return idi, nil
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
