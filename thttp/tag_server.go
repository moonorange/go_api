package thttp

import (
	"encoding/json"
	"net/http"
)

// TasksUpdate implements gen.ServerInterface.
func (t *Server) ListTags(w http.ResponseWriter, r *http.Request) {
	tasks, err := t.tagService.ListTags(r.Context())
	if err != nil {
		sendError(w, http.StatusBadRequest, err.Error())
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(tasks)
}
