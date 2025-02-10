package users_test

//go:generate mockgen -source .\generate_test.go -destination .\mocks_test.go  -package users_test

import (
	"context"
	"database/sql"
	"log/slog"
	"net/http"
	"without-alias/internal/dtos"
	"without-alias/internal/lib"
	"without-alias/internal/lib/auth"
	"time"

	"github.com/julienschmidt/httprouter"
)



type (
		IUsers interface {
			GetUser(ctx context.Context, arg *dtos.GetUserParams) (todo *dtos.User, err error)
			FindUsers(ctx context.Context, arg *dtos.FindUsersParams) (users []*dtos.User, err error)
			CreateUser(ctx context.Context, arg *dtos.CreateUserParams) (id int64, err error)
			UpdateUser(ctx context.Context, arg *dtos.UpdateUserParams) error
			DeleteUser(ctx context.Context, arg *dtos.DeleteUserParams) error
		}

		IUsersRepository interface {
			GetUser(ctx context.Context, arg *dtos.GetUserParams) (user *dtos.User, err error)
			FindUsers(ctx context.Context, arg *dtos.FindUsersParams) (users []*dtos.User, err error)
			CreateUser(ctx context.Context, arg *dtos.CreateUserParams) (id int64, err error)
			UpdateUser(ctx context.Context, arg *dtos.UpdateUserParams) error
			DeleteUser(ctx context.Context, arg *dtos.DeleteUserParams) error
		}

		IApp interface {
		Postgres() *sql.DB
		Logger() lib.ILogger

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
		GetUserById(id int) (*lib.IUser, error)
		SetUserByIDMethod(fn func(id int) (*lib.IUser, error))
	
		// rest helpers
		BadRequest(w http.ResponseWriter, r *http.Request, err error)         // 400
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

)

func addrOfStr(s string) *string { return &s }