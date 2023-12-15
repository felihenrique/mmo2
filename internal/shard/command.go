package shard

import (
	"log"
	"mmo2/pkg/game"
	"mmo2/pkg/gsp"
	"mmo2/pkg/payloads"
	"mmo2/pkg/serialization"
)

type ICommand interface {
	Execute()
}
type Broadcast = func(event serialization.ISerializable)
type BroadcastFiltered = func(event serialization.ISerializable, filterPeer *gsp.TcpPeer)
type AckRequest = func(eventId int16)

/*
TODO:
- Checa se a posição de destino da entidade contem: outra entidade, colisor do mapa
- Caso não tenha adiciona dx, dy à posição do jogador
- Caso tenha se movido, lê o spatial map para detectar as outras entidades do tipo player
próximo do jogador
- Envia o evento EntityMoved para todos os jogadores em um radio de 6 seções
(a tela do jogador 4 x 4 seções)
*/
type MoveCommand struct {
	event     payloads.MoveRequest
	eventId   int16
	player    *Player
	broadcast BroadcastFiltered
}

func (c *MoveCommand) Execute() {
	tc, tok := c.player.entity.Get(game.TypeTransform)
	if !tok {
		log.Printf("move command error: entity %d doesn't have transform", c.player.entity.ID())
		return
	}
	transform := tc.(*game.Transform)
	transform.X += c.event.Dx
	transform.Y += c.event.Dy
	c.player.peer.AckEvent(c.eventId)
}

type JoinShardCommand struct {
	event     payloads.JoinShardRequest
	eventId   int16
	player    *Player
	world     *game.World
	broadcast BroadcastFiltered
}

func (c *JoinShardCommand) Execute() {
	entity := c.world.NewEntity()
	c.player.entity = entity
	transform := game.Transform{}
	switch c.event.Portal {
	case 0:
		transform.X = 0
		transform.Y = 0
	case 1:
		transform.X = 100
		transform.Y = 100
	case 2:
		transform.X = 200
		transform.Y = 200
	default:
		transform.X = 0
		transform.Y = 0
	}
	entity.Add(&transform)
	c.player.peer.AckEvent(c.eventId)
}

/*
TODO: Haverá um evento para pedir a lista de entidades ao redor de uma seção x (raio default 6)
O jogador enviará esse comando a cada 1 segundo e o servidor olhará no spatial map e enviará
para o jogador
*/

/*
Haverá um commando para cada tipo de evento. Na instancia do command (no metodo OnEvent)
serão colocados todas as referências que aquele comando precisa, como por exemplo o estado
do jogo, jogador que executou, entre outras coisas.
No comando de movimento, serão checados também as colisões com obstaculos e entidades.
Primeiramente serão processados todos os inputs, depois será processada a lógica do jogo
Por exemplo, checar se um monstro está próximo do jogador e caso esteja atacar ele.
*/
