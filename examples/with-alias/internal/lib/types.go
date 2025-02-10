package lib

import (
	"database/sql"
	"log/slog"
	"net/http"
	"time"

	"with-alias/internal/lib/auth"

	"github.com/julienschmidt/httprouter"
)

type (
	IApp interface {
		Sqlite3() *sql.DB
		Logger() ILogger

		Mux() *httprouter.Router
		ServeHTTP() error
		
		// context
		ContextTime() time.Duration

		// JSON helpers
		JSON(w http.ResponseWriter, status int, data any) error
		JSONWithHeaders(w http.ResponseWriter, status int, data any, headers http.Header) error
		DecodeJSON(w http.ResponseWriter, r *http.Request, dst interface{}) error
		DecodeJSONStrict(w http.ResponseWriter, r *http.Request, dst interface{}) error

	
		// Auth helpers
		Hash(plaintextPassword string) (string, error)
		CompareHashAndPassword(plaintextPassword, hashedPassword string) (bool, error)
		VerifyToken(token string) (string, error)
		CreateToken(userID int) (*auth.Token, error)
		RequireAuthenticatedUser(next http.HandlerFunc) http.HandlerFunc
		GetUserById(id int) (*IUser, error)
		SetUserByIDMethod(fn func(id int) (*IUser, error))
	
		// rest helpers
		BadRequest(w http.ResponseWriter, r *http.Request, data any)           // 400
		AuthenticationRequired(w http.ResponseWriter, r *http.Request)        // 401
		InvalidAuthenticationToken(w http.ResponseWriter, r *http.Request)    // 401
		NotFound(w http.ResponseWriter, r *http.Request)                      // 404
		MethodNotAllowed(w http.ResponseWriter, r *http.Request)              // 405
		UnprocessableEntity(w http.ResponseWriter, r *http.Request, data any) // 422
		ServerError(w http.ResponseWriter, r *http.Request, err error)        // 500
		
		ParseIntFromRequest(name string, r *http.Request) (int64, error)
	}

	ILogger interface {
		Debug(msg string, keysAndValues ...interface{})
		Info(msg string, keysAndValues ...interface{})
		Warn(msg string, keysAndValues ...interface{})
		Error(msg string, keysAndValues ...interface{})
		Fatal(msg string, keysAndValues ...interface{})
		Handler() slog.Handler
	}

	IUser struct {
		ID int64
	}
)