new.service.rest:
	mkdir -p app/service/definition
	goctl api -o app/service/definition/service.api

gen.service.rest:
	cd app/service; [ -f go.mod ] || go mod init service
	goctl api go -api app/service/definition/service.api -dir ./app/service/api/rest
	cd app/service; go mod tidy

run.service.rest:
	cd app/service; wire ./...; go mod tidy; go run api/rest/service.go api/rest/wire_gen.go -f api/rest/etc/service-api.yaml

install.golang.migrate.mysql:
	go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

create.mysql.migration:
	cd ./app/service; migrate create -ext sql -dir databases/mysqldb -seq ${NAME}

migrate.mysql.migration:
	cd .; migrate -database "mysql://${USERDB}:${PASSDB}@tcp(localhost:3306)/user" -path app/service/databases/mysqldb up

run.mysql.db:
	cd app/infra; docker-compose up -d