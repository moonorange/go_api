package main

import (
	"net/http"

	"github.com/moonorange/go_api/openapi"
)

type TodoServer struct {
}

func (t TodoServer) TasksGetAll(w http.ResponseWriter, r *http.Request) {
	// our logic to retrieve all todos from a persistent layer
}

func (t TodoServer) TasksCreate(w http.ResponseWriter, r *http.Request) {
	// our logic to store the todo into a persistent layer
}

func (t TodoServer) TasksDelete(w http.ResponseWriter, r *http.Request, taskId string) {
	// our logic to delete a todo from the persistent layer
}

func (t TodoServer) TasksRead(w http.ResponseWriter, r *http.Request, taskId string) {
	// our logic to read the todo.
}

func (t TodoServer) TasksUpdate(w http.ResponseWriter, r *http.Request, taskId string) {
	// our logic to update the todo.
}

func main() {
	s := TodoServer{}
	h := openapi.Handler(s)

	http.ListenAndServe(":8000", h)
}
