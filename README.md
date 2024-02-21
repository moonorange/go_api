# Go API

Generate server and client stubs using `https://github.com/deepmap/oapi-codegen`

## Generate code

```sh
make gen-todo
```

## Directory structure

`gen`: auto-generated code from the YAML file that follows the swagger specification

`services` Business logic

`repositories`: Database interaction logic

`handlers`: Handlers for each route
