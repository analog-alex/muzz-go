package dao

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"muzz-service/pkg/types"
	"strings"
)

func GetAllUsers() ([]types.User, error) {
	query := `
		SELECT id, email, username, password, gender, age 
		FROM application_users
	`
	r, err := conn.Query(context.Background(), query)

	if err != nil {
		return nil, err
	}

	return rowMapper(r)
}

func GetAllUsersExcludingSwipes(id int, filters UsersFilter) ([]types.User, error) {
	query := `
		SELECT u.id, u.email, u.username, u.password, u.gender, u.age
		FROM application_users u
		    
		-- filter out the users that have been swiped by the current user
		LEFT JOIN swipes s ON u.id = s.to_id AND s.from_id = $1
		
		WHERE u.id <> $1 
			AND s.to_id IS NULL
	`

	enrichedQuery, params := applyFilters(query, []interface{}{id}, filters)

	r, err := conn.Query(context.Background(), enrichedQuery, params...)
	if err != nil {
		fmt.Println(err)
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

// private functions

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

func applyFilters(query string, params []interface{}, filters UsersFilter) (string, []interface{}) {
	var conditions []string

	if filters.MinAge != "" {
		conditions = append(conditions, "u.age >= $"+fmt.Sprint(len(params)+1))
		params = append(params, filters.MinAge)
	}

	if filters.MaxAge != "" {
		conditions = append(conditions, "u.age <= $"+fmt.Sprint(len(params)+1))
		params = append(params, filters.MaxAge)
	}

	if filters.Gender != "" {
		conditions = append(conditions, "u.gender = $"+fmt.Sprint(len(params)+1))
		params = append(params, filters.Gender)
	}

	// Combine conditions if there are any
	if len(conditions) > 0 {
		query += " AND " + strings.Join(conditions, " AND ")
	}

	return query, params
}
