package database

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
)

func Init(pool *pgxpool.Pool) error {
	err := CreateCemTable(pool)
	if err != nil {
		return err
	}
	return nil
}

func CreateCemTable(pool *pgxpool.Pool) error {
	// SQL statement to create the cem_score table
	createTableSQL := `
	CREATE TABLE IF NOT EXISTS cem_score (
		id SERIAL PRIMARY KEY,
		timescale VARCHAR,
		osat NUMERIC,
		taste NUMERIC,
		speed NUMERIC,
		ace NUMERIC,
		cleanliness NUMERIC,
		accuracy NUMERIC
	);
	`

	// Get a connection from the pool
	conn, err := pool.Acquire(context.Background())
	if err != nil {
		return fmt.Errorf("error acquiring connection from pool: %v", err)
	}
	defer conn.Release()

	// Execute the SQL statement to create the table
	_, err = conn.Exec(context.Background(), createTableSQL)
	if err != nil {
		return fmt.Errorf("error creating cem_score table: %v", err)
	}

	// Insert rows if the table was just created
	if _, err := conn.Exec(context.Background(), "SELECT COUNT(*) FROM cem_score"); err != nil {
		return fmt.Errorf("error checking row count: %v", err)
	}

	var rowCount int
	if err := conn.QueryRow(context.Background(), "SELECT COUNT(*) FROM cem_score").Scan(&rowCount); err != nil {
		return fmt.Errorf("error getting row count: %v", err)
	}

	if rowCount == 0 {
		// Insert rows
		_, err := conn.Exec(context.Background(), `
		INSERT INTO cem_score (timescale, osat, taste, speed, ace, cleanliness, accuracy)
		VALUES 
			('current month', 100, 100, 100, 100, 100, 100),
			('three month rolling', 100, 100, 100, 100, 100, 100),
			('year to date', 100, 100, 100, 100, 100, 100),
			('annual goal', 100, 100, 100, 100, 100, 100);
		`)
		if err != nil {
			return fmt.Errorf("error inserting rows: %v", err)
		}
	}

	return nil
}

