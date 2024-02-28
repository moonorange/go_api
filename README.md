# Go API

Generate server stubs using `https://github.com/deepmap/oapi-codegen`

## Generate code

```sh
make gen-todo
```

## Directory structure

`gen`: auto-generated code from the YAML file that follows the swagger specification

`services`: Business logic interface. Actual implementation lies in *service_impl.go

`handlers`: Handlers for each routing

`models`: Model object

## DB Migration

Using [pressly/goose](https://github.com/pressly/goose) for db migration

```sh
goose up
goose down
```

## API Call

Create Task

```sh
curl -d '{"description":"task3", "completed":false}' -H "Content-Type: application/json" -X POST http://localhost:8080/task
```

List Task

```sh
curl -H "Content-Type: application/json" -X GET http://localhost:8080/task
```

## Hot reloading

```sh
make air
```
