package gui

import (
	"github.com/magicmonkey/go-streamdeck"
)

type SceneAction struct{
	TargetName string
	SceneRegistry *SceneRegistry
	Display *streamdeck.StreamDeck
}

func (action *SceneAction) Pressed(btn streamdeck.Button) {
	action.SceneRegistry.Registry[action.TargetName].Write(*action.Display)
}