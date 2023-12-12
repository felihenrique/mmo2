shard:
	GOMAXPROCS=1 go run cmd/shard/shard.go
simc:
	GOMAXPROCS=1 go run cmd/sim/client/client.go
sims
	GOMAXPROCS=1 go run cmd/sim/server/server.go