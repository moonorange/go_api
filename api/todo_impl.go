package api

import (
	"encoding/json"
	"net/http"
	"sync"

	"github.com/google/uuid"
	"github.com/moonorange/go_api/domain"
	"github.com/moonorange/go_api/domain/services"
)

type (
	TodoStore struct {
		TODOs map[domain.TaskId]domain.Todo
		Lock  sync.Mutex
	}
	todoAPI struct {
		s     services.TodoService
		Store *TodoStore
	}
)

// Make sure we conform to ServerInterface
var _ ServerInterface = (*todoAPI)(nil)

func NewTodoAPI(service services.TodoService) ServerInterface {
	return &todoAPI{
		s: service,
		Store: &TodoStore{
			TODOs: make(map[domain.TaskId]domain.Todo)},
	}
}

// TasksCreate implements ServerInterface.
func (t *todoAPI) TasksCreate(w http.ResponseWriter, r *http.Request) {
	var newTODO domain.Todo
	err := json.NewDecoder(r.Body).Decode(&newTODO)
	if err != nil {
		sendTodoError(w, http.StatusBadRequest, "Invalid format for newTODO")
		return
	}
	t.Store.Lock.Lock()
	defer t.Store.Lock.Unlock()

	newID := domain.TaskId(uuid.New())
	newTODO.ID = newID
	t.Store.TODOs[newID] = newTODO

	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(newTODO)
}

// TasksDelete implements ServerInterface.
func (t *todoAPI) TasksDelete(w http.ResponseWriter, r *http.Request, taskId string) {
	panic("unimplemented")
}

// TasksGetAll implements ServerInterface.
func (t *todoAPI) TasksGetAll(w http.ResponseWriter, r *http.Request) {
	t.Store.Lock.Lock()
	defer t.Store.Lock.Unlock()

	var tasks []domain.Todo

	for _, todo := range t.Store.TODOs {
		// Add all pets if we're not filtering
		tasks = append(tasks, todo)
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(tasks)
}

// TasksRead implements ServerInterface.
func (t *todoAPI) TasksRead(w http.ResponseWriter, r *http.Request, taskId string) {
	panic("unimplemented")
}

// TasksUpdate implements ServerInterface.
func (t *todoAPI) TasksUpdate(w http.ResponseWriter, r *http.Request, taskId string) {
	panic("unimplemented")
}

// sendTodoError wraps sending of an error in the Error format, and
// handling the failure to marshal that.
func sendTodoError(w http.ResponseWriter, code int, message string) {
	castedCode := int32(code)
	todoErr := Error{
		Code:    &castedCode,
		Message: &message,
	}
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(todoErr)
}
