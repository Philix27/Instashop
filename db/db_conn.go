package db

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
)

func Connect(dbUrl string) (*Queries, context.Context, *pgx.Conn) {
	ctx := context.Background()
	// dsn := os.Getenv("DB_URL")
	conn, err := pgx.Connect(ctx, dbUrl)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// defer conn.Close(ctx)

	fmt.Println("Connected to the database!")

	queries := New(conn)

	return queries, ctx, conn

}
