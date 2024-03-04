package thttp

import (
	"encoding/json"
	"net/http"

	"github.com/moonorange/go_api/models"
)

// TasksCreate implements gen.ServerInterface.
func (t *Server) TasksCreate(w http.ResponseWriter, r *http.Request) {
	var newTask models.Task
	err := json.NewDecoder(r.Body).Decode(&newTask)
	if err != nil {
		sendError(w, http.StatusBadRequest, "Invalid format for newTask")
		return
	}

	err = t.TaskService.TasksCreate(r.Context(), &newTask)
	if err != nil {
		sendError(w, http.StatusBadRequest, err.Error())
		return
	}

	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(newTask)
}

// TasksDelete implements gen.ServerInterface.
func (t *Server) TasksDelete(w http.ResponseWriter, r *http.Request, taskId string) {
	panic("unimplemented")
}

// TasksGetAll implements gen.ServerInterface.
func (t *Server) TasksGetAll(w http.ResponseWriter, r *http.Request) {
	tasks, err := t.TaskService.TasksGetAll(r.Context())
	if err != nil {
		Error(w, r, err)
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(tasks)
}

// TasksRead implements gen.ServerInterface.
func (t *Server) TasksRead(w http.ResponseWriter, r *http.Request, taskId string) {
	panic("unimplemented")
}

// TasksUpdate implements gen.ServerInterface.
func (t *Server) TasksUpdate(w http.ResponseWriter, r *http.Request, taskId string) {
	panic("unimplemented")
}
