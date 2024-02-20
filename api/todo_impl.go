package api

import (
	"encoding/json"
	"net/http"
	"sync"

	"github.com/google/uuid"
	"github.com/moonorange/go_api/api/gen"
	"github.com/moonorange/go_api/application/usecase"
)

type (
	TodoStore struct {
		TODOs map[string]gen.Task
		Lock  sync.Mutex
	}
	todoAPI struct {
		uc    usecase.TodoUseCase
		Store *TodoStore
	}
)

// Make sure we conform to gen.ServerInterface
var _ gen.ServerInterface = (*todoAPI)(nil)

func NewTodoAPI(usecase usecase.TodoUseCase) gen.ServerInterface {
	return &todoAPI{
		uc: usecase,
		Store: &TodoStore{
			TODOs: make(map[string]gen.Task)},
	}
}

// TasksCreate implements gen.ServerInterface.
func (t *todoAPI) TasksCreate(w http.ResponseWriter, r *http.Request) {
	var newTODO gen.Task
	err := json.NewDecoder(r.Body).Decode(&newTODO)
	if err != nil {
		sendTodoError(w, http.StatusBadRequest, "Invalid format for newTODO")
		return
	}
	t.Store.Lock.Lock()
	defer t.Store.Lock.Unlock()

	newID := uuid.New().String()
	newTODO.Id = &newID
	t.Store.TODOs[newID] = newTODO

	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(newTODO)
}

// TasksDelete implements gen.ServerInterface.
func (t *todoAPI) TasksDelete(w http.ResponseWriter, r *http.Request, taskId string) {
	panic("unimplemented")
}

// TasksGetAll implements gen.ServerInterface.
func (t *todoAPI) TasksGetAll(w http.ResponseWriter, r *http.Request) {
	t.Store.Lock.Lock()
	defer t.Store.Lock.Unlock()

	var tasks []gen.Task

	for _, todo := range t.Store.TODOs {
		// Add all pets if we're not filtering
		tasks = append(tasks, todo)
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(tasks)
}

// TasksRead implements gen.ServerInterface.
func (t *todoAPI) TasksRead(w http.ResponseWriter, r *http.Request, taskId string) {
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
func (t *todoAPI) TasksUpdate(w http.ResponseWriter, r *http.Request, taskId string) {
	panic("unimplemented")
}

// sendTodoError wraps sending of an error in the Error format, and
// handling the failure to marshal that.
func sendTodoError(w http.ResponseWriter, code int, message string) {
	castedCode := int32(code)
	todoErr := gen.Error{
		Code:    &castedCode,
		Message: &message,
	}
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(todoErr)
}
