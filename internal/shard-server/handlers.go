package shard

import (
	"fmt"
	"log"
	"mmo2/game/ecs"
	"mmo2/game/packets"
	"mmo2/pkg/events"
)

/*
TODO:
- Checa se a posição de destino da entidade contem: outra entidade, colisor do mapa
- Caso não tenha adiciona dx, dy à posição do jogador
- Caso tenha se movido, lê o spatial map para detectar as outras entidades do tipo player
próximo do jogador
- Envia o evento EntityMoved para todos os jogadores em um radio de 6 seções
(a tela do jogador 4 x 4 seções)
*/
func (s *Server) moveRequest(player *Player, event events.Raw) {
	request := packets.ParseMoveRequest(event)
	if !player.entity.Has(ecs.TypeTransform) {
		log.Printf("wrong: entity %d doesn't have position", player.entity.ID())
		return
	}
	transform := ecs.Get[*ecs.Transform](player.entity, ecs.TypeTransform)
	transform.X = request.Move.FinalX
	transform.Y = request.Move.FinalY
	player.peer.SendResponse(event, packets.NewAckRequest())
	fmt.Printf("Move: %f, %f, Final: %f, %f \n",
		request.Move.QuantityX, request.Move.QuantityY, request.Move.FinalX, request.Move.FinalY)
	s.BroadcastFiltered(
		packets.NewEntityMoved(player.entity.ID(), request.Move),
		player.peer,
	)
}

func (s *Server) joinShardRequest(player *Player, event events.Raw) {
	if player.entity != nil {
		player.peer.SendResponse(event, packets.NewRequestError("You already joined this shard"))
		return
	}
	request := packets.ParseJoinShardRequest(event)
	entity := ecs.MainWorld.NewEntity()
	player.entity = entity
	transform := ecs.NewTransform(0, 0, 0)
	living := ecs.NewLiving(request.Name, 10)
	playerCircle := ecs.NewCircle(40, request.Color)
	entity.Add(transform, living, playerCircle)
	player.peer.SendResponse(event, packets.NewJoinShardResponse(
		entity.ID(), transform, living, playerCircle,
	))
	for _, entity := range ecs.MainWorld.Entities() {
		if entity.ID() == player.entity.ID() {
			continue
		}
		player.peer.SendBytes(packets.NewPlayerJoined(
			entity.ID(), ecs.Get[*ecs.Transform](entity, ecs.TypeTransform),
			ecs.Get[*ecs.Living](entity, ecs.TypeLiving), ecs.Get[*ecs.Circle](entity, ecs.TypeCircle),
		).ToBytes(0))
	}
	s.BroadcastFiltered(packets.NewPlayerJoined(
		entity.ID(), transform, living, playerCircle,
	), player.peer)
}
