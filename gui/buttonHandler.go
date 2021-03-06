package gui

import (
	"fmt"
	"github.com/magicmonkey/go-streamdeck"
	"github.com/magicmonkey/go-streamdeck/actionhandlers"
	"github.com/magicmonkey/go-streamdeck/buttons"
	"streamdeckOpenHab/openhab"
)

type TempButton struct {
	Room      string
	ItemNames []string
	sd        *streamdeck.StreamDeck
	curState  string
	nextState string
}

func (receiver TempButton) GenerateButton() *buttons.ImageFileButton {

	if receiver.curState == "" {
		receiver.curState = "cold"
		receiver.nextState = "warm"
	}

	path := fmt.Sprintf("%s/%s_%s_%s.png", FileBasePath, receiver.Room, receiver.curState, receiver.nextState)
	btn, _ := buttons.NewImageFileButton(path)
	handler := actionhandlers.CustomAction{}
	handler.SetHandler(receiver.Pressed)
	btn.SetActionHandler(&handler)
	return btn
}

func (receiver TempButton) Pressed(btn streamdeck.Button) {

	switch receiver.nextState {
	case "warm":
		openhab.SetTemps(fmt.Sprintf("%.2f", openhab.TempWarmLimit), receiver.ItemNames)
		receiver.nextState = "cold"
		receiver.curState = "warm"
	case "cold":
		openhab.SetTemps(fmt.Sprintf("%.2f", openhab.TempColdLimit), receiver.ItemNames)
		receiver.nextState = "warm"
		receiver.curState = "cold"
	default:
		openhab.SetTemps(fmt.Sprintf("%.2f", openhab.TempWarmLimit), receiver.ItemNames)
		receiver.nextState = "cold"
		receiver.curState = "off"
	}

	buttonIndex := btn.GetButtonIndex()
	button := receiver.GenerateButton()

	receiver.sd.AddButton(buttonIndex, button)
}

type LightButton struct {
	Room      string
	ItemNames []string
	sd        *streamdeck.StreamDeck
	active    bool
}

func (receiver LightButton) GenerateButton() *buttons.ImageFileButton {

	postfix := "off"
	if receiver.active {
		postfix = "on"
	}

	path := fmt.Sprintf("%s/light_%s_%s.png", FileBasePath, receiver.Room, postfix)
	btn, _ := buttons.NewImageFileButton(path)
	handler := actionhandlers.CustomAction{}
	handler.SetHandler(receiver.Pressed)
	btn.SetActionHandler(&handler)
	return btn
}

func (receiver LightButton) Pressed(btn streamdeck.Button) {

	openhab.SetLightStates(receiver.ItemNames, !receiver.active)
	receiver.active = !receiver.active

	buttonIndex := btn.GetButtonIndex()
	button := receiver.GenerateButton()

	receiver.sd.AddButton(buttonIndex, button)
}
