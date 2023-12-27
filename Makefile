shard:
	go run cmd/shard/main.go
game:
	go run cmd/game/main.go
simc:
	go run cmd/sim/client/client.go
sims:
	go run cmd/sim/server/server.go
test:
	go test ./...
builds:
	go build -o bin/serialize-generator cmd/generator/serialize/main.go
gen:
	go generate ./...