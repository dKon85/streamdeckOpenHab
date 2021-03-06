package actionHandler

import (
	"fmt"
	"github.com/magicmonkey/go-streamdeck"
	"streamdeckOpenHab/openhab"
)

type OpenHabAction struct{
	//TextButton *streamdeck.Button
	Cmd func()
}

func (action *OpenHabAction) Pressed(btn streamdeck.Button) {
	if action.Cmd != nil {
		action.Cmd()
	/*} else {
		com := openhab.OpenHabCommunicator{}

		state, _ := com.GetItemState()

		action.TextButton.SetText(state)*/
	}
}

func (action *OpenHabAction) listItems(btn streamdeck.Button) {

	com := openhab.OpenHabCommunicator{}

	items, err := com.ListItems()

	if( err != nil ){
		fmt.Errorf("Error in call %s", err)
		return
	}

	//action.TextButton.SetText(fmt.Sprintf("%d", len(*items)))

	fmt.Printf("Found %d Items", len(*items) )
}