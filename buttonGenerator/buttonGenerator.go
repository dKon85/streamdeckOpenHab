package buttonGenerator

import (
	"github.com/magicmonkey/go-streamdeck"
	"github.com/magicmonkey/go-streamdeck/actionhandlers"
	"github.com/magicmonkey/go-streamdeck/buttons"
	"streamdeckOpenHab/actionHandler"
	"time"
)

type SingleActionButton struct{
	url string
}

func (*SingleActionButton) Pressed(){
	// OpenHabCom.sendCommand( url )
}

func GenerateButtons( sd *streamdeck.StreamDeck, stopFunc func()){
	button1 := buttons.NewTextButton("1")
	button1.SetActionHandler(&actionhandlers.TextLabelChangeAction{NewLabel: "ONE"})
	sd.AddButton(0, button1)

	button2 := buttons.NewTextButton("2")
	button2.SetActionHandler(&actionhandlers.TextLabelChangeAction{NewLabel: "TWO"})
	sd.AddButton(1, button2)

	button3 := buttons.NewTextButton("3")
	button3.SetActionHandler(&actionhandlers.TextLabelChangeAction{NewLabel: "THREE"})
	sd.AddButton(2, button3)

	button4 := buttons.NewTextButton("4")
	button4.SetActionHandler(&actionhandlers.TextLabelChangeAction{NewLabel: "FOUR"})
	sd.AddButton(3, button4)

	button5 := buttons.NewTextButton("5")
	button5.SetActionHandler(&actionhandlers.TextLabelChangeAction{NewLabel: "FIVE"})
	sd.AddButton(4, button5)

	button6 := buttons.NewTextButton("6")
	button6.SetActionHandler(&actionhandlers.TextLabelChangeAction{NewLabel: "SIX"})
	sd.AddButton(5, button6)

	button7 := buttons.NewTextButton("7")
	button7.SetActionHandler(&actionhandlers.TextLabelChangeAction{NewLabel: "SEVEN"})
	sd.AddButton(6, button7)

	button8 := buttons.NewTextButton("8")
	button8.SetActionHandler(&actionhandlers.TextLabelChangeAction{NewLabel: "EIGHT"})
	sd.AddButton(7, button8)

	button9 := buttons.NewTextButton("9")
	button9.SetActionHandler(&actionhandlers.TextLabelChangeAction{NewLabel: "NINE"})
	sd.AddButton(8, button9)

	button10 := buttons.NewTextButton("10")
	button10.SetActionHandler(&actionhandlers.TextLabelChangeAction{NewLabel: "TEN"})
	sd.AddButton(9, button10)

	button11 := buttons.NewTextButton("11")
	button11.SetActionHandler(&actionhandlers.TextLabelChangeAction{NewLabel: "ELEVEN"})
	sd.AddButton(10, button11)

	button12 := buttons.NewTextButton("12")
	button12.SetActionHandler(&actionhandlers.TextLabelChangeAction{NewLabel: "TWELVE"})
	sd.AddButton(11, button12)

	button13 := buttons.NewTextButton("13")
	button13.SetActionHandler(&actionhandlers.TextLabelChangeAction{NewLabel: "THIRTEEN"})
	sd.AddButton(12, button13)

	button14 := buttons.NewTextButton("14")
	button14.SetActionHandler(&actionHandler.StopAppAction{
		func(){
			time.Sleep(5 * time.Second)
			RemoveButtons(sd)
			stopFunc()
		}})
	sd.AddButton(13, button14)

	button15 := buttons.NewTextButton("15")
	button15.SetActionHandler(&actionHandler.StopAppAction{stopFunc} )
	sd.AddButton(14, button15)
}

func RemoveButtons( sd *streamdeck.StreamDeck ){
	sd.AddButton(0, buttons.NewTextButton("") )
	sd.AddButton(1, buttons.NewTextButton(""))
	sd.AddButton(2, buttons.NewTextButton(""))
	sd.AddButton(3, buttons.NewTextButton(""))
	sd.AddButton(4, buttons.NewTextButton(""))
	sd.AddButton(5, buttons.NewTextButton(""))
	sd.AddButton(6, buttons.NewTextButton(""))
	sd.AddButton(7, buttons.NewTextButton(""))
	sd.AddButton(8, buttons.NewTextButton(""))
	sd.AddButton(9, buttons.NewTextButton(""))
	sd.AddButton(10, buttons.NewTextButton(""))
	sd.AddButton(11, buttons.NewTextButton(""))
	sd.AddButton(12, buttons.NewTextButton(""))
	sd.AddButton(13, buttons.NewTextButton(""))
	sd.AddButton(14, buttons.NewTextButton(""))
	sd.SetBrightness(0)
}
