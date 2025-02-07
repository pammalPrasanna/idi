package lib

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/json-iterator/go"
)

var jsonit = jsoniter.ConfigFastest

func (i *idi) JSON(w http.ResponseWriter, status int, data any) error {
	return i.JSONWithHeaders(w, status, data, nil)
}

func (i *idi) JSONWithHeaders(w http.ResponseWriter, status int, data any, headers http.Header) error {
	// js, err := json.Marshal(data)
	js, err := jsonit.Marshal(data)
	if err != nil {
		return err
	}

	// js = append(js, '\n')

	for key, value := range headers {
		w.Header()[key] = value
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)

	return nil
}

func (i *idi) DecodeJSON(w http.ResponseWriter, r *http.Request, dst interface{}) error {
	return i.decodeJSON(w, r, dst, false)
}

func (i *idi) DecodeJSONStrict(w http.ResponseWriter, r *http.Request, dst interface{}) error {
	return i.decodeJSON(w, r, dst, true)
}

func (i *idi) decodeJSON(w http.ResponseWriter, r *http.Request, dst interface{}, disallowUnknownFields bool) error {
	maxBytes := 1_048_576
	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))

	// dec := json.NewDecoder(r.Body)
	dec := jsonit.NewDecoder(r.Body)

	if disallowUnknownFields {
		dec.DisallowUnknownFields()
	}

	err := dec.Decode(dst)
	if err != nil {
		var syntaxError *json.SyntaxError
		var unmarshalTypeError *json.UnmarshalTypeError
		var invalidUnmarshalError *json.InvalidUnmarshalError

		switch {
		case errors.As(err, &syntaxError):
			return fmt.Errorf("body contains badly-formed JSON (at character %d)", syntaxError.Offset)

		case errors.Is(err, io.ErrUnexpectedEOF):
			return errors.New("body contains badly-formed JSON")

		case errors.As(err, &unmarshalTypeError):
			if unmarshalTypeError.Field != "" {
				return fmt.Errorf("body contains incorrect JSON type for field %q", unmarshalTypeError.Field)
			}
			return fmt.Errorf("body contains incorrect JSON type (at character %d)", unmarshalTypeError.Offset)

		case errors.Is(err, io.EOF):
			return errors.New("body must not be empty")

		case strings.HasPrefix(err.Error(), "json: unknown field "):
			fieldName := strings.TrimPrefix(err.Error(), "json: unknown field ")
			return fmt.Errorf("body contains unknown key %s", fieldName)

		case err.Error() == "http: request body too large":
			return fmt.Errorf("body must not be larger than %d bytes", maxBytes)

		case errors.As(err, &invalidUnmarshalError):
			i.logger.Error("unable to decode json", "decodeJSON", err)
			panic(err)

		default:
			return err
		}
	}

	err = dec.Decode(&struct{}{})
	if !errors.Is(err, io.EOF) {
		return errors.New("body must only contain a single JSON value")
	}

	return nil
}
