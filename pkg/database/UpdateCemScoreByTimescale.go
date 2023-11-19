package database

import (
	"context"
	"fluffy-duck/pkg/models"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
)

// UpdateCemScoreByTimescale updates the CemScore in the database for the specified timescale.
func UpdateCemScoreByTimescale(pool *pgxpool.Pool, cemScore *models.CemScore) error {
	// Get a connection from the pool
	conn, err := pool.Acquire(context.Background())
	if err != nil {
		return fmt.Errorf("error acquiring connection from pool: %v", err)
	}
	defer conn.Release()

	// Update the cem_score table for the specified timescale
	_, err = conn.Exec(context.Background(), `
		UPDATE cem_score
		SET osat = $2, taste = $3, speed = $4, ace = $5, cleanliness = $6, accuracy = $7
		WHERE timescale = $1;
	`, cemScore.Timescale, cemScore.Osat, cemScore.Taste, cemScore.Speed, cemScore.Ace, cemScore.Cleanliness, cemScore.Accuracy)

	if err != nil {
		return fmt.Errorf("error updating cem_score: %v", err)
	}

	return nil
}
