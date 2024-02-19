gen-todo:
	go run github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen --config api/oapi-codegen.yaml todo.yaml

run:
	go run main/main.go
