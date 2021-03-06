package gui

import (
	"fmt"
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
	buttons [Cols][Rows]streamdeck.Button
}

func (scene *Scene) init() {
	button := GetEmptyButton()
	for x := 0; x < Cols; x++ {
		for y := 0; y < Rows; y++ {
			scene.AddButton(button, x, y)
		}
	}
}

func (scene *Scene) AddButton(button streamdeck.Button, x, y int) {
	if x > Cols-1 || y > Rows -1 || x < 0 || y < 0{
		fmt.Printf("Button index out of bounds (X: 0-%d|Y: 0-%d). Is (%d|%d)", Cols-1, Rows-1, x, y)
		return
	}
	scene.buttons[x][y] = button
}

func (scene *Scene) Write(sd streamdeck.StreamDeck) {
	sd.SetBrightness(50)
	for x := 0; x < Cols; x++ {
		for y := 0; y < Rows; y++ {
			if scene.buttons[x][y] != nil {
				sd.AddButton(x+y*5, scene.buttons[x][y])
			}
		}
	}
}

func GetEmptyButton() streamdeck.Button {
	return buttons.NewTextButton("")
}
