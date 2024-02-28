gen-todo:
	go get github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen
	go run github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen --config gen/oapi-codegen.yaml todo.yaml

run:
	go run main/main.go

air:
	air --build.cmd "go build -o tmp/main cmd/server/main.go" --build.bin "./tmp/main"

setup/test:
	mysql -h ${MYSQL_HOST} -u root -P ${MYSQL_PORT} -p${MYSQL_ROOT_PASSWORD} -e "CREATE DATABASE IF NOT EXISTS ${MYSQL_TEST_DATABASE};"
	mysql -h ${MYSQL_HOST} -u root -P ${MYSQL_PORT} -p${MYSQL_ROOT_PASSWORD} -e "CREATE USER IF NOT EXISTS '${MYSQL_TEST_USER}'@'%' IDENTIFIED BY '${MYSQL_PASSWORD}';"
	mysql -h ${MYSQL_HOST} -u root -P ${MYSQL_PORT} -p${MYSQL_ROOT_PASSWORD} -e "GRANT ALL PRIVILEGES ON ${MYSQL_TEST_DATABASE}.* TO '${MYSQL_TEST_USER}'@'%';"
	GOOSE_DBSTRING="${MYSQL_USER}:${MYSQL_PASSWORD}@tcp(${MYSQL_HOST}:${MYSQL_PORT})/${MYSQL_TEST_DATABASE}?parseTime=true" && goose up
