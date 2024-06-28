package repository

import "muzz-service/pkg/types"

var users []types.User

func GetAll() []types.User {

	return users
}

func Create(user types.User) types.User {
	user.ID = len(users) + 1
	users = append(users, user)
	return user
}

func GetByEmail(email string) (types.User, bool) {
	for _, user := range users {
		if user.Email == email {
			return user, true
		}
	}
	return types.User{}, false
}
