package repository

import (
	"context"
	"github.com/jackc/pgx/v5"
	"muzz-service/pkg/types"
)

func GetAllUsers() ([]types.User, error) {
	query := "SELECT id, email, username, password, gender, age FROM application_users"
	r, err := conn.Query(context.Background(), query)

	if err != nil {
		return nil, err
	}

	return rowMapper(r)
}

func GetAllUsersExcludingSwipes(id int) ([]types.User, error) {
	query := `
		SELECT u.id, u.email, u.username, u.password, u.gender, u.age
		FROM application_users u
		-- filter out the users that have been swiped by the current user
		LEFT JOIN swipes s ON u.id = s.to_id AND s.from_id = $1
		WHERE u.id <> $1 
			AND s.to_id IS NULL
	`

	r, err := conn.Query(context.Background(), query, id)
	if err != nil {
		return nil, err
	}

	return rowMapper(r)
}

func GetUsersByEmail(email string) ([]types.User, error) {
	query := `
		SELECT id, email, username, password, gender, age FROM application_users
		WHERE email = $1
	`

	r, err := conn.Query(context.Background(), query, email)
	if err != nil {
		return nil, err
	}

	return rowMapper(r)
}

func CreateUser(user types.User) (types.User, error) {
	insert := `
		INSERT INTO application_users (email, username, password, gender, age)
		VALUES ($1, $2, $3, $4, $5) 
		RETURNING id
	`

	err := conn.QueryRow(context.Background(), insert, user.Email, user.Name, user.Password, user.Gender, user.Age).Scan(&user.ID)
	if err != nil {
		return types.User{}, err
	}

	return user, nil
}

func rowMapper(r pgx.Rows) ([]types.User, error) {
	users := make([]types.User, 0)

	for r.Next() {
		var user types.User

		err := r.Scan(&user.ID, &user.Email, &user.Name, &user.Password, &user.Gender, &user.Age)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}
