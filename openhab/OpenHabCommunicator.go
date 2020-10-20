package openhab

import (
	"bytes"
	"fmt"
	"net/http"
	"strings"
)

type OpenHabCommunicator struct{

}

func (*OpenHabCommunicator) ListItems() (*[]Item, error){

	c := NewClient("")

	//ctx := context.Background()

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/items", c.BaseURL), nil)
	if err != nil {
		return nil, err
	}

	//req = req.WithContext(ctx)


	res := [] Item{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil

}

func (*OpenHabCommunicator) GetItemState() (string, error){

	c := NewClient("")

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/items/XiaomiBedroomHTPSensor_Humidity/state", c.BaseURL), nil)
	if err != nil {
		return "", err
	}

	res, err := c.sendPlainRequest(req)

	if err != nil {
		return "", err
	}

	return res, nil

}
func ToggleLights( itemNames ... string ){

	if len(itemNames) == 0 {
		return
	}

	state := GetLightState(itemNames[0])

	if len(state) != 3 {
		return
	}

	value := "0"

	if state[2] == "0" {
		value = "60"
	}

	for _, itemName := range itemNames {
		changeBrightness( itemName, value )
	}
}

func ToggleLight( itemName string ){

	state := GetLightState(itemName)

	if len(state) != 3 {
		return
	}

	value := "0"

	if state[2] == "0" {
		value = "60"
	}

	changeBrightness(itemName, value)

}

func changeBrightness(itemName string, value string) {
	c := NewClient("")

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/items/%s", c.BaseURL, itemName), bytes.NewBufferString(fmt.Sprintf("30,60,%s", value)))
	if err != nil {
		return
	}

	c.sendPlainRequest(req)
}

func GetLightState( itemName string ) []string{
	c := NewClient("")

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/items/%s/state", c.BaseURL, itemName), nil)
	if err != nil {
		return []string{}
	}

	res, _ := c.sendPlainRequest(req)

	return strings.Split(res, ",")

}

func SetTemp( targetTemp, itemName string ){
	c := NewClient("")

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/items/%s_4_SetTemperature", c.BaseURL, itemName), bytes.NewBufferString(targetTemp))
	if err != nil {
		return
	}

	c.sendPlainRequest(req)
}

func SetTemps( targetTemp string, itemNames ... string){
	if len(itemNames) == 0 {
		return
	}

	for _, itemName := range itemNames {
		SetTemp(targetTemp, itemName)
	}
}