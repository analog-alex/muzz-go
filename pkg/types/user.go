package types

import (
	"fmt"
	"strconv"
)

type User struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Gender   string `json:"gender"`
	Age      uint8  `json:"age"`

	// operational data, not part of the "presentation" data
	Distance float64 `json:"-"`
	Likes    int     `json:"-"`
}

type UserCredentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserProfile struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Gender   string  `json:"gender"`
	Age      uint8   `json:"age"`
	Distance float64 `json:"distanceFromMe,omitempty"`
}

func (u *User) ToUserProfile() UserProfile {
	return UserProfile{
		ID:       u.ID,
		Name:     u.Name,
		Gender:   u.Gender,
		Age:      u.Age,
		Distance: u.Distance,
	}
}

func (up *UserProfile) HumanizeDistance() {
	// convert meters to kilometers and round to 2 decimal places
	up.Distance, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", up.Distance/1000), 64)
}
