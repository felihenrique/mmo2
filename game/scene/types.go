package scene

import "mmo2/internal/shard-client"

type IScene interface {
	Init()
	Finalize()
	RenderGUI(client *shard.Client)
}
