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

	// A simple yellow button in position 26
	// cButton := buttons.NewColourButton(color.RGBA{255, 255, 0, 255})
	// sd.AddButton(9, cButton)

	// A button with text on it in position 2, which echoes to the console when presesd
	// myButton := buttons.NewTextButton("Hi world")
	// myButton.SetActionHandler(&actionhandlers.TextPrintAction{Label: "You pressed me"})
	// sd.AddButton(2, myButton)

	sd.SetBrightness(50)

	registry := gui.SceneRegistry{}
	// A button with text on it which changes when pressed
	//buttonGenerator.GenerateButtons(sd, stop)
	scene := gui.GetTestScene(sd, &registry, stop)
	registry.Init(scene, gui.GetMainScene(sd, &registry), gui.GetSettingsScene(sd, &registry), gui.GetSleepScene(sd, &registry) )
	scene.Write(*sd)

	for run {
		time.Sleep( 1 * time.Second )
	}

	gui.GetEmptyScene().Write(*sd)
	sd.SetBrightness(0)

	// A button which performs multiple actions when pressed
	// multiActionButton := buttons.NewColourButton(color.RGBA{255, 0, 255, 255})
	// thisActionHandler := &actionhandlers.ChainedAction{}
	// thisActionHandler.AddAction(&actionhandlers.TextPrintAction{Label: "Purple press"})
	// thisActionHandler.AddAction(&actionhandlers.ColourChangeAction{NewColour: color.RGBA{255, 0, 0, 255}})
	// multiActionButton.SetActionHandler(thisActionHandler)
	// sd.AddButton(12, multiActionButton)

	// time.Sleep(40 * time.Second)



}

func stop(){
	run = false
}
