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

func (MoveCommand) Execute() {

}
