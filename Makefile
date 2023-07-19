postgres:
	docker run --name postgreSql -e POSTGRES_USER=root -e POSTGRES_PASSWORD=123456 -p 5432:5432 -d postgres:latest 

createdb: 
	docker exec -it postgreSql createdb --username=root --owner=root simple_bank

dropdb: 
	docker exec -it postgreSql dropdb simple_bank

migrateup: 
	migrate -path db/migration -database "postgresql://root:123456@localhost:5432/simple_bank?sslmode=disable" -verbose up
migratedown: 
	migrate -path db/migration -database "postgresql://root:123456@localhost:5432/simple_bank?sslmode=disable" -verbose down

sqlc: 
	sqlc generate

test: 
	go test -v -cover ./...

.PHONY : postgres createdb dropdb migrateup migratedown sqlc test