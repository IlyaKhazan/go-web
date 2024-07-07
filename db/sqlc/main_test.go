package db

import (
	"context"
	_ "github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/jackc/pgx/v5/pgxpool"
	"log"
	"os"
	"testing"
)

var testQueries *Queries

func TestMain(m *testing.M) {

	connPool, err := pgxpool.New(context.Background(), "postgresql://root:secret@localhost:5432/webserver?sslmode=disable")

	if err != nil {
		log.Fatal("Невозможно подключиться к БД", err)
	}
	testQueries = New(connPool)

	os.Exit(m.Run())
}
