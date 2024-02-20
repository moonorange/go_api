# Go API

Generate server and client stubs using `https://github.com/deepmap/oapi-codegen`

## Generate code

```sh
make gen-todo
```

## Directory structure

`api/gen` has auto-generated code from the YAML file that follows the swagger specification

`api/**impl` has the implementation of the server interface of the generated code

You write the application's use case logic in the `application/usecase`

You write data-related code in the `domain/repository`

You return data to the client in the `interface/*/handler`
