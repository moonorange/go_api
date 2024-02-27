package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"

	"github.com/go-chi/chi/v5"
	"github.com/moonorange/go_api/configs"
	"github.com/moonorange/go_api/gen"
	"github.com/moonorange/go_api/infra/mysql"
	"github.com/moonorange/go_api/thttp"
	middleware "github.com/oapi-codegen/nethttp-middleware"
)

func main() {
	// Instantiate a new type to represent our application.
	// This type lets us shared setup code with our end-to-end tests.
	ctx, cancel := context.WithCancel(context.Background())
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() { <-c; cancel() }()

	m := NewMain()

	// Execute program.
	if err := m.Run(ctx); err != nil {
		m.Close()
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

// Main represents the program.
type Main struct {
	// Configuration path and parsed config data.
	Config     configs.Config
	HTTPServer *http.Server

	// SQLite database used by SQLite service implementations.
	DB *mysql.DB
}

// NewMain returns a new instance of Main.
func NewMain() *Main {
	return &Main{
		Config:     configs.DefaultConfig(),
		HTTPServer: &http.Server{},
		DB:         mysql.NewDB(configs.GetDefaultDSN()),
	}
}

// Run executes the program. The configuration should already be set up before
// calling this function.
func (m *Main) Run(ctx context.Context) (err error) {
	// Open the database. This will instantiate the MySQL connection
	if err := m.DB.Open(); err != nil {
		return fmt.Errorf("cannot open db: %w", err)
	}

	// Instantiate MySQL-backed services.
	taskService := mysql.NewTaskService(m.DB)
	tagService := mysql.NewTagService(m.DB)
	todoServer := thttp.NewServer(taskService, tagService)

	// This is how you set up a basic chi router
	r := chi.NewRouter()

	// Use our validation middleware to check all requests against the
	// OpenAPI schema.
	swagger, err := gen.GetSwagger()
	// Clear out the servers array in the swagger spec, that skips validating
	// that server names match. We don't know how this thing will be run.
	swagger.Servers = nil
	if err != nil {
		log.Fatal(err.Error())
	}
	r.Use(middleware.OapiRequestValidator(swagger))

	// We now register our todoServer above as the handler for the interface
	gen.HandlerFromMux(todoServer, r)

	m.HTTPServer.Addr = net.JoinHostPort("0.0.0.0", "8080")
	m.HTTPServer.Handler = r
	fmt.Printf("Server listening on %s", m.HTTPServer.Addr)

	// And we serve HTTP until the world ends.
	log.Fatal(m.HTTPServer.ListenAndServe())

	return nil
}

// Close gracefully stops the program.
func (m *Main) Close() error {
	if m.HTTPServer != nil {
		if err := m.HTTPServer.Close(); err != nil {
			return err
		}
	}
	if m.DB != nil {
		if err := m.DB.Close(); err != nil {
			return err
		}
	}
	return nil
}
