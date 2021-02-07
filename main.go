package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"strconv"
)

// Concurrency safe Postgres connection pool. Going global here for such
// a simple application.
var dbcp *pgxpool.Pool
var dsn PostgresDSN

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Printf("Error loading .env file, ignoring...")
	}

	dsn = NewPostgresDSN()
	if !dsn.Valid() {
		log.Fatalf("Unable to Generate Postgresql DSN. Missing environment variables?")
	}
}

func main() {
	var err error

	dbcp, err = pgxpool.Connect(context.Background(), dsn.ToString())
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer dbcp.Close()

	// Run migration
	conn, err := dbcp.Acquire(context.Background())
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Unable to get DB connection: %v\n", err)
	}
	defer conn.Release()

	_, err = conn.Query(context.Background(), PersonMigration)
	if err != nil {
		panic(err)
	}

	port := 8080
	http.HandleFunc("/status", statusHandler)
	http.HandleFunc("/persons", getPersonsHandler)
	http.HandleFunc("/create", createPersonHandler)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(port), nil))

}
