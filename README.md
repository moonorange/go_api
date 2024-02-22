# Go API

Generate server stubs using `https://github.com/deepmap/oapi-codegen`

## Generate code

```sh
make gen-todo
```

## Directory structure

`gen`: auto-generated code from the YAML file that follows the swagger specification

`services` Business logic

`repositories`: Database interaction logic

`handlers`: Handlers for each route

## DB Migration

Using pressly/goose[https://github.com/pressly/goose] for db migration
