gen-todo:
	openapi-generator generate \
    -i todo.yaml \
    -g go \
    -o gen/todo \
	--additional-properties packageName=todo
