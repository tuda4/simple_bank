postgres:
	docker run --name postgreSql -e POSTGRES_USER=root -e POSTGRES_PASSWORD=123456 -p 5432:5432 -d postgres:latest 

createdb: 
	docker exec -it postgreSql createdb --username=root --owner=root simple_bank

dropdb: 
	docker exec -it postgreSql dropdb simple_bank

migrateup: 
	migrate -path db/migration -database "postgresql://root:123456@localhost:5432/simple_bank?sslmode=disable" -verbose up

migrateup1: 
	migrate -path db/migration -database "postgresql://root:123456@localhost:5432/simple_bank?sslmode=disable" -verbose up 1

migrateup2: 
	migrate -path db/migration -database "postgresql://root:123456@localhost:5432/simple_bank?sslmode=disable" -verbose up 2

migratedown: 
	migrate -path db/migration -database "postgresql://root:123456@localhost:5432/simple_bank?sslmode=disable" -verbose down

migratedown1: 
	migrate -path db/migration -database "postgresql://root:123456@localhost:5432/simple_bank?sslmode=disable" -verbose down 1

migratedown2: 
	migrate -path db/migration -database "postgresql://root:123456@localhost:5432/simple_bank?sslmode=disable" -verbose down 2

sqlc: 
	sqlc generate

server: 
	air -c .air.toml

mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/tuda4/simple_bank/db/sqlc Store

test: 
	go test -v -cover ./...

.PHONY : postgres createdb dropdb migrateup migratedown sqlc test mock migrateup1 migratedown1 migrateup2 migratedown2