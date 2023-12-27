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
build-sim:
	go build -o bin/simc ./cmd/sim/client/client.go && go build -o bin/sims ./cmd/sim/server/server.go && chmod +x bin/sims && chmod +x bin/simc
sim-docker:
	make build-sim && docker-compose up --build --remove-orphans