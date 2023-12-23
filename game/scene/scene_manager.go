package scene

var currentScene IScene

func Start(scene IScene) {
	if currentScene == nil {
		currentScene = scene
	}
	scene.Init()
}
