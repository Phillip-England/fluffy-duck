package database

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
)

func GetCemScores(pool *pgxpool.Pool) ([][]interface{}, error) {
	// Get a connection from the pool
	conn, err := pool.Acquire(context.Background())
	if err != nil {
		return nil, fmt.Errorf("error acquiring connection from pool: %v", err)
	}
	defer conn.Release()

	// Query the cem_score table
	rows, err := conn.Query(context.Background(), `
		SELECT timescale, osat, taste, speed, ace, cleanliness, accuracy
		FROM cem_score;
	`)
	if err != nil {
		return nil, fmt.Errorf("error querying cem_score table: %v", err)
	}
	defer rows.Close()

	// Iterate through the rows and populate a slice of slices
	var cemScores [][]interface{}
	for rows.Next() {
		var timescale string
		var osat, taste, speed, ace, cleanliness, accuracy float64
		if err := rows.Scan(
			&timescale,
			&osat,
			&taste,
			&speed,
			&ace,
			&cleanliness,
			&accuracy,
		); err != nil {
			return nil, fmt.Errorf("error scanning row: %v", err)
		}
		rowData := []interface{}{timescale, osat, taste, speed, ace, cleanliness, accuracy}
		cemScores = append(cemScores, rowData)
	}

	// Check for errors from iterating over rows
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over rows: %v", err)
	}

	return cemScores, nil
}
