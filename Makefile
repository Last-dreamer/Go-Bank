postgres:
	docker run --name postgres01 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=dreamer -d postgres:15-alpine

createdb:
	docker exec -it postgres01 createdb --username=root --owner=root bank

dropdb:
	docker exec -it postgres01 dropdb bank

migrateup:
	migrate -path db/migration -database "postgresql://root:dreamer@localhost:5432/bank?sslmode=disable" --verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:dreamer@localhost:5432/bank?sslmode=disable" --verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

.PHONY: postgres createdb migrateup migratedown sqlc test server