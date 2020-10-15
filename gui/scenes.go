package gui

import (
	"github.com/magicmonkey/go-streamdeck"
	"github.com/magicmonkey/go-streamdeck/actionhandlers"
	"github.com/magicmonkey/go-streamdeck/buttons"
	"streamdeckOpenHab/openhab/actionHandler"
	"time"
)

const (
	testSceneName = "TestScene"
	mainSceneName = "MainScene"
	emptySceneName = "EmptyScene"
)

func GetTestScene( sd *streamdeck.StreamDeck, registry *SceneRegistry, stopFunc func()) (*Scene){

	result := Scene{name: testSceneName}

	button1 := buttons.NewTextButton("1")
	button1.SetActionHandler(&actionHandler.OpenHabAction{button1})
	result.AddButton( button1, 0,0 )

	button2 := buttons.NewTextButton("2")
	button2.SetActionHandler(&actionhandlers.TextLabelChangeAction{NewLabel: "TWO"})
	result.AddButton( button2, 1,0)

	button3 := buttons.NewTextButton("3")
	button3.SetActionHandler(&actionhandlers.TextLabelChangeAction{NewLabel: "THREE"})
	result.AddButton(button3, 2, 0)

	button4 := buttons.NewTextButton("4")
	button4.SetActionHandler(&actionhandlers.TextLabelChangeAction{NewLabel: "FOUR"})
	result.AddButton( button4,3,0)

	button5 := buttons.NewTextButton(">")
	button5.SetActionHandler(&SceneAction{mainSceneName, registry, sd})
	result.AddButton(button5, 4,0 )

	button6 := buttons.NewTextButton("6")
	button6.SetActionHandler(&actionhandlers.TextLabelChangeAction{NewLabel: "SIX"})
	result.AddButton(button6, 0,1)

	button7 := buttons.NewTextButton("7")
	button7.SetActionHandler(&actionhandlers.TextLabelChangeAction{NewLabel: "SEVEN"})
	result.AddButton( button7, 1, 1)

	button8 := buttons.NewTextButton("8")
	button8.SetActionHandler(&actionhandlers.TextLabelChangeAction{NewLabel: "EIGHT"})
	result.AddButton(button8,2 ,1)

	button9 := buttons.NewTextButton("9")
	button9.SetActionHandler(&actionhandlers.TextLabelChangeAction{NewLabel: "NINE"})
	result.AddButton(button9, 3,1)

	button10 := buttons.NewTextButton("NXT")
	/*custom := actionhandlers.CustomAction{}
	custom.SetHandler( func(streamdeck.Button) {

		scene := Scene{name: "initScene" }
		scene.init()
		scene.Write( *sd )} )*/

	button10.SetActionHandler(&SceneAction{mainSceneName, registry, sd})
	result.AddButton(button10,4,1)

	button11 := buttons.NewTextButton("11")
	button11.SetActionHandler(&actionhandlers.TextLabelChangeAction{NewLabel: "ELEVEN"})
	result.AddButton(button11,0,2)

	button12 := buttons.NewTextButton("12")
	button12.SetActionHandler(&actionhandlers.TextLabelChangeAction{NewLabel: "TWELVE"})
	result.AddButton(button12, 1,2)

	button13 := buttons.NewTextButton("13")
	button13.SetActionHandler(&actionhandlers.TextLabelChangeAction{NewLabel: "THIRTEEN"})
	result.AddButton(button13,2,2)

	button14 := buttons.NewTextButton("14")
	button14.SetActionHandler(&StopAppAction{
		func() {
			time.Sleep(5 * time.Second)
			stopFunc()
		}})
	result.AddButton(button14, 3,2)

	button15 := buttons.NewTextButton("15")
	button15.SetActionHandler(&StopAppAction{stopFunc})
	result.AddButton(button15, 4,2)

	return &result
}

func GetMainScene( sd *streamdeck.StreamDeck, registry *SceneRegistry) (*Scene) {

	result := Scene{name: mainSceneName}
	result.init()

	buttonBwd := buttons.NewTextButton("<")
	buttonBwd.SetActionHandler(&SceneAction{testSceneName, registry, sd})
	result.AddButton(buttonBwd, 0, 0)

	buttonFwd := buttons.NewTextButton(">")
	buttonFwd.SetActionHandler(&actionHandler.OpenHabAction{buttonFwd})
	result.AddButton(buttonFwd, 4, 0)

	return &result
}

func GetEmptyScene( ) (*Scene) {

	result := Scene{name: emptySceneName}
	result.init()

	return &result
}