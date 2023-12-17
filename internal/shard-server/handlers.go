package shard

import (
	"log"
	"mmo2/pkg/events"
	"mmo2/pkg/game"
	"mmo2/pkg/gsp"
	"mmo2/pkg/packets"
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
	move := packets.MoveInput{}
	events.Unserialize(pe.Event, &move)
	tc, tok := player.entity.Get(game.TypePosition)
	if !tok {
		log.Printf("move command error: entity %d doesn't have position", player.entity.ID())
		return
	}
	position := tc.(*game.Position)
	position.X += move.Dx
	position.Y += move.Dy
	s.ackInput(pe.Event, pe.Peer)
}

func (s *Server) joinShardRequest(player *Player, pe gsp.PeerEvent) {
	event := packets.JoinShardRequest{}
	events.Unserialize(pe.Event, &event)
	entity := s.world.NewEntity()
	player.entity = entity
	position := game.Position{}
	switch event.Portal {
	case 0:
		position.X = 0
		position.Y = 0
	case 1:
		position.X = 100
		position.Y = 100
	case 2:
		position.X = 200
		position.Y = 200
	default:
		position.X = 0
		position.Y = 0
	}
	entity.Add(&position)
	s.ackInput(pe.Event, pe.Peer)
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
