package database

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
)

// Clear deletes all tables in the specified PostgreSQL database.
func Clear(pool *pgxpool.Pool) error {
	ctx := context.Background()

	// Get a list of all tables in the current schema.
	tablesQuery := "SELECT table_name FROM information_schema.tables WHERE table_schema = current_schema() AND table_type = 'BASE TABLE'"
	rows, err := pool.Query(ctx, tablesQuery)
	if err != nil {
		return fmt.Errorf("error querying tables: %v", err)
	}
	defer rows.Close()

	// Drop each table in the list.
	for rows.Next() {
		var tableName string
		if err := rows.Scan(&tableName); err != nil {
			return fmt.Errorf("error scanning table name: %v", err)
		}

		dropTableQuery := fmt.Sprintf("DROP TABLE IF EXISTS %s CASCADE", tableName)
		if _, err := pool.Exec(ctx, dropTableQuery); err != nil {
			return fmt.Errorf("error dropping table %s: %v", tableName, err)
		}
	}

	if err := rows.Err(); err != nil {
		return fmt.Errorf("error iterating over tables: %v", err)
	}

	fmt.Println("All tables dropped successfully.")
	return nil
}
