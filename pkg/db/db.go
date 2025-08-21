package db

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

var db_url string = "DB_URL"
var connection *pgxpool.Pool

func GetConnection() *pgxpool.Pool {
	if connection != nil {
		return connection
	}

	var err error
	connection, err = pgxpool.New(context.Background(), os.Getenv(db_url))
	if err != nil {
		fmt.Printf("Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}

	return connection
}

func Close() {
	connection.Close()
}
