package types

type Swipe struct {
	Accept bool `json:"success"` // i.e. right swipe
	From   int  `json:"from"`
	To     int  `json:"to"`
}

type SwipeRequest struct {
	UserId     int    `json:"userId"`
	Preference string `json:"preference"`
}

type SwipeInnerResponse struct {
	Matched bool `json:"matched"`
	MatchId *int `json:"matchID,omitempty"`
}

type SwipeResponse struct {
	Result SwipeInnerResponse `json:"results"`
}
