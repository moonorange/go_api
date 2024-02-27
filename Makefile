gen-todo:
	go get github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen
	go run github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen --config gen/oapi-codegen.yaml todo.yaml

run:
	go run main/main.go

air:
	air --build.cmd "go build -o tmp/main cmd/server/main.go" --build.bin "./tmp/main"
