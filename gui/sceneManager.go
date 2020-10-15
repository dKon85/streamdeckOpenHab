package gui

import (
	"github.com/magicmonkey/go-streamdeck"
	"github.com/magicmonkey/go-streamdeck/buttons"
)

const (
	Rows = 3
	Cols = 5
)

type SceneRegistry struct{
	Registry map[string]*Scene
}

func (reg *SceneRegistry) Init( scenes ... *Scene){
	reg.Registry = make(map[string]*Scene)
	for _, scene := range scenes{
		reg.Registry[scene.name] = scene
	}
}

type Scene struct {
	name    string
	buttons [Cols][Rows]*buttons.TextButton
}

func (scene *Scene) init() {
	for x := 0; x < Cols; x++ {
		for y := 0; y < Rows; y++ {
			scene.AddButton(GetEmptyButton(), x, y)
		}
	}
}

func (scene *Scene) AddButton(button *buttons.TextButton, x, y int) {
	scene.buttons[x][y] = button
}

func (scene *Scene) Write(sd streamdeck.StreamDeck) {
	for x := 0; x < Cols; x++ {
		for y := 0; y < Rows; y++ {
			sd.AddButton(x+y*5, scene.buttons[x][y])
		}
	}
}

func GetEmptyButton() *buttons.TextButton {
	return buttons.NewTextButton("")
}
