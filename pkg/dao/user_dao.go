package dao

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"muzz-service/pkg/types"
	"strings"
)

func GetUserById(id int) (types.User, error) {
	query := `
		SELECT id, email, username
		FROM application_users
		WHERE id = $1
	`

	var user types.User
	err := conn.QueryRow(context.Background(), query, id).Scan(&user.ID, &user.Email, &user.Name)
	if err != nil {
		return types.User{}, err
	}

	return user, nil
}

func GetAllUsersExcludingSwipes(id int, filters UsersFilter, sort UsersSort) ([]types.User, error) {
	query := `
		SELECT 
			u.id, u.email, u.username, u.password, u.gender, u.age,
			ST_Distance(location, (select location from application_users where id = $1)) AS distance
		FROM application_users u
		    
		-- filter out the users that have been swiped by the current user
		LEFT JOIN swipes s ON u.id = s.to_id AND s.from_id = $1
		
		WHERE u.id <> $1 
			AND s.to_id IS NULL
	`

	enrichedQuery, params := applyFilters(query, []interface{}{id}, filters)
	enrichedQuery = applySorters(enrichedQuery, sort)

	r, err := conn.Query(context.Background(), enrichedQuery, params...)
	if err != nil {
		return nil, err
	}

	return rowMapperWithDistance(r)
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
		INSERT INTO application_users (email, username, password, gender, age, location)
		VALUES ($1, $2, $3, $4, $5, ST_SetSRID(
					-- insert random point for location			
                    ST_MakePoint(
                        random() * (180 - (-180)) + (-180),
                        random() * (90 - (-90)) + (-90)
                    ),
                    4326
                )) 
		RETURNING id
	`

	err := conn.QueryRow(context.Background(), insert, user.Email, user.Name, user.Password, user.Gender, user.Age).Scan(&user.ID)
	if err != nil {
		return types.User{}, err
	}

	return user, nil
}

func IncrementsLikesForUser(id int) error {
	query := `
		UPDATE application_users
		SET likes = likes + 1
		WHERE id = $1
	`

	_, err := conn.Exec(context.Background(), query, id)
	if err != nil {
		return err
	}

	return nil
}

// private functions

func rowMapperWithDistance(r pgx.Rows) ([]types.User, error) {
	users := make([]types.User, 0)

	for r.Next() {
		var user types.User

		err := r.Scan(&user.ID, &user.Email, &user.Name, &user.Password, &user.Gender, &user.Age, &user.Distance)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
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

func applySorters(query string, sort UsersSort) string {
	var orderClauses []string

	if sort.DistanceSort {
		orderClauses = append(orderClauses, "distance ASC")
	}

	if sort.AttractivenessSort {
		orderClauses = append(orderClauses, "likes DESC")
	}

	if len(orderClauses) > 0 {
		query += " ORDER BY " + strings.Join(orderClauses, ", ")
	}

	return query
}
