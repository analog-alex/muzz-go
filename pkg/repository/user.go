package repository

import "muzz-service/pkg/entities"

var users []entities.User

func GetAll() []entities.User {
	return users

}

func Create(user entities.User) entities.User {
	user.ID = len(users) + 1
	users = append(users, user)
	return user
}

func GetByEmail(email string) (entities.User, bool) {
	for _, user := range users {
		if user.Email == email {
			return user, true
		}
	}
	return entities.User{}, false
}
