gen-petstore:
	openapi-generator generate \
		-i petstore.yaml \
		-g go \
		--additional-properties packageName=petstore \
		--type-mappings number=float64 \
		-o gen/petstore
