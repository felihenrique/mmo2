package shard

import (
	"log"
	"mmo2/pkg/events"
	"mmo2/pkg/game"
	"mmo2/pkg/gsp"
)

type ICommand interface {
	Execute()
}
type BroadcastFunc = func(event events.ISerializable, eventId int16)
type BroadcastFilteredFunc = func(event events.ISerializable, eventId int16, filterPeer *gsp.TcpPeer)

type MoveCommand struct {
	event             events.Move
	eventId           int16
	world             *game.World
	player            Player
	broadcastFiltered BroadcastFilteredFunc
}

/*
TODO:
- Checa se a posição de destino da entidade contem: outra entidade, colisor do mapa
- Caso não tenha adiciona dx, dy à posição do jogador
- Caso tenha se movido, lê o spatial map para detectar as outras entidades do tipo player
próximo do jogador
- Envia o evento EntityMoved para todos os jogadores em um radio de 6 seções
(a tela do jogador 4 x 4 seções)
*/
func (c *MoveCommand) Execute() {
	entity := c.world.GetEntity(c.player.entityId)
	if !entity.Has(game.TypeTransform) || !entity.Has(game.TypeMovable) {
		log.Printf("move command error: entity %d doesn't have transform or movable", entity.ID())
		return
	}
	transform := entity.Get(game.TypeTransform).(*game.Transform)
	movable := entity.Get(game.TypeMovable).(*game.Movable)
	transform.X += c.event.Dx
	transform.Y += c.event.Dy
	movedEvent := events.EntityMoved{
		NewX:     transform.X,
		NewY:     transform.Y,
		EntityId: c.player.entityId,
		Velocity: movable.Velocity,
	}
	c.broadcastFiltered(&movedEvent, 0, c.player.peer)
	c.player.peer.AckEvent(c.eventId)
}

/*
TODO: Implementar escrita assincrona para o peer, para evitar que bloqueie a thread principal
do servidor
*/

/*
TODO: Haverá um evento para pedir a lista de entidades ao redor de uma seção x (raio default 6)
O jogador enviará esse comando a cada 1 segundo e o servidor olhará no spatial map e enviará
para o jogador
*/
