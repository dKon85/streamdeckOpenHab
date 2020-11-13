package gui

import (
	"github.com/magicmonkey/go-streamdeck"
	"github.com/magicmonkey/go-streamdeck/actionhandlers"
	"github.com/magicmonkey/go-streamdeck/buttons"
	"streamdeckOpenHab/openhab"
	"streamdeckOpenHab/openhab/actionHandler"
	"time"
)

const (
	testSceneName = "TestScene"
	mainSceneName = "MainScene"
	emptySceneName = "EmptyScene"
	settingsSceneName = "SettingsScene"
	sleepSceneName = "SleepScene"
	tempSceneName = "tempScene"
)

func GetTestScene( sd *streamdeck.StreamDeck, registry *SceneRegistry, stopFunc func()) (*Scene){

	result := Scene{name: testSceneName}

	button1 := buttons.NewTextButton("1")
	button1.SetActionHandler(&actionHandler.OpenHabAction{})
	result.AddButton( button1, 0,0 )

	/*button2 := buttons.NewTextButton("2")
	button2.SetActionHandler(&actionhandlers.TextLabelChangeAction{NewLabel: "TWO"})
	result.AddButton( button2, 1,0)*/

	tempButton := &TempButton{
		Room:     "bedroom",
		ItemNames: []string{"HeaterBedroom"},
		sd: sd,
		curState: "cold",
		nextState: "warm",
	}
	result.AddButton(tempButton.GenerateButton(), 1, 0)

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

	imgTest, _ := buttons.NewImageFileButton("images/light_bedroom_on.png")
	result.AddButton(imgTest, 2, 1)

	//imgTest.SetActionHandler(&actionhandlers.TextLabelChangeAction{NewLabel: "EIGHT"})

	/*button8 := buttons.NewTextButton("8")
	button8.SetActionHandler(&actionhandlers.TextLabelChangeAction{NewLabel: "EIGHT"})
	result.AddButton(button8,2 ,1)*/

	button9 := buttons.NewTextButton("9")
	button9.SetActionHandler(&actionhandlers.TextLabelChangeAction{NewLabel: "NINE"})
	result.AddButton(button9, 3,1)

	button10 := buttons.NewTextButton("NXT")
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


	buttonBedroom, _ := buttons.NewImageFileButton("images/light_bedroom_on.png")
	buttonBedroom.SetActionHandler(&actionHandler.OpenHabAction{ func(){openhab.ToggleLight("LightBedRoom_Color")}})
	result.AddButton(buttonBedroom, 1, 0)

	buttonLivingroom, _ := buttons.NewImageFileButton("images/light_livingroom_on.png")
	buttonLivingroom.SetActionHandler(&actionHandler.OpenHabAction{ func(){openhab.ToggleLights("LightLivingRoom_Color", "LightPlayLivingroomDoor_Color", "LightPlayLivingroomWindow_Color", "LightstripBedroom_Color")}})
	result.AddButton(buttonLivingroom, 1, 1)

	buttonKitchen, _ := buttons.NewImageFileButton("images/light_kitchen_on.png")
	buttonKitchen.SetActionHandler(&actionHandler.OpenHabAction{ func(){openhab.ToggleLight("LightKitchen_Color")}})
	result.AddButton(buttonKitchen, 3, 1)

	buttonFwd := buttons.NewTextButton(">")
	buttonFwd.SetActionHandler(&SceneAction{tempSceneName, registry, sd})
	result.AddButton(buttonFwd, 4, 0)

	return &result
}


func GetTempScene( sd *streamdeck.StreamDeck, registry *SceneRegistry) (*Scene) {

	result := Scene{name: tempSceneName}
	result.init()

	bedroomTempButton := &TempButton{
		Room:     "bedroom",
		ItemNames: []string{"HeaterBedroom"},
		sd: sd,
		curState: "cold",
		nextState: "warm",
	}
	result.AddButton(bedroomTempButton.GenerateButton(), 1, 0)

	livingroomTempButton := &TempButton{
		Room:     "livingroom",
		ItemNames: []string{"HeaterLivingroomWindow", "HeaterLivingroomDoor"},
		sd: sd,
		curState: "cold",
		nextState: "warm",
	}
	result.AddButton(livingroomTempButton.GenerateButton(), 1, 1)

	bathroomTempButton := &TempButton{
		Room:     "bathroom",
		ItemNames: []string{"HeaterBathroom"},
		sd: sd,
		curState: "cold",
		nextState: "warm",
	}
	result.AddButton(bathroomTempButton.GenerateButton(), 2, 1)

	kitchenTempButton := &TempButton{
		Room:     "kitchen",
		ItemNames: []string{"HeaterKitchen"},
		sd: sd,
		curState: "cold",
		nextState: "warm",
	}
	result.AddButton(kitchenTempButton.GenerateButton(), 3, 1)

	buttonFwd := buttons.NewTextButton(">")
	buttonFwd.SetActionHandler(&SceneAction{settingsSceneName, registry, sd})
	result.AddButton(buttonFwd, 4, 0)

	buttonBwd := buttons.NewTextButton("<")
	buttonBwd.SetActionHandler(&SceneAction{mainSceneName, registry, sd})
	result.AddButton(buttonBwd, 0, 0)

	return &result
}


func GetEmptyScene( ) (*Scene) {

	result := Scene{name: emptySceneName}
	result.init()

	return &result
}

func GetSettingsScene( sd *streamdeck.StreamDeck, registry *SceneRegistry, shutdown func()) (*Scene) {
	result := Scene{name: settingsSceneName}
	result.init()

	buttonBwd := buttons.NewTextButton("<")
	buttonBwd.SetActionHandler(&SceneAction{mainSceneName, registry, sd})
	result.AddButton(buttonBwd, 0, 0)

	buttonSleep , _ := buttons.NewImageFileButton("images/sleep.png")
	thisActionHandler := &actionhandlers.ChainedAction{}
	thisActionHandler.AddAction(&actionHandler.OpenHabAction{func(){sd.SetBrightness(0)}})
	thisActionHandler.AddAction(&SceneAction{sleepSceneName, registry, sd})
	buttonSleep.SetActionHandler(thisActionHandler)
	result.AddButton(buttonSleep, 1, 1)

	buttonStop := buttons.NewTextButton("Shutdown")
	buttonStop.SetActionHandler(&StopAppAction{shutdown})
	result.AddButton(buttonStop, 3,2)

	buttonFwd := buttons.NewTextButton(">")
	buttonFwd.SetActionHandler(&SceneAction{testSceneName, registry, sd})
	result.AddButton(buttonFwd, 4, 0)

	return &result
}

func GetSleepScene( sd *streamdeck.StreamDeck, registry *SceneRegistry) (*Scene) {
	result := Scene{name: sleepSceneName}
	result.init()


	wakeUpButton := GetEmptyButton()

	thisActionHandler := &actionhandlers.ChainedAction{}
	thisActionHandler.AddAction(&actionHandler.OpenHabAction{func(){sd.SetBrightness(50)}})
	thisActionHandler.AddAction(&SceneAction{mainSceneName, registry, sd})
	wakeUpButton.(*buttons.TextButton).SetActionHandler(thisActionHandler)

	for x := 0; x < Cols; x++ {
		for y := 0; y < Rows; y++ {
			result.AddButton(wakeUpButton, x, y)
		}
	}

	sd.SetBrightness(0)

	return &result
}