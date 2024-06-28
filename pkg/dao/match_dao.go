package dao

import (
	"context"
	"muzz-service/pkg/types"
)

func CreateMatch(match types.Match) (types.Match, error) {
	insert := `
		INSERT INTO matches (user1_id, user2_id)
		VALUES ($1, $2) 
	`

	_, err := conn.Exec(context.Background(), insert, match.UserOneID, match.UserTwoID)
	if err != nil {
		return types.Match{}, err
	}

	return match, nil
}
