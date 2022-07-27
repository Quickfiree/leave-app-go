package dbConnection

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
)

func DBConn() *pgxpool.Pool {
	dbURL := "postgres://postgres:yourpasswordhere@localhost:5432/postgres"
	pgxpool, error := pgxpool.Connect(context.Background(), dbURL) // returns either a pointer or error
	if error != nil {
		fmt.Println("An error has occurred: ", error)
	}

	return pgxpool
}
