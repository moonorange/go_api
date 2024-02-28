package thttp

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/moonorange/go_api/gen"
	"github.com/moonorange/go_api/services"
	"github.com/moonorange/go_api/terrors"
)

type (
	Server struct {
		HTTPServer *http.Server

		// Services used by the various HTTP routes.
		TaskService services.TaskService
		TagService  services.TagService
	}
)

// Make sure we conform to gen.ServerInterface
var _ gen.ServerInterface = (*Server)(nil)

// Wraps sending of an error in the Error format, and
// handling the failure to marshal that.
func sendError(w http.ResponseWriter, code int, message string) {
	castedCode := int32(code)
	todoErr := gen.Error{
		Code:    &castedCode,
		Message: &message,
	}
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(todoErr)
}

// LogError logs an error with the HTTP route information.
func LogError(r *http.Request, err error) {
	log.Printf("[http] error: %s %s: %s", r.Method, r.URL.Path, err)
}

// Error prints & optionally logs an error message.
func Error(w http.ResponseWriter, r *http.Request, err error) {
	// Extract error code & message.
	code, message := terrors.ErrorCode(err), terrors.ErrorMessage(err)

	// Print user message to response based on request accept header.
	switch r.Header.Get("Accept") {
	case "application/json":
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(ErrorStatusCode(code))
		json.NewEncoder(w).Encode(&ErrorResponse{Error: message})
	}
}

// lookup of application error codes to HTTP status codes.
var codes = map[string]int{
	terrors.ECONFLICT:       http.StatusConflict,
	terrors.EINVALID:        http.StatusBadRequest,
	terrors.ENOTFOUND:       http.StatusNotFound,
	terrors.ENOTIMPLEMENTED: http.StatusNotImplemented,
	terrors.EUNAUTHORIZED:   http.StatusUnauthorized,
	terrors.EINTERNAL:       http.StatusInternalServerError,
}

// ErrorStatusCode returns the associated HTTP status code for a WTF error code.
func ErrorStatusCode(code string) int {
	if v, ok := codes[code]; ok {
		return v
	}
	return http.StatusInternalServerError
}
