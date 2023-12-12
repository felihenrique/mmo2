package shard

import "mmo2/pkg/events"

type ICommand interface {
	Execute()
}

type MoveCommand struct {
	payload events.MoveEvent
	// acesso a posicao do jogador
	// acesso aos dados de colisao
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
