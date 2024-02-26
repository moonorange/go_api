package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"sync"

	"github.com/moonorange/go_api/api/services"
	"github.com/moonorange/go_api/domain"
	"github.com/moonorange/go_api/gen"
	"github.com/moonorange/go_api/terrors"
	"github.com/moonorange/go_api/thttp"
)

type (
	TodoStore struct {
		TODOs map[string]gen.Task
		Lock  sync.Mutex
	}
	todoHandler struct {
		s     services.TodoService
		Store *TodoStore
	}
)

// Make sure we conform to gen.ServerInterface
var _ gen.ServerInterface = (*todoHandler)(nil)

func NewTodoHandler(services services.TodoService) gen.ServerInterface {
	return &todoHandler{
		s: services,
		Store: &TodoStore{
			TODOs: make(map[string]gen.Task)},
	}
}

// TasksCreate implements gen.ServerInterface.
func (t *todoHandler) TasksCreate(w http.ResponseWriter, r *http.Request) {
	var newTask domain.Todo
	err := json.NewDecoder(r.Body).Decode(&newTask)
	if err != nil {
		sendError(w, http.StatusBadRequest, "Invalid format for newTask")
		return
	}

	err = t.s.TasksCreate(context.Background(), &newTask)
	if err != nil {
		sendError(w, http.StatusBadRequest, err.Error())
		return
	}

	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(newTask)
}

// TasksDelete implements gen.ServerInterface.
func (t *todoHandler) TasksDelete(w http.ResponseWriter, r *http.Request, taskId string) {
	panic("unimplemented")
}

// TasksGetAll implements gen.ServerInterface.
func (t *todoHandler) TasksGetAll(w http.ResponseWriter, r *http.Request) {
	tasks, err := t.s.TasksGetAll(context.Background())
	if err != nil {
		sendError(w, http.StatusBadRequest, err.Error())
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(tasks)
}

// TasksRead implements gen.ServerInterface.
func (t *todoHandler) TasksRead(w http.ResponseWriter, r *http.Request, taskId string) {
	t.Store.Lock.Lock()
	defer t.Store.Lock.Unlock()

	v, ok := t.Store.TODOs[taskId]
	if !ok {
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(nil)
		return
	}

	_ = json.NewEncoder(w).Encode(v)
}

// TasksUpdate implements gen.ServerInterface.
func (t *todoHandler) TasksUpdate(w http.ResponseWriter, r *http.Request, taskId string) {
	panic("unimplemented")
}

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
		json.NewEncoder(w).Encode(&thttp.ErrorResponse{Error: message})
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
