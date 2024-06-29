package types

type Swipe struct {
	Accept bool `json:"success"` // i.e. right swipe
	From   int  `json:"from"`
	To     int  `json:"to"`
}

type SwipeRequest struct {
	UserId     int    `json:"user_id"`
	Preference string `json:"preference"`
}

type SwipeResponse struct {
	Matched bool `json:"matched"`
	MatchId *int `json:"match_id,omitempty"`
}
