package lib

import (
	"fmt"
	"log/slog"
	"net/http"
	"runtime/debug"
	"strconv"
	"strings"

	"github.com/julienschmidt/httprouter"
)

func (i *rootApp) reportServerError(r *http.Request, err error) {
	var (
		message = err.Error()
		method  = r.Method
		url     = r.URL.String()
		trace   = string(debug.Stack())
	)

	requestAttrs := slog.Group("request", "method", method, "url", url)
	i.logger.Error(message, requestAttrs, "trace", trace)
}

func (i *rootApp) errorMessage(w http.ResponseWriter, r *http.Request, status int, message string, headers http.Header) {
	message = strings.ToUpper(message[:1]) + message[1:]

	err := i.JSONWithHeaders(w, status, map[string]string{"error": message}, headers)
	if err != nil {
		i.reportServerError(r, err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (i *rootApp) ServerError(w http.ResponseWriter, r *http.Request, err error) {
	i.reportServerError(r, err)

	message := "The server encountered a problem and could not process your request"
	i.errorMessage(w, r, http.StatusInternalServerError, message, nil)
}

func (i *rootApp) NotFound(w http.ResponseWriter, r *http.Request) {
	message := "The requested resource could not be found"
	i.errorMessage(w, r, http.StatusNotFound, message, nil)
}

func (i *rootApp) MethodNotAllowed(w http.ResponseWriter, r *http.Request) {
	message := fmt.Sprintf("The %s method is not supported for this resource", r.Method)
	i.errorMessage(w, r, http.StatusMethodNotAllowed, message, nil)
}

func (i *rootApp) BadRequest(w http.ResponseWriter, r *http.Request, data any) {
	jerr := i.JSON(w, http.StatusBadRequest, data)
	if jerr != nil {
		i.ServerError(w, r, jerr)
	}
}

func (i *rootApp) UnprocessableEntity(w http.ResponseWriter, r *http.Request, data any) {
	jerr := i.JSON(w, http.StatusUnprocessableEntity, data)
	if jerr != nil {
		i.ServerError(w, r, jerr)
	}
}

func (i *rootApp) InvalidAuthenticationToken(w http.ResponseWriter, r *http.Request) {
	headers := make(http.Header)
	headers.Set("WWW-Authenticate", "Bearer")

	i.errorMessage(w, r, http.StatusUnauthorized, "Invalid authentication token", headers)
}

func (i *rootApp) AuthenticationRequired(w http.ResponseWriter, r *http.Request) {
	i.errorMessage(w, r, http.StatusUnauthorized, "You must be authenticated to access this resource", nil)
}

func (i *rootApp) ParseIntFromRequest(name string, r *http.Request) (int64, error) {
	params := httprouter.ParamsFromContext(r.Context())
	id, err := strconv.ParseInt(params.ByName("id"), 10, 64)
	if err != nil || id < 1 {
		return 0, ErrInvalidParameterID
	}
	return id, nil
}
