package shard

import (
	"log"
	"mmo2/game/ecs"
	"mmo2/game/packets"
	"mmo2/pkg/event_utils"
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
func (s *Server) moveRequest(player *Player, event event_utils.Raw) {
	request, _ := packets.ParseMoveRequest(event)
	if !player.entity.Has(ecs.TypeTransform) {
		log.Printf("wrong: entity %d doesn't have position", player.entity.ID())
		return
	}
	transform := ecs.Get[*ecs.Transform](player.entity, ecs.TypeTransform)
	transform.X += request.Dx
	transform.Y += request.Dy
	player.peer.SendResponse(event, packets.NewAckRequest())
	s.BroadcastFiltered(
		packets.NewEntityMoved(player.entity.ID(), ecs.NewMoveTo(transform.X, transform.Y)),
		player.peer,
	)
}

func (s *Server) joinShardRequest(player *Player, event event_utils.Raw) {
	if player.entity != nil {
		player.peer.SendResponse(event, packets.NewRequestError("You already joined this shard"))
		return
	}
	request, _ := packets.ParseJoinShardRequest(event)
	entity := ecs.MainWorld.NewEntity()
	player.entity = entity
	entity.Add(
		ecs.NewTransform(32, 32, 0),
		ecs.NewLiving(request.Name, 128),
		ecs.NewCircle(32, request.Color),
	)
	player.peer.SendResponse(event, packets.NewJoinShardResponse(entity.ToBytes()))
	for _, entity := range ecs.MainWorld.Entities() {
		if entity.ID() == player.entity.ID() {
			continue
		}
		player.peer.SendBytes(entity.ToBytes())
	}
	s.BroadcastFiltered(packets.NewPlayerJoined(entity.ToBytes()), player.peer)
}
