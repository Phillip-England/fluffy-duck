package database

import (
	"context"
	"fluffy-duck/pkg/models"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
)

func GetCemScoresByTimescale(pool *pgxpool.Pool, timescale string) (*models.CemScore, error) {
	// Get a connection from the pool
	conn, err := pool.Acquire(context.Background())
	if err != nil {
		return nil, fmt.Errorf("error acquiring connection from pool: %v", err)
	}
	defer conn.Release()

	// Query the cem_score table for the specified timescale
	row := conn.QueryRow(context.Background(), `
		SELECT timescale, osat, taste, speed, ace, cleanliness, accuracy
		FROM cem_score
		WHERE timescale = $1;
	`, timescale)

	// Populate the struct with the retrieved values
	var cemScore models.CemScore
	err = row.Scan(
		&cemScore.Timescale,
		&cemScore.Osat,
		&cemScore.Taste,
		&cemScore.Speed,
		&cemScore.Ace,
		&cemScore.Cleanliness,
		&cemScore.Accuracy,
	)

	if err != nil {
		return nil, fmt.Errorf("error scanning row: %v", err)
	}

	return &cemScore, nil
}
