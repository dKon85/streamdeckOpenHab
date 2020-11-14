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

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/items", c.BaseURL), nil)
	if err != nil {
		return nil, err
	}

	res := [] Item{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil

}

func GetItemStateWithDefault(itemName, defaultValue string) (string){
	state, err := GetItemState(itemName)

	if err != nil {
		return defaultValue
	}
	return state
}

func GetItemState(itemName string) (string, error){

	c := NewClient("")

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/items/%s/state", c.BaseURL, itemName), nil)
	if err != nil {
		return "", err
	}

	res, err := c.sendPlainRequest(req)

	if err != nil {
		return "", err
	}

	return res, nil

}

func IsLightActive( itemName string ) bool{
	state := GetLightState(itemName)

	if len(state) != 3 {
		return false
	}

	if state[2] == "0" {
		return false
	} else {
		return true
	}
}

func SetLightStates( itemNames []string, activate bool ){

	if len(itemNames) == 0 {
		return
	}

	value := "0"

	if activate {
		value = "60"
	}

	for _, itemName := range itemNames {
		changeBrightness( itemName, value )
	}
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

func SetTemps( targetTemp string, itemNames []string){
	if len(itemNames) == 0 {
		return
	}

	for _, itemName := range itemNames {
		SetTemp(targetTemp, itemName)
	}
}