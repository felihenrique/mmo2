package shard

import (
	"log"
	"mmo2/pkg/events"
	"mmo2/pkg/game"
	"mmo2/pkg/gsp"
	"mmo2/pkg/payloads"
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

func (s *Server) moveRequest(player *Player, pe gsp.PeerEvent) {
	move := payloads.MoveRequest{}
	events.Unserialize(pe.Event, &move)
	tc, tok := player.entity.Get(game.TypeTransform)
	if !tok {
		log.Printf("move command error: entity %d doesn't have transform", player.entity.ID())
		return
	}
	transform := tc.(*game.Transform)
	transform.X += move.Dx
	transform.Y += move.Dy
	player.peer.AckEvent(events.GetId(pe.Event))
}

func (s *Server) joinShardRequest(player *Player, pe gsp.PeerEvent) {
	event := payloads.JoinShardRequest{}
	events.Unserialize(pe.Event, &event)
	entity := s.world.NewEntity()
	player.entity = entity
	transform := game.Transform{}
	switch event.Portal {
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
	player.peer.AckEvent(events.GetId(pe.Event))
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
