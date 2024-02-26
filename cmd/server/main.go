package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/moonorange/go_api/api/handlers"
	"github.com/moonorange/go_api/configs"
	"github.com/moonorange/go_api/gen"
	"github.com/moonorange/go_api/infra/mysql"
	middleware "github.com/oapi-codegen/nethttp-middleware"
)

func main() {
	port := flag.String("port", "8080", "Port for test HTTP server")
	flag.Parse()

	swagger, err := gen.GetSwagger()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading swagger spec\n: %s", err)
		os.Exit(1)
	}

	// Clear out the servers array in the swagger spec, that skips validating
	// that server names match. We don't know how this thing will be run.
	swagger.Servers = nil

	dsn := configs.GetDefaultDSN()
	db := mysql.NewDB(dsn)
	err = db.Open()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening db\n: %s", err)
		os.Exit(1)
	}

	// Create an instance of our handler which satisfies the generated interface
	s := mysql.NewTODOService(db)
	todoServer := handlers.NewTodoHandler(s)

	// This is how you set up a basic chi router
	r := chi.NewRouter()

	// Use our validation middleware to check all requests against the
	// OpenAPI schema.
	r.Use(middleware.OapiRequestValidator(swagger))

	// We now register our todoServer above as the handler for the interface
	gen.HandlerFromMux(todoServer, r)

	server := &http.Server{
		Handler: r,
		Addr:    net.JoinHostPort("0.0.0.0", *port),
	}
	fmt.Printf("Server listening on %s", server.Addr)

	// And we serve HTTP until the world ends.
	log.Fatal(server.ListenAndServe())
}
