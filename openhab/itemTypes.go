package openhab

type ItemList struct{
	Collection []Item
}

type Item struct {
	Link             string `json:"link"`
	State            string `json:"state"`
	StateDescription struct {
		Pattern  string        `json:"pattern"`
		ReadOnly bool          `json:"readOnly"`
		Options  []interface{} `json:"options"`
	} `json:"stateDescription"`
	Editable   bool          `json:"editable"`
	Type       string        `json:"type"`
	Name       string        `json:"name"`
	Label      string        `json:"label"`
	Category   string        `json:"category"`
	Tags       []interface{} `json:"tags"`
	GroupNames []interface{} `json:"groupNames"`
}