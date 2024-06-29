package types

type User struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Gender   string `json:"gender"`
	Age      uint8  `json:"age"`
}

type UserCredentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserProfile struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Gender string `json:"gender"`
	Age    uint8  `json:"age"`
}

func (u *User) ToUserProfile() UserProfile {
	return UserProfile{
		ID:     u.ID,
		Name:   u.Name,
		Gender: u.Gender,
		Age:    u.Age,
	}
}
