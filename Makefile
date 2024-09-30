postgres:
	docker run --name postgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=password -d postgres:alpine

redis:
	docker run --name redis -d -p 6379:6379 redis:alpine redis-server --save 60 1

createdb:
	docker exec -it postgres createdb --username=root --owner=root murmur

dropdb:
	docker exec -it postgres dropdb murmur

recreate:
	make dropdb && make createdb

start:
	docker start postgres && docker start redis

# test:
# 	go test -v -cover ./service/... ./handler/...

lint:
	golangci-lint run

mock:
	mockery --all

fmt:
	go fmt murmur-server/...

swag:
	swag init

workflow:
	make fmt && make lint && make test