package openhab

import (
	"fmt"
	"net/http"
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
