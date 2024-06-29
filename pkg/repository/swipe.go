package repository

import (
	"context"
	"muzz-service/pkg/types"
)

func CreateSwipe(swipe types.Swipe) (types.Swipe, error) {
	insert := `
		INSERT INTO swipes (from_id, to_id, accept)
		VALUES ($1, $2, $3) 
	`

	_, err := conn.Exec(context.Background(), insert, swipe.From, swipe.To, swipe.Accept)
	if err != nil {
		return types.Swipe{}, err
	}

	return swipe, nil
}

func CheckSwipeRight(from int, to int) (bool, error) {
	query := `
		SELECT accept FROM swipes
		WHERE from_id = $1 AND to_id = $2
	`

	r, err := conn.Query(context.Background(), query, from, to)
	if err != nil {
		return false, err
	}

	var accept bool
	for r.Next() {
		err = r.Scan(&accept)
		if err != nil {
			return false, err
		}
	}

	return accept, nil
}
