package db

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"muzz-service/config"
)

var pool *pgxpool.Pool

func GetDB() *pgxpool.Pool {
	return pool
}

func init() {
	c := config.GetApplicationConfig()
	connectionString := fmt.Sprintf(
		"postgresql://%s:%s@%s:%s/%s",
		c.DbUser,
		c.DbPassword,
		c.DbHost,
		c.DbPort,
		c.DbName,
	)

	p, err := pgxpool.New(context.Background(), connectionString)
	if err != nil {
		panic(err)
	}

	pool = p

	// add tables
	// statements are already idempotent
	// in a real world scenario, use a migration tool
	AddTables()
}

func AddTables() {
	var err error

	// create application users table
	//
	//
	_, err = pool.Exec(context.Background(), `
		CREATE TABLE IF NOT EXISTS application_users
		(
			id         SERIAL PRIMARY KEY,
			username   VARCHAR(255) NOT NULL,
			password   VARCHAR(255) NOT NULL,
			email      VARCHAR(255) NOT NULL UNIQUE,
			gender     VARCHAR(1) NOT NULL,
			age        INT NOT NULL,
			location   GEOGRAPHY(POINT, 4326),	
			likes 	   INT NOT NULL DEFAULT 0,
			created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
		);
	`)
	if err != nil {
		panic(err)
	}

	// create swipes table
	//
	//
	_, err = pool.Exec(context.Background(), `
		CREATE TABLE IF NOT EXISTS swipes
		(
			from_id    INT NOT NULL,
			to_id      INT NOT NULL,
			accept     BOOLEAN NOT NULL,
			created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
			PRIMARY KEY (from_id, to_id)
		);
	`)
	if err != nil {
		panic(err)
	}

	// create matches table
	//
	//
	_, err = pool.Exec(context.Background(), `
		CREATE TABLE IF NOT EXISTS matches
		(
			user1_id   INT NOT NULL,
			user2_id   INT NOT NULL,
			created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
			PRIMARY KEY (user1_id, user2_id)
		);
	`)
	if err != nil {
		panic(err)
	}
}
