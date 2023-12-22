package shard

import (
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
	move := packets.ParseMoveRequest(event)
	tc, tok := player.entity.Get(ecs.TypePosition)
	if !tok {
		log.Printf("wrong: entity %d doesn't have position", player.entity.ID())
		return
	}
	position := tc.(*ecs.Position)
	position.X += move.Dx
	position.Y += move.Dy
	player.peer.SendResponse(event, packets.NewAckRequest())
	s.BroadcastFiltered(
		packets.NewEntityMoved(player.entity.ID(), position), player.peer,
	)
}

func (s *Server) joinShardRequest(player *Player, event events.Raw) {
	if player.entity != nil {
		player.peer.SendResponse(event, packets.NewRequestError("You already joined this shard"))
		return
	}
	request := packets.ParseJoinShardRequest(event)
	entity := s.world.NewEntity()
	player.entity = entity
	position := ecs.NewPosition(0, 0)
	movable := ecs.NewMovable(10)
	name := ecs.NewName(request.Name)
	playerCircle := ecs.NewPlayerCircle(40, request.Color)
	entity.Add(position, movable, name, playerCircle)
	player.peer.SendResponse(event, packets.NewJoinShardResponse(
		entity.ID(), position, movable, name, playerCircle,
	))
	s.BroadcastFiltered(packets.NewPlayerJoined(
		entity.ID(), position, name, movable, playerCircle,
	), player.peer)
}
