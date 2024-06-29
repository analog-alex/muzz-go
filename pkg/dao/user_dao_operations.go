package dao

import (
	"fmt"
	"strconv"
)

type UsersFilter struct {
	MinAge string
	MaxAge string
	Gender string
}

type UsersSort struct {
	DistanceSort bool
}

func (uf *UsersFilter) Validate() error {
	if uf.MinAge != "" {
		if _, err := strconv.Atoi(uf.MinAge); err != nil {
			return fmt.Errorf("min_age must be a number")
		}
	}

	if uf.MaxAge != "" {
		if _, err := strconv.Atoi(uf.MaxAge); err != nil {
			return fmt.Errorf("max_age must be a number")
		}
	}

	if uf.MinAge != "" && uf.MaxAge != "" {
		minAge, _ := strconv.Atoi(uf.MinAge)
		maxAge, _ := strconv.Atoi(uf.MaxAge)

		if minAge > maxAge {
			return fmt.Errorf("min_age must be less than max_age")
		}
	}

	if uf.Gender != "" {
		if len(uf.Gender) > 1 {
			return fmt.Errorf("gender params is invalid")
		}
	}

	return nil
}
