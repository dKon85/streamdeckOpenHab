package main

import (
	streamdeck "github.com/magicmonkey/go-streamdeck"
	_ "github.com/magicmonkey/go-streamdeck/devices"
	"streamdeckOpenHab/gui"
	"time"
)

//LIB & Example taken from https://github.com/magicmonkey/go-streamdeck

var run bool = true

func main() {
	sd, err := streamdeck.New()
	if err != nil {
		panic(err)
	}

	sd.SetBrightness(50)

	registry := &gui.SceneRegistry{}
	scene := gui.GetTestScene(sd, registry, stop)
	registry.Init(scene, gui.GetMainScene(sd, registry), gui.GetSettingsScene(sd, registry, stop), gui.GetSleepScene(sd, registry), gui.GetTempScene(sd, registry) )
	scene.Write(*sd)

	for run {
		time.Sleep( 1 * time.Second )
	}

	gui.GetEmptyScene().Write(*sd)
	sd.SetBrightness(0)

}

func stop(){
	run = false
}
