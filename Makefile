gen-todo:
	oapi-codegen --package=openapi --generate types,chi-server todo.yaml > oapi/openapi.gen.go
