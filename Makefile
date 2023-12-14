shard:
	GOMAXPROCS=1 go run cmd/shard/shard.go
simc:
	GOMAXPROCS=4 go run cmd/sim/client/client.go
sims:
	GOMAXPROCS=4 go run cmd/sim/server/server.go
test:
	go test ./...
build-serialize:
	go build -o bin/serialize-generator cmd/generator/serialize/main.go
generate:
	go generate ./...