postgres:
	docker run --name postgres-12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

createdb:
	docker exec -it postgres-12 createdb --username=root --owner=root webserver

dropdb:
	sudo docker exec -it postgres-12 dropdb webserver

migrateup:
	migrate -path ./db/migrations -database "postgresql://root:secret@localhost:5432/webserver?sslmode=disable" up

migratedown:
	migrate -path ./db/migrations -database "postgresql://root:secret@localhost:5432/webserver?sslmode=disable" down

sqlc:
	sqlc generate

.PHONY: createdb dropdb postgres migrateup migratedown sqlc