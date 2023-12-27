package scene

import "mmo2/internal/shard-client"

var currentScene IScene

func ChangeTo(scene IScene) {
	if currentScene != nil {
		currentScene.Finalize()
	}
	currentScene = scene
	scene.Init()
}

func RenderGUI(client *shard.Client) {
	if currentScene != nil {
		currentScene.RenderGUI(client)
	}
}
