package gui

import (
	"github.com/magicmonkey/go-streamdeck"
	"github.com/magicmonkey/go-streamdeck/buttons"
)


type StopAppAction struct {
	StopFunc func()
}

func (action *StopAppAction) Pressed(btn streamdeck.Button) {
	mybtn := btn.(*buttons.TextButton)
	mybtn.SetText("BYE")
	action.StopFunc()
}

