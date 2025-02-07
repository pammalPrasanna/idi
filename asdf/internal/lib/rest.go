package lib

import (
	"fmt"
	"log/slog"
	"net/http"
	"runtime/debug"
	"strings"

	"errors"
	"github.com/julienschmidt/httprouter"
	"strconv"
)

func (i *idi) reportServerError(r *http.Request, err error) {
	var (
		message = err.Error()
		method  = r.Method
		url     = r.URL.String()
		trace   = string(debug.Stack())
	)

	requestAttrs := slog.Group("request", "method", method, "url", url)
	i.logger.Error(message, requestAttrs, "trace", trace)
}

func (i *idi) errorMessage(w http.ResponseWriter, r *http.Request, status int, message string, headers http.Header) {
	message = strings.ToUpper(message[:1]) + message[1:]

	err := i.JSONWithHeaders(w, status, map[string]string{"Error": message}, headers)
	if err != nil {
		i.reportServerError(r, err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (i *idi) ServerError(w http.ResponseWriter, r *http.Request, err error) {
	i.reportServerError(r, err)

	message := "The server encountered a problem and could not process your request"
	i.errorMessage(w, r, http.StatusInternalServerError, message, nil)
}

func (i *idi) NotFound(w http.ResponseWriter, r *http.Request) {
	message := "The requested resource could not be found"
	i.errorMessage(w, r, http.StatusNotFound, message, nil)
}

func (i *idi) MethodNotAllowed(w http.ResponseWriter, r *http.Request) {
	message := fmt.Sprintf("The %s method is not supported for this resource", r.Method)
	i.errorMessage(w, r, http.StatusMethodNotAllowed, message, nil)
}

func (i *idi) BadRequest(w http.ResponseWriter, r *http.Request, err error) {
	i.errorMessage(w, r, http.StatusBadRequest, err.Error(), nil)
}

func (i *idi) UnprocessableEntity(w http.ResponseWriter, r *http.Request, data any) {
	jerr := i.JSON(w, http.StatusUnprocessableEntity, data)
	if jerr != nil {
		i.ServerError(w, r, jerr)
	}
}

func (i *idi) InvalidAuthenticationToken(w http.ResponseWriter, r *http.Request) {
	headers := make(http.Header)
	headers.Set("WWW-Authenticate", "Bearer")

	i.errorMessage(w, r, http.StatusUnauthorized, "Invalid authentication token", headers)
}

func (i *idi) AuthenticationRequired(w http.ResponseWriter, r *http.Request) {
	i.errorMessage(w, r, http.StatusUnauthorized, "You must be authenticated to access this resource", nil)
}

func (i *idi) ParseIntFromRequest(name string, r *http.Request) (int64, error) {
	params := httprouter.ParamsFromContext(r.Context())
	id, err := strconv.ParseInt(params.ByName("id"), 10, 64)
	if err != nil || id < 1 {
		return 0, errors.New("invalid id parameter")
	}
	return id, nil
}
