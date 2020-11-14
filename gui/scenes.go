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
	testSceneName     = "TestScene"
	mainSceneName     = "MainScene"
	emptySceneName    = "EmptyScene"
	settingsSceneName = "SettingsScene"
	sleepSceneName    = "SleepScene"
	tempSceneName     = "tempScene"
)

func GetTestScene(sd *streamdeck.StreamDeck, registry *SceneRegistry, stopFunc func()) *Scene {

	result := Scene{name: testSceneName}

	button1 := buttons.NewTextButton("1")
	button1.SetActionHandler(&actionHandler.OpenHabAction{})
	result.AddButton(button1, 0, 0)

	tempButton := &TempButton{
		Room:      "bedroom",
		ItemNames: []string{"HeaterBedroom"},
		sd:        sd,
		curState:  "cold",
		nextState: "warm",
	}
	result.AddButton(tempButton.GenerateButton(), 1, 0)

	button3 := buttons.NewTextButton("3")
	button3.SetActionHandler(&actionhandlers.TextLabelChangeAction{NewLabel: "THREE"})
	result.AddButton(button3, 2, 0)

	button4 := buttons.NewTextButton("4")
	button4.SetActionHandler(&actionhandlers.TextLabelChangeAction{NewLabel: "FOUR"})
	result.AddButton(button4, 3, 0)

	button5 := buttons.NewTextButton(">")
	button5.SetActionHandler(&SceneAction{mainSceneName, registry, sd})
	result.AddButton(button5, 4, 0)

	button6 := buttons.NewTextButton("6")
	button6.SetActionHandler(&actionhandlers.TextLabelChangeAction{NewLabel: "SIX"})
	result.AddButton(button6, 0, 1)

	button7 := buttons.NewTextButton("7")
	button7.SetActionHandler(&actionhandlers.TextLabelChangeAction{NewLabel: "SEVEN"})
	result.AddButton(button7, 1, 1)

	imgTest, _ := buttons.NewImageFileButton("images/light_bedroom_on.png")
	result.AddButton(imgTest, 2, 1)

	button9 := buttons.NewTextButton("9")
	button9.SetActionHandler(&actionhandlers.TextLabelChangeAction{NewLabel: "NINE"})
	result.AddButton(button9, 3, 1)

	button10 := buttons.NewTextButton("NXT")
	button10.SetActionHandler(&SceneAction{mainSceneName, registry, sd})
	result.AddButton(button10, 4, 1)

	button11 := buttons.NewTextButton("11")
	button11.SetActionHandler(&actionhandlers.TextLabelChangeAction{NewLabel: "ELEVEN"})
	result.AddButton(button11, 0, 2)

	button12 := buttons.NewTextButton("12")
	button12.SetActionHandler(&actionhandlers.TextLabelChangeAction{NewLabel: "TWELVE"})
	result.AddButton(button12, 1, 2)

	button13 := buttons.NewTextButton("13")
	button13.SetActionHandler(&actionhandlers.TextLabelChangeAction{NewLabel: "THIRTEEN"})
	result.AddButton(button13, 2, 2)

	button14 := buttons.NewTextButton("14")
	button14.SetActionHandler(&StopAppAction{
		func() {
			time.Sleep(5 * time.Second)
			stopFunc()
		}})
	result.AddButton(button14, 3, 2)

	button15 := buttons.NewTextButton("15")
	button15.SetActionHandler(&StopAppAction{stopFunc})
	result.AddButton(button15, 4, 2)

	return &result
}

func GetMainScene(sd *streamdeck.StreamDeck, registry *SceneRegistry) *Scene {

	result := Scene{name: mainSceneName}
	result.init()

	buttonBwd := buttons.NewTextButton("<")
	buttonBwd.SetActionHandler(&SceneAction{testSceneName, registry, sd})
	result.AddButton(buttonBwd, 0, 0)

	lightButton, _ := buttons.NewImageFileButton("images/light.png")
	lightButton.SetActionHandler(&SceneAction{mainSceneName, registry, sd})
	result.AddButton(lightButton, 1, 0)

	tempButton, _ := buttons.NewImageFileButton("images/temp_dark.png")
	tempButton.SetActionHandler(&SceneAction{tempSceneName, registry, sd})
	result.AddButton(tempButton, 2, 0)

	settingsButton, _ := buttons.NewImageFileButton("images/settings_dark.png")
	settingsButton.SetActionHandler(&SceneAction{settingsSceneName, registry, sd})
	result.AddButton(settingsButton, 3, 0)

	lightBedroom := &LightButton{
		Room:      "bedroom",
		ItemNames: []string{"LightBedRoom_Color"},
		sd:        sd,
		active:    openhab.IsLightActive("LightBedRoom_Color"),
	}
	result.AddButton(lightBedroom.GenerateButton(), 1, 1)

	lightLivingroom := &LightButton{
		Room:      "livingroom",
		ItemNames: []string{"LightLivingRoom_Color", "LightPlayLivingroomDoor_Color", "LightPlayLivingroomWindow_Color", "LightstripBedroom_Color"},
		sd:        sd,
		active:    openhab.IsLightActive("LightLivingRoom_Color" ),
	}
	result.AddButton(lightLivingroom.GenerateButton(), 1, 2)

	lightKitchen := &LightButton{
		Room:      "kitchen",
		ItemNames: []string{"LightKitchen_Color"},
		sd:        sd,
		active:    openhab.IsLightActive("LightKitchen_Color" ),
	}
	result.AddButton(lightKitchen.GenerateButton(), 3, 2)

	buttonFwd := buttons.NewTextButton(">")
	buttonFwd.SetActionHandler(&SceneAction{tempSceneName, registry, sd})
	result.AddButton(buttonFwd, 4, 0)

	return &result
}

func GetTempScene(sd *streamdeck.StreamDeck, registry *SceneRegistry) *Scene {

	result := Scene{name: tempSceneName}
	result.init()

	lightButton, _ := buttons.NewImageFileButton("images/light_dark.png")
	lightButton.SetActionHandler(&SceneAction{mainSceneName, registry, sd})
	result.AddButton(lightButton, 1, 0)

	tempButton, _ := buttons.NewImageFileButton("images/temp.png")
	tempButton.SetActionHandler(&SceneAction{tempSceneName, registry, sd})
	result.AddButton(tempButton, 2, 0)

	settingsButton, _ := buttons.NewImageFileButton("images/settings_dark.png")
	settingsButton.SetActionHandler(&SceneAction{settingsSceneName, registry, sd})
	result.AddButton(settingsButton, 3, 0)

	bedroomTempState := openhab.TempToName(openhab.ConvertTemperatureToFloat(openhab.GetItemStateWithDefault("HeaterBedroom", "0.0")))
	bedroomTempButton := &TempButton{
		Room:      "bedroom",
		ItemNames: []string{"HeaterBedroom"},
		sd:        sd,
		curState:  bedroomTempState,
		nextState: openhab.DetermineNextTempState(bedroomTempState),
	}
	result.AddButton(bedroomTempButton.GenerateButton(), 1, 1)

	livingroomTempState := openhab.TempToName(openhab.ConvertTemperatureToFloat(openhab.GetItemStateWithDefault("HeaterLivingroomWindow", "0.0")))
	livingroomTempButton := &TempButton{
		Room:      "livingroom",
		ItemNames: []string{"HeaterLivingroomWindow", "HeaterLivingroomDoor"},
		sd:        sd,
		curState:  livingroomTempState,
		nextState: openhab.DetermineNextTempState(livingroomTempState),
	}
	result.AddButton(livingroomTempButton.GenerateButton(), 1, 2)

	bathroomTempState := openhab.TempToName(openhab.ConvertTemperatureToFloat(openhab.GetItemStateWithDefault("HeaterBathroom", "0.0")))
	bathroomTempButton := &TempButton{
		Room:      "bathroom",
		ItemNames: []string{"HeaterBathroom"},
		sd:        sd,
		curState:  bathroomTempState,
		nextState: openhab.DetermineNextTempState(bathroomTempState),
	}
	result.AddButton(bathroomTempButton.GenerateButton(), 2, 2)

	kitchenTempState := openhab.TempToName(openhab.ConvertTemperatureToFloat(openhab.GetItemStateWithDefault("HeaterKitchen", "0.0")))
	kitchenTempButton := &TempButton{
		Room:      "kitchen",
		ItemNames: []string{"HeaterKitchen"},
		sd:        sd,
		curState:  kitchenTempState,
		nextState: openhab.DetermineNextTempState(kitchenTempState),
	}
	result.AddButton(kitchenTempButton.GenerateButton(), 3, 2)

	buttonFwd := buttons.NewTextButton(">")
	buttonFwd.SetActionHandler(&SceneAction{settingsSceneName, registry, sd})
	result.AddButton(buttonFwd, 4, 0)

	buttonBwd := buttons.NewTextButton("<")
	buttonBwd.SetActionHandler(&SceneAction{mainSceneName, registry, sd})
	result.AddButton(buttonBwd, 0, 0)

	return &result
}

func GetEmptyScene() *Scene {

	result := Scene{name: emptySceneName}
	result.init()

	return &result
}

func GetSettingsScene(sd *streamdeck.StreamDeck, registry *SceneRegistry, shutdown func()) *Scene {
	result := Scene{name: settingsSceneName}
	result.init()

	buttonBwd := buttons.NewTextButton("<")
	buttonBwd.SetActionHandler(&SceneAction{tempSceneName, registry, sd})
	result.AddButton(buttonBwd, 0, 0)

	lightButton, _ := buttons.NewImageFileButton("images/light_dark.png")
	lightButton.SetActionHandler(&SceneAction{mainSceneName, registry, sd})
	result.AddButton(lightButton, 1, 0)

	tempButton, _ := buttons.NewImageFileButton("images/temp_dark.png")
	tempButton.SetActionHandler(&SceneAction{tempSceneName, registry, sd})
	result.AddButton(tempButton, 2, 0)

	settingsButton, _ := buttons.NewImageFileButton("images/settings.png")
	settingsButton.SetActionHandler(&SceneAction{settingsSceneName, registry, sd})
	result.AddButton(settingsButton, 3, 0)

	buttonSleep, _ := buttons.NewImageFileButton("images/sleep.png")
	thisActionHandler := &actionhandlers.ChainedAction{}
	thisActionHandler.AddAction(&actionHandler.OpenHabAction{func() { sd.SetBrightness(0) }})
	thisActionHandler.AddAction(&SceneAction{sleepSceneName, registry, sd})
	buttonSleep.SetActionHandler(thisActionHandler)
	result.AddButton(buttonSleep, 1, 1)

	buttonStop := buttons.NewTextButton("Shutdown")
	buttonStop.SetActionHandler(&StopAppAction{shutdown})
	result.AddButton(buttonStop, 3, 2)

	buttonFwd := buttons.NewTextButton(">")
	buttonFwd.SetActionHandler(&SceneAction{testSceneName, registry, sd})
	result.AddButton(buttonFwd, 4, 0)

	return &result
}

func GetSleepScene(sd *streamdeck.StreamDeck, registry *SceneRegistry) *Scene {
	result := Scene{name: sleepSceneName}
	result.init()

	wakeUpButton := GetEmptyButton()

	thisActionHandler := &actionhandlers.ChainedAction{}
	thisActionHandler.AddAction(&actionHandler.OpenHabAction{func() { sd.SetBrightness(50) }})
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
