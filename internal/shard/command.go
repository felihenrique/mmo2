package shard

import (
	"log"
	"mmo2/pkg/events"
	"mmo2/pkg/game"
)

type ICommand interface {
	Execute()
}
type BroadcastFunc = func(event events.ISerializable, eventId int16)

type MoveCommand struct {
	event     events.Move
	eventId   int16
	world     *game.World
	player    Player
	broadcast BroadcastFunc
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
	if !entity.Has(game.TypeTransform) {
		log.Printf("move command error: entity %d doesn't have transform", entity.ID())
		return
	}
	transform := entity.Get(game.TypeTransform).(game.Transform)
	transform.X += c.event.Dx
	transform.Y += c.event.Dy
	entity.Add(transform)
	c.event.Dx = transform.X
	c.event.Dy = transform.Y
	c.broadcast(&c.event, c.eventId)
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
