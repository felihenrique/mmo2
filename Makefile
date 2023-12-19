shards:
	go run cmd/shard/server/main.go
shardc:
	go run cmd/shard/client/main.go
simc:
	go run cmd/sim/client/client.go
sims:
	go run cmd/sim/server/server.go
test:
	go test ./...
build-serialize:
	go build -o bin/serialize-generator cmd/generator/serialize/main.go
generate:
	go generate ./...