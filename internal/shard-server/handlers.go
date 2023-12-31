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
*/
func (s *Server) moveRequest(player *Player, event event_utils.Raw) {
	request, _ := packets.ParseMoveRequest(event)
	if !player.entity.Has(ecs.TypeTransform) {
		log.Printf("wrong: entity %d doesn't have transform", player.entity.ID())
		return
	}
	transform := ecs.Get[*ecs.Transform](player.entity, ecs.TypeTransform)
	var moveTo *ecs.MoveTo
	if player.entity.Has(ecs.TypeMoveTo) {
		moveTo = ecs.Get[*ecs.MoveTo](player.entity, ecs.TypeMoveTo)
	} else {
		moveTo = ecs.NewMoveTo(transform.X, transform.Y)
	}
	moveTo.X += request.Dx
	moveTo.Y += request.Dy
	player.entity.Add(moveTo)
	player.peer.SendResponse(event, packets.NewAckRequest())
	s.BroadcastFiltered(packets.NewEntityMoved(player.entity.ID(), moveTo), player.peer)
}

func (s *Server) joinShardRequest(player *Player, event event_utils.Raw) {
	if player.entity != nil {
		player.peer.SendResponse(event, packets.NewRequestError("You already joined this shard"))
		return
	}
	request, _ := packets.ParseJoinShardRequest(event)
	playerEntity := ecs.MainWorld.NewEntity()
	player.entity = playerEntity
	playerEntity.Add(
		ecs.NewTransform(32, 32, 0),
		ecs.NewLiving(request.Name, 90),
		ecs.NewCircle(32, request.Color),
	)
	player.peer.SendResponse(event, packets.NewJoinShardResponse(playerEntity.ToBytes()))

	for _, entity := range ecs.MainWorld.Entities() {
		if entity.ID() == player.entity.ID() {
			continue
		}
		player.peer.SendBytes(
			packets.NewEntityCreated(entity.ToBytes()).ToBytes(0),
		)
	}
	s.BroadcastFiltered(packets.NewEntityCreated(playerEntity.ToBytes()), player.peer)
}
