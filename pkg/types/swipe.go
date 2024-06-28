package types

type Swipe struct {
	Accept bool `json:"success"` // i.e. right swipe
	From   int  `json:"from"`
	To     int  `json:"to"`
}
