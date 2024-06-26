## Go API

This is a REST API server written in Go.

## Generate server stubs

Generate server stubs using `https://github.com/deepmap/oapi-codegen`

```sh
make gen-todo
```

## Directory structure

`gen`: auto-generated code from the YAML file that follows the swagger specification. This contains handlers for API

`services`: Business logic interface. Actual implementation lies in `*service_impl.go`

`thttp`:  Implementation of `ServerInterface` defined in the `gen` folder

`domain`: The application domain. Domain describes what the application does.

`db`: Migration files, db settings

## DB Migration

Using [pressly/goose](https://github.com/pressly/goose) for db migration

```sh
goose up
goose down
```

## Run server

Start necessary services on Docker

```sh
make up
```

Run API server

```sh
make run
```

Run server with Hot reloading

This command runs the server with hot reloading enabled for faster development iterations.

```sh
make air
```

## API Call Examples

Create Task

```sh
curl -d '{"description":"task3", "completed":false}' -H "Content-Type: application/json" -X POST http://localhost:8080/task
```

List Task

```sh
curl -H "Content-Type: application/json" -X GET http://localhost:8080/task
```

## Setting up a test environment

Create a test database and a user.

Apply migration to the test database.

```sh
make setup/test
```

## Monitoring

This project uses `Prometheus` and `Grafana` for monitoring.

### Prometheus

After running `make up`, go to http://localhost:9090/metrics to access and check the metrics collected by Prometheus.

### Grafana

After running `make up`, go to http://localhost:3000 to access Grafana's interactive dashboards and explore the metrics

You can sign in with the following credentials:
- username: admin
- password admin
